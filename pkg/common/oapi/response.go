package oapi

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Headers    map[string]string
	Data       interface{}
	Code       int
	ErrMessage string
	Response   *http.Response
}

func NewResponse(data interface{}) *APIResponse {
	return &APIResponse{Data: data}
}

func SendResp(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func SendFormError(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(data)
}

func ForwardResponse(w http.ResponseWriter, apiResp *APIResponse) {
	// Copy headers from the response to our relay
	w.Header().Set("Content-Type", apiResp.Response.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", apiResp.Response.Header.Get("Content-Length"))
	w.Header().Set("Error-Code", apiResp.Response.Header.Get("Error-Code"))

	// Copy status
	w.WriteHeader(apiResp.Response.StatusCode)

	// Copy the body
	w.Write([]byte(apiResp.ErrMessage))
}

func Redirect(w http.ResponseWriter, url string) error {
	w.WriteHeader(http.StatusSeeOther)
	return SendResp(w, url)
}

func (resp *APIResponse) Send(w http.ResponseWriter) error {
	for key, value := range resp.Headers {
		w.Header().Set(key, value)
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp.Data)
}
