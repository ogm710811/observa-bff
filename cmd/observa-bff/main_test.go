package main

import (
	"context"
	"testing"
)

func TestHandler_ReturnsExpectedGreeting(t *testing.T) {
	// Arrange: set up the input event and expected output
	event := Event{Message: "observa-developer"}
	expected := "Hello and welcome, observa-developer!"

	// Act: call the handler directly
	resp, err := handler(context.Background(), event)
	if err != nil {
		t.Fatalf("handler returned an error: %v", err)
	}

	// Assert: check output matches expectation
	if resp.Greeting != expected {
		t.Errorf("expected greeting %q, got %q", expected, resp.Greeting)
	}
}
