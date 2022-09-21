package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/",helloWorld).Methods("GET")
	route.HandleFunc("/home", home).Methods("GET")
	route.HandleFunc("/contact",contact).Methods("GET")
	route.HandleFunc("/project",project).Methods("GET")
	route.HandleFunc("/form-project",AddProject).Methods("POST")
	route.HandleFunc("/form-contact",AddContact).Methods("POST")

	fmt.Println("server running port 7000")
	http.ListenAndServe("localhost:7000",route)
}
func helloWorld(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}
func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("home.html")

	if err != nil{
		w.Write([]byte("web tidak tersedia" + err.Error()))
		return
	}
	tmpl.Execute(w,nil)
}
func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("contact.html")

	if err != nil{
		w.Write([]byte("web tidak tersedia" + err.Error()))
		return
	}
	tmpl.Execute(w,nil)
}
func project(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("project.html")

	if err != nil{
		w.Write([]byte("web tidak tersedia" + err.Error()))
		return
	}
	tmpl.Execute(w,nil)
}
func AddProject(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Add Project : " + r.PostForm.Get("input-project"))
	fmt.Println("start Date : " + r.PostForm.Get("input-start"))
	fmt.Println("End Date : " + r.PostForm.Get("input-end"))
	fmt.Println("Description : " + r.PostForm.Get("input-description"))
	fmt.Println("Node Js : " + r.PostForm.Get("nodejs"))
	fmt.Println("Vue JS : " + r.PostForm.Get("vuejs"))
	fmt.Println("React Js : " + r.PostForm.Get("reactjs"))
	fmt.Println("java : " + r.PostForm.Get("java"))

	http.Redirect(w,r,"/home",http.StatusMovedPermanently)
}
func AddContact(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nama : " + r.PostForm.Get("input-nama"))
	fmt.Println("email : " + r.PostForm.Get("input-email"))
	fmt.Println("phone Number : " + r.PostForm.Get("input-phone"))
	fmt.Println("subject : " + r.PostForm.Get("input-subject"))
	fmt.Println("Description : " + r.PostForm.Get("input-description"))
	http.Redirect(w,r,"/home",http.StatusMovedPermanently)
}
 
