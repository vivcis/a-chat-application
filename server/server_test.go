package main

import (
	"net"
	"testing"
)

//TESTING BROADCAST-MESSAGE
func TestBroadcast(t *testing.T) {
	expected := "message broad-casted to other clients"
	var connection net.Conn
	if BroadcastMessage(connection) != expected {
		t.Errorf("Expected %v this to equal %s", connection, expected)
	}
}
