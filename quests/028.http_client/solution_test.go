package http_client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInventorySyncClient(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		prods := []Product{{ID: "1", Name: "Shoes", Price: 50.0}}
		json.NewEncoder(w).Encode(prods)
	})

	mux.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "1" {
			json.NewEncoder(w).Encode(Product{ID: "1", Name: "Shoes", Price: 50.0})
			return
		}
		http.Error(w, "not found", http.StatusNotFound)
	})

	mux.HandleFunc("POST /products", func(w http.ResponseWriter, r *http.Request) {
		var p Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if p.ID == "new" {
			w.WriteHeader(http.StatusCreated)
			return
		}
		http.Error(w, "conflict", http.StatusConflict)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	t.Run("FetchProducts", func(t *testing.T) {
		prods, err := FetchProducts(ts.URL)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(prods) != 1 || prods[0].ID != "1" {
			t.Errorf("expected 1 product with ID '1', got %+v", prods)
		}
	})

	t.Run("CreateProduct_Success", func(t *testing.T) {
		err := CreateProduct(ts.URL, Product{ID: "new", Name: "Hat", Price: 10.0})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("CreateProduct_Failure", func(t *testing.T) {
		err := CreateProduct(ts.URL, Product{ID: "conflict"})
		if err == nil {
			t.Errorf("expected error for bad POST, got nil")
		}
	})

	t.Run("FetchProduct_Found", func(t *testing.T) {
		p, found, err := FetchProduct(ts.URL, "1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !found {
			t.Errorf("expected found=true")
		}
		if p.Name != "Shoes" {
			t.Errorf("expected Shoes, got %s", p.Name)
		}
	})

	t.Run("FetchProduct_NotFound", func(t *testing.T) {
		_, found, err := FetchProduct(ts.URL, "999")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if found {
			t.Errorf("expected found=false")
		}
	})

	t.Run("FetchMultipleProducts", func(t *testing.T) {
		ids := []string{"1", "999", "1", "1"}
		results, err := FetchMultipleProducts(ts.URL, ids)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(results) != 1 {
			t.Errorf("expected 1 product in map, got %d", len(results))
		}
		if p, ok := results["1"]; !ok || p.Name != "Shoes" {
			t.Errorf("expected to find Shoes with ID 1 in results")
		}
	})
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the http_client Quest 🎉", colorReset)
	}
	os.Exit(code)
}
