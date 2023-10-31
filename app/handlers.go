package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/insuajuan/GO-url-shortener/utils"
	"html/template"
	"net/http"
)

var urls = make(map[string]string)

type urlData struct {
	OriginalURL  string
	ShortenedURL string
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
		return
	}

	// Serve the HTML form
	htmlFile := "static/homepage.html"
	http.ServeFile(w, r, htmlFile)
}

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Generate a unique shortened key for the original URL
	shortKey := utils.GenerateShortKey()
	urls[shortKey] = originalURL

	// Construct the full shortened URL
	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	// Parse the HTML template
	tmpl, err := template.ParseFiles("static/shorten_response.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Data for template
	data := urlData{
		OriginalURL:  originalURL,
		ShortenedURL: shortenedURL,
	}

	// Execute the template with the data
	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := mux.Vars(r)["shortKey"]

	// Retrieve the original URL from the `urls` map using the shortened key
	originalURL, found := urls[shortKey]
	fmt.Printf("original url is %v", originalURL)
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// Redirect the user to the original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
