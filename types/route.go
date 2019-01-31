package types

import (
	"net/http"
)

// Route describes the actual information used by the router to create endpoints.
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Name	    string
}