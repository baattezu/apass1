package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/foradmins", app.foradmins)
	mux.HandleFunc("/foradmins/createnews", app.createNews)
	mux.HandleFunc("/foradmins/deletenews", app.deleteNews)
	mux.HandleFunc("/forstudents", app.forStudents)
	mux.HandleFunc("/forstaff", app.forStaff)
	mux.HandleFunc("/forapplicants", app.forApplicants)
	return mux
}
