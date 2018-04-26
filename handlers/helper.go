package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func getURL(r *http.Request) string {
	ctx := r.Context()
	return ctx.Value("URL").(string)
}

func joinStatusCode(code int, status string) string {
	return strconv.Itoa(code) + " " + status
}

func httpError(w http.ResponseWriter, err error) bool {
	if err != nil {
		// print the error to the stdout
		log.Println(err.Error())

		// write the error to the response
		http.Error(w, err.Error(), 500)
		return true
	}
	return false
}

func httpWrite(res response) {
	t := &http.Response{
		Status:        res.status,
		StatusCode:    res.code,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString("res.body")),
		ContentLength: int64(len(res.body)),
		Request:       res.request,
		Header:        make(http.Header, 0),
	}

	buff := bytes.NewBuffer(nil)
	t.Write(buff)

	fmt.Println(buff)
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	// encode the response data to json
	jData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// set the json response content-type header
	w.Header().Set("Content-Type", "application/json")
	// write the data to the response
	w.Write(jData)
	return nil
}
