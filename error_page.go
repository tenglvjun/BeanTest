package main

import (
	"net/http"
)

type ErrorPage struct {
	Common PageCommon
}

func (e ErrorPage) Render(w http.ResponseWriter, req *http.Request) {
	RenderPage(e, w, req, false)
}

func (e ErrorPage) GetPath() string {
	return e.Common.Path
}

func (e ErrorPage) GetTplFile() string {
	return e.Common.TplFile
}
