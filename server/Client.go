package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

type Client struct {
	port int
}

// Function returns a full URL address of the server.
func (c *Client) url() string {
	return fmt.Sprintf("http://localhost:%d", c.port)
}

// Function starts a browser in windows environment.
func (c *Client) startBrowser() error {
	return exec.Command("cmd", "/c", "start", c.url()).Start()
}

// Function starts the server.
func (c *Client) startServer() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", c.port), nil)
}

// Function returns a html page template created from string.
func getTemplate(pageBody string) (*template.Template, error) {
	t, err := template.New("").Parse(pageBody)
	return t, err
}
