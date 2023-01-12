package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lalitlab/go-microservices/details"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")

	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")

	hostname, err := details.GetHostName()
	if err != nil {
		panic(err)
	}

	myIP := details.GetIP()

	response := map[string]string{
		"hostname": hostname,
		"ip":       myIP,
	}

	json.NewEncoder(w).Encode(response)
}

func listPost(w http.ResponseWriter, r *http.Request) {
	log.Println("List posts")

	vars := mux.Vars(r)
	ZipCode := vars["zip_code"]         // the page
	CountryCode := vars["country_code"] // the page

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts" + CountryCode + "/" + ZipCode)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(body))
}

func getPost(w http.ResponseWriter, r *http.Request) {
	log.Println("INSIDE Get post with ID")
	vars := mux.Vars(r)
	postID := vars["post_id"]
	log.Println("Get post with ID", postID)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + postID)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(body))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/details", detailsHandler)
	r.HandleFunc("/posts", listPost)
	r.HandleFunc("/posts/{post_id}", getPost)

	// Start the server
	log.Println("Web server has started!!!")
	http.ListenAndServe(":80", r)
}
