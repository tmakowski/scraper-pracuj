package server

import (
	"fmt"
	"log"
	"net/http"
)

// Function allows PageClient to serve as a handler function for server. This
// function is used as a / endpoint.
func (c *Client) mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := getTemplate(c.mainPageBody())
	if err != nil {
		log.Printf("[ERROR] Couldn't open the page: %s.\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, nil)
}

// Function returns main page body.
func (c *Client) mainPageBody() string {
	return fmt.Sprintf(`
	<script>
		function disableEmptyInputs(form) {
			let controls = form.elements;
			for (let i = 0; i < controls.length; i++) {
				controls[i].disabled = controls[i].value === '';
			}
		}
	</script>
	
	<form action="http://localhost:%d/pracuj" method="get" autocomplete="off"
	onsubmit="disableEmptyInputs(this)">
		<div>
			<label for="city">Miasto (brak = cała Polska):</label><br>
			<input id="city" name="city">
		</div><br>
		<div>
			<label for="keyword">Słowa kluczowe:</label><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword"><br>
			<input id="keyword" name="keyword">
		</div><br>
		<button>Szukaj</button>
	</form>
	`, c.port)
}
