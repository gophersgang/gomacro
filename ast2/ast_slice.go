/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * ast_slice.go
 *
 *  Created on Feb 25, 2017
 *      Author Massimiliano Ghilardi
 */

package ast2

import (
	"go/ast"
	"go/token"
)

// Ast wrappers for variable-length slices of ast.Nodes - they are not full-blown ast.Node

func (x NodeSlice) Interface() interface{}  { return x.X }
func (x ExprSlice) Interface() interface{}  { return x.X }
func (x FieldSlice) Interface() interface{} { return x.X }
func (x DeclSlice) Interface() interface{}  { return x.X }
func (x IdentSlice) Interface() interface{} { return x.X }
func (x SpecSlice) Interface() interface{}  { return x.X }
func (x StmtSlice) Interface() interface{}  { return x.X }

func (x NodeSlice) Op() token.Token  { return token.COMMA }     // FIXME
func (x ExprSlice) Op() token.Token  { return token.COMMA }     // FIXME
func (x FieldSlice) Op() token.Token { return token.SEMICOLON } // FIXME
func (x DeclSlice) Op() token.Token  { return token.SEMICOLON } // FIXME
func (x IdentSlice) Op() token.Token { return token.COMMA }     // FIXME
func (x SpecSlice) Op() token.Token  { return token.SEMICOLON } // FIXME
func (x StmtSlice) Op() token.Token  { return token.SEMICOLON } // FIXME

func (x NodeSlice) New() Ast  { return NodeSlice{X: []ast.Node{}} }
func (x ExprSlice) New() Ast  { return ExprSlice{X: []ast.Expr{}} }
func (x FieldSlice) New() Ast { return FieldSlice{X: []*ast.Field{}} }
func (x DeclSlice) New() Ast  { return DeclSlice{X: []ast.Decl{}} }
func (x IdentSlice) New() Ast { return IdentSlice{X: []*ast.Ident{}} }
func (x SpecSlice) New() Ast  { return SpecSlice{X: []ast.Spec{}} }
func (x StmtSlice) New() Ast  { return StmtSlice{X: []ast.Stmt{}} }

func (x NodeSlice) Size() int  { return len(x.X) }
func (x ExprSlice) Size() int  { return len(x.X) }
func (x FieldSlice) Size() int { return len(x.X) }
func (x DeclSlice) Size() int  { return len(x.X) }
func (x IdentSlice) Size() int { return len(x.X) }
func (x SpecSlice) Size() int  { return len(x.X) }
func (x StmtSlice) Size() int  { return len(x.X) }

func (x NodeSlice) Get(i int) Ast  { return ToAst(x.X[i]) }
func (x ExprSlice) Get(i int) Ast  { return ToAst(x.X[i]) }
func (x FieldSlice) Get(i int) Ast { return ToAst(x.X[i]) }
func (x DeclSlice) Get(i int) Ast  { return ToAst(x.X[i]) }
func (x IdentSlice) Get(i int) Ast { return ToAst(x.X[i]) }
func (x SpecSlice) Get(i int) Ast  { return ToAst(x.X[i]) }
func (x StmtSlice) Get(i int) Ast  { return ToAst(x.X[i]) }

func (x NodeSlice) Set(i int, child Ast)  { x.X[i] = ToNode(child) }
func (x ExprSlice) Set(i int, child Ast)  { x.X[i] = ToExpr(child) }
func (x FieldSlice) Set(i int, child Ast) { x.X[i] = ToField(child) }
func (x DeclSlice) Set(i int, child Ast)  { x.X[i] = ToDecl(child) }
func (x IdentSlice) Set(i int, child Ast) { x.X[i] = ToIdent(child) }
func (x SpecSlice) Set(i int, child Ast)  { x.X[i] = ToSpec(child) }
func (x StmtSlice) Set(i int, child Ast)  { x.X[i] = ToStmt(child) }

func (x NodeSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }
func (x ExprSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }
func (x FieldSlice) Slice(lo, hi int) AstWithSlice { x.X = x.X[lo:hi]; return x }
func (x DeclSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }
func (x IdentSlice) Slice(lo, hi int) AstWithSlice { x.X = x.X[lo:hi]; return x }
func (x SpecSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }
func (x StmtSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }

func (x NodeSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToNode(child)); return x }
func (x ExprSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToExpr(child)); return x }
func (x FieldSlice) Append(child Ast) AstWithSlice { x.X = append(x.X, ToField(child)); return x }
func (x DeclSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToDecl(child)); return x }
func (x IdentSlice) Append(child Ast) AstWithSlice { x.X = append(x.X, ToIdent(child)); return x }
func (x SpecSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToSpec(child)); return x }
func (x StmtSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToStmt(child)); return x }

