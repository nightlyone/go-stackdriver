package stackdriver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Stackdriver API endpoint.
var ApiEndpoint = "https://custom-gateway.stackdriver.com/v1"

// Your API, generated in the Stackdriver console.
var ApiKey = os.ExpandEnv("${STACKDRIVER_API_KEY}")

// If we don't get a positive HTTP code, this is is the error returned with details about the request and the received HTTP status code
type SubmissionError struct {
	Request    *http.Request
	StatusCode int
}

// Pretty print error
func (e *SubmissionError) Error() string {
	return fmt.Sprint("Stackdriver submission to ", e.Request.URL.String(), "returned HTTP status code", e.StatusCode)
}

// Reported on missing api key
var ErrApiKeyMissing = errors.New("Stackdriver ApiKey missing. STACKDRIVER_API_KEY environment variable not set?")

// Submit implements common submission logic for Metric, Deploy and Annotation types.
func submit(data interface{}) (err error) {
	var api string
	switch data.(type) {
	case *Metrics, Metrics:
		api = "custom"
	case *Deploy, Deploy:
		api = "deployevent"
	case *Annotation, Annotation:
		api = "annotationevent"
	default:
		panic(fmt.Sprintf("Don't know how to submit %T", data))
	}

	if ApiKey == "" {
		return ErrApiKeyMissing
	}

	message, err := json.Marshal(data)
	if err != nil {
		return err
	}

	body := bytes.NewReader(message)
	req, err := http.NewRequest("POST", ApiEndpoint+"/"+api, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Stackdriver-Apikey", ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return &SubmissionError{
			Request:    resp.Request,
			StatusCode: resp.StatusCode,
		}
	}
	defer resp.Body.Close()
	return err
}
