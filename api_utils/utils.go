// Author: Paweł Konopko
// License: MIT

package api_utils

import (
	"net/http"
	"encoding/json"
)

func ParseJSONRequest(r *http.Request, container interface{}) error {
	err := json.NewDecoder(r.Body).Decode(container)
	if err != nil {
		return err
	}
	return nil
}

func WriteJSONResponse(w http.ResponseWriter, content interface{}) {
	json.NewEncoder(w).Encode(content)
}

func WriteErrorResponse(w http.ResponseWriter, err *HandlerError) {
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(err.Code)
	w.Write([]byte(err.Message))
}
