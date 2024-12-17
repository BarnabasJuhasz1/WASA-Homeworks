//go:build !webui

package main

import (
	"net/http"
)

// registerWebUI is an empty stub because `webui` tag has not been specified.
func RegisterWebUI(hdl http.Handler) (http.Handler, error) {
	return hdl, nil
}
