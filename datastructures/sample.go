package datastructures

import (
	"fmt"
	"strings"
)

type SampleDS[T int | float32 | float64 | string] struct {
	size  uint32
	items []T
}

func (s *SampleDS[T]) AddItem(item T) {
	s.items = append(s.items, item)
	s.size++
}

func (s *SampleDS[T]) Size() uint32 {
	return s.size
}

func (s *SampleDS[T]) PrintItems() string {
	sb := strings.Builder{}
	for i := 0; i < len(s.items); i++ {
		item := s.items[i]
		sb.WriteString(fmt.Sprintf("%v", item))
	}

	return sb.String()
}

func (s *SampleDS[T]) Clear() {
	s.size = 0
	s.items = nil
}
