package languages

const (
	// LanguagesGetLanguages is a string representation of the current endpoint for getting languages
	LanguagesGetLanguages = "v1/metadata/getLanguages"
)

// Language is a struct containing matching data for a language found in text
type Language struct {
	// Name - the language identified
	Name string `json:"name"`
	// Confidence - a float value from 0.0 to 1.0 of our trust in the result
	Confidence float32 `json:"confidence"`
}
