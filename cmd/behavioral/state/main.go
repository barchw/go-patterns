package main

import (
	"fmt"

	"claude-test/internal/behavioral/state"
)

func main() {
	fmt.Println("=== State Pattern ===")

	vm := state.NewVendingMachine()

	fmt.Println("Attempting operations without coin:")
	fmt.Printf("  SelectProduct: %s\n", vm.SelectProduct("Cola"))
	fmt.Printf("  Dispense:      %s\n", vm.Dispense())

	fmt.Println("Normal purchase flow:")
	fmt.Printf("  InsertCoin:    %s\n", vm.InsertCoin())
	fmt.Printf("  InsertCoin:    %s\n", vm.InsertCoin())
	fmt.Printf("  SelectProduct: %s\n", vm.SelectProduct("Cola"))
	fmt.Printf("  Dispense:      %s\n", vm.Dispense())

	fmt.Println("Second purchase:")
	fmt.Printf("  InsertCoin:    %s\n", vm.InsertCoin())
	fmt.Printf("  SelectProduct: %s\n", vm.SelectProduct("Water"))
	fmt.Printf("  Dispense:      %s\n", vm.Dispense())
}
