package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates  *template.Template
type person struct {
	FirstName  string   `json:"firstName"`
	SecondName string   `json:"secondName"`
	Age        int      `json:"age"`
	Graduation []string `json:"graduation"`
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!")) // Its necessary to deliver in byte slice format.
}

func homepage(w http.ResponseWriter, r *http.Request) {

	templates = template.Must(template.ParseGlob("*.html")) // will take all the html files in the current directory and parse them into a template

	user := person{
		FirstName:  "Leonardo",
		SecondName: "Suave",
		Age:        30,
		Graduation: []string{"Bachelor's Degree", "Master's Degree"}}

	w.WriteHeader(http.StatusOK)
	templates.ExecuteTemplate(w, "home.html", user) // Execute the template and write it to the response writer (nil (parameter 3) can be datas to pass to html file)
}	

func user(w http.ResponseWriter, r *http.Request) {

	// create a person struct and return it as json
	user := person{
		FirstName:  "John",
		SecondName: "Doe",
		Age:        30,
		Graduation: []string{"Bachelor's Degree", "Master's Degree"}}
	
	w.Header().Set("Content-Type", "application/json") // To set the content type of the response to json
	w.WriteHeader(http.StatusOK)

	// Encode the person struct to json and write it to the response writer
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Erro ao retornar JSON:", err)
	}
}

func users(w http.ResponseWriter, r *http.Request) {

	// Read data from json file and unmarshal it into a slice of person structs
	file, error := os.ReadFile("../class-18-json/unmarshal/example.json")
	if error != nil {
		log.Fatal(error)
	}

	var users []person
	if error := json.Unmarshal(file, &users); error != nil {
		log.Fatal(error)
	}

	w.Header().Set("Content-Type", "application/json") // To set the content type of the response to json
	w.WriteHeader(http.StatusOK)

	// Encode the person struct to json and write it to the response writer
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println("Erro ao retornar JSON:", err)
	}
}

// HTTP is a package that provides HTTP server and client implementations.
// Clint do request and server handle the request and response.
func main() {

	// r is a request and w is a response writer. The request is sent by the client and the response is sent by the server.
	http.HandleFunc("/home", home)

	http.HandleFunc("/homepage", homepage)

	http.HandleFunc("/user", user)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":3000", nil)) // To start the server
}
