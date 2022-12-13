package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleUserCreate(t *testing.T) {
	req, err := http.NewRequest("POST", "/crete", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleUserCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
