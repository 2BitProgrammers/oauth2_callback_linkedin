package html

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// RequestInfo contains all the data about the initial request
type requestInfo struct {
	URI     string `json:"uri"`
	Method  string `json:"method"`
	Payload string `json:"payload"`
}

// Response is what we will send back to the user
type response struct {
	Date       time.Time   `json:"date"`
	StatusCode int         `json:"statusCode"`
	StatusText string      `json:"statusText"`
	Data       string      `json:"data"`
	Errors     string      `json:"errors"`
	Request    requestInfo `json:"request"`
}

// returnResponse - this will return a json response to the web client
func returnResponseJson(w http.ResponseWriter, method string, uri string, requestPayload string, status int, statusText string, data string, errors string) {
	sResponse := response{Date: time.Now(), StatusCode: status, StatusText: statusText, Errors: errors}
	sResponse.Data = data
	sResponse.Request.URI = uri
	sResponse.Request.Method = method
	sResponse.Request.Payload = requestPayload

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if method == "OPTIONS" {
		log.Printf("%s %s %d %s", method, uri, http.StatusOK, requestPayload)
		return
	}

	log.Printf("%s %s %d %s", method, uri, status, requestPayload)
	if errors != "" {
		log.Printf("[ERROR] %s - %s", uri, errors)
	}

	joResponse, err := json.Marshal(sResponse)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("[ERROR] Internal Server Error - Failed to parse Json.  %s", err), http.StatusInternalServerError)
		return
	}

	if status == 200 {
		w.Write(joResponse)
		w.Write([]byte("\n\n\n"))
	} else {
		http.Error(w, string(joResponse), status)
	}
}

// returnResponseHtml - this will return a html response to the web client
func returnResponseHtml(w http.ResponseWriter, r *http.Request, htmlData string) {
	method := r.Method
	uri := r.RequestURI
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	// Allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	log.Printf("%s %s %d %s", method, uri, http.StatusOK, "OK")
	w.Write([]byte(htmlData))
}
