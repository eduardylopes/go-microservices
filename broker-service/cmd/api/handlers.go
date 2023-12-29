package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload, nil)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	requestPayload := RequestPayload{}

	if err := app.readJSON(w, r, &requestPayload); err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	case "log":
		app.logItem(w, requestPayload.Log)
	default:
		app.errorJSON(w, errors.New("unknown action"), http.StatusBadRequest)
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	JSONData, _ := json.MarshalIndent(a, "", "\t")

	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(JSONData))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"), http.StatusBadRequest)
		return
	}

	JSONFromService := JSONResponse{}

	if err := json.NewDecoder(response.Body).Decode(&JSONFromService); err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if JSONFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	payload := JSONResponse{
		Error:   false,
		Message: "authenticated",
		Data:    JSONFromService.Data,
	}

	app.writeJSON(w, http.StatusAccepted, payload, nil)
}

func (app *Config) logItem(w http.ResponseWriter, entry LogPayload) {
	JSONData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	req, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(JSONData))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := JSONResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, payload, nil)
}
