package state

import "testing"

func TestVendingMachine_FullFlow(t *testing.T) {
	vm := NewVendingMachine()

	msg := vm.InsertCoin()
	if msg != "Coin inserted" {
		t.Errorf("InsertCoin: got %q, want %q", msg, "Coin inserted")
	}

	msg = vm.SelectProduct("Cola")
	if msg != "Product selected: Cola" {
		t.Errorf("SelectProduct: got %q, want %q", msg, "Product selected: Cola")
	}

	msg = vm.Dispense()
	if msg != "Dispensing: Cola" {
		t.Errorf("Dispense: got %q, want %q", msg, "Dispensing: Cola")
	}
}

func TestVendingMachine_IdleState(t *testing.T) {
	vm := NewVendingMachine()

	msg := vm.SelectProduct("Water")
	if msg != "Please insert a coin first" {
		t.Errorf("SelectProduct in idle: got %q", msg)
	}

	msg = vm.Dispense()
	if msg != "Please insert a coin and select a product" {
		t.Errorf("Dispense in idle: got %q", msg)
	}
}

func TestVendingMachine_HasCoinState(t *testing.T) {
	vm := NewVendingMachine()
	vm.InsertCoin()

	msg := vm.InsertCoin()
	if msg != "Coin already inserted" {
		t.Errorf("double InsertCoin: got %q", msg)
	}

	msg = vm.Dispense()
	if msg != "Please select a product first" {
		t.Errorf("Dispense without select: got %q", msg)
	}
}

func TestVendingMachine_DispensingState(t *testing.T) {
	vm := NewVendingMachine()
	vm.InsertCoin()
	vm.SelectProduct("Chips")

	msg := vm.InsertCoin()
	if msg != "Please wait, dispensing in progress" {
		t.Errorf("InsertCoin while dispensing: got %q", msg)
	}

	msg = vm.SelectProduct("Water")
	if msg != "Already dispensing, please wait" {
		t.Errorf("SelectProduct while dispensing: got %q", msg)
	}
}

func TestVendingMachine_StateTransitions(t *testing.T) {
	vm := NewVendingMachine()

	if _, ok := vm.CurrentState().(*IdleState); !ok {
		t.Error("initial state should be IdleState")
	}

	vm.InsertCoin()
	if _, ok := vm.CurrentState().(*HasCoinState); !ok {
		t.Error("after InsertCoin should be HasCoinState")
	}

	vm.SelectProduct("Cola")
	if _, ok := vm.CurrentState().(*DispensingState); !ok {
		t.Error("after SelectProduct should be DispensingState")
	}

	vm.Dispense()
	if _, ok := vm.CurrentState().(*IdleState); !ok {
		t.Error("after Dispense should be IdleState")
	}
}

func TestVendingMachine_MultipleCycles(t *testing.T) {
	vm := NewVendingMachine()

	for _, product := range []string{"Cola", "Water", "Chips"} {
		vm.InsertCoin()
		vm.SelectProduct(product)
		msg := vm.Dispense()
		expected := "Dispensing: " + product
		if msg != expected {
			t.Errorf("got %q, want %q", msg, expected)
		}
	}
}
