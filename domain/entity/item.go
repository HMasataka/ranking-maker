package entity

import (
	"github.com/goccy/go-json"
)

type Item struct {
	ID string `json:"id" yaml:"id" toml:"id"`
}

func (r *Item) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Item) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, r)
}
