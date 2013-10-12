package stackdriver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var ApiEndpoint = "https://custom-gateway.stackdriver.com/v1"
var ApiKey = os.ExpandEnv("${STACKDRIVER_API_KEY}")

// Common submission logic
func Submit(data interface{}) (err error) {
	var api string
	switch data.(type) {
	case *Metrics:
		api = "custom"
	case Metrics:
		api = "custom"
	case *Deploy:
		api = "deployevent"
	case Deploy:
		api = "deployevent"
	case *Annotation:
		api = "annotationevent"
	case Annotation:
		api = "annotationevent"
	default:
		log.Panic("Don't know how to submit %T", data)
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
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		log.Print("FATAL: query ", data)
		log.Print("FATAL: request ", resp.Request)
		log.Print("FATAL: messsage", message)
		log.Fatal("FATAL: Status != 200, got ", resp.Status)
	}
	defer resp.Body.Close()
	return err
}
