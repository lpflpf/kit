package tpl

import (
	"encoding/json"
	"errors"
	"github.com/flosch/pongo2"
	"github.com/lpflpf/kit/util"
	"math"
	"strconv"
	"strings"
)

// 向下取整
func floor(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	val := util.Float64(in.Interface(), 2)
	return pongo2.AsValue(math.Floor(val)), nil
}

func index(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	val := in.Interface()
	switch v := val.(type) {
	case map[string]interface{}:
		return pongo2.AsValue(v[param.String()]), nil
	default:
		if in.CanSlice() {
			return in.Index(param.Integer()), nil
		}
		return pongo2.AsValue(""), nil
	}
}

// 字符串替换
// 参数用逗号分割
// replace str,old,times
func replace(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	val := in.String()
	paramString := param.String()

	data := strings.Split(paramString, ",")

	switch len(data) {
	case 2:
		return pongo2.AsValue(strings.ReplaceAll(val, data[0], data[1])), nil
	case 3:
		return pongo2.AsValue(strings.Replace(val, data[0], data[1], util.Int(data[2]))), nil
	default:
		return nil, &pongo2.Error{
			OrigError: errors.New("params failed"),
		}
	}
}

// 追加
func cat(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	return pongo2.AsValue(in.String() + param.String()), nil
}

// 添加前缀
func prefix(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	val := in.String()
	prefix := param.String()

	return pongo2.AsValue(prefix + val), nil
}

// 字串 支持负数
func substr(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	val := in.String()
	token := strings.Split(param.String(), ":")

	if len(token) != 2 {
		return pongo2.AsValue(""), nil
	}

	begin, end := util.Int(token[0]), util.Int(token[1])

	if end < 0 {
		end = len(val) + end
	}

	if begin < 0 {
		begin = len(val) + begin
	}

	if begin < 0 || end < 0 || begin > end || end > len(val) || begin > len(val) {
		return pongo2.AsValue(""), nil
	}

	return pongo2.AsValue(val[begin:end]), nil
}

// json 格式化
func jsonFormat(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	data, e := json.Marshal(in.Interface())

	if e != nil {
		return pongo2.AsValue("{}"), nil
	}
	return pongo2.AsValue(string(data)), nil
}

// md5
func md5Sum(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	return pongo2.AsValue(util.Md5([]byte(in.String()))), nil
}

// 四舍五入
func round(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	p := 2
	if param != nil {
		p = in.Integer()
	}
	return pongo2.AsValue(strconv.FormatFloat(in.Float(), 'f', p, 64)), nil
}

var commonFilterFuncs = map[string]pongo2.FilterFunction{
	"floor":   floor,
	"index":   index,
	"replace": replace,
	"cat":     cat,
	"prefix":  prefix,
	"substr":  substr,
	"md5":     md5Sum,
	"json":    jsonFormat,
	"round":   round,
}
