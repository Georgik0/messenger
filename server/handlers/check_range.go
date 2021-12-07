package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HndlInput []struct {
	name   string
	method string
	target string
	json   string
	input  http.Handler
	want   string
}

func CheckTestsRange(cases HndlInput, t *testing.T) {
	for _, current_case := range cases {
		request := httptest.NewRequest(current_case.method, current_case.target, bytes.NewReader([]byte(current_case.json)))
		responseRecoder := httptest.NewRecorder()

		current_case.input.ServeHTTP(responseRecoder, request)
		if responseRecoder.Body.String() != current_case.want {
			t.Errorf("received: %v	expected: %v\n", responseRecoder.Body.String(), current_case.want)
		} else {
			t.Logf("Ok: %v\n", responseRecoder.Body.String())
		}
	}
}
