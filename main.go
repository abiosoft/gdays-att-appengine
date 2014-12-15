package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"appengine"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/mark", markHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func handler(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)

	data := map[string]interface{}{
		"attendees": getAttendeesCount(c),
	}

	b, err := ioutil.ReadFile("index.html")
	if err != nil {
		serveErr(rw, req, err)
		return
	}

	t := template.Must(template.New("index").Parse(string(b)))
	t.Execute(rw, data)
}

func markHandler(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)

	err := markAttendance(c)
	if err != nil {
		serveErr(rw, req, err)
		return
	}
	http.Redirect(rw, req, "/", 301)
}

func serveErr(rw http.ResponseWriter, req *http.Request, err error) {
	rw.WriteHeader(500)
	rw.Write([]byte(fmt.Sprint("Internal Server Error", err)))
}
