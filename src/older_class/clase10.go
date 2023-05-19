package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["josue"] = 14
	m["jose"] = 20

	fmt.Println(m)

	//recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//obtener valor
	value, ok := m["josue"]
	fmt.Println(value, ok)

}
