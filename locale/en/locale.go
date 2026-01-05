package en

import (
	localeInterface "github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/locale"
)

func New() localeInterface.Interface {
	return locale.Locale{
		"types.url.http":                   "The %s format is not HTTP URL.",
		"types.uint8":                      "The %s must be an uint8.",
		"types.uint16":                     "The %s must be an uint16.",
		"types.uint32":                     "The %s must be an uint32.",
		"types.uint64":                     "The %s must be an uint64.",
		"types.str":                        "The %s must be an string.",
		"types.object":                     "The %s must be an object.",
		"types.ip.v4":                      "The %s must be a valid IPv4 address.",
		"types.int8":                       "The %s must be an int8.",
		"types.int16":                      "The %s must be an int16.",
		"types.int32":                      "The %s must be an int32.",
		"types.int64":                      "The %s must be an int64.",
		"types.geojson.valid":              "The %s must be an valid GeoJSON.",
		"types.geojson.type":               "The %s has wrong GeoJSON type.",
		"types.geojson.feature.collection": "The %s must be an valid GeoJSON FeatureCollection.",
		"types.geojson.point":              "The %s must be an valid GeoJSON Point.",
		"types.geojson.line.string":        "The %s must be an valid GeoJSON LineString.",
		"types.geojson.polygon":            "The %s must be an valid GeoJSON Polygon.",
		"types.geojson.multi.point":        "The %s must be an valid GeoJSON MultiPoint.",
		"types.geojson.multi.line.string":  "The %s must be an valid GeoJSON MultiLineString.",
		"types.geojson.multi.polygon":      "The %s must be an valid GeoJSON MultiPolygon.",
		"types.float":                      "The %s must be an float.",
		"types.date":                       "The %s does not match the format %s.",
		"types.boolean":                    "The %s must be an boolean.",
		"types.array":                      "The %s must be an array.",
		"required":                         "The %s field is required.",
		"not.empty":                        "The %s must not be empty.",
		"min.array":                        "The %s must have at least %d items.",
		"min.string":                       "The %s must be at least %d characters.",
		"min.numeric":                      "The %s must be at least %d.",
		"max.array":                        "The %s must not have more than %d items.",
		"max.string":                       "The %s must not be greater than %d characters.",
		"max.numeric":                      "The %s must not be greater than %d.",
		"in":                               "The selected %s is invalid.",
		"empty":                            "The %s must be empty.",
		"each":                             "The %s must be an array.",
	}
}
