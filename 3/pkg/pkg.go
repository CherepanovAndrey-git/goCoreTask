package pkg

type StringIntMap struct {
	data map[string]int
}

// NewMap возвращает ссылку на экземпляр структуры StringIntMap
func NewMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Add добавляет ключ и значение в структуру
func (m *StringIntMap) Add(key string, value int) {
	if m.data == nil {
		m.data = make(map[string]int)
	}
	m.data[key] = value
}

// Get возвращает значение по ключу и
func (m *StringIntMap) Get(key string) (int, bool) {
	if m.data == nil {
		return 0, false
	}
	value, ok := m.data[key]
	return value, ok
}

// Remove удаляет ключ и значение по ключу
func (m *StringIntMap) Remove(key string) {
	if m.data == nil {
		return
	}
	delete(m.data, key)
}

// Exist проверяет, существует ли ключ в карте
func (m *StringIntMap) Exist(key string) bool {
	if m.data == nil {
		return false
	}
	_, exists := m.data[key]
	return exists
}

// MapCopy копирует и возвращает карту
func (m *StringIntMap) MapCopy() map[string]int {
	mapCopy := make(map[string]int)
	if m.data != nil {
		for k, v := range m.data {
			mapCopy[k] = v
		}
	}
	return mapCopy
}
