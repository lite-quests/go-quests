package maps

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the maps Quest 🎉")
	}
	os.Exit(code)
}

func TestCacheOperations(t *testing.T) {
	c := NewCache()
	c.Set("a", 10)
	c.Set("b", 20)
	c.Set("c", 5)

	testsGet := []struct {
		key      string
		expected int
		ok       bool
	}{
		{"a", 10, true},
		{"b", 20, true},
		{"c", 5, true},
		{"d", 0, false},
	}

	for _, tt := range testsGet {
		got, ok := c.Get(tt.key)
		if got != tt.expected || ok != tt.ok {
			t.Fatalf("Get(%q) = (%v, %v); expected (%v, %v)", tt.key, got, ok, tt.expected, tt.ok)
		}
	}
	if c.Count() != 3 {
		t.Fatalf("Count() = %v; expected 3", c.Count())
	}

	keys := c.AllKeys()
	expectedKeys := map[string]bool{"a": true, "b": true, "c": true}
	if len(keys) != len(expectedKeys) {
		t.Fatalf("AllKeys() = %v; expected length %v", keys, len(expectedKeys))
	}
	for _, k := range keys {
		if !expectedKeys[k] {
			t.Fatalf("AllKeys() contains unexpected key %q", k)
		}
	}
	c.Delete("b")
	if _, ok := c.Get("b"); ok {
		t.Fatalf("Delete('b') failed; key still exists")
	}
	if c.Count() != 2 {
		t.Fatalf("Count() after Delete = %v; expected 2", c.Count())
	}
	c.RemoveBelow(10)
	if _, ok := c.Get("c"); ok {
		t.Fatalf("RemoveBelow(10) failed; key 'c' still exists")
	}
	if _, ok := c.Get("a"); !ok {
		t.Fatalf("RemoveBelow(10) removed key 'a' incorrectly")
	}
	if c.Count() != 1 {
		t.Fatalf("Count() after RemoveBelow = %v; expected 1", c.Count())
	}
	empty := NewCache()
	if empty.Count() != 0 {
		t.Fatalf("NewCache() should start empty; got Count() = %v", empty.Count())
	}
	if len(empty.AllKeys()) != 0 {
		t.Fatalf("NewCache() AllKeys() should be empty; got %v", empty.AllKeys())
	}
	if val, ok := empty.Get("x"); ok || val != 0 {
		t.Fatalf("Get() on empty cache should return (0, false); got (%v, %v)", val, ok)
	}
}
