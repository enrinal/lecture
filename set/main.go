package main

import "fmt"

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *set) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *set) Clear() {
	s.m = make(map[string]struct{})
}

func (s *set) Size() int {
	return len(s.m)
}

func (s *set) Values() []string {
	values := make([]string, 0, len(s.m))
	for k := range s.m {
		values = append(values, k)
	}
	return values
}

// Union The union of two sets contains all the elements contained in either set (or both sets).
// Union returns a new set which contains all the elements of s and other.
func (s *set) Union(s2 *set) *set {
	s3 := NewSet()
	for k := range s.m {
		s3.Add(k)
	}
	for k := range s2.m {
		s3.Add(k)
	}
	return s3
}

// Intersect The intersection of two sets contains only the elements that are in both sets.
// Intersect returns a new set which is the intersection of s and s2
func (s *set) Intersect(s2 *set) *set {
	s3 := NewSet()
	for k := range s.m {
		if s2.Contains(k) {
			s3.Add(k)
		}
	}
	return s3
}

// Difference The difference of two sets is all the elements in s that are not in s2.
// Difference returns a new set which is the difference of s and s2
func (s *set) Difference(s2 *set) *set {
	s3 := NewSet()
	for k := range s.m {
		if !s2.Contains(k) {
			s3.Add(k)
		}
	}
	return s3
}

// Complement The complement of a set A contains everything that is not in the set A.
// Complement returns a new set which is the complement of s and the universe
func (s *set) Complement(s2 *set) *set {
	s3 := NewSet()
	for k := range s2.m {
		if !s.Contains(k) {
			s3.Add(k)
		}
	}
	return s3
}

func main() {
	s := NewSet()

	s.Add("Peter")
	s.Add("David")

	fmt.Println(s.Contains("Peter"))  // True
	fmt.Println(s.Contains("George")) // False

	s.Remove("David")
	fmt.Println(s.Contains("David")) // False

	animal := NewSet()
	animal.Add("cat")
	animal.Add("dog")
	animal.Add("bird")
	mamalia := NewSet()
	mamalia.Add("dog")
	mamalia.Add("cat")
	mamalia.Add("cow")

	fmt.Println(animal.Intersect(mamalia).Values())
	fmt.Println(animal.Union(mamalia).Values())
	fmt.Println(animal.Difference(mamalia).Values())
	fmt.Println(animal.Complement(mamalia).Values())
}
