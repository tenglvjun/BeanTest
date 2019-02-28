package main

import (
	"net/http"
)

type IndexPage struct {
	Common PageCommon
}

func (i IndexPage) Render(w http.ResponseWriter, req *http.Request) {
	RenderPage(i, w, req, true)
}

func (i IndexPage) GetPath() string {
	return i.Common.Path
}

func (i IndexPage) GetTplFile() string {
	return i.Common.TplFile
}
