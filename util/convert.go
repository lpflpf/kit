package util

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"math"
	"reflect"
	"strconv"
	"time"
)

func Int(src interface{}) int {
	return int(Int64(src))
}

func Int64(src interface{}) (dst int64) {
	var ok bool
	var err error
	if dst, ok = src.(int64); !ok {
		switch src := src.(type) {
		case int32:
			return int64(src)
		case int:
			return int64(src)
		case int16:
			return int64(src)
		case int8:
			return int64(src)
		case uint16:
			return int64(src)
		case uint32:
			return int64(src)
		case uint64:
			return int64(src)
		case uint8:
			return int64(src)
		case []byte:
			dst, err = strconv.ParseInt(string(src), 10, 64)
			if err != nil {
				panic(fmt.Sprintf("cannot convert %s to int64, err:%s", src, err.Error()))
			}
		case string:
			dst, err = strconv.ParseInt(src, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("cannot convert %s to int64, err:%s", src, err.Error()))
			}
		case float32:
			return int64(math.Round(float64(src)))
		case float64:
			return int64(math.Round(src))
		case nil:
			dst = 0
		case bool:
			if src {
				dst = 1
			}
		case complex64:
			panic(fmt.Sprintf("cannot convert complex %g to int64", src))
		case complex128:
			panic(fmt.Sprintf("cannot convert complex %g to int64", src))
		default:
			str := fmt.Sprintf("%v", src)
			dst, err = strconv.ParseInt(str, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("cannot convert %s to int64, err:%s", src, err.Error()))
			}
		}
	}
	return
}

func Boolean(src interface{}) bool {
	return Int64(src) != 0
}

func Float64(src interface{}, prec int) (dst float64) {
	var err error
	switch src := src.(type) {
	case float32:
		return float64(src)
	case float64:
		return src
	case nil:
		return 0
	case []byte:
		dst, err = strconv.ParseFloat(string(src), 64)
		if err != nil {
			panic(fmt.Sprintf("cannot convert %s to float64, err:%s", src, err.Error()))
		}
	case bool:
		if src {
			dst = 1
		}
	case complex64:
		panic(fmt.Sprintf("cannot convert complex %g to float64", src))
	case complex128:
		panic(fmt.Sprintf("cannot convert complex %g to float64", src))
	default:
		str := fmt.Sprintf("%v", src)
		dst, err = strconv.ParseFloat(str, 64)
		if err != nil {
			panic(fmt.Sprintf("cannot convert %s to float64, err:%s", src, err.Error()))
		}
	}
	str := strconv.FormatFloat(dst, 'f', prec, 64)
	dst, _ = strconv.ParseFloat(str, 64)
	return dst
}

func String(src interface{}) string {
	switch val := src.(type) {
	case int:
		return strconv.Itoa(val)
	case []byte:
		return string(val)
	case nil:
		return ""
	case string:
		return val
	default:
		rt := reflect.TypeOf(val)

		if rt.Kind() == reflect.Ptr {
			return String(reflect.ValueOf(val).Elem().Interface())
		}
		return fmt.Sprintf("%v", src)
	}
}

func StringSlice(src interface{}) (dst []string) {
	for _, val := range src.([]interface{}) {
		dst = append(dst, String(val))
	}
	return
}

func GBK2U8(src string) string {
	result, _ := simplifiedchinese.GBK.NewDecoder().String(src)
	return result
}

func U82GBK(src string) string {
	result, _ := simplifiedchinese.GBK.NewEncoder().String(src)
	return result
}
