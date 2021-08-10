package oapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type APIRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Data    interface{}
	Result  interface{}
}

func NewRequest(method, url string) *APIRequest {
	return &APIRequest{
		Method: method,
		URL:    url,
	}
}

func (apiReq *APIRequest) Do() (*APIResponse, error) {
	var request *http.Request
	apiResp := new(APIResponse)

	if apiReq.Data != nil {
		dataJSON, err := json.Marshal(apiReq.Data)
		if err != nil {
			return apiResp, fmt.Errorf("Failed marshal data of %s %s request. Error: %v\n", apiReq.Method, apiReq.URL, err)
		}
		request, _ = http.NewRequest(apiReq.Method, apiReq.URL, bytes.NewBuffer(dataJSON))
	} else {
		request, _ = http.NewRequest(apiReq.Method, apiReq.URL, nil)
	}

	request.Header.Set("Content-Type", "application/json")
	for key, value := range apiReq.Headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResp, fmt.Errorf("%s %s request failed with error: %v\n", apiReq.Method, apiReq.URL, err)
	}

	apiResp.Response = response

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		code, _ := strconv.Atoi(response.Header.Get("Error-Code"))
		apiResp.Code = code
		apiResp.ErrMessage = string(body)
		return apiResp, fmt.Errorf("%s %s request failed. Err: %s", apiReq.Method, apiReq.URL, string(body))
	}

	if apiReq.Result != nil {
		apiResp.Data = &apiReq.Result
		if err := json.NewDecoder(response.Body).Decode(&apiResp.Data); err != nil {
			return apiResp, err
		}
	}

	return apiResp, nil
}
