// this file was generated by gomacro command: import "unicode/utf8"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"unicode/utf8"
)

func init() {
	Packages["unicode/utf8"] = Package{
	Binds: map[string]Value{
		"DecodeLastRune":	ValueOf(utf8.DecodeLastRune),
		"DecodeLastRuneInString":	ValueOf(utf8.DecodeLastRuneInString),
		"DecodeRune":	ValueOf(utf8.DecodeRune),
		"DecodeRuneInString":	ValueOf(utf8.DecodeRuneInString),
		"EncodeRune":	ValueOf(utf8.EncodeRune),
		"FullRune":	ValueOf(utf8.FullRune),
		"FullRuneInString":	ValueOf(utf8.FullRuneInString),
		"MaxRune":	ValueOf(utf8.MaxRune),
		"RuneCount":	ValueOf(utf8.RuneCount),
		"RuneCountInString":	ValueOf(utf8.RuneCountInString),
		"RuneError":	ValueOf(utf8.RuneError),
		"RuneLen":	ValueOf(utf8.RuneLen),
		"RuneSelf":	ValueOf(utf8.RuneSelf),
		"RuneStart":	ValueOf(utf8.RuneStart),
		"UTFMax":	ValueOf(utf8.UTFMax),
		"Valid":	ValueOf(utf8.Valid),
		"ValidRune":	ValueOf(utf8.ValidRune),
		"ValidString":	ValueOf(utf8.ValidString),
	},
	Types: map[string]Type{
	},
	Proxies: map[string]Type{
	} }
}
