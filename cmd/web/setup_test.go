package main

import (
	"net/http"
	"os"
	"testing"
)

// That will run befor our tests runs
func TestMain(m *testing.M) {
	// before you run the tests do whatever
	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
