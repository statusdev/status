package status

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Remove our server from a subscription list of a remote server
func RemoveSubscription(profile Profile, server string) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(profile)
	res, _ := http.NewRequest(http.MethodDelete, server+"/subscribers", b)
	if res.Response.StatusCode > 200 {
		return fmt.Errorf("Received status %s from %s", res.Response.Status, server)
	}
	return nil
}

// Add our server to a subscription list of a remote server
func AddSubscription(profile Profile, server string) error {
	res := postRequest(profile, server+"/subscribers")
	if res.StatusCode > 200 {
		return fmt.Errorf("received status %s from %s", res.Status, server)
	}
	return nil
}

func Notify(status ProfileStatus, subscribers map[string]Profile) error {
	errors := make([]string, 0)
	for _, subscriber := range subscribers {
		res := postRequest(status, subscriber.URL+"/subscribers")
		if res.StatusCode > 200 {
			message := fmt.Sprintf("received status %s from %s", res.Status, subscriber.URL)
			errors = append(errors, message)
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("errors during notify %v", errors)
	}
	return nil
}

func postRequest(obj interface{}, url string) *http.Response {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(obj)
	res, _ := http.Post(url, "application/json; charset=utf-8", b)
	return res
}
