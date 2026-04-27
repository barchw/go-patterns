package state

type State interface {
	InsertCoin(machine *VendingMachine) string
	SelectProduct(machine *VendingMachine) string
	Dispense(machine *VendingMachine) string
}

type VendingMachine struct {
	state   State
	product string
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{state: &IdleState{}}
}

func (vm *VendingMachine) SetState(s State) {
	vm.state = s
}

func (vm *VendingMachine) CurrentState() State {
	return vm.state
}

func (vm *VendingMachine) InsertCoin() string {
	return vm.state.InsertCoin(vm)
}

func (vm *VendingMachine) SelectProduct(product string) string {
	vm.product = product
	return vm.state.SelectProduct(vm)
}

func (vm *VendingMachine) Dispense() string {
	return vm.state.Dispense(vm)
}

type IdleState struct{}

func (s *IdleState) InsertCoin(machine *VendingMachine) string {
	machine.SetState(&HasCoinState{})
	return "Coin inserted"
}

func (s *IdleState) SelectProduct(machine *VendingMachine) string {
	return "Please insert a coin first"
}

func (s *IdleState) Dispense(machine *VendingMachine) string {
	return "Please insert a coin and select a product"
}

type HasCoinState struct{}

func (s *HasCoinState) InsertCoin(machine *VendingMachine) string {
	return "Coin already inserted"
}

func (s *HasCoinState) SelectProduct(machine *VendingMachine) string {
	machine.SetState(&DispensingState{})
	return "Product selected: " + machine.product
}

func (s *HasCoinState) Dispense(machine *VendingMachine) string {
	return "Please select a product first"
}

type DispensingState struct{}

func (s *DispensingState) InsertCoin(machine *VendingMachine) string {
	return "Please wait, dispensing in progress"
}

func (s *DispensingState) SelectProduct(machine *VendingMachine) string {
	return "Already dispensing, please wait"
}

func (s *DispensingState) Dispense(machine *VendingMachine) string {
	machine.SetState(&IdleState{})
	product := machine.product
	machine.product = ""
	return "Dispensing: " + product
}
