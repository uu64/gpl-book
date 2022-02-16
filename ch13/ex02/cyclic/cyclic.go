package cyclic

import (
	"reflect"
	"unsafe"
)

func isCyclic(x reflect.Value, seen map[checked]bool) bool {
	if !x.IsValid() {
		return false
	}

	// cycle check
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := checked{xptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCyclic(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCyclic(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCyclic(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCyclic(x.MapIndex(k), seen) {
				return true
			}
		}
		return false

	default:
		return false
	}
}

func IsCyclic(x interface{}) bool {
	seen := make(map[checked]bool)
	return isCyclic(reflect.ValueOf(x), seen)
}

type checked struct {
	x unsafe.Pointer
	t reflect.Type
}
