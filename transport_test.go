package basicauth

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestTransport(t *testing.T) {
	transport := &Transport{
		Username: "bored-engineer",
		Password: "hunter2",
	}

	resp, err := (&http.Client{
		Transport: transport,
	}).Get("https://httpbin.org/basic-auth/bored-engineer/hunter2")
	if err != nil {
		t.Fatalf("(*http.Client).Do failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("(*http.Response).Body.Read failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("http.StatusCode(%d) != http.StatusOK", resp.StatusCode)
	}

	var parsed struct {
		Authenticated bool   `json:"authenticated"`
		User          string `json:"user"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if !parsed.Authenticated {
		t.Fatalf("parsed.Authenticated != true")
	}
	if parsed.User != "bored-engineer" {
		t.Fatalf("parsed.User != \"bored-engineer\"")
	}
}
