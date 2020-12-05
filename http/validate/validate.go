package validate

type Map interface {
	Validate(field string) bool
	Fields() []string
}

func ValidateMap(m Map) bool {
	fields := m.Fields()
	for _, field := range fields {
		if !m.Validate(field) {
			return false
		}
	}
	return true
}

type Array interface {
	Validate(i int) bool
	Length() int
}

func ValidateArray(a Array) bool {
	l := a.Length()
	for i := 0; i < l; i++ {
		if !a.Validate(i) {
			return false
		}
	}
	return true
}
