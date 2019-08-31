package server

import (
	"../scrapers"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Handler for /pracuj endpoint.
func (c *Client) pracujHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("[ERROR] Only GET method is defined.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Reading parameters
	city := r.URL.Query().Get("city")
	keywords := r.URL.Query()["keyword"]
	URL := fmt.Sprintf("https://www.pracuj.pl/praca/%s;wp", city)
	log.Printf("[START] Searching in <%s> for: %s.\n", city, strings.Join(keywords, ", "))

	// Creating collector
	sc := scrapers.CreateClient(city, keywords)
	sc.AddHandlersPracuj()

	sc.Visit(URL)
	sc.Wait()

	// Creating the results' page template.
	t, errPage := getTemplate(c.pracujPageBody())
	if errPage != nil {
		log.Printf("[ERROR] Couldn't open the page: %s.\n", errPage)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Putting the results into the template.
	errExecute := t.Execute(w, sc)
	if errExecute != nil {
		log.Printf("[ERROR] Couldn't print results: %s.\n", errExecute)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	log.Println("[END] Searching finished.")
}

// Function returns main page body.
func (c *Client) pracujPageBody() string {
	return fmt.Sprint(`
	<font face='verdana'>
	<h2>Parametry</h2>
	<div style="margin-left: 0.5em;">
		<strong>Miasto:</strong> {{.City}}<br><br>
		<strong>Słowa kluczowe:</strong>
			<ul style="margin-left: 0.5em;">
				{{range $i, $k := .Keywords}}
					<li>{{ $k }}</li>
				{{ end }}
			</ul>
	</div>
	<h2>Wyniki</h2>
	<div style="margin-left: 0.5em;">
		<strong>Liczba przetworzonych ofert:</strong> {{.OffersCounter}}
		{{range $keyword, $values := .Data}}
			<h3>{{ $keyword }}</h3>
			{{range $i, $val := $values}}
				<a href="{{ $val }}" style="margin-left: 0.5em;">{{ $val }}</a><br>
			{{ end }}
		{{ end }}
	</div>
	</font>
	`)
}
