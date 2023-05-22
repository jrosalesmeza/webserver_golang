package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jrosalesmeza/webserver_golang/utils"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Este es el Home!")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata utils.MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "Error decoding %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user utils.User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error decoding %v", err)
		return
	}
	response, err := user.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	return
}
