package handlers

import (
	"fmt"
	"net/http"

	"now-go-mod-child/handlers/child"
)

func Child(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, child.Name())
}
