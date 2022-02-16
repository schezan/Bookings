package main

//template data stores data from templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	Csrftoken string
	Flash     string
	Warning   string
	Error     string
}
