package main

import (
	"fmt"

	linkedlist "github.com/renatospaka/linked-list/linkedList"
)

func main() {
	list := linkedlist.List{}
	tiago := linkedlist.Pessoa{"Tiago", "Souza", 32}
	joao := linkedlist.Pessoa{"João", "Batista", 28}
	dani := linkedlist.Pessoa{"Dani", "Silva", 19}
	maria := linkedlist.Pessoa{"Maria", "José", 56}
	marcos := linkedlist.Pessoa{"Marcos", "Souza", 37}
	carla := linkedlist.Pessoa{"Carla", "Pereira", 43}

	list.Append(tiago)
	list.Append(joao)
	list.Append(dani)
	list.Append(maria)
	list.Append(marcos)
	list.Append(carla)

	list.Display()
	fmt.Println("-------------------------")

	pessoa := list.Search("Maria")
	fmt.Println(pessoa)
	fmt.Println("-------------------------")

	// list.Delete("Dani")
	// list.Display()
	// fmt.Println("-------------------------")

	list.Delete("Tiago")
	list.Display()
	fmt.Println("-------------------------")

	list.Delete("Carla")
	list.Display()
	fmt.Println("-------------------------")
}
