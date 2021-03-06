// this file was generated by gomacro command: import "encoding/binary"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"encoding/binary"
)

func init() {
	Packages["encoding/binary"] = Package{
	Binds: map[string]Value{
		"BigEndian":	ValueOf(&binary.BigEndian).Elem(),
		"LittleEndian":	ValueOf(&binary.LittleEndian).Elem(),
		"MaxVarintLen16":	ValueOf(binary.MaxVarintLen16),
		"MaxVarintLen32":	ValueOf(binary.MaxVarintLen32),
		"MaxVarintLen64":	ValueOf(binary.MaxVarintLen64),
		"PutUvarint":	ValueOf(binary.PutUvarint),
		"PutVarint":	ValueOf(binary.PutVarint),
		"Read":	ValueOf(binary.Read),
		"ReadUvarint":	ValueOf(binary.ReadUvarint),
		"ReadVarint":	ValueOf(binary.ReadVarint),
		"Size":	ValueOf(binary.Size),
		"Uvarint":	ValueOf(binary.Uvarint),
		"Varint":	ValueOf(binary.Varint),
		"Write":	ValueOf(binary.Write),
	},
	Types: map[string]Type{
		"ByteOrder":	TypeOf((*binary.ByteOrder)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"ByteOrder":	TypeOf((*ByteOrder_encoding_binary)(nil)).Elem(),
	} }
}

// --------------- proxy for encoding/binary.ByteOrder ---------------
type ByteOrder_encoding_binary struct {
	Object	interface{}
	PutUint16_	func([]byte, uint16) 
	PutUint32_	func([]byte, uint32) 
	PutUint64_	func([]byte, uint64) 
	String_	func() string
	Uint16_	func([]byte) uint16
	Uint32_	func([]byte) uint32
	Uint64_	func([]byte) uint64
}
func (Proxy ByteOrder_encoding_binary) PutUint16(unnamed0 []byte, unnamed1 uint16)  {
	Proxy.PutUint16_(unnamed0, unnamed1)
}
func (Proxy ByteOrder_encoding_binary) PutUint32(unnamed0 []byte, unnamed1 uint32)  {
	Proxy.PutUint32_(unnamed0, unnamed1)
}
func (Proxy ByteOrder_encoding_binary) PutUint64(unnamed0 []byte, unnamed1 uint64)  {
	Proxy.PutUint64_(unnamed0, unnamed1)
}
func (Proxy ByteOrder_encoding_binary) String() string {
	return Proxy.String_()
}
func (Proxy ByteOrder_encoding_binary) Uint16(unnamed0 []byte) uint16 {
	return Proxy.Uint16_(unnamed0)
}
func (Proxy ByteOrder_encoding_binary) Uint32(unnamed0 []byte) uint32 {
	return Proxy.Uint32_(unnamed0)
}
func (Proxy ByteOrder_encoding_binary) Uint64(unnamed0 []byte) uint64 {
	return Proxy.Uint64_(unnamed0)
}
