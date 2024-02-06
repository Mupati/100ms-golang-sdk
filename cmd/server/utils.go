package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

// Helper method to make all api calls to 100ms
func PerformHTTPCall(url, method string, payload *bytes.Buffer) (string, error) {

	var requestBody io.Reader
	managementToken, err := GenerateManagementToken(os.Getenv("APP_ACCESS_KEY"), os.Getenv("APP_SECRET"))

	if err != nil {
		return "", err
	}

	client := &http.Client{}

	if payload == nil {
		requestBody = nil
	} else {
		requestBody = payload
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return "", err
	}
	// Add Authorization header
	req.Header.Add("Authorization", "Bearer "+managementToken)
	req.Header.Add("Content-Type", "application/json")

	// Send HTTP request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	resp, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	return string(resp), nil

}

func StructToJSONPayload(input interface{}) (*bytes.Buffer, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonBytes), nil
}

// encodeStructToQueryString encodes the fields of a struct into a query string
func EncodeStructToQueryString(s interface{}) (string, error) {
	values, err := EncodeStructToValues(s)
	if err != nil {
		return "", err
	}
	return values.Encode(), nil
}

// encodeStructToValues encodes the fields of a struct into url.Values
func EncodeStructToValues(s interface{}) (url.Values, error) {
	values := url.Values{}
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		// Add the field value to the url.Values
		values.Add(fieldName, fmt.Sprintf("%v", val.Field(i).Interface()))
	}

	return values, nil
}
