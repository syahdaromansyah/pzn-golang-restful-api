package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, anyCreateRequest any) error {
	jsonDecoder := json.NewDecoder(r.Body)
	return jsonDecoder.Decode(anyCreateRequest)
}

func WriteToResponseBody(w http.ResponseWriter, webResponse any) error {
	jsonEncoder := json.NewEncoder(w)
	return jsonEncoder.Encode(webResponse)
}
