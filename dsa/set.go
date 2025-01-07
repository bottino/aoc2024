package dsa

import (
	"fmt"
	"slices"
	"strings"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values ...T) Set[T] {
	s := make(Set[T], len(values))
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s *Set[T]) Add(v T) {
	(*s)[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	if _, ok := (*s)[v]; ok {
		delete(*s, v)
	}
}

func (a Set[T]) Equal(b Set[T]) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}

	return true
}

func (a Set[T]) Union(b Set[T]) Set[T] {
	c := NewSet[T]()
	for k := range a {
		c.Add(k)
	}
	for k := range b {
		c.Add(k)
	}

	return c
}

func (a Set[T]) Intersection(b Set[T]) Set[T] {
	c := NewSet[T]()
	for k := range a {
		if _, ok := b[k]; ok {
			c.Add(k)
		}
	}

	return c
}

func (s Set[T]) String() string {
	var setStr []string
	for k := range s {
		setStr = append(setStr, fmt.Sprintf("%v", k))
	}
	slices.Sort(setStr)
	return fmt.Sprintf("[%s]", strings.Join(setStr, " "))
}
