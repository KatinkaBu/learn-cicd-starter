// go
package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_MissingHeader(t *testing.T) {
	h := http.Header{} // don’t set Authorization
	key, err := GetAPIKey(h)

	if key != "" {
		t.Fatalf("expected empty key, got %q", key)
	}
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestGetAPIKey_Malformed_WrongScheme(t *testing.T) {
	h := http.Header{}                   // create headers map
	h.Set("Authorization", "Bearer abc") // set the header value here

	key, err := GetAPIKey(h) // call the function under test

	if key != "" {
		t.Fatalf("expected empty key, got %q", key)
	}
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetAPIKey_Malformed_MissingToken(t *testing.T) {
	h := http.Header{}
	// TODO: set Authorization to "ApiKey" (no token)
	key, err := GetAPIKey(h)

	if key != "" {
		t.Fatalf("expected empty key, got %q", key)
	}
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetAPIKey_Valid(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey secret123")

	key, err := GetAPIKey(h)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key != "secret123" {
		t.Fatalf("expected key %q, got %q", "secret123", key)
	}
}
