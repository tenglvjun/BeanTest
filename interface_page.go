package main

import (
	"net/http"
)

type PageCommon struct {
	Title   string
	Path    string
	TplFile string
}

type InterfacePage interface {
	Render(w http.ResponseWriter, req *http.Request)
	GetPath() string
	GetTplFile() string
}
