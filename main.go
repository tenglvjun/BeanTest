package main

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

func Redirect(path string, w http.ResponseWriter, req *http.Request) {
	inter := router.GetPage(path)
	if inter == nil {
		log.Panicln("can't find the page interface")
	}

	http.Redirect(w, req, inter.GetPath(), http.StatusFound)
}

func RouterHandler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	if url == "/" {
		// if the url is root, redirect to the index page
		Redirect("welcome", w, req)
		return
	}

	s := strings.Split(url, "/")

	// get the page interface
	inter := router.GetPage(s[1])
	if inter != nil {
		l := len(s)

		// Render page pattern
		if l == 2 {
			inter.Render(w, req)
			return
		}

		// method function pattern
		if l == 3 {
			methodName := s[2]
			method := reflect.ValueOf(inter).MethodByName(methodName)

			in := make([]reflect.Value, 2)
			in[0] = reflect.ValueOf(w)
			in[1] = reflect.ValueOf(req)
			method.Call(in)
			return
		}
	}

	// something wrong
	Redirect("error", w, req)
}

func InitSetting() {
	setting = Setting{
		Host: "",
		Port: "6767",

		DBHost: "",
		DBName: "bean_test",
		DBUser: "root",
		DBPwd:  "tlj789520",
	}
}

func InitRouter() {
	router = Router{}
	router.Init()
}

func InitStaticFileServer() {
	http.Handle("/css/", http.FileServer(http.Dir(".")))
	http.Handle("/img/", http.FileServer(http.Dir(".")))
	http.Handle("/js/", http.FileServer(http.Dir(".")))
}

func InitRouterHandler() {
	http.HandleFunc("/", RouterHandler)
}

func InitEnv() {
	InitSetting()
	InitStaticFileServer()
	InitRouter()
	InitRouterHandler()
}

func main() {
	InitEnv()

	err := http.ListenAndServe(setting.GetListenAndServe(), nil)
	if err != nil {
		log.Fatal("start server failed: ", err)
	}
}
