package main

import "fmt"

// import "github.com/lpflpf/kit/tpl"
import "github.com/lpflpf/kit/util"

func main() {
	//	tpl.Init("./view")
	//	output, err := tpl.Render2String("index.html", map[string]interface{}{
	//		"head":       "hello world",
	//		"floatdata":  12.222222,
	//		"mapdata":    map[string]interface{}{"abc": "def"},
	//		"slicedata":  []int{1, 2, 3},
	//		"stringdata": "begin | abc abc abc abc | end",
	//		"jsondata":   map[string]interface{}{"abc": []int{1, 2, 3}},
	//	})
	//
	//	fmt.Println(output, err)
	//
	//	util.Dump(true)

	fmt.Println(util.Int64(4))
	fmt.Println(util.Int64(6.9))
}
