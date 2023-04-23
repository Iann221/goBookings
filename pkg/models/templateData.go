package models

// TemplateData: holds data sent from handler to template
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{} // buat objek
	CSRFToken string
	Flash string
	Warning string
	Error string
}
