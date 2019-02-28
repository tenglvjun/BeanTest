package main

import (
	"net/http"
)

type WikipediaPage struct {
	Common PageCommon
}

func (w WikipediaPage) Render(rw http.ResponseWriter, req *http.Request) {
	RenderPage(w, rw, req, true)
}

func (w WikipediaPage) GetPath() string {
	return w.Common.Path
}

func (w WikipediaPage) GetTplFile() string {
	return w.Common.TplFile
}
