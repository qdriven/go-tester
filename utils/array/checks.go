package array

import (
	"go-tester/utils/math"
	"reflect"
	"strings"
)

// IntsHas check the []int contains the given value
func IntsHas(ints []int, val int) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

// Int64sHas check the []int64 contains the given value
func Int64sHas(ints []int64, val int64) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

// StringsHas check the []string contains the given element
func StringsHas(ss []string, val string) bool {
	for _, ele := range ss {
		if ele == val {
			return true
		}
	}
	return false
}

// HasValue check array(strings, intXs, uintXs) should be contains the given value(int(X),string).
func HasValue(arr, val interface{}) bool {
	return Contains(arr, val)
}

// Contains check array(strings, intXs, uintXs) should be contains the given value(int(X),string).
func Contains(arr, val interface{}) bool {
	if val == nil || arr == nil {
		return false
	}

	// if is string value
	if strVal, ok := val.(string); ok {
		if ss, ok := arr.([]string); ok {
			return StringsHas(ss, strVal)
		}

		rv := reflect.ValueOf(arr)
		if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
			for i := 0; i < rv.Len(); i++ {
				if v, ok := rv.Index(i).Interface().(string); ok && strings.EqualFold(v, strVal) {
					return true
				}
			}
		}

		return false
	}

	// as int value
	intVal, err := math.Int64(val)
	if err != nil {
		return false
	}

	if int64s, err := ToInt64s(arr); err == nil {
		return Int64sHas(int64s, intVal)
	}
	return false
}

// NotContains check array(strings, ints, uints) should be not contains the given value.
func NotContains(arr, val interface{}) bool {
	return false == Contains(arr, val)
}
