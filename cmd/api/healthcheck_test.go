package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHealthCheckHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	app := &application{}
	app.HealthCheckHandler(rr, request)
	rs := rr.Result()
	if rs.StatusCode != http.StatusOK {
		t.Errorf(cmp.Diff(rs.StatusCode, http.StatusOK))
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	var result map[string]map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatalf(err.Error())
	}
	content := map[string]string{
		"status":      "Available",
		"environment": app.config.env,
	}
	if result["status"]["status"] != content["status"] {
		t.Errorf(cmp.Diff(result["status"]["status"], content["status"]))
	}
	if result["status"]["environment"] != content["environment"] {
		t.Errorf(cmp.Diff(result["status"]["environment"], content["environment"]))
	}
}
