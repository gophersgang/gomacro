// this file was generated by gomacro command: import "regexp"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	pkg "regexp"
)

func Package_regexp() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Compile":          ValueOf(pkg.Compile),
			"CompilePOSIX":     ValueOf(pkg.CompilePOSIX),
			"Match":            ValueOf(pkg.Match),
			"MatchReader":      ValueOf(pkg.MatchReader),
			"MatchString":      ValueOf(pkg.MatchString),
			"MustCompile":      ValueOf(pkg.MustCompile),
			"MustCompilePOSIX": ValueOf(pkg.MustCompilePOSIX),
			"QuoteMeta":        ValueOf(pkg.QuoteMeta),
		}, map[string]Type{
			"Regexp": TypeOf((*pkg.Regexp)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_regexp()
	Binds["regexp"] = binds
	Types["regexp"] = types
}