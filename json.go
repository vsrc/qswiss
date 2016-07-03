package qswiss
 
import (
	"encoding/json"
)

// GetJSON transforms map to json array of byte
// function trims trailing space too
func GetJSON(code map[string]interface{}) []byte {

	js, err := json.MarshalIndent(code, "", "    ")

	if err != nil {
		PanicIf(err)
	}

	return js
}
