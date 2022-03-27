package main

import (
	"net"
	"testing"
)

//TEST READING FROM SERVER
func TestRead(t *testing.T) {
	expected := "network not reading message"
	var connection net.Conn
	if Reader(connection) != expected {
		t.Errorf("Expected %v this to equal %s", connection, expected)
	}
}

//TEST WRITING TO SERVER
func TestWrite(t *testing.T) {
	expected := "network not writing message"
	var connection net.Conn
	var clientName string
	if Writer(connection, clientName) != expected {
		t.Errorf("Expected %v this to equal %s", connection, expected)
	}
}
