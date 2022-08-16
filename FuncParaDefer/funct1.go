package main

import "fmt"

func do(m1 *map[int]int) {
	(*m1)[3] = 0
	*m1 = make(map[int]int)
	(*m1)[4] = 4
	fmt.Println("m1", *m1)

}

func main() {
	m := map[int]int{4: 1, 8: 2, 12: 3}
	fmt.Println("m", m)
	do(&m)
	fmt.Println("m", m)

}
