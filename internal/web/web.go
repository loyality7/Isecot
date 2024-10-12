package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/loyality7/Isecot/internal/scanner"
)

func StartServer() {
	fmt.Println("Starting web server on http://localhost:8080")
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/scan", handleScan)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/templates/index.html")
	tmpl.Execute(w, nil)
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	scanner.ScanNetwork()
	fmt.Fprintf(w, "Scan complete. Check console for results.")
}
