// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser implements a parser for Go source files. Input may be
// provided in a variety of forms (see the various Parse* functions); the
// output is an abstract syntax tree (AST) representing the Go source. The
// parser is invoked through one of the Parse* functions.
//
// The parser accepts a larger language than is syntactically permitted by
// the Go spec, for simplicity, and for improved robustness in the presence
// of syntax errors. For instance, in method declarations, the receiver is
// treated like an ordinary parameter list and thus may contain multiple
// entries where the spec permits exactly one. Consequently, the corresponding
// field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.
//
package parser

import (
	"fmt"
	"go/ast"
	"go/token"

	mt "github.com/cosmos72/gomacro/token"
)

// patch: quote and friends
func (p *parser) parseQuote() ast.Expr {
	if p.trace {
		defer un(trace(p, "Quote"))
	}

	saveDepth := p.quasiquoteDepth
	saveQuote := p.quote

	switch p.tok {
	case mt.QUOTE:
		p.quote = true
	case mt.QUASIQUOTE:
		p.quasiquoteDepth++
	case mt.UNQUOTE, mt.UNQUOTE_SPLICE:
		p.quasiquoteDepth--
	}
	defer func() {
		p.quasiquoteDepth = saveDepth
		p.quote = saveQuote
	}()

	op := p.tok
	opPos := p.pos
	opName := mt.String(op) // use the actual name QUOTE/QUASIQUOTE/UNQUOTE/UNQUOTE_SPLICE even if we found ~' ~` ~, ~,@
	p.next()

	var node ast.Node

	// QUOTE, QUASIQUOTE, UNQUOTE and UNQUOTE_SLICE must be followed by one of:
	// * a basic literal
	// * an identifier
	// * a block statement - containing relaxed syntax, see parseQuotedBlockStmt()
	// * another QUOTE, QUASIQUOTE or UNQUOTE (not UNQUOTE_SPLICE, it must be wrapped in {})
	switch p.tok {
	case token.EOF, token.RPAREN, token.RBRACK, token.RBRACE,
		token.COMMA, token.PERIOD, token.SEMICOLON, token.COLON:

		// no applicable expression after QUOTE/QUASIQUOTE/...: just return the keyword itself
		return &ast.Ident{NamePos: opPos, Name: opName}

	case token.IDENT:
		node = &ast.Ident{NamePos: p.pos, Name: p.lit}
		p.next()

	case token.INT, token.FLOAT, token.IMAG, token.CHAR, token.STRING:
		node = &ast.BasicLit{ValuePos: p.pos, Kind: p.tok, Value: p.lit}
		p.next()

	case mt.QUOTE, mt.QUASIQUOTE, mt.UNQUOTE, mt.UNQUOTE_SPLICE:
		node = p.parseQuote()

	case token.LBRACE:
		// FIXME ideally, we do NOT want to parse macros when inside some depth of QUOTE and QUASIQUOTE
		// because we may not know about the macro yet, and because the macro could be inside fragments like
		//   quasiquote{some_macro unquote_splice{...}}
		// which means there is no way know how many arguments are present,
		// because they will be generated by splicing the results of user code.
		//
		// also, we should allow things like
		//   quasiquote{for unquote_splice{...}}
		// i.e. even the number of arguments to reserved keywords could vary depending on user code
		node = p.parseBlockStmt()

	default:
		p.errorExpected(p.pos, "one of: '{', 'IDENT', 'INT', 'STRING', 'QUOTE', 'QUASIQUOTE', 'UNQUOTE' or 'UNQUOTE_SPLICE'")
	}

	expr, _ := MakeQuote(p, op, opPos, node)
	return expr
}

// MakeQuote creates an ast.UnaryExpr representing quote{node}.
// Returns both the unaryexpr and the blockstmt containing its body
func MakeQuote(p_or_nil *parser, op token.Token, pos token.Pos, node ast.Node) (*ast.UnaryExpr, *ast.BlockStmt) {
	var body *ast.BlockStmt
	var stmt ast.Stmt
	switch node := node.(type) {
	case nil:
		break
	case *ast.BlockStmt:
		body = node
	case ast.Stmt:
		stmt = node
	case ast.Expr:
		stmt = &ast.ExprStmt{X: node}
	default:
		msg := fmt.Sprintf("%v: expecting statement or expression, found %T %#v", op, node)
		if p_or_nil != nil {
			p_or_nil.error(node.Pos(), msg)
		} else {
			panic(msg)
		}
	}
	if body == nil {
		list := make([]ast.Stmt, 0)
		if stmt != nil {
			list = append(list, stmt)
		}
		body = &ast.BlockStmt{Lbrace: stmt.Pos(), List: list, Rbrace: stmt.End()}
	}

	// due to go/ast strictly typed model, there is only one mechanism
	// to insert a statement inside an expression: use a closure.
	// so we return a unary expression: QUOTE (func() { /*block*/ })
	typ := &ast.FuncType{Func: token.NoPos, Params: &ast.FieldList{}}
	fun := &ast.FuncLit{Type: typ, Body: body}
	return &ast.UnaryExpr{OpPos: pos, Op: op, X: fun}, body
}

// macro calls syntax is "foo [;] bar [;] baz"... recognize it
func (p *parser) expectSemiOrSpace() {
	// semicolon is optional before a closing ')' or '}'
	// we make it optional also between two identifiers
	switch p.tok {
	case token.RPAREN, token.RBRACK, token.RBRACE:
		break
	case token.COMMA:
		// permit a ',' instead of a ';' but complain
		p.errorExpected(p.pos, "';'")
		fallthrough
	case token.SEMICOLON:
		p.next()
	default:
		if p.tok == token.IDENT && p.tok0 == token.IDENT {
			break
		}
		p.errorExpected(p.pos, "';'")
		syncStmt(p)
	}
}

// parseExprBlock parses a block statement inside an expression.
func (p *parser) parseExprBlock() ast.Expr {
	if p.trace {
		defer un(trace(p, "ExprBlock"))
	}

	pos := p.pos
	block := p.parseBlockStmt()

	// due to go/ast strictly typed model, there is only one mechanism
	// to insert a block statement (or any statement) inside an expression:
	// use a closure. so we return the unary expression:
	// MACRO func() { /*block*/ }
	typ := &ast.FuncType{Params: &ast.FieldList{}}
	fun := &ast.FuncLit{Type: typ, Body: block}
	return &ast.UnaryExpr{OpPos: pos, Op: mt.MACRO, X: fun}
}
