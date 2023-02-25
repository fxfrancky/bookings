package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fxfrancky/bookings/internal/config"
	"github.com/fxfrancky/bookings/internal/driver"
	"github.com/fxfrancky/bookings/internal/handlers"
	"github.com/fxfrancky/bookings/internal/models"
	"github.com/fxfrancky/bookings/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

const portNumber = ":8080"

// main is the main application function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	defer close(app.MainChan)
	// Starting mail listener
	log.Println("Starting mail listener")
	listenForMail()

	// app.MainChan <- msg

	fmt.Printf("Starting application on port %s", portNumber)
	// Use the previous function

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// What am going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})
	// read flags
	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "", "Database name")
	dbUser := flag.String("dbuser", "", "Database user")
	dbPass := flag.String("dbpass", "", "Database password")
	dbPort := flag.String("dbport", "5432", "Database Port")
	dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require)")
	// // Email Flags
	serverHost := flag.String("serverhost", "", "Server host")
	serverPort := flag.String("serverport", "", "Server Port")

	serverUsername := flag.String("serverusername", "", "Server Username")

	serverPassword := flag.String("serverpassword", "", "Server Password")

	flag.Parse()
	if *dbName == "" || *dbUser == "" {
		log.Println("Missing required flags")
		os.Exit(1)
	}
	mailChan := make(chan models.MailData)
	app.MainChan = mailChan

	// Using Email flags
	app.ServerHost = *serverHost
	app.ServerPort = *serverPort
	app.ServerUsername = *serverUsername
	app.ServerPassword = *serverPassword

	// Change this to true when in production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	// Create an info Logger (An info log to print to the console)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	// Create an error Logger (An error log to print to the console) Lshortfile gives more info about the error
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Initialize session scs
	session = scs.New()
	// Set The session duration (24 H)
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// Update the config session field
	app.Session = session

	// Connect database
	log.Println("Connection to database...")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Cannot connect to the database! Dying")
	}
	log.Println("Connected to the database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)

	return db, nil
}
