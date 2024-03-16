package entity

import (
	"time"

	"github.com/goccy/go-json"
)

type Item struct {
	ID        string    `json:"id" yaml:"id" toml:"id"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at" toml:"created_at"`
}

func (r *Item) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Item) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, r)
}
