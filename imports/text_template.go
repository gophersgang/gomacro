// this file was generated by gomacro command: import "text/template"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"text/template"
)

func init() {
	Packages["text/template"] = Package{
	Binds: map[string]Value{
		"HTMLEscape":	ValueOf(template.HTMLEscape),
		"HTMLEscapeString":	ValueOf(template.HTMLEscapeString),
		"HTMLEscaper":	ValueOf(template.HTMLEscaper),
		"IsTrue":	ValueOf(template.IsTrue),
		"JSEscape":	ValueOf(template.JSEscape),
		"JSEscapeString":	ValueOf(template.JSEscapeString),
		"JSEscaper":	ValueOf(template.JSEscaper),
		"Must":	ValueOf(template.Must),
		"New":	ValueOf(template.New),
		"ParseFiles":	ValueOf(template.ParseFiles),
		"ParseGlob":	ValueOf(template.ParseGlob),
		"URLQueryEscaper":	ValueOf(template.URLQueryEscaper),
	},
	Types: map[string]Type{
		"ExecError":	TypeOf((*template.ExecError)(nil)).Elem(),
		"FuncMap":	TypeOf((*template.FuncMap)(nil)).Elem(),
		"Template":	TypeOf((*template.Template)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	} }
}
