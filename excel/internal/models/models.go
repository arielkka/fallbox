package models

import "encoding/json"

type (
	Message struct {
		Text    string
		ID      string
		Body    []byte
		ReplyTo string
	}
	Response struct {
		Body []byte `json:"body,omitempty"`
		ID   int    `json:"id,omitempty"`
		Err  string `json:"err,omitempty"`
	}
	Request struct {
		ID     int    `json:"id,omitempty"`
		UserID string `json:"user_id,omitempty"`
		Body   []byte `json:"body,omitempty"`
	}
	IsDeleted struct {
		Flag bool `json:"is_deleted"`
	}
)

func NewMessage(text, ID, sender string, body []byte) *Message {
	return &Message{Text: text, ID: ID, Body: body, ReplyTo: sender}
}

func (r *Response) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Request) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (i *IsDeleted) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
