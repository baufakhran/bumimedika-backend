package util

import (
	"reflect"
	"sort"
)

func DeleteMapByKey[T any](m map[uint64]T, key uint64) {
	if _, exists := m[key]; exists {
		delete(m, key)
	}
}

func GetIDsFromMapKeys[T any](mapIDs map[uint64]T) []uint64 {
	ids := make([]uint64, 0, len(mapIDs))
	for id := range mapIDs {
		ids = append(ids, id)
	}
	return ids
}

func DistinctSliceUint(in []uint64) []uint64 {
	seen := make(map[uint64]struct{}, len(in))
	out := make([]uint64, 0, len(in))
	for _, v := range in {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

func ContainsKeyUint(m map[uint64]struct{}, key uint64) bool {
	_, found := m[key]
	return found
}

func KeysFromMapUint(m map[uint64]struct{}) []uint64 {
	keys := make([]uint64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapUintToSlice[T any](m map[uint64]T) []T {
	result := make([]T, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// stringOrEmpty menghindari nil pointer pada string
func StringOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
func DistinctSlice[T comparable](in []T) []T {
	seen := make(map[T]struct{})
	out := make([]T, 0, len(in))
	for _, v := range in {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

func IsExistInSlice(value interface{}, slice interface{}) bool {
	sliceValue := reflect.ValueOf(slice)
	valueToFind := reflect.ValueOf(value)
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		if reflect.DeepEqual(item.Interface(), valueToFind.Interface()) {
			return true
		}
	}

	return false
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func SortNumericSlice[T Number](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func FindByID[T any](slice []T, id uint64) (T, bool) {
	for _, item := range slice {
		if v, ok := any(item).(interface {
			GetID() uint64
		}); ok && v.GetID() == id {
			return item, true
		}
	}
	var zero T
	return zero, false
}
