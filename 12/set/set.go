package set

type Set map[string]struct{}

func (s Set) Add(data string) {
	s[data] = struct{}{}
}

func (s Set) Remove(data string) {
	delete(s, data)
}

func (s Set) Has(data string) bool {
	_, ok := s[data]
	return ok
}