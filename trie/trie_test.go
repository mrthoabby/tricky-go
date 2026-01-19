package trie

import "testing"

func TestTrie(t *testing.T) {
	tr := New[int]()

	tr.Insert("apple", 1)
	if val, ok := tr.Search("apple"); !ok || val != 1 {
		t.Error("Search should find 'apple' with value 1")
	}
	if _, ok := tr.Search("app"); ok {
		t.Error("Search should not find 'app' yet")
	}
	if !tr.StartsWith("app") {
		t.Error("StartsWith should find 'app'")
	}

	tr.Insert("app", 2)
	if val, ok := tr.Search("app"); !ok || val != 2 {
		t.Error("Search should find 'app' now with value 2")
	}

	// Test Remove
	tr.Remove("apple")
	if _, ok := tr.Search("apple"); ok {
		t.Error("Search should not find 'apple' after removal")
	}
	if val, ok := tr.Search("app"); !ok || val != 2 {
		t.Error("Search should still find 'app' with value 2")
	}

	tr.Remove("app")
	if _, ok := tr.Search("app"); ok {
		t.Error("Search should not find 'app' after removal")
	}
}

func TestRemoveUnicodeWord(t *testing.T) {
	instance := New[bool]()
	instance.Insert("hola", true)
	instance.Insert("holístico", true)
	instance.Insert("niño", true)
	instance.Insert("niña", true)

	removed := instance.Remove("holístico")
	if !removed {
		t.Fatalf("expected Remove to return true for existing word")
	}
	if _, ok := instance.Search("holístico"); ok {
		t.Fatalf("expected removed unicode word to be absent")
	}
	if !instance.StartsWith("hol") {
		t.Fatalf("expected prefix to remain after removal")
	}

	if _, ok := instance.Search("niño"); !ok {
		t.Fatalf("expected to find the word 'niño'")
	}
	if _, ok := instance.Search("niña"); !ok {
		t.Fatalf("expected to find the word 'niña'")
	}
	if !instance.StartsWith("niñ") {
		t.Fatalf("expected prefix 'niñ' to be recognized")
	}

	instance.Remove("niño")
	if _, ok := instance.Search("niño"); ok {
		t.Fatalf("expected removed word 'niño' to be absent")
	}
	if _, ok := instance.Search("niña"); !ok {
		t.Fatalf("expected 'niña' to remain after removing 'niño'")
	}
}

func TestRemovePrefixKeepsLongerWord(t *testing.T) {
	instance := New[string]()
	instance.Insert("go", "lang")
	instance.Insert("golang", "google")

	removed := instance.Remove("go")
	if !removed {
		t.Fatalf("expected Remove to return true for existing word")
	}
	if _, ok := instance.Search("go"); ok {
		t.Fatalf("expected removed prefix word to be absent")
	}
	if _, ok := instance.Search("golang"); !ok {
		t.Fatalf("expected longer word to remain")
	}
}

func TestTrie_All_IteratesKeyValues(t *testing.T) {
	tr := New[int]()
	tr.Insert("a", 1)
	tr.Insert("b", 2)

	got := make(map[string]int)
	for k, v := range tr.All() {
		got[k] = v
	}

	if len(got) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(got))
	}
	if got["a"] != 1 {
		t.Errorf("Expected value 1 for key 'a', got %d", got["a"])
	}
	if got["b"] != 2 {
		t.Errorf("Expected value 2 for key 'b', got %d", got["b"])
	}
}
