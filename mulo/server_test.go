package mulo

import (
	"testing"
)

func TestStartHealthServer(t *testing.T) {
	healthServer := StartHealthServer()
	if healthServer == nil {
		t.Errorf("Expected health server to != nil")
	}
	state := healthServer.Query()
	if state.Ok != false {
		t.Errorf("Initial health should be false")
	}
}

func TestHealthServer_UpdateAndQuery(t *testing.T) {
	healthServer := StartHealthServer()

	name := "test"
	ok := false
	msg := "Test message"

	healthServer.Update(name, ok, msg)
	state := healthServer.Query()
	if state.Ok != ok {
		t.Errorf("Systemstate ok was %v, expected %v", state.Ok, ok)
	}

	ok = true
	healthServer.Update(name, ok, msg)
	state = healthServer.Query()
	if state.Ok != ok {
		t.Errorf("Systemstate ok was %v, expected %v after update", state.Ok, ok)
	}

}

func TestStopHealthServer(t *testing.T) {
	server := StartHealthServer()
	server.stop <- true
	//Not sure how to test

}
