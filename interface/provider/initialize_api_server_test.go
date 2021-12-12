package provider

import "testing"

func TestInitializeAPIServer(t *testing.T) {
	t.Helper()
	_, err := InitializeAPIServer()
	if err != nil {
		t.Errorf("Failed Initialize Server! 1: %v", err.Error())
	}
}
