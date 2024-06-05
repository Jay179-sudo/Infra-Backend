package main

import (
	"testing"
)

var validCases = map[string]struct {
	input map[string]string
}{
	"No items": {
		input: make(map[string]string),
	},
	"One Item": {
		input: map[string]string{
			"status": "Available",
		},
	},
	"Two Items": {
		input: map[string]string{
			"status":      "Available",
			"environment": "Development",
		},
	},
}

var InvalidCases = map[string]struct {
	input interface{}
}{
	"Channel Input": {
		input: make(chan string),
	},
}

// valid inputs --> inputs parsable by Go
func TestJsonValidForValidInputs(t *testing.T) {
	for name, tc := range validCases {
		t.Run(name, func(t *testing.T) {
			_, err := dataToJson(envelope{"status": tc.input})
			if err != nil {
				t.Errorf("Could not parse JSON data properly.")
			}
		})
	}
}

func TestJsonInvalidForInvalidInputs(t *testing.T) {
	for name, tc := range InvalidCases {
		t.Run(name, func(t *testing.T) {
			_, err := dataToJson(envelope{"status": tc.input})
			if err == nil {
				t.Errorf("Could not parse JSON data properly.")
			}
		})
	}
}
