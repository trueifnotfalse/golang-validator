package geojson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "Point",
				"coordinates": []any{4, 3},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestPointNegativeType(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "Point",
				"coordinates": []any{4, 3},
			},
		},
	}
	r := New(Polygon)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The %s has wrong GeoJSON type.", err.Error())
	}
}

func TestPointNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "Point",
				"coordinates": []any{4, 7, 33},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestPointSingleCoordinateNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "Point",
				"coordinates": []any{4},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestLineStringPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "LineString",
				"coordinates": []any{[]any{3, 4}, []any{7, 8}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestLineStringSinglePointNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "LineString",
				"coordinates": []any{[]any{3, 4}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestLineStringStringCoordinateNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "LineString",
				"coordinates": []any{[]any{"qwe", "asd"}, []any{7, 8}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestLineStringEmptyNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "LineString",
				"coordinates": []any{},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestPolygonPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "Polygon",
				"coordinates": []any{[]any{[]any{3, 4}, []any{7, 8}, []any{8, 9}, []any{3, 4}}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestPolygonWithCutPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type": "Polygon",
				"coordinates": []any{
					[]any{[]any{1, 1}, []any{1, 10}, []any{10, 10}, []any{10, 1}, []any{1, 1}},
					[]any{[]any{3, 4}, []any{7, 8}, []any{8, 9}, []any{3, 4}},
				},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestPolygonFirstAndLastPointsNotEqualNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "Polygon",
				"coordinates": []any{[]any{[]any{3, 4}, []any{7, 8}, []any{8, 9}, []any{3, 41}}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestPolygonLengthLessThanFourNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "Polygon",
				"coordinates": []any{[]any{[]any{3, 4}, []any{7, 8}, []any{3, 4}}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestPolygonWithCutLengthLessThanFourNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type": "Polygon",
				"coordinates": []any{
					[]any{[]any{1, 1}, []any{1, 10}, []any{10, 10}, []any{10, 1}, []any{1, 1}},
					[]any{[]any{3, 4}, []any{7, 8}, []any{3, 4}},
				},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestPolygonWithCutFirstAndLastPointsNotEqualNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type": "Polygon",
				"coordinates": []any{
					[]any{[]any{1, 1}, []any{1, 10}, []any{10, 10}, []any{10, 1}, []any{1, 1}},
					[]any{[]any{3, 4}, []any{7, 8}, []any{8, 8}, []any{32, 4}},
				},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestPolygonEmptyNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "Polygon",
				"coordinates": []any{},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestMultiPointPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "MultiPoint",
				"coordinates": []any{[]any{3, 4}, []any{7, 8}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestMultiPointSinglePointPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "MultiPoint",
				"coordinates": []any{[]any{3, 4}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestMultiPointPolygonNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "MultiPoint",
				"coordinates": []any{[]any{[]any{3, 4}, []any{7, 8}, []any{8, 9}, []any{3, 4}}},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestMultiPointEmptyNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "MultiPoint",
				"coordinates": []any{},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestMultiLineStringPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type": "MultiLineString",
				"coordinates": []any{
					[]any{[]any{1, 1}, []any{1, 10}, []any{10, 10}, []any{10, 1}, []any{11, 14}},
					[]any{[]any{3, 4}, []any{7, 8}, []any{8, 9}, []any{3, 4}},
				},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestMultiLineStringEmptyNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "MultiLineString",
				"coordinates": []any{},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestMultiPolygonPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type": "MultiPolygon",
				"coordinates": []any{
					[]any{
						[]any{[]any{1, 1}, []any{1, 10}, []any{10, 10}, []any{10, 1}, []any{1, 1}},
						[]any{[]any{3, 4}, []any{7, 8}, []any{8, 9}, []any{3, 4}},
					},
					[]any{
						[]any{[]any{100, 14}, []any{112, 14}, []any{110, 17}, []any{100, 1}, []any{100, 14}},
						[]any{[]any{33, 41}, []any{71, 88}, []any{88, 90}, []any{33, 41}},
					},
				},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestMultiPolygonEmptyNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "Feature",
			"geometry": map[string]any{
				"type":        "MultiPolygon",
				"coordinates": []any{},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an valid GeoJSON.", err.Error())
	}
}

func TestFeatureCollectionPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"type": "FeatureCollection",
			"features": []map[string]any{
				{
					"type": "Feature",
					"geometry": map[string]any{
						"type":        "Point",
						"coordinates": []any{4, 3},
					},
				},
				{
					"type": "Feature",
					"geometry": map[string]any{
						"type":        "Polygon",
						"coordinates": []any{[]any{[]any{3, 4}, []any{7, 8}, []any{8, 9}, []any{3, 4}}},
					},
				},
			},
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}
