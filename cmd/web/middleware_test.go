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

// Generate Password
// You can edit this code!
// Click here and start typing.
// package main

// import (
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

// func main() {
// 	pwd := "password"

// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pwd), 12)

// 	fmt.Println(string(hashedPassword))
// }
