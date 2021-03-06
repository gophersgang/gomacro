// this file was generated by gomacro command: import "encoding/xml"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"encoding/xml"
)

func init() {
	Packages["encoding/xml"] = Package{
	Binds: map[string]Value{
		"CopyToken":	ValueOf(xml.CopyToken),
		"Escape":	ValueOf(xml.Escape),
		"EscapeText":	ValueOf(xml.EscapeText),
		"HTMLAutoClose":	ValueOf(&xml.HTMLAutoClose).Elem(),
		"HTMLEntity":	ValueOf(&xml.HTMLEntity).Elem(),
		"Header":	ValueOf(xml.Header),
		"Marshal":	ValueOf(xml.Marshal),
		"MarshalIndent":	ValueOf(xml.MarshalIndent),
		"NewDecoder":	ValueOf(xml.NewDecoder),
		"NewEncoder":	ValueOf(xml.NewEncoder),
		"Unmarshal":	ValueOf(xml.Unmarshal),
	},
	Types: map[string]Type{
		"Attr":	TypeOf((*xml.Attr)(nil)).Elem(),
		"CharData":	TypeOf((*xml.CharData)(nil)).Elem(),
		"Comment":	TypeOf((*xml.Comment)(nil)).Elem(),
		"Decoder":	TypeOf((*xml.Decoder)(nil)).Elem(),
		"Directive":	TypeOf((*xml.Directive)(nil)).Elem(),
		"Encoder":	TypeOf((*xml.Encoder)(nil)).Elem(),
		"EndElement":	TypeOf((*xml.EndElement)(nil)).Elem(),
		"Marshaler":	TypeOf((*xml.Marshaler)(nil)).Elem(),
		"MarshalerAttr":	TypeOf((*xml.MarshalerAttr)(nil)).Elem(),
		"Name":	TypeOf((*xml.Name)(nil)).Elem(),
		"ProcInst":	TypeOf((*xml.ProcInst)(nil)).Elem(),
		"StartElement":	TypeOf((*xml.StartElement)(nil)).Elem(),
		"SyntaxError":	TypeOf((*xml.SyntaxError)(nil)).Elem(),
		"TagPathError":	TypeOf((*xml.TagPathError)(nil)).Elem(),
		"Token":	TypeOf((*xml.Token)(nil)).Elem(),
		"UnmarshalError":	TypeOf((*xml.UnmarshalError)(nil)).Elem(),
		"Unmarshaler":	TypeOf((*xml.Unmarshaler)(nil)).Elem(),
		"UnmarshalerAttr":	TypeOf((*xml.UnmarshalerAttr)(nil)).Elem(),
		"UnsupportedTypeError":	TypeOf((*xml.UnsupportedTypeError)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"Marshaler":	TypeOf((*Marshaler_encoding_xml)(nil)).Elem(),
		"MarshalerAttr":	TypeOf((*MarshalerAttr_encoding_xml)(nil)).Elem(),
		"Token":	TypeOf((*Token_encoding_xml)(nil)).Elem(),
		"Unmarshaler":	TypeOf((*Unmarshaler_encoding_xml)(nil)).Elem(),
		"UnmarshalerAttr":	TypeOf((*UnmarshalerAttr_encoding_xml)(nil)).Elem(),
	} }
}

// --------------- proxy for encoding/xml.Marshaler ---------------
type Marshaler_encoding_xml struct {
	Object	interface{}
	MarshalXML_	func(e *xml.Encoder, start xml.StartElement) error
}
func (Proxy Marshaler_encoding_xml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return Proxy.MarshalXML_(e, start)
}

// --------------- proxy for encoding/xml.MarshalerAttr ---------------
type MarshalerAttr_encoding_xml struct {
	Object	interface{}
	MarshalXMLAttr_	func(name xml.Name) (xml.Attr, error)
}
func (Proxy MarshalerAttr_encoding_xml) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return Proxy.MarshalXMLAttr_(name)
}

// --------------- proxy for encoding/xml.Token ---------------
type Token_encoding_xml struct {
	Object	interface{}
}

// --------------- proxy for encoding/xml.Unmarshaler ---------------
type Unmarshaler_encoding_xml struct {
	Object	interface{}
	UnmarshalXML_	func(d *xml.Decoder, start xml.StartElement) error
}
func (Proxy Unmarshaler_encoding_xml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return Proxy.UnmarshalXML_(d, start)
}

// --------------- proxy for encoding/xml.UnmarshalerAttr ---------------
type UnmarshalerAttr_encoding_xml struct {
	Object	interface{}
	UnmarshalXMLAttr_	func(attr xml.Attr) error
}
func (Proxy UnmarshalerAttr_encoding_xml) UnmarshalXMLAttr(attr xml.Attr) error {
	return Proxy.UnmarshalXMLAttr_(attr)
}
