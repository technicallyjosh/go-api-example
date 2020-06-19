package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// This is just a placeholder to check against. This would normally be hashed with a salt in a
// database and then verified against the salt and hash based on input... For example purposes this
// will do
const (
	dbUsername = "technicallyjosh"
	dbPassword = "testing123"
)

// VerifyUserRequest represents the request body of the verifyUserHandler function
type VerifyUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ErrorResponse represents the response body when an error happens
type ErrorResponse struct {
	Message string `json:"message"`
}

// Validate validates the data on an instance of VerifyUserRequest
// Normally I would do validation using something like https://github.com/go-playground/validator
// but for simplicity and showing patterns, this will do.
func (req *VerifyUserRequest) Validate() error {
	if strings.TrimSpace(req.Username) == "" {
		return errors.New("username is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

// sendJSON is a helper function to send back responses without too much repetition
// I do know that the preferred way to marshal json with less overhead is to do something like
// `json.NewEncoder(w).Encode(body)` but it's a bit harder to test in this scenario as the encoder
// returns a new line
func sendJSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	bytes, _ := json.Marshal(body)
	w.Write(bytes)
}

// PUT /users/verify
// In application, this would most likely not be called with the username as that can be pulled from
// the current JWT issued and being passed in. For this, it just includes it. This route is
// basically just a simpler login function...
func verifyUserHandler(w http.ResponseWriter, r *http.Request) {
	var body VerifyUserRequest

	// ran into this when testing and you send a nil value
	if r.Body == nil {
		sendJSON(w, http.StatusBadRequest, &ErrorResponse{
			Message: "Bad request",
		})
		return
	}

	// decode incoming JSON and throw 400 if error decoding it
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendJSON(w, http.StatusBadRequest, &ErrorResponse{
			Message: "Bad request",
		})
		return
	}

	// simple validation on the body
	if err := body.Validate(); err != nil {
		sendJSON(w, http.StatusUnprocessableEntity, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// verify username
	if strings.ToLower(body.Username) != dbUsername {
		sendJSON(w, http.StatusNotFound, &ErrorResponse{
			Message: "User not found",
		})
		return
	}

	// verify password
	// Depending on the middleware for the service, if accepting a JWT or anything of the sort, we'd
	// probably want to reserve the 401 for that authorization failing and instead return a `true`
	// or `false` here with a 200. Either way, this is just an example :shrug:
	// note: you could also return a 409 conflict meaning (conflict of input and data)
	if body.Password != dbPassword {
		sendJSON(w, http.StatusUnauthorized, &ErrorResponse{
			Message: "Incorrect password",
		})
		return
	}

	// just send back an empty object
	sendJSON(w, http.StatusOK, map[string]interface{}{})
}
