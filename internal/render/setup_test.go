package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fxfrancky/bookings/internal/config"
	"github.com/fxfrancky/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// What am going to put in the session
	gob.Register(models.Reservation{})

	// Change this to true when in production
	testApp.InProduction = false

	// Create an info Logger (An info log to print to the console)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	// Create an error Logger (An error log to print to the console) Lshortfile gives more info about the error
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	// Initialize session scs
	session = scs.New()
	// Set The session duration (24 H)
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// Update the config session field
	testApp.Session = session
	app = &testApp

	// The above test will be executed before any other test in the package
	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
