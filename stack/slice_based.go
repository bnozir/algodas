package stack

type SliceBased struct {
	items []int
}

func NewSliceBased() *SliceBased {
	return &SliceBased{}
}

func (s *SliceBased) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *SliceBased) Push(i int) {
	s.items = append(s.items, i)
}

func (s *SliceBased) Top() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

func (s *SliceBased) Pop() (int, bool) {
	i, ok := s.Top()
	if ok {
		s.items = s.items[:len(s.items)-1]
	}

	return i, ok
}
