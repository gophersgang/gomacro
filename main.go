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
 * main.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"fmt"
	"os"

	gmi "github.com/cosmos72/gomacro/interpreter"
	gmp "github.com/cosmos72/gomacro/parser"
)

func main() {
	args := os.Args[1:]

	// args := []string{"macro add3(a, b, c interface{}) interface{} { ~`{~,a + ~,b + ~,c } }\nMacroExpand1(quote{add3;1;2;3})\nadd3;4;5;6"}

	// args := []string{"x:=~'{var z interface{} = 42}"}

	// generated with: find [a-u]* -type f -name \*.go | grep -v internal | grep -v testdata | grep -v cmd/ | grep -v builtin | xargs -d'\n' dirname | sort -u | while read i; do echo -n "\"$i\"; "; done
	// plus some hand-made tweaks
	// args = `import ( "archive/tar"; "archive/zip"; "bufio"; "bytes"; "compress/bzip2"; "compress/flate"; "compress/gzip"; "compress/lzw"; "compress/zlib"; "container/heap"; "container/list"; "container/ring"; "context"; "crypto"; "crypto/aes"; "crypto/cipher"; "crypto/des"; "crypto/dsa"; "crypto/ecdsa"; "crypto/elliptic"; "crypto/hmac"; "crypto/md5"; "crypto/rand"; "crypto/rc4"; "crypto/rsa"; "crypto/sha1"; "crypto/sha256"; "crypto/sha512"; "crypto/subtle"; "crypto/tls"; "crypto/x509"; "crypto/x509/pkix"; "database/sql"; "database/sql/driver"; "debug/dwarf"; "debug/elf"; "debug/gosym"; "debug/macho"; "debug/pe"; "debug/plan9obj"; "encoding"; "encoding/ascii85"; "encoding/asn1"; "encoding/base32"; "encoding/base64"; "encoding/binary"; "encoding/csv"; "encoding/gob"; "encoding/hex"; "encoding/json"; "encoding/pem"; "encoding/xml"; "errors"; "expvar"; "flag"; "fmt"; "go/ast"; "go/build"; "go/constant"; "go/doc"; "go/format"; "go/importer"; "go/parser"; "go/printer"; "go/scanner"; "go/token"; "go/types"; "hash"; "hash/adler32"; "hash/crc32"; "hash/crc64"; "hash/fnv"; "html"; "html/template"; "image"; "image/color"; "image/color/palette"; "image/draw"; "image/gif"; "image/jpeg"; "image/png"; "index/suffixarray"; "io"; "io/ioutil"; "log"; "log/syslog"; "math"; "math/big"; "math/cmplx"; "math/rand"; "mime"; "mime/multipart"; "mime/quotedprintable"; "net"; "net/http"; "net/http/cgi"; "net/http/cookiejar"; "net/http/fcgi"; "net/http/httptest"; "net/http/httptrace"; "net/http/httputil"; "net/http/pprof"; "net/mail"; "net/rpc"; "net/rpc/jsonrpc"; "net/smtp"; "net/textproto"; "net/url"; "os"; "os/exec"; "os/signal"; "os/user"; "path"; "path/filepath"; "plugin"; "reflect"; "regexp"; "regexp/syntax"; "runtime"; "runtime/debug"; "runtime/pprof"; "runtime/trace"; "sort"; "strconv"; "strings"; "sync"; "sync/atomic"; "syscall"; "testing"; "testing/iotest"; "testing/quick"; "text/scanner"; "text/tabwriter"; "text/template"; "text/template/parse"; "time"; "unicode"; "unicode/utf16"; "unicode/utf8"; "unsafe" )`
	// args = `import "github.com/pquerna/ffjson/ffjson"`

	var cmd gmi.Cmd
	cmd.Init()

	cmd.ParserMode |= gmp.Trace & 0
	cmd.Options |= gmi.OptTrapPanic // | gmi.OptShowAfterMacroExpansion // | gmi.OptShowAfterParse // | gmi.OptDebugMacroExpand // |  gmi.OptDebugQuasiquote  // | gmi.OptShowEvalDuration

	err := cmd.Main(args)
	if err != nil {
		fmt.Fprintln(cmd.Stderr, err)
	}
}
