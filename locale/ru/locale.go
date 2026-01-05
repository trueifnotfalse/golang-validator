package ru

import (
	localeInterface "github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/locale"
)

func New() localeInterface.Interface {
	return locale.Locale{
		"types.url.http":                   "Формат %s не соответствует HTTP-URL.",
		"types.uint8":                      "%s должен быть типом uint8.",
		"types.uint16":                     "%s должен быть типом uint16.",
		"types.uint32":                     "%s должен быть типом uint32.",
		"types.uint64":                     "%s должен быть типом uint64.",
		"types.str":                        "%s должен быть типом string.",
		"types.object":                     "%s должен быть типом object.",
		"types.ip.v4":                      "%s должен быть действительным IPv4-адресом.",
		"types.int8":                       "%s должен быть типом int8.",
		"types.int16":                      "%s должен быть типом int16.",
		"types.int32":                      "%s должен быть типом int32.",
		"types.int64":                      "%s должен быть типом int64.",
		"types.geojson.valid":              "%s должен представлять собой допустимый GeoJSON.",
		"types.geojson.type":               "В переменной %s указан неверный тип GeoJSON.",
		"types.geojson.feature.collection": "%s должен представлять собой допустимую коллекцию объектов GeoJSON.",
		"types.geojson.point":              "%s должен представлять собой допустимую точку GeoJSON.",
		"types.geojson.line.string":        "%s должен представлять собой допустимую строку линии GeoJSON.",
		"types.geojson.polygon":            "%s должен представлять собой допустимый полигон GeoJSON.",
		"types.geojson.multi.point":        "%s должен представлять собой допустимый объект GeoJSON MultiPoint.",
		"types.geojson.multi.line.string":  "%s должен быть допустимым объектом GeoJSON MultiLineString.",
		"types.geojson.multi.polygon":      "%s должен представлять собой допустимый многоугольник GeoJSON.",
		"types.float":                      "%s должен быть числом с плавающей запятой.",
		"types.date":                       "%s не соответствует формату %s.",
		"types.boolean":                    "%s должен быть логическим значением.",
		"types.array":                      "%s должен быть массивом.",
		"required":                         "%s обязательно для заполнения.",
		"not.empty":                        "%s не должно быть пустым.",
		"min.array":                        "В блоке %s должно быть не менее %d элементов.",
		"min.string":                       "Символы %s должны состоять как минимум из %d символов.",
		"min.numeric":                      "Значение %s должно быть не менее %d.",
		"max.array":                        "В блоке %s не должно быть более %d элементов.",
		"max.string":                       "Количество символов %s не должно превышать %d.",
		"max.numeric":                      "Значение %s не должно превышать %d.",
		"in":                               "%s недействителен.",
		"empty":                            "%s должно быть пустым.",
		"each":                             "%s должен быть массивом.",
	}
}
