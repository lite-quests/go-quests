package solutions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Read README.md for the instructions

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// TODO: Implement FetchProducts
func FetchProducts(apiURL string) ([]Product, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(apiURL + "/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}

// TODO: Implement CreateProduct
func CreateProduct(apiURL string, p Product) error {
	client := &http.Client{Timeout: 5 * time.Second}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(p); err != nil {
		return err
	}

	resp, err := client.Post(
		apiURL+"/products",
		"application/json",
		&buf,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return nil
}

// TODO: Implement FetchProduct
func FetchProduct(apiURL string, id string) (Product, bool, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(apiURL + "/products/" + id)
	if err != nil {
		return Product{}, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Product{}, false, nil
	}

	if resp.StatusCode != http.StatusOK {
		return Product{}, false, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var p Product
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return Product{}, false, err
	}

	return p, true, nil
}

// TODO: Implement FetchMultipleProducts
func FetchMultipleProducts(apiURL string, ids []string) (map[string]Product, error) {
	results := make(map[string]Product)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)

		go func(productID string) {
			defer wg.Done()

			p, found, err := FetchProduct(apiURL, productID)
			if err != nil || !found {
				return
			}

			mu.Lock()
			results[productID] = p
			mu.Unlock()
		}(id)
	}

	wg.Wait()
	return results, nil
}
