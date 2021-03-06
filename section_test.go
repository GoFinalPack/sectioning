package section

import "testing"

var objSlice = []*obj{
	{ID: 1, Name: "5", Age: 1},
	{ID: 2, Name: "4", Age: 2},
	{ID: 3, Name: "3", Age: 3},
	{ID: 4, Name: "2", Age: 4},
	{ID: 5, Name: "5", Age: 5},
	{ID: 1, Name: "1", Age: 6},
}

type obj struct {
	ID   int
	Name string
	Age  int
}

func TestIndex(t *testing.T) {
	b := Index(objSlice, func(e *obj) bool {
		return e.ID == 100
	})
	t.Log(b) // -1
}

func TestUnique(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Unique(intSlice, func(i, j int) bool {
		return i == j
	})
	t.Log(b)
}

func TestCount(t *testing.T) {
	b := Count(objSlice, func(e *obj) bool {
		return e.ID == 1
	})
	t.Log(b)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	c := Count(intSlice, func(i int) bool {
		return i == 7
	})
	t.Log(c)
}

func TestCut(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Cut(intSlice, func(i int) bool {
		return i == 7
	})
	t.Log(b)
}

func TestDelete(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Delete(intSlice, func(i int) bool {
		return i == 7
	})
	t.Log(b)
}

func TestFirst(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := First(intSlice)
	t.Log(b)
}

func TestIsEmpty(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := IsEmpty(intSlice)
	t.Log(b)
}

func TestPop(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Pop(intSlice)
	t.Log(b)
}

func TestPush(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Push(intSlice, 100)
	t.Log(b)
}

func TestShuffle(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Shuffle(intSlice)
	t.Log(b)

	c := Shuffle(objSlice)
	t.Log(c)
}

func TestSort(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Sort(intSlice, func(i, j int) bool {
		return i > j
	})
	t.Log(b)

	c := Sort(objSlice, func(i, j *obj) bool {
		return i.ID > j.ID
	})
	t.Log(c)
}

func TestMin(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := MinOrMax(intSlice, func(i, j int) bool {
		return i > j
	})
	t.Log(b)
}
