package  easyjson

import (
  jwriter "github.com/mailru/easyjson/jwriter"
  jlexer "github.com/mailru/easyjson/jlexer"
  json "encoding/json"
)

var _ = json.RawMessage{} // suppress unused package warning

func easyjson_decode_github_com_dimiro1_experiments_easyjson_Person(in *jlexer.Lexer, out *Person) {
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
    case "name":
      out.Name = in.String()
    case "age":
      out.Age = in.Int()
    default:
      in.SkipRecursive()
    }
    in.WantComma()
  }
  in.Delim('}')
}
func easyjson_encode_github_com_dimiro1_experiments_easyjson_Person(out *jwriter.Writer, in *Person) {
  out.RawByte('{')
  first := true
  _ = first
  if !first { out.RawByte(',') }
  first = false
  out.RawString("\"name\":")
  out.String(in.Name)
  if !first { out.RawByte(',') }
  first = false
  out.RawString("\"age\":")
  out.Int(in.Age)
  out.RawByte('}')
}
func (v *Person) MarshalJSON() ([]byte, error) {
  w := jwriter.Writer{}
  easyjson_encode_github_com_dimiro1_experiments_easyjson_Person(&w, v)
  return w.Buffer.BuildBytes(), w.Error
}
func (v *Person) MarshalEasyJSON(w *jwriter.Writer) {
  easyjson_encode_github_com_dimiro1_experiments_easyjson_Person(w, v)
}
func (v *Person) UnmarshalJSON(data []byte) error {
  r := jlexer.Lexer{Data: data}
  easyjson_decode_github_com_dimiro1_experiments_easyjson_Person(&r, v)
  return r.Error()
}
func (v *Person) UnmarshalEasyJSON(l *jlexer.Lexer) {
  easyjson_decode_github_com_dimiro1_experiments_easyjson_Person(l, v)
}
