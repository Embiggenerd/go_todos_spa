package main

import (
	_ "context"
	"fmt"
	_ "html/template"
	"net/http"
	_ "postGres/models"
	_ "postGres/utils"
	_ "strconv"

	_ "github.com/lib/pq"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div id=root></div>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
