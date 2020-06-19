package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// createRequest is a simple abstraction to return a new request to test with
func createRequest(body io.Reader, t *testing.T) *http.Request {
	req, err := http.NewRequest("PUT", "/users/verify", body)

	if err != nil {
		t.Fatal(err)
	}

	return req
}

// createRecorder creates a recorder from the specified request, creates the handler, and returns
// the recorder
func createRecorder(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(verifyUserHandler)

	handler.ServeHTTP(rr, req)

	return rr
}

// expectStatusCode runs got/expected code for less repetition
func expectStatusCode(t *testing.T, got int, expected int) {
	if got != expected {
		t.Errorf("handler returned wrong status code: got %v want %v", got, expected)
	}
}

// expectBody runs got/expected body for less repetition
func expectBody(t *testing.T, got interface{}, expected interface{}) {
	if got != expected {
		t.Errorf("handler returned wrong body: got %v want %v", got, expected)
	}
}

func TestVerifyUserMissingBody(t *testing.T) {
	req := createRequest(nil, t)

	rr := createRecorder(req)

	expectStatusCode(t, rr.Code, http.StatusBadRequest)

	expected := `{"message":"Bad request"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned wrong body: got %v want %v", rr.Body, expected)
	}
}

func TestVerifyUserMissingUsername(t *testing.T) {
	req := createRequest(strings.NewReader(`{}`), t)

	rr := createRecorder(req)

	expectStatusCode(t, rr.Code, http.StatusUnprocessableEntity)

	expected := `{"message":"username is required"}`
	expectBody(t, rr.Body.String(), expected)
}

func TestVerifyUserMissingPassword(t *testing.T) {
	body := strings.NewReader(`{"username":"technicallyjosh"}`)
	req := createRequest(body, t)

	rr := createRecorder(req)

	expectStatusCode(t, rr.Code, http.StatusUnprocessableEntity)

	expected := `{"message":"password is required"}`
	expectBody(t, rr.Body.String(), expected)
}

func TestVerifyUserMissingUser(t *testing.T) {
	body := strings.NewReader(`{"username":"josh","password":"testing123"}`)
	req := createRequest(body, t)

	rr := createRecorder(req)

	expectStatusCode(t, rr.Code, http.StatusNotFound)

	expected := `{"message":"User not found"}`
	expectBody(t, rr.Body.String(), expected)
}

func TestVerifyUserInvalidPassword(t *testing.T) {
	body := strings.NewReader(`{"username":"technicallyjosh","password":"testing"}`)
	req := createRequest(body, t)

	rr := createRecorder(req)

	expectStatusCode(t, rr.Code, http.StatusUnauthorized)

	expected := `{"message":"Incorrect password"}`
	expectBody(t, rr.Body.String(), expected)
}

func TestVerifyUserValid(t *testing.T) {
	body := strings.NewReader(`{"username":"technicallyjosh","password":"testing123"}`)
	req := createRequest(body, t)

	rr := createRecorder(req)

	expectStatusCode(t, rr.Code, http.StatusOK)

	expected := `{}`
	expectBody(t, rr.Body.String(), expected)
}
