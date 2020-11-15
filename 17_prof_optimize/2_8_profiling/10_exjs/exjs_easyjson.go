// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package exjs

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF1fd6b82DecodeProfilingLesson10Exjs(in *jlexer.Lexer, out *A) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "I":
			out.I = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF1fd6b82EncodeProfilingLesson10Exjs(out *jwriter.Writer, in A) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"I\":"
		out.RawString(prefix[1:])
		out.Int(int(in.I))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v A) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF1fd6b82EncodeProfilingLesson10Exjs(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v A) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF1fd6b82EncodeProfilingLesson10Exjs(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *A) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF1fd6b82DecodeProfilingLesson10Exjs(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *A) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF1fd6b82DecodeProfilingLesson10Exjs(l, v)
}