package status

import (
	"bytes"
	"encoding/json"
	"net/http"
)
// Remove our server from a subscription list of a remote server
func RemoveSubscription(profile Profile, server string) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(profile)
	res, _ := http.NewRequest(http.MethodDelete, server + "/subscribers", b)
}

// Add our server to a subscription list of a remote server
func AddSubscription(profile Profile, server string) error {
	postRequest(profile, server + "/subscribers")
}


func Notify(status ProfileStatus) error {
	panic("implement me")
}

func postRequest(obj interface{}, url string) *http.Response {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(obj)
	res, _ := http.Post(url, "application/json; charset=utf-8", b)
	return res
}