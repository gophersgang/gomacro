/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * literal.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"
	"strconv"
	"strings"
)

func (env *Env) evalLiteral0(node *ast.BasicLit) interface{} {
	kind := node.Kind
	str := node.Value
	var ret interface{}

	switch kind {

	case token.INT:
		if strings.HasPrefix(str, "-") {
			i64, err := strconv.ParseInt(str, 0, 64)
			if err != nil {
				return error_(err)
			}
			// prefer int to int64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(i64)
			if int64(i) == i64 {
				return i
			}
			return i64
		} else {
			u64, err := strconv.ParseUint(str, 0, 64)
			if err != nil {
				return error_(err)
			}
			// prefer, in order: int, int64, uint, uint64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(u64)
			if i >= 0 && uint64(i) == u64 {
				return i
			}
			i64 := int64(u64)
			if i64 >= 0 && uint64(i64) == u64 {
				return i64
			}
			u := uint(u64)
			if uint64(u) == u64 {
				return u
			}
			return u64
		}

	case token.FLOAT:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return error_(err)
		}
		ret = f

	case token.IMAG:
		if strings.HasSuffix(str, "i") {
			str = str[:len(str)-1]
		}
		im, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return error_(err)
		}
		ret = complex(0.0, im)
		// env.Debugf("evalLiteral(): parsed IMAG %s -> %T %#v -> %T %#v", str, im, im, ret, ret)

	case token.CHAR:
		return unescapeChar(str)

	case token.STRING:
		return unescapeString(str)

	default:
		env.errorf("unimplemented basic literal: %v", node)
		ret = nil
	}
	return ret
}

func (env *Env) evalCompositeLiteral(node *ast.CompositeLit) (r.Value, []r.Value) {
	t := env.evalType(node.Type)
	obj := Nil
	switch t.Kind() {
	case r.Map:
		obj = r.MakeMap(t)
		kt := t.Key()
		vt := t.Elem()
		for _, elt := range node.Elts {
			switch elt := elt.(type) {
			case *ast.KeyValueExpr:
				key := env.valueToType(env.evalExpr1(elt.Key), kt)
				val := env.valueToType(env.evalExpr1(elt.Value), vt)
				obj.SetMapIndex(key, val)
			default:
				env.errorf("map literal: invalid element, expecting <*ast.KeyValueExpr>, found: %v <%v>", elt, r.TypeOf(elt))
			}
		}
	case r.Array, r.Slice:
		vt := t.Elem()
		idx := -1
		val := Nil
		zero := Nil
		if t.Kind() == r.Array {
			obj = r.New(t).Elem()
		} else {
			zero = r.Zero(vt)
			obj = r.MakeSlice(t, 0, len(node.Elts))
		}
		for _, elt := range node.Elts {
			switch elt := elt.(type) {
			case *ast.KeyValueExpr:
				idx = int(env.valueToType(env.evalExpr1(elt.Key), typeOfInt).Int())
				val = env.valueToType(env.evalExpr1(elt.Value), vt)
			default:
				// golang specs:
				// "An element without a key uses the previous element's index plus one.
				// If the first element has no key, its index is zero."
				idx++
				val = env.valueToType(env.evalExpr1(elt), vt)
			}
			if zero != Nil { // is slice
				for obj.Len() <= idx {
					obj = r.Append(obj, zero)
				}
			}
			obj.Index(idx).Set(val)
		}
	case r.Struct:
		obj = r.New(t).Elem()
		var pairs, elts bool
		var field r.Value
		var expr ast.Expr
		for idx, elt := range node.Elts {
			switch elt := elt.(type) {
			case *ast.KeyValueExpr:
				if elts {
					return env.errorf("cannot mix keyed and non-keyed initializers in struct composite literal: %v", node)
				}
				pairs = true
				name := elt.Key.(*ast.Ident).Name
				field = obj.FieldByName(name)
				expr = elt.Value
			default:
				if pairs {
					return env.errorf("cannot mix keyed and non-keyed initializers in struct composite literal: %v", node)
				}
				elts = true
				field = obj.Field(idx)
				expr = elt
			}
			val := env.valueToType(env.evalExpr1(expr), field.Type())
			field.Set(val)
		}
	default:
		env.errorf("unexpected composite literal: %v", node)
	}
	return obj, nil
}

// lambda()
func (env *Env) evalFunctionLiteral(node *ast.FuncLit) (r.Value, []r.Value) {
	// env.Debugf("func() at position %v", node.Type.Func)

	ret, _ := env.evalDeclFunction(nil, node.Type, node.Body)
	return ret, nil
}
