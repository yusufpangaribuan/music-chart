package util

import (
	"net/http"

	models "github.com/lp/music-chart/internal/model"
	"github.com/mholt/binding"
)

func BindQueryParams(req *http.Request) (params models.BasicSelectParams) {
	binding.Bind(req, &params)
	if params.Limit == 0 {
		params.Limit = 10
	}
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Page > 1 {
		params.Offset = params.Limit * (params.Page - 1)
	}
	return
}
