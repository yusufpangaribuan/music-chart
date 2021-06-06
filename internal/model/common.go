package models

import (
	"net/http"

	"github.com/mholt/binding"
)

type Response struct {
	Data        interface{} `json:"data,omitempty"`
	ProcessTime string      `json:"process_time"`
	Status      int         `json:"status"`
	Error       interface{} `json:"error,omitempty"`
}

type BasicSelectParams struct {
	Limit  uint64
	Page   uint64
	Offset uint64
}

func (b *BasicSelectParams) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&b.Limit: "limit",
		&b.Page:  "page",
	}
}
