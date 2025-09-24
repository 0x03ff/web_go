package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}







func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)

    // Use json.Marshal to marshal the value into JSON
    jsonData, err := json.Marshal(v)
    if err != nil {
        return err
    }

    // Write the JSON data to the response writer
    _, err = w.Write(jsonData)
    return err
}




func WriteError(w http.ResponseWriter, status int, err error){
	WriteJSON(w, status, map[string]string{"error":err.Error()})
}