/*
Package state -- State

What is it?
State is a behavioral design pattern that allows an object to change its behavior
depending on its internal state. From the client's perspective, the object appears as if its
class has changed. Each state is encapsulated in a separate struct with its own logic.

When to use?
  - When an object has many states and its behavior differs significantly depending on the state.
  - When the state transition logic is complex and scattered across conditional statements.
  - When you want to add new states without modifying existing code.
  - When each state has its own rules -- which actions are allowed and which state they lead to.

When NOT to use?
  - When an object has only 2-3 simple states with little logic -- a flag and an if will be more readable.
  - When state transitions are trivial and rare -- the overhead of additional structs is not justified.
  - When states don't differ in behavior -- if all react the same way, the pattern is unnecessary.

Tips and pitfalls:
  - State vs Strategy: in Strategy, the client changes the algorithm; in State, transitions happen automatically.
  - Stateless states (IdleState, HasCoinState, DispensingState) can be shared (Flyweight pattern).
  - Each state must implement all interface methods, even those not allowed in that state.
  - Adding a new state only requires a new struct implementing State -- existing code doesn't need changes.
*/
package state

// State defines the state interface with vending machine actions.
type State interface {
	InsertCoin(machine *VendingMachine) string
	SelectProduct(machine *VendingMachine) string
	Dispense(machine *VendingMachine) string
}

// VendingMachine is the vending machine context that delegates actions to the current state.
type VendingMachine struct {
	state   State
	product string
}

// NewVendingMachine creates a new vending machine in the idle state (IdleState).
func NewVendingMachine() *VendingMachine {
	return &VendingMachine{state: &IdleState{}}
}

// SetState changes the current state of the machine.
func (vm *VendingMachine) SetState(s State) {
	vm.state = s
}

// CurrentState returns the current state of the machine.
func (vm *VendingMachine) CurrentState() State {
	return vm.state
}

// InsertCoin delegates the coin insertion action to the current state.
func (vm *VendingMachine) InsertCoin() string {
	return vm.state.InsertCoin(vm)
}

// SelectProduct sets the selected product and delegates the selection action to the current state.
func (vm *VendingMachine) SelectProduct(product string) string {
	vm.product = product
	return vm.state.SelectProduct(vm)
}

// Dispense delegates the product dispensing action to the current state.
func (vm *VendingMachine) Dispense() string {
	return vm.state.Dispense(vm)
}

// IdleState represents the idle state -- the machine is waiting for a coin to be inserted.
type IdleState struct{}

// InsertCoin accepts a coin and transitions to HasCoinState.
func (s *IdleState) InsertCoin(machine *VendingMachine) string {
	machine.SetState(&HasCoinState{})
	return "Coin inserted"
}

// SelectProduct informs that a coin must be inserted first.
func (s *IdleState) SelectProduct(machine *VendingMachine) string {
	return "Please insert a coin first"
}

// Dispense informs that a coin must be inserted and a product selected first.
func (s *IdleState) Dispense(machine *VendingMachine) string {
	return "Please insert a coin and select a product"
}

// HasCoinState represents the state after a coin has been inserted -- the machine is waiting for product selection.
type HasCoinState struct{}

// InsertCoin informs that a coin has already been inserted.
func (s *HasCoinState) InsertCoin(machine *VendingMachine) string {
	return "Coin already inserted"
}

// SelectProduct accepts the product selection and transitions to DispensingState.
func (s *HasCoinState) SelectProduct(machine *VendingMachine) string {
	machine.SetState(&DispensingState{})
	return "Product selected: " + machine.product
}

// Dispense informs that a product must be selected first.
func (s *HasCoinState) Dispense(machine *VendingMachine) string {
	return "Please select a product first"
}

// DispensingState represents the product dispensing state.
type DispensingState struct{}

// InsertCoin informs that dispensing is in progress.
func (s *DispensingState) InsertCoin(machine *VendingMachine) string {
	return "Please wait, dispensing in progress"
}

// SelectProduct informs that dispensing is in progress.
func (s *DispensingState) SelectProduct(machine *VendingMachine) string {
	return "Already dispensing, please wait"
}

// Dispense dispenses the product and transitions back to IdleState.
func (s *DispensingState) Dispense(machine *VendingMachine) string {
	machine.SetState(&IdleState{})
	product := machine.product
	machine.product = ""
	return "Dispensing: " + product
}
