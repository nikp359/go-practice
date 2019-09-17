package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	Product{ID: 1, Name: "Hover Shooters", Slug: "hover-shooters",
		Description: "Shoot your way to the top on 14 different hoverboards"},
	Product{ID: 2, Name: "Ocean Explorer", Slug: "ocean-explorer",
		Description: "Explore the depths of the sea in this one of a kind"},
	Product{ID: 3, Name: "Dinosaur Park", Slug: "dinosaur-park",
		Description: "Go back 65 million years in the past and ride a T-Rex"},
	Product{ID: 4, Name: "Cars VR", Slug: "cars-vr",
		Description: "Get behind the wheel of the fastest cars in the world."},
	Product{ID: 5, Name: "Robin Hood", Slug: "robin-hood",
		Description: "Pick up the bow and arrow and master the art of archery"},
	Product{ID: 6, Name: "Real World VR", Slug: "real-world-vr",
		Description: "Explore the seven wonders of the world in VR"},
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	//API
	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	http.ListenAndServe(":3000", r)
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
