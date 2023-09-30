package errs

import (
	"encoding/json"
	"net/http"
)

func ReturnError(w http.ResponseWriter, code Code) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code.Status)

	return json.NewEncoder(w).Encode(code)
}
