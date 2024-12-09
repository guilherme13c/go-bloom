package bloom

import (
	"testing"
)

type TestStruct struct {
	ID   int
	Name string
}

type Color int
const (
	Red Color = iota
	Green
	Blue
)

func TestInsertAndFindStruct(t *testing.T) {
	bf := NewBloomFilter(200, 3)

	item := TestStruct{ID: 1, Name: "Test"}

	bf.Insert(item)
	if !bf.Find(item) {
		t.Errorf("O item %v deveria estar no filtro", item)
	}

	if bf.Find(TestStruct{ID: 2, Name: "Other"}) {
		t.Errorf("O item %v não deveria estar no filtro", TestStruct{ID: 2, Name: "Other"})
	}
}

func TestInsertAndFindEnum(t *testing.T) {
	bf := NewBloomFilter(200, 3)

	bf.Insert(Red)
	bf.Insert(Green)

	if !bf.Find(Red) {
		t.Errorf("A cor Red deveria estar no filtro")
	}

	if !bf.Find(Green) {
		t.Errorf("A cor Green deveria estar no filtro")
	}

	if bf.Find(Blue) {
		t.Errorf("A cor Blue não deveria estar no filtro")
	}
}

func TestInsertAndFindSlice(t *testing.T) {
	bf := NewBloomFilter(200, 3)

	item := []int{1, 2, 3}

	bf.Insert(item)
	if !bf.Find(item) {
		t.Errorf("O slice %v deveria estar no filtro", item)
	}

	if bf.Find([]int{4, 5, 6}) {
		t.Errorf("O slice %v não deveria estar no filtro", []int{4, 5, 6})
	}
}

func TestInsertMultipleStructs(t *testing.T) {
	bf := NewBloomFilter(200, 3)

	items := []TestStruct{
		{ID: 1, Name: "Alpha"},
		{ID: 2, Name: "Beta"},
		{ID: 3, Name: "Gamma"},
	}

	for _, item := range items {
		bf.Insert(item)
	}

	for _, item := range items {
		if !bf.Find(item) {
			t.Errorf("O item %v deveria estar no filtro", item)
		}
	}

	if bf.Find(TestStruct{ID: 4, Name: "Delta"}) {
		t.Errorf("O item Delta não deveria estar no filtro")
	}
}

func TestInsertAndFindEmptySlice(t *testing.T) {
	bf := NewBloomFilter(200, 3)

	item := []int{}

	bf.Insert(item)
	if !bf.Find(item) {
		t.Errorf("O slice vazio deveria estar no filtro")
	}

	if bf.Find([]int{0}) {
		t.Errorf("O slice %v não deveria estar no filtro", []int{0})
	}
}

func TestInsertAndFindStrings(t *testing.T) {
	bf := NewBloomFilter(200, 3)

	items := []string{"foo", "bar", "baz"}

	for _, item := range items {
		bf.Insert(item)
	}

	for _, item := range items {
		if !bf.Find(item) {
			t.Errorf("A string %s deveria estar no filtro", item)
		}
	}

	if bf.Find("qux") {
		t.Errorf("A string 'qux' não deveria estar no filtro")
	}
}

func TestInsertAndFindIntegers(t *testing.T) {
	bf := NewBloomFilter(200, 3)

	items := []int{42, 100, 256}

	for _, item := range items {
		bf.Insert(item)
	}

	for _, item := range items {
		if !bf.Find(item) {
			t.Errorf("O número %d deveria estar no filtro", item)
		}
	}

	if bf.Find(512) {
		t.Errorf("O número 512 não deveria estar no filtro")
	}
}
