package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, requeat *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", requeat.URL.Path[1:])

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
