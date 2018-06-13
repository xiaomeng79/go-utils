//这是一个context设置了值之后取值的一些函数
//设置值:
// ctx = context.Background()
// ctx = context.WithValue(ctx, key, val)
// 然后通过ctx.Value(key)取值
package context

import (
	"context"
	"errors"
	"time"
)

var (
	typeError = errors.New("type error")
)

//取字符串值
func GetString(ctx context.Context, key string) (string, error) {
	if v, ok := ctx.Value(key).(string); ok {
		return v, nil
	}
	return "", typeError
}

//取设置的bool值
func GetBool(ctx context.Context, key string) (bool, error) {
	if v, ok := ctx.Value(key).(bool); ok {
		return v, nil
	}
	return false, typeError
}

//取int类型的值
func GetInt(ctx context.Context, key string) (int, error) {
	if v, ok := ctx.Value(key).(int); ok {
		return v, nil
	}
	return 0, typeError
}

//取int64的值，注意设置的时候一定为int64
func GetInt64(ctx context.Context, key string) (int64, error) {
	if v, ok := ctx.Value(key).(int64); ok {
		return v, nil
	}
	return 0, typeError
}
//取float64的值
func GetFloat64(ctx context.Context, key string) (float64, error) {
	if v, ok := ctx.Value(key).(float64); ok {
		return v, nil
	}
	return 0.00, typeError
}
//取出设置的时间值
func GetTime(ctx context.Context, key string) (time.Time, error) {
	if v, ok := ctx.Value(key).(time.Time); ok {
		return v, nil
	}
	return time.Now(), typeError
}
//取出字符串slice
func GetStringSlice(ctx context.Context, key string) ([]string, error) {
	if v, ok := ctx.Value(key).([]string); ok {
		return v, nil
	}
	return nil, typeError
}
//取出key为string，值为interface的值
func GetStringMap(ctx context.Context, key string) (map[string]interface{}, error) {
	if v, ok := ctx.Value(key).(map[string]interface{}); ok {
		return v, nil
	}
	return nil, typeError
}
//取出mapstring类型的值
func GetStringMapString(ctx context.Context, key string) (map[string]string, error) {
	if v, ok := ctx.Value(key).(map[string]string); ok {
		return v, nil
	}
	return nil, typeError
}
//可以再扩展,记得加单元测试