// variable-length ast.Nodes

func (x BlockStmt) Interface() interface{}  { return x.X }
func (x FieldList) Interface() interface{}  { return x.X }
func (x File) Interface() interface{}       { return x.X }
func (x GenDecl) Interface() interface{}    { return x.X }
func (x ReturnStmt) Interface() interface{} { return x.X }

func (x BlockStmt) Node() ast.Node  { return x.X }
func (x FieldList) Node() ast.Node  { return x.X }
func (x File) Node() ast.Node       { return x.X }
func (x GenDecl) Node() ast.Node    { return x.X }
func (x ReturnStmt) Node() ast.Node { return x.X }

func (x BlockStmt) Op() token.Token  { return token.LBRACE }
func (x FieldList) Op() token.Token  { return token.ELLIPSIS }
func (x File) Op() token.Token       { return token.EOF }
func (x GenDecl) Op() token.Token    { return x.X.Tok }
func (x ReturnStmt) Op() token.Token { return token.RETURN }

func (x BlockStmt) New() Ast { return BlockStmt{&ast.BlockStmt{Lbrace: x.X.Lbrace, Rbrace: x.X.Rbrace}} }
func (x FieldList) New() Ast { return FieldList{&ast.FieldList{}} }
func (x File) New() Ast {
	return File{&ast.File{Doc: x.X.Doc, Package: x.X.Package, Name: x.X.Name, Scope: x.X.Scope, Imports: x.X.Imports, Comments: x.X.Comments}}
}
func (x GenDecl) New() Ast {
	return GenDecl{&ast.GenDecl{Doc: x.X.Doc, TokPos: x.X.TokPos, Tok: x.X.Tok, Lparen: x.X.Lparen, Rparen: x.X.Rparen}}
}
func (x ReturnStmt) New() Ast { return ReturnStmt{&ast.ReturnStmt{Return: x.X.Return}} }

func (x BlockStmt) Size() int  { return len(x.X.List) }
func (x FieldList) Size() int  { return len(x.X.List) }
func (x File) Size() int       { return len(x.X.Decls) }
func (x GenDecl) Size() int    { return len(x.X.Specs) }
func (x ReturnStmt) Size() int { return len(x.X.Results) }

func (x BlockStmt) Get(i int) Ast  { return ToAst(x.X.List[i]) }
func (x FieldList) Get(i int) Ast  { return ToAst(x.X.List[i]) }
func (x File) Get(i int) Ast       { return ToAst(x.X.Decls[i]) }
func (x GenDecl) Get(i int) Ast    { return ToAst(x.X.Specs[i]) }
func (x ReturnStmt) Get(i int) Ast { return ToAst(x.X.Results[i]) }

func (x BlockStmt) Set(i int, child Ast)  { x.X.List[i] = ToStmt(child) }
func (x FieldList) Set(i int, child Ast)  { x.X.List[i] = ToField(child) }
func (x File) Set(i int, child Ast)       { x.X.Decls[i] = ToDecl(child) }
func (x GenDecl) Set(i int, child Ast)    { x.X.Specs[i] = ToSpec(child) }
func (x ReturnStmt) Set(i int, child Ast) { x.X.Results[i] = ToExpr(child) }

func (x BlockStmt) Slice(lo, hi int) AstWithSlice  { x.X.List = x.X.List[lo:hi]; return x }
func (x FieldList) Slice(lo, hi int) AstWithSlice  { x.X.List = x.X.List[lo:hi]; return x }
func (x File) Slice(lo, hi int) AstWithSlice       { x.X.Decls = x.X.Decls[lo:hi]; return x }
func (x GenDecl) Slice(lo, hi int) AstWithSlice    { x.X.Specs = x.X.Specs[lo:hi]; return x }
func (x ReturnStmt) Slice(lo, hi int) AstWithSlice { x.X.Results = x.X.Results[lo:hi]; return x }

func (x BlockStmt) Append(child Ast) AstWithSlice {
	x.X.List = append(x.X.List, ToStmt(child))
	return x
}
func (x FieldList) Append(child Ast) AstWithSlice {
	x.X.List = append(x.X.List, ToField(child))
	return x
}
func (x File) Append(child Ast) AstWithSlice {
	x.X.Decls = append(x.X.Decls, ToDecl(child))
	return x
}
func (x GenDecl) Append(child Ast) AstWithSlice {
	x.X.Specs = append(x.X.Specs, ToSpec(child))
	return x
}
func (x ReturnStmt) Append(child Ast) AstWithSlice {
	x.X.Results = append(x.X.Results, ToExpr(child))
	return x
}
