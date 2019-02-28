package main

import (
	"html/template"
	"log"
	"net/http"
)

func OutputJson(r Result, w http.ResponseWriter) {
	data, err := MarshalToJson(r)
	if err != nil {
		log.Println("OutputJson marshall result failed: ", err)
		errResult := Result{
			Code: RESULT_FAILED,
			Msg:  ERROR_MSG_SERVER_ERROR,
			Data: nil,
		}

		data, _ = MarshalToJson(errResult)
	}

	w.Write(data)
}

// Render page method, exclude error page
func RenderPage(page InterfacePage, w http.ResponseWriter, req *http.Request, doRedirect bool) error {
	tplFile := page.GetTplFile()

	t, err := template.ParseFiles(tplFile, "pages/common_header.html", "pages/header.html")
	if err != nil {
		log.Println("parse template file falied: " + tplFile)
		if doRedirect {
			Redirect("error", w, req)
		}
		return err
	}

	err = t.Execute(w, page)
	if err != nil {
		log.Println("execute page template file failed: " + tplFile)

		if doRedirect {
			Redirect("error", w, req)
		}
		return err
	}

	// success
	return nil
}
