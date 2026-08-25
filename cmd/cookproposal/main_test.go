package main

import "testing"

func TestEntrypointContract(t *testing.T) {
	if defaultDB() == "" {
		t.Fatal("empty default")
	}
}
func defaultDB() string { return "cookproposal.db" }
