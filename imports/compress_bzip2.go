// this file was generated by gomacro command: import "compress/bzip2"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"compress/bzip2"
)

func init() {
	Packages["compress/bzip2"] = Package{
	Binds: map[string]Value{
		"NewReader":	ValueOf(bzip2.NewReader),
	},
	Types: map[string]Type{
		"StructuralError":	TypeOf((*bzip2.StructuralError)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	} }
}
