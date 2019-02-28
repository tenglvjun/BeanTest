package main

import (
	"net/http"
)

type GeometryPage struct {
	Common PageCommon
}

func (g GeometryPage) Render(w http.ResponseWriter, req *http.Request) {
	RenderPage(g, w, req, true)
}

func (g GeometryPage) GetPath() string {
	return g.Common.Path
}

func (g GeometryPage) GetTplFile() string {
	return g.Common.TplFile
}
