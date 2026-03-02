package http_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func waitForServer(url string) error {
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("server did not start in time at %s", url)
}

func TestStartCatalogServer(t *testing.T) {
	port, err := getFreePort()
	if err != nil {
		t.Fatalf("failed to find free port: %v", err)
	}

	go func() {
		_ = StartCatalogServer(port)
	}()

	url := fmt.Sprintf("http://localhost:%d/products", port)
	if err := waitForServer(url); err != nil {
		t.Fatalf("StartCatalogServer failed to start: %v", err)
	}

	t.Run("post_product", func(t *testing.T) {
		prod := Product{ID: "p1", Name: "Laptop", Price: 999.99}
		b, _ := json.Marshal(prod)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
		if err != nil {
			t.Fatalf("POST failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 201 Created or 200 OK, got %v", resp.StatusCode)
		}
	})

	t.Run("get_products", func(t *testing.T) {
		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("GET failed: %v", err)
		}
		defer resp.Body.Close()

		var prods []Product
		if err := json.NewDecoder(resp.Body).Decode(&prods); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		found := false
		for _, p := range prods {
			if p.ID == "p1" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected to find product p1 in list")
		}
	})

	t.Run("get_product_by_id", func(t *testing.T) {
		resp, err := http.Get(url + "/p1")
		if err != nil {
			t.Fatalf("GET /p1 failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200 OK, got %v", resp.StatusCode)
			return
		}

		var prod Product
		if err := json.NewDecoder(resp.Body).Decode(&prod); err != nil {
			t.Fatalf("failed to decode: %v", err)
		}
		if prod.Name != "Laptop" {
			t.Errorf("expected name 'Laptop', got %q", prod.Name)
		}
	})

	t.Run("get_product_not_found", func(t *testing.T) {
		resp, err := http.Get(url + "/invalid-id")
		if err != nil {
			t.Fatalf("GET /invalid-id failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404 Not Found, got %v", resp.StatusCode)
		}
	})
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the http_server Quest 🎉", colorReset)
	}
	os.Exit(code)
}
