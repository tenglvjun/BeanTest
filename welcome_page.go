package main

import (
	"net/http"
)

type WelcomePage struct {
	Common PageCommon
}

func (w WelcomePage) Render(rw http.ResponseWriter, req *http.Request) {
	RenderPage(w, rw, req, true)
}

func (w WelcomePage) GetPath() string {
	return w.Common.Path
}

func (w WelcomePage) GetTplFile() string {
	return w.Common.TplFile
}
