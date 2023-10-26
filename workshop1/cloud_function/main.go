package p

import (
	"encoding/json"
	"net/http"
)

type HelloWorldRequestBody struct {
	Message string `json:"message"`
}

type HelloWorldResponseBody struct {
	NewMessage string `json:"new_message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var body HelloWorldRequestBody

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Write([]byte("Error parsing request body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Do something with the request body
	response := HelloWorldResponseBody{
		NewMessage: body.Message + ". Make GDSC great again!",
	}

	// Encode the response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Error encoding response"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
