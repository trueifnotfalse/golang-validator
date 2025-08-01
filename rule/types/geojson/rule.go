package geojson

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/utils"
)

const (
	FeatureCollection = "FeatureCollection"
	feature           = "Feature"
	Point             = "Point"
	LineString        = "LineString"
	Polygon           = "Polygon"
	MultiPoint        = "MultiPoint"
	MultiLineString   = "MultiLineString"
	MultiPolygon      = "MultiPolygon"
)

type Rule struct {
	typeList []string
	message  string
}

func New(typeList ...string) *Rule {
	return &Rule{
		message:  "The %s must be an valid GeoJSON.",
		typeList: typeList,
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	r.message = fmt.Sprintf(r.message, key)
	m, ok := utils.ToMap(v)
	if !ok {
		return errors.New(r.message)
	}
	v, ok = m["type"]
	if !ok {
		return errors.New(r.message)
	}
	s := utils.ToString(v)
	switch s {
	case FeatureCollection:
		return r.validateFeatureCollection(m)
	case feature:
		return r.validateFeature(m)
	}

	return errors.New(r.message)
}

func (r *Rule) validateFeatureCollection(v map[string]any) error {
	a, ok := v["features"]
	if !ok {
		return errors.New(r.message)
	}
	s, ok := utils.ToSlice(a)
	if !ok {
		return errors.New(r.message)
	}
	var err error
	for i := 0; i < len(s); i++ {
		v, ok = utils.ToMap(s[i])
		if !ok {
			return errors.New(r.message)
		}
		err = r.validateFeature(v)
		if err != nil {
			return errors.New(r.message)
		}
	}

	return nil
}

func (r *Rule) validateFeature(v map[string]any) error {
	a, ok := v["geometry"]
	if !ok {
		return errors.New(r.message)
	}

	return r.validateGeometry(a)
}

func (r *Rule) validateGeometry(v any) error {
	m, ok := utils.ToMap(v)
	if !ok {
		return errors.New(r.message)
	}
	v, ok = m["type"]
	if !ok {
		return errors.New(r.message)
	}
	t := utils.ToString(v)
	v, ok = m["coordinates"]
	if !ok {
		return errors.New(r.message)
	}
	switch t {
	case Point:
		return r.validatePoint(v)
	case LineString:
		return r.validateLineString(v)
	case Polygon:
		return r.validatePolygon(v)
	case MultiPoint:
		return r.validateMultiPoint(v)
	case MultiLineString:
		return r.validateMultiLineString(v)
	case MultiPolygon:
		return r.validateMultiPolygon(v)
	}

	return errors.New(r.message)
}

func (r *Rule) validatePoint(v any) error {
	s, ok := utils.ToFloat64Slice(v)
	if !ok {
		return errors.New(r.message)
	}
	if len(s) != 2 {
		return errors.New(r.message)
	}
	return nil
}

func (r *Rule) validateLineString(v any) error {
	s, err := r.getCoordinatesAsSlice(v)
	if err != nil {
		return err
	}
	if len(s) < 2 {
		return errors.New(r.message)
	}
	for i := 0; i < len(s); i++ {
		err = r.validatePoint(s[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Rule) validatePolygon(v any) error {
	s, err := r.getCoordinatesAsSlice(v)
	if err != nil {
		return err
	}
	var (
		ls []any
		ok bool
	)
	for i := 0; i < len(s); i++ {
		ls, ok = utils.ToSlice(s[i])
		if !ok {
			return errors.New(r.message)
		}
		if len(ls) < 4 {
			return errors.New(r.message)
		}
		err = r.validatePolygonPart(ls)
		if err != nil {
			return errors.New(r.message)
		}
	}
	return nil
}

func (r *Rule) validatePolygonPart(v []any) error {
	var (
		s, first, last []float64
		ok             bool
	)
	for i := 0; i < len(v); i++ {
		s, ok = utils.ToFloat64Slice(v[i])
		if !ok {
			return errors.New(r.message)
		}
		if len(first) == 0 {
			first = s
		}
		if len(s) != 2 {
			return errors.New(r.message)
		}
		last = s
	}
	if !r.isPointsEqual(first, last) {
		return errors.New(r.message)
	}

	return nil
}

func (r *Rule) isPointsEqual(first, second []float64) bool {
	if first[0] == second[0] && first[1] == second[1] {
		return true
	}
	return false
}

func (r *Rule) validateMultiPoint(v any) error {
	s, err := r.getCoordinatesAsSlice(v)
	if err != nil {
		return err
	}
	for i := 0; i < len(s); i++ {
		err = r.validatePoint(s[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Rule) validateMultiLineString(v any) error {
	s, err := r.getCoordinatesAsSlice(v)
	if err != nil {
		return err
	}
	for i := 0; i < len(s); i++ {
		err = r.validateLineString(s[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Rule) validateMultiPolygon(v any) error {
	s, err := r.getCoordinatesAsSlice(v)
	if err != nil {
		return err
	}
	for i := 0; i < len(s); i++ {
		err = r.validatePolygon(s[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Rule) getCoordinatesAsSlice(v any) ([]any, error) {
	s, ok := utils.ToSlice(v)
	if !ok || len(s) == 0 {
		return nil, errors.New(r.message)
	}
	return s, nil
}

func NewFeatureCollection() *Rule {
	return &Rule{
		message:  "The %s must be an valid GeoJSON FeatureCollection.",
		typeList: []string{FeatureCollection},
	}
}

func NewPoint() *Rule {
	return &Rule{
		message:  "The %s must be an valid GeoJSON Point.",
		typeList: []string{Point},
	}
}

func NewLineString() *Rule {
	return &Rule{
		message:  "The %s must be an valid GeoJSON LineString.",
		typeList: []string{LineString},
	}
}

func NewPolygon() *Rule {
	return &Rule{
		message:  "The %s must be an valid GeoJSON Polygon.",
		typeList: []string{Polygon},
	}
}

func NewMultiPoint() *Rule {
	return &Rule{
		message:  "The %s must be an valid GeoJSON MultiPoint.",
		typeList: []string{MultiPoint},
	}
}

func NewMultiLineString() *Rule {
	return &Rule{
		message:  "The %s must be an valid GeoJSON MultiLineString.",
		typeList: []string{MultiLineString},
	}
}

func NewMultiPolygon() *Rule {
	return &Rule{
		message:  "The %s must be an valid GeoJSON MultiPolygon.",
		typeList: []string{MultiPolygon},
	}
}
