package geojson

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"slices"
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
	loc              locale.Interface
	typeList         []string
	message          string
	wrongTypeMessage string
}

func New(typeList ...string) rule.Interface {
	if len(typeList) == 0 {
		typeList = []string{
			Point,
			LineString,
			Polygon,
			MultiPoint,
			MultiLineString,
			MultiPolygon,
		}
	}
	return &Rule{
		message:          "types.geojson.valid",
		wrongTypeMessage: "types.geojson.type",
		typeList:         typeList,
	}
}

func (r *Rule) SetLocale(v locale.Interface) rule.Interface {
	r.loc = v
	return r
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	r.message = r.getErrorMessage(r.message, key)
	r.wrongTypeMessage = r.getErrorMessage(r.wrongTypeMessage, key)
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
	if !slices.Contains(r.typeList, t) {
		return errors.New(r.wrongTypeMessage)
	}
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

func NewFeatureCollection() rule.Interface {
	return &Rule{
		message:  "types.geojson.feature.collection",
		typeList: []string{FeatureCollection},
	}
}

func NewPoint() rule.Interface {
	return &Rule{
		message:  "types.geojson.point",
		typeList: []string{Point},
	}
}

func NewLineString() rule.Interface {
	return &Rule{
		message:  "types.geojson.line.string",
		typeList: []string{LineString},
	}
}

func NewPolygon() rule.Interface {
	return &Rule{
		message:  "types.geojson.polygon",
		typeList: []string{Polygon},
	}
}

func NewMultiPoint() rule.Interface {
	return &Rule{
		message:  "types.geojson.multi.point",
		typeList: []string{MultiPoint},
	}
}

func NewMultiLineString() rule.Interface {
	return &Rule{
		message:  "types.geojson.multi.line.string",
		typeList: []string{MultiLineString},
	}
}

func NewMultiPolygon() rule.Interface {
	return &Rule{
		message:  "types.geojson.multi.polygon",
		typeList: []string{MultiPolygon},
	}
}

func (r *Rule) getErrorMessage(message, key string) string {
	if r.loc == nil {
		return message
	}

	return fmt.Sprintf(r.loc.Translate(message), key)
}
