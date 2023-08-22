package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/Naturewoker/go-microservices.git/details"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	localIp, err := details.GetLocalIP()
	if err != nil {
		panic(err)
	}
	response := map[string]string{
		"hostname": hostname,
		"local IP": localIp.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("Web server has started")
	log.Fatal(http.ListenAndServe(":80", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("Web server has started")
// 	http.ListenAndServe(":80", nil)
// }

// import (
// 	"fmt"
// 	"unsafe"

// 	geo "github.com/Naturewoker/go-microservices.git/geometry"
// 	"rsc.io/quote"
// )

// func main() {
// 	x := 30
// 	y, z := 10, 20
// 	name := "naturewoker"
// 	isYou := true
// 	fmt.Println(quote.Go())
// 	fmt.Println(x, y, z, name, isYou)
// 	fmt.Printf("Type of name is %T and size is %d", name, unsafe.Sizeof(name))
// 	fmt.Println()
// 	a, p := rectProps(1, 2)
// 	fmt.Printf("Area is %.1f and Perimeter is %.1f", a, p)

// 	// var daysOfTheMonth map[string]int
// 	// daysOfTheMonth["Jan"] = 31
// 	var daysOfTheMonth = map[string]int{"Jan": 31, "Feb": 28}
// 	fmt.Println()
// 	fmt.Println(daysOfTheMonth)

// 	fmt.Println(geo.Area(2, 2))

// }

// func rectProps(length, width float64) (area, perimeter float64) {
// 	area = length * width
// 	perimeter = 2 * (length + width)
// 	return
// }
