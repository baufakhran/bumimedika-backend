package util

import (
	"reflect"
	"strings"
)

func GetJSONFieldName(obj interface{}, field string) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if f, ok := t.FieldByName(field); ok {
		tag := f.Tag.Get("json")
		if tag != "" && tag != "-" {
			return tag
		}
	}

	return field
}
func Chain(funcs ...func() error) error {
	for _, fn := range funcs {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func ContainsKey(m map[uint64]struct{}, key uint64) bool {
	_, found := m[key]
	return found
}

func MapToSlice[T any](m map[uint64]T) []T {
	result := make([]T, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func MapUint64ToString(id uint64, m map[uint64]string) string {
	if name, ok := m[id]; ok {
		return name
	}
	return ""
}

func FindFieldByJSONTag(structValue reflect.Value, jsonTagName string) reflect.Value {
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		jsonTag := field.Tag.Get("json")

		// Parse JSON tag (handle cases like "field_name,omitempty")
		if jsonTag != "" {
			tagParts := strings.Split(jsonTag, ",")
			if len(tagParts) > 0 && tagParts[0] == jsonTagName {
				return structValue.Field(i)
			}
		}
	}

	return reflect.Value{}
}

func CastNullableInt64toUint64(val *int64) *uint64 {
	if val == nil {
		return nil
	}
	u := uint64(*val)
	return &u
}

func CastNullableUint64(val *uint64) uint64 {
	if val == nil {
		return 0
	}
	u := uint64(*val)
	return u
}

func DefaultString(val *string, defaultVal string) string {
	if val == nil || *val == "" {
		return defaultVal
	}
	return *val
}

func DefaultRawString(val string, defaultVal string) string {
	if val == "" {
		return defaultVal
	}
	return val
}
