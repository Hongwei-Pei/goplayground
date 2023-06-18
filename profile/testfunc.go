package profile

import "fmt"

func AULF() {
	type T struct {
		Name string
	}
	for i := 0; i < 10000; i++ {
		n := fmt.Sprintf("test: %d", i)
		t := T{Name: n}
		fmt.Println(t.Name)
	}
}
