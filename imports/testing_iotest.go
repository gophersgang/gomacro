// this file was generated by gomacro command: import "testing/iotest"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"testing/iotest"
)

func init() {
	Packages["testing/iotest"] = Package{
	Binds: map[string]Value{
		"DataErrReader":	ValueOf(iotest.DataErrReader),
		"ErrTimeout":	ValueOf(&iotest.ErrTimeout).Elem(),
		"HalfReader":	ValueOf(iotest.HalfReader),
		"NewReadLogger":	ValueOf(iotest.NewReadLogger),
		"NewWriteLogger":	ValueOf(iotest.NewWriteLogger),
		"OneByteReader":	ValueOf(iotest.OneByteReader),
		"TimeoutReader":	ValueOf(iotest.TimeoutReader),
		"TruncateWriter":	ValueOf(iotest.TruncateWriter),
	},
	Types: map[string]Type{
	},
	Proxies: map[string]Type{
	} }
}
