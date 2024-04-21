// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http

import (
	json "encoding/json"
	models "github.com/SlavaShagalov/my-trello-backend/internal/models"
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

func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp(in *jlexer.Lexer, out *partialUpdateRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				if out.Title == nil {
					out.Title = new(string)
				}
				*out.Title = string(in.String())
			}
		case "position":
			if in.IsNull() {
				in.Skip()
				out.Position = nil
			} else {
				if out.Position == nil {
					out.Position = new(int)
				}
				*out.Position = int(in.Int())
			}
		case "board_id":
			if in.IsNull() {
				in.Skip()
				out.BoardID = nil
			} else {
				if out.BoardID == nil {
					out.BoardID = new(int)
				}
				*out.BoardID = int(in.Int())
			}
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp(out *jwriter.Writer, in partialUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		if in.Title == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Title))
		}
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		if in.Position == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Position))
		}
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		if in.BoardID == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.BoardID))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v partialUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v partialUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *partialUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *partialUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp(l, v)
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp1(in *jlexer.Lexer, out *listSimpleResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "lists":
			if in.IsNull() {
				in.Skip()
				out.Lists = nil
			} else {
				in.Delim('[')
				if out.Lists == nil {
					if !in.IsDelim(']') {
						out.Lists = make([]models.List, 0, 0)
					} else {
						out.Lists = []models.List{}
					}
				} else {
					out.Lists = (out.Lists)[:0]
				}
				for !in.IsDelim(']') {
					var v1 models.List
					easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalModels(in, &v1)
					out.Lists = append(out.Lists, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp1(out *jwriter.Writer, in listSimpleResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"lists\":"
		out.RawString(prefix[1:])
		if in.Lists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Lists {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalModels(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v listSimpleResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v listSimpleResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *listSimpleResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *listSimpleResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp1(l, v)
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalModels(in *jlexer.Lexer, out *models.List) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "board_id":
			out.BoardID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "position":
			out.Position = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalModels(out *jwriter.Writer, in models.List) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		out.Int(int(in.BoardID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		out.Int(int(in.Position))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp2(in *jlexer.Lexer, out *listResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "lists":
			if in.IsNull() {
				in.Skip()
				out.Lists = nil
			} else {
				in.Delim('[')
				if out.Lists == nil {
					if !in.IsDelim(']') {
						out.Lists = make([]itemResponse, 0, 0)
					} else {
						out.Lists = []itemResponse{}
					}
				} else {
					out.Lists = (out.Lists)[:0]
				}
				for !in.IsDelim(']') {
					var v4 itemResponse
					(v4).UnmarshalEasyJSON(in)
					out.Lists = append(out.Lists, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp2(out *jwriter.Writer, in listResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"lists\":"
		out.RawString(prefix[1:])
		if in.Lists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Lists {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v listResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v listResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *listResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *listResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp2(l, v)
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp3(in *jlexer.Lexer, out *itemResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "board_id":
			out.BoardID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "position":
			out.Position = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "cards":
			if in.IsNull() {
				in.Skip()
				out.Cards = nil
			} else {
				in.Delim('[')
				if out.Cards == nil {
					if !in.IsDelim(']') {
						out.Cards = make([]models.Card, 0, 0)
					} else {
						out.Cards = []models.Card{}
					}
				} else {
					out.Cards = (out.Cards)[:0]
				}
				for !in.IsDelim(']') {
					var v7 models.Card
					easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalModels1(in, &v7)
					out.Cards = append(out.Cards, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp3(out *jwriter.Writer, in itemResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		out.Int(int(in.BoardID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		out.Int(int(in.Position))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"cards\":"
		out.RawString(prefix)
		if in.Cards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Cards {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalModels1(out, v9)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v itemResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v itemResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *itemResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *itemResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp3(l, v)
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalModels1(in *jlexer.Lexer, out *models.Card) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "list_id":
			out.ListID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "content":
			out.Content = string(in.String())
		case "position":
			out.Position = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalModels1(out *jwriter.Writer, in models.Card) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"list_id\":"
		out.RawString(prefix)
		out.Int(int(in.ListID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		out.Int(int(in.Position))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp4(in *jlexer.Lexer, out *getResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "board_id":
			out.BoardID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "position":
			out.Position = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp4(out *jwriter.Writer, in getResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		out.Int(int(in.BoardID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		out.Int(int(in.Position))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp4(l, v)
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp5(in *jlexer.Lexer, out *createResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "board_id":
			out.BoardID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "position":
			out.Position = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp5(out *jwriter.Writer, in createResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		out.Int(int(in.BoardID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		out.Int(int(in.Position))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v createResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v createResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *createResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *createResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp5(l, v)
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp6(in *jlexer.Lexer, out *createRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "title":
			out.Title = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp6(out *jwriter.Writer, in createRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v createRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v createRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *createRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *createRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovMyTrelloBackendInternalListsDeliveryHttp6(l, v)
}
