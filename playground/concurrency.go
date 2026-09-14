package main

import (
	"fmt"
	"sync"
)

/*
*
I don't yet understand concurrency :/
*/
func main() {
	var wg sync.WaitGroup
	names := []string{"Lorraine", "Dennis", "Bundi", "Mithamo"}

	for _, name := range append(names, names...) {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			fmt.Println(capitalizeNameManually(name))
		}(name)
	}

	wg.Wait()

	fmt.Println(names)
}

func capitalizeNameManually(name string) string {
	capitalized := ""
	for i := 0; i < len(name); i++ {
		if name[i] > 'Z' {
			capitalized += string(name[i] + 'A' - 'a')
		} else {
			capitalized += string(name[i])
		}
	}
	return capitalized
}
