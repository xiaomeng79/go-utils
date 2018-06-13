package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	ctx           = context.Background()
	contextValues = []struct {
		key string
		val interface{}
	}{
		{"string", "this is string"},
		{"nil", nil},
		{"string_nil", ""},
		{"bool_true", true},
		{"bool_false", false},
		{"int", int(-123)},
		{"uint", uint(123)},
		{"int64", int64(9223372036854775807)},
		{"float64", float64(1.23)},
		{"float32", float32(1.23)},
		{"time", time.Now()},
		{"time_duration", time.Second},
		{"stringslice", []string{"foo", "bar"}},
		{"intslice", []int{123, -123}},
		{"stringmap", map[string]interface{}{"foo": "123", "bar": 123}},
		{"stringmapstring", map[string]string{"foo": "bar", "bar": "foo"}},
	}
)

func init() {
	for _, v := range contextValues {
		ctx = context.WithValue(ctx, v.key, v.val)
	}
}

func TestGetString(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetString(ctx, v.key)
		if v.key == "string" {
			assert.Equalf(t, v.val, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else if v.key == "string_nil" {
			assert.Zerof(t, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

func TestGetBool(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetBool(ctx, v.key)
		if v.key == "bool_true" {
			assert.Truef(t, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else if v.key == "bool_false" {
			assert.Falsef(t, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

func TestGetInt(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetInt(ctx, v.key)
		if v.key == "int" {
			assert.Exactlyf(t, v.val, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

func TestGetInt64(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetInt64(ctx, v.key)
		if v.key == "int64" {
			assert.Exactlyf(t, v.val, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

func TestGetFloat64(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetFloat64(ctx, v.key)
		if v.key == "float64" {
			assert.Exactlyf(t, v.val, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

func TestGetTime(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetTime(ctx, v.key)
		if v.key == "time" {
			assert.Exactlyf(t, v.val, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

func TestGetStringSlice(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetStringSlice(ctx, v.key)
		if v.key == "stringslice" {
			assert.Exactlyf(t, v.val, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

func TestGetStringMap(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetStringMap(ctx, v.key)
		if v.key == "stringmap" {
			assert.Exactlyf(t, v.val, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

func TestGetStringMapString(t *testing.T) {
	for _, v := range contextValues {
		res, err := GetStringMapString(ctx, v.key)
		if v.key == "stringmapstring" {
			assert.Exactlyf(t, v.val, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}
