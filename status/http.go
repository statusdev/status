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
	res, err := http.NewRequest(http.MethodDelete, server+"/subscribers", b)
	if err != nil {
		return err
	}

	if res.Response.StatusCode > 200 {
		return fmt.Errorf("Received status %s from %s", res.Response.Status, server)
	}
	return nil
}

// Add our server to a subscription list of a remote server
func AddSubscription(profile Profile, server string) error {
	fmt.Printf(server)
	res, err := postRequest(profile, server+"/subscribers")
	if err != nil {
		return err
	}
	if res.StatusCode > 200 {
		return fmt.Errorf("received status %s from %s", res.Status, server)
	}
	return nil
}

func Notify(status ProfileStatus, subscribers map[string]Profile) error {
	errors := make([]error, 0)
	for _, subscriber := range subscribers {
		res, err := postRequest(status, subscriber.URL+"/subscribers")
		if err != nil || res.StatusCode > 200 {
			errors = append(errors, err)
			continue
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("errors during notify %v", errors)
	}
	return nil
}

func postRequest(obj interface{}, url string) (	*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(obj)
	res, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return nil, err
	}
	return res, nil
}
