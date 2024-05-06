package entity

import (
	"time"

	"github.com/goccy/go-json"
)

func newItem(item *UnixItem) *Item {
	return &Item{
		ID:        item.ID,
		CreatedAt: time.Unix(item.CreatedAt, 0),
	}
}

type Item struct {
	ID        string    `json:"id" yaml:"id" toml:"id"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at" toml:"created_at"`
}

func (r *Item) MarshalBinary() ([]byte, error) {
	item := newUnixItem(r)
	return json.Marshal(item)
}

func (r *Item) UnmarshalBinary(data []byte) error {
	var unixItem UnixItem
	if err := json.Unmarshal(data, &unixItem); err != nil {
		return err
	}

	r = newItem(&unixItem)

	return nil
}

type UnixItem struct {
	ID        string `json:"id" yaml:"id" toml:"id"`
	CreatedAt int64  `json:"created_at" yaml:"created_at" toml:"created_at"`
}

func newUnixItem(item *Item) *UnixItem {
	return &UnixItem{
		ID:        item.ID,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
