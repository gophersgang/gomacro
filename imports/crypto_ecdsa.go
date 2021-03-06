// this file was generated by gomacro command: import "crypto/ecdsa"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/ecdsa"
)

func init() {
	Packages["crypto/ecdsa"] = Package{
	Binds: map[string]Value{
		"GenerateKey":	ValueOf(ecdsa.GenerateKey),
		"Sign":	ValueOf(ecdsa.Sign),
		"Verify":	ValueOf(ecdsa.Verify),
	},
	Types: map[string]Type{
		"PrivateKey":	TypeOf((*ecdsa.PrivateKey)(nil)).Elem(),
		"PublicKey":	TypeOf((*ecdsa.PublicKey)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	} }
}
