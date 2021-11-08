package helper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func WriteJson(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(response))
	if err != nil {
		log.Println(err)
		return errors.New("Can't send response")
	}
	return nil
}

func ReadContent(w http.ResponseWriter, r *http.Request, model interface{}) error {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("Can't read content")
	}

	if err := json.Unmarshal(content, model); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("Can't parse model")
	}

	return nil
}

func BodyClose(r *http.Request) {
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}
}

func BodyCloseResponse(r *http.Response) {
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}
}
