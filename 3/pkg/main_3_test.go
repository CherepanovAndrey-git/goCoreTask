package pkg

import (
	"testing"
)

func TestNewMap(t *testing.T) {
	t.Run("create new map", func(t *testing.T) {
		m := NewMap()
		if m.data == nil {
			t.Fatal("Map should be initialized")
		}
		if len(m.data) != 0 {
			t.Fatal("New map should be empty")
		}
	})
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value int
	}{
		{"add first element", "test", 42},
		{"add with empty key", "", 0},
		{"overwrite existing", "test", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMap()
			m.Add(tt.key, tt.value)

			if val, ok := m.Get(tt.key); !ok || val != tt.value {
				t.Errorf("Expected (%d, true), got (%d, %v)", tt.value, val, ok)
			}
		})
	}
}

func TestGet(t *testing.T) {
	m := NewMap()
	m.Add("exists", 123)

	tests := []struct {
		name     string
		key      string
		expected int
		exists   bool
	}{
		{"existing key", "exists", 123, true},
		{"non-existent key", "missing", 0, false},
		{"empty key", "", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := m.Get(tt.key)
			if val != tt.expected || ok != tt.exists {
				t.Errorf("Expected (%d, %v), got (%d, %v)", tt.expected, tt.exists, val, ok)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	m := NewMap()
	m.Add("toRemove", 456)
	m.Add("keep", 789)

	t.Run("remove existing key", func(t *testing.T) {
		m.Remove("toRemove")
		if _, ok := m.Get("toRemove"); ok {
			t.Error("Key should be removed")
		}
		if _, ok := m.Get("keep"); !ok {
			t.Error("Other keys should remain")
		}
	})

	t.Run("remove non-existent key", func(t *testing.T) {
		initialLen := len(m.data)
		m.Remove("missing")
		if len(m.data) != initialLen {
			t.Error("Map length should not change")
		}
	})
}

func TestExist(t *testing.T) {
	m := NewMap()
	m.Add("check", 999)

	tests := []struct {
		name   string
		key    string
		exists bool
	}{
		{"existing key", "check", true},
		{"non-existent key", "unknown", false},
		{"empty key", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exists := m.Exist(tt.key); exists != tt.exists {
				t.Errorf("Expected %v, got %v", tt.exists, exists)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	t.Run("copy full map", func(t *testing.T) {
		m := NewMap()
		m.Add("one", 1)
		m.Add("two", 2)

		copyMap := m.MapCopy()

		if len(copyMap) != len(m.data) {
			t.Fatal("Copied map size mismatch")
		}

		for k, v := range m.data {
			if copyMap[k] != v {
				t.Errorf("Value mismatch for key %s", k)
			}
		}
	})

	t.Run("modify copy doesn't affect original", func(t *testing.T) {
		m := NewMap()
		m.Add("original", 5)

		copyMap := m.MapCopy()
		copyMap["original"] = 10
		copyMap["new"] = 20

		if val, _ := m.Get("original"); val != 5 {
			t.Error("Original map should not be modified")
		}
		if m.Exist("new") {
			t.Error("Original map should not have new key")
		}
	})
}

func TestEdgeCases(t *testing.T) {
	t.Run("nil map operations", func(t *testing.T) {
		var m *StringIntMap = &StringIntMap{} // data == nil

		t.Run("get from nil map", func(t *testing.T) {
			if val, ok := m.Get("any"); val != 0 || ok != false {
				t.Error("Should return (0, false) for nil map")
			}
		})

		t.Run("exists in nil map", func(t *testing.T) {
			if m.Exist("any") {
				t.Error("Should return false for nil map")
			}
		})

		t.Run("add to nil map", func(t *testing.T) {
			m.Add("nil", 123)
			if m.data == nil {
				t.Fatal("Map should be initialized after Add")
			}
			if val, _ := m.Get("nil"); val != 123 {
				t.Error("Should store value in initialized map")
			}
		})
	})
}
