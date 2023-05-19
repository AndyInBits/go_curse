package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping(p *pc) {
	fmt.Println(myPC.brand, "pong")
	p.brand = "apple"
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a
	fmt.Println(a, *b)

	*b = 100
	fmt.Println(a, *b)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	// myPC.ping(&myPC)

	// myPC.ram = 20

	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)

	myPC.duplicateRAM()
	fmt.Println(myPC)

}
