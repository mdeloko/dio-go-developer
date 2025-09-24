package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var last, antepenult, current int
	return func() int{
		
		switch current{
			case 0:
				current += 1
				return 0
			case 1:
				if last != 0{
					current += last
					antepenult = last
					return last
				}else{
					last = current
					return current
				}
			default:
				current = last + antepenult
				antepenult = last
				last = current
				return current
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

