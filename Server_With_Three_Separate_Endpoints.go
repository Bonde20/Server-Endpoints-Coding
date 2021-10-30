
package main

import (
	"fmt"
	"log"
	"net/http"
)

var P Person
var data []int

func homePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "<title>You're Welcome to Go</title>")

	fmt.Fprintf(w, "<h1>Hello World</h1>")
	fmt.Fprintf(w, "<img src='images/smiley.gif' alt='smiley' style='width:200px;height:200px;'>")

}

func post(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>This is a Post Request</h1>")

	P = Person{id: 2356, name: "weaver", age: 29}
	data = append(data, 102, 248, 456, 7570)

}

func get(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>View The Posted Values</h1>")
	fmt.Fprintf(w, "<img src='images/weaver.jpg' alt='weaver' style='width:200px;height:200px;'>")

	fmt.Fprintln(w, "<p> The appended slice is:", data)

	fmt.Fprintf(w, "<p> The ID is: %d | The Name is: %s | The Age is:%d", P.id, P.name, P.age)

}

func main() {

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/post", post)
	http.HandleFunc("/get", get)
	fmt.Println("Server starting...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Person struct {
	id   int
	name string
	age  int
}
