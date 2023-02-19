package main

import (
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("Type is not http.Handler %T", v)

	}
}
func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("Type is not http.Handler %T", v)

	}
}
