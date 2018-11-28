package wannabe

import (
	"os"
	"reflect"
	"strings"
)

// BoolEnv makes an environment variaable be bool.
func BoolEnv(key string) bool {
	if v, ok := os.LookupEnv(key); ok {
		return Bool(v)
	}
	return false
}

// Bool makes numeric, string and other types be bool.
func Bool(v interface{}) bool {
	switch v := v.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case int8:
		return v != 0
	case int16:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case uint:
		return v != 0
	case uint8:
		return v != 0
	case uint16:
		return v != 0
	case uint32:
		return v != 0
	case uint64:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	case string:
		switch strings.TrimSpace(v) {
		case "1", "t", "T", "true", "TRUE", "on", "ON", "y", "Y", "yes", "YES":
			return true
		case "0", "f", "F", "false", "FALSE", "off", "OFF", "n", "N", "no", "NO":
			return false
		}
		return false
	case nil:
		return false
	default:
		return !isZero(reflect.ValueOf(v))
	}
	panic("unreachable")
}

func isZero(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Array, reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Map, reflect.Slice:
		return rv.IsNil() || rv.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return rv.IsNil()
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			if !isZero(rv.Field(i)) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}
