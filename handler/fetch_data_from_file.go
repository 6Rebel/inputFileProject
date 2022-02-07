package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"inputFileProject/helper"
	"inputFileProject/pojo"
	"io/ioutil"
	"net/http"
)

const req_api  = "/api/topTenWords"

func FetchDataFromFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `GoLang_Test`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, header, err := r.FormFile("GoLang_Test")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile("tempFile.txt", fileBytes, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("tempFile.txt")
	if err != nil {
		fmt.Println(err)
	}

	var dataFromFile pojo.TextInput
	dataFromFile.Text = string(data)

	req, err := json.Marshal(dataFromFile)
	if err != nil {
		fmt.Println(err)
	}

	// Post Call To API
	responseFromAPI,err := helper.PostCallService(req_api, req)
	if err != nil {
		logrus.Errorf("CallService: error in calling service: %v", err)
	}

	outputData, _ := ioutil.ReadAll(responseFromAPI.Body)
	err = responseFromAPI.Body.Close()
	if err != nil {
		logrus.Errorf("CallService: error in closing response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//converting byte data into Struct
	var outputDataINJSON []pojo.WordFrequency
	err = json.Unmarshal(outputData, &outputDataINJSON)
	if err != nil {
		logrus.Errorf("CallService: error in converting to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	helper.EncodeJSONBody(w, http.StatusOK, outputDataINJSON)
}