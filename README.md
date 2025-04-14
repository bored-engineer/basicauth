# basicauth [![Go Reference](https://pkg.go.dev/badge/github.com/bored-engineer/basicauth.svg)](https://pkg.go.dev/github.com/bored-engineer/basicauth)
Provides a simple [http.RoundTripper](https://pkg.go.dev/net/http#RoundTripper) implementation for injecting HTTP Basic Authentication to each request:

```go
package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bored-engineer/basicauth"
)

func main() {
	transport := &basicauth.Transport{
		Username: "bored-engineer",
		Password: "hunter2",
	}

	resp, err := (&http.Client{
		Transport: transport,
	}).Get("https://httpbin.org/basic-auth/bored-engineer/hunter2")
	if err != nil {
		log.Fatalf("(*http.Client).Do failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("(*http.Response).Body.Read failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("http.StatusCode(%d) != http.StatusOK", resp.StatusCode)
	}
	if _, err := os.Stdout.Write(body); err != nil {
		log.Fatalf("os.Stdout.Write failed: %v", err)
	}
}
```
