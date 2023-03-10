package models

import "github.com/fxfrancky/bookings/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	// CSRFTOKEN Cross Site Request Forgery Token
	CSRFToken string
	// Messages types to send users
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
