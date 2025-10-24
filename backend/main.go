package main

// import (
// "fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func Index[T comparable](s []T, x T) int {
// 	for i, v := range s {
// 		if v == x {
// 			return i
// 		}
// 	}

// 	return -1
// }

// type List[T comparable] struct {
// 	next *List[T]
// 	val  T
// }

// func (l *List[T]) Append(v T) {
// 	// Copy the pointer of first node of the list to the variable current
// 	current := l

// 	// Check if current is the last node in the list
// 	for current.next != nil {
// 		// swap current to next node in the list
// 		current = current.next
// 	}
// 	// attach new node to the last node in the list
// 	current.next = &List[T]{val: v}
// }

// func (l *List[T]) Length() int {
// 	// copy the pointer of first node of the list to the variable current
// 	current := l
// 	count := 0

// 	// Check if current is the last node in the list
// 	for current.next != nil {
// 		count++
// 		// swap current to next node in the list
// 		current = current.next
// 	}

// 	return count
// }

// func (l *List[T]) Find(target T) *List[T] {
// 	// copy the pointer of first node of the list to the variable current
// 	current := l

// 	// Check if current is the last node in the list
// 	for current.next != nil {
// 		// Check if target equals current value
// 		if current.val == target {
// 			return current
// 		}
// 		// swap current to next node in the list
// 		current = current.next
// 	}

// 	return nil
// }

// func (l *List[T]) Delete(target T) {
// 	// Create new dummy node and set the next value to the root node of the list
// 	dummy := &List[T]{next: l}
// 	prev := dummy
// 	current := l

// 	// Check if current is the last node in the list
// 	for current.next != nil {
// 		// Check if target equals current value
// 		if current.val == target {
// 			// Copy prev node next value to current node next value
// 			prev.next = current.next
// 		}
// 		// Swap prev and current nodes with next nodes
// 		prev, current = current, current.next
// 	}
// }

// func (l *List[T]) Treverse(f func(v T)) {
// 	current := l

// 	for current != nil {
// 		// Call function f on all node values
// 		f(current.val)
// 		current = current.next
// 	}
// }

// func TestGoroutine(c chan string) {

// 	c <- "Hello"
// 	c <- "World!"

// }

func Movedtostack(x int) int {
	return x * x
}

func Escapedtoheap() *int {
	x := 42
	return &x
}

func main() {
	// v := Vertex{3, 4}

	// fmt.Println(v.Abs())

	// si := []int{1, 2, 3, 4}
	// fmt.Println(Index(si, 4))

	// ss := []string{"hello", "hi", "how", "are", "you"}
	// fmt.Println(Index(ss, "you"))

	// c := make(chan string, 2)

	// go TestGoroutine(c)

	// msg, msg1 := <-c, <-c

	// fmt.Print(msg + " " + msg1)

	Movedtostack(2)
	Escapedtoheap()
}
