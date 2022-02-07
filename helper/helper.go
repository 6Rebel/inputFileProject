package helper

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func PostCallService(api string, request []byte) (*http.Response, error){
	res, err := http.Post("http://127.0.0.1:8098"+api, "application/json; content=UTF-8", bytes.NewBuffer(request))
	if err != nil {
		logrus.Errorf("CallService: error in calling service: %v", err)
		return nil, err
	}
	return res, nil
}

func EncodeJSONBody(resp http.ResponseWriter, statusCode int, data interface{}) {
	resp.WriteHeader(statusCode)
	err := json.NewEncoder(resp).Encode(data)
	if err != nil {
		logrus.Errorf("Error encoding response %v", err)
	}
}