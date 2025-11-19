package main

import (
	"slices"

	"golang.org/x/exp/constraints"
)

/*
 In a scratch file, implement these from scratch (no scaffolding):

  // 1. Generic Min/Max
  func Min[T constraints.Ordered](a, b T) T
  func Max[T constraints.Ordered](a, b T) T

  // 2. Generic Contains
  func Contains[T comparable](slice []T, value T) bool

  // 3. Generic Map
  func Map[T, U any](items []T, fn func(T) U) []U

  // 4. Generic Filter
  func Filter[T any](items []T, predicate func(T) bool) []T

  // 5. Generic Set (struct)
  type Set[T comparable] struct { design this  }
  func (s *Set[T]) Add(value T)
  func (s *Set[T]) Contains(value T) bool
*/

func Min[T constraints.Ordered](a, b T) T {
	return min(a, b)
}

func Max[T constraints.Ordered](a, b T) T {
	return max(a, b)
}

func Contains[T comparable](slice []T, value T) bool {
	return slices.Contains(slice, value)
}

func Map[T, U any](items []T, fn func(T) U) []U {

	result := []U{}

	for i := range items {
		result = append(result, fn(items[i]))
	}
	return result
}

func Filter[T any](items []T, predicate func(T) bool) []T {

	result := []T{}

	for _, v := range items {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

type Set[T comparable] struct {
	Elements map[T]struct{}
}

func (s *Set[T]) Add(value T) {
	s.Elements[value] = struct{}{}
}

func (s *Set[T]) Contains(value T) bool {

	if _, ok := s.Elements[value]; ok {
		return true
	}

	return false
}

func main() {

}
