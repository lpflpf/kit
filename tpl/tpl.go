package tpl

import (
	"bytes"
	"github.com/flosch/pongo2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 模板相关

var defaultTemplate map[string]*pongo2.Template
var viewPath string

func Init(path string) {
	viewPath = path
	defaultTemplate = make(map[string]*pongo2.Template)
	for name, funcs := range commonFilterFuncs {
		_ = pongo2.RegisterFilter(name, funcs)
	}

	viewPath = filepath.Clean(viewPath)
	_ = filepath.Walk(viewPath, func(path string, info os.FileInfo, _ error) (err error) {
		if filepath.Ext(path) == ".html" || filepath.Ext(path) == ".css" {
			var viewKey = path[len(viewPath):]
			defaultTemplate[viewKey] = pongo2.Must(pongo2.FromFile(path))
		}
		return nil
	})

}

// 模板渲染
func Render(w io.Writer, tmpl string, p map[string]interface{}) error {
	if tmpl[0] != '/' {
		tmpl = "/" + tmpl
	}

	if v, ok := defaultTemplate[tmpl]; ok {
		return v.ExecuteWriter(p, w)
	} else {
		return nil
	}
}

func Render2Bytes(tmpl string, p map[string]interface{}) (output []byte, err error) {
	buffer := new(bytes.Buffer)
	err = Render(buffer, tmpl, p)
	return buffer.Bytes(), err
}

func Render2String(tmpl string, p map[string]interface{}) (output string, err error) {
	buffer := new(bytes.Buffer)
	err = Render(buffer, tmpl, p)
	return buffer.String(), err
}

func GetRawData(tmpl string) ([]byte, error) {
	return ioutil.ReadFile(viewPath + "/" + tmpl)
}
