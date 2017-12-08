package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Serve Http:", err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path[1:] {
	case "check":
		fmt.Fprint(w, "If you are reading this that means your application is UP")
	case "api":
		fmt.Fprint(w, "API is working")
	default:
		fmt.Fprint(w, "can't answer that!")
	}
}
