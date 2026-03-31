package json

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the JSON Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestMarshalProduct(t *testing.T) {
	p := Product{
		ID:      1,
		Name:    "Gopher Plushie",
		Price:   19.99,
		InStock: true,
		Tags:    []string{"toy", "go"},
	}

	data, err := MarshalProduct(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := string(data)
	checks := []string{`"id":1`, `"name":"Gopher Plushie"`, `"price":19.99`, `"in_stock":true`, `"tags":["toy","go"]`}
	for _, c := range checks {
		if !strings.Contains(got, c) {
			t.Errorf("expected JSON to contain %q, got: %s", c, got)
		}
	}
}

func TestMarshalProduct_OmitEmptyTags(t *testing.T) {
	p := Product{
		ID:      2,
		Name:    "Go Book",
		Price:   39.99,
		InStock: false,
	}

	data, err := MarshalProduct(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := string(data)
	if strings.Contains(got, `"tags"`) {
		t.Errorf("expected 'tags' to be omitted when nil, got: %s", got)
	}
}

func TestUnmarshalProduct(t *testing.T) {
	input := `{"id":3,"name":"Keyboard","price":79.99,"in_stock":true,"tags":["hardware","input"]}`

	p, err := UnmarshalProduct([]byte(input))
if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.ID != 3 {
		t.Errorf("expected ID 3, got %d", p.ID)
	}
	if p.Name != "Keyboard" {
		t.Errorf("expected Name 'Keyboard', got %q", p.Name)
	}
	if p.Price != 79.99 {
		t.Errorf("expected Price 79.99, got %f", p.Price)
	}
	if !p.InStock {
		t.Errorf("expected InStock true, got false")
	}
	if len(p.Tags) != 2 || p.Tags[0] != "hardware" || p.Tags[1] != "input" {
		t.Errorf("unexpected Tags: %v", p.Tags)
	}
}

func TestUnmarshalProduct_InvalidJSON(t *testing.T) {
	_, err := UnmarshalProduct([]byte(`not-json`))
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestMarshalProducts(t *testing.T) {
	products := []Product{
		{ID: 1, Name: "Mouse", Price: 29.99, InStock: true},
		{ID: 2, Name: "Monitor", Price: 299.99, InStock: false},
	}

	data, err := MarshalProducts(products)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got :=string(data)
	if !strings.HasPrefix(got, "[") || !strings.HasSuffix(got, "]") {
		t.Errorf("expected a JSON array, got: %s", got)
	}
	if !strings.Contains(got, `"name":"Mouse"`) {
		t.Errorf("expected JSON to contain Mouse, got: %s", got)
	}
	if !strings.Contains(got, `"name":"Monitor"`) {
		t.Errorf("expected JSON to contain Monitor, got: %s", got)
	}
}

func TestUnmarshalProducts(t *testing.T) {
	input := `[{"id":1,"name":"Mouse","price":29.99,"in_stock":true},{"id":2,"name":"Monitor","price":299.99,"in_stock":false}]`

	products, err := UnmarshalProducts([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(products) != 2 {
		t.Fatalf("expected 2 products, got %d", len(products))
	}
	if products[0].Name != "Mouse" {
		t.Errorf("expected first product 'Mouse', got %q", products[0].Name)
	}
	if products[1].Name != "Monitor" {
		t.Errorf("expected second product 'Monitor', got %q", products[1].Name)
	}
}

func TestPrettyMarshal(t *testing.T) {
	p := Product{
		ID:      1,
		Name:    "Gopher Plushie",
		Price:   19.99,
		InStock:true,
	}

	data, err := PrettyMarshal(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := string(data)
	if !strings.Contains(got, "\t") {
		t.Errorf("expected indented JSON with tabs, got: %s", got)
	}
	if !strings.Contains(got, "\n") {
		t.Errorf("expected multi-line JSON, got: %s", got)
	}
}
