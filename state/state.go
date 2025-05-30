package state

import "fmt"

// State 状态接口
type State interface {
	InsertCoin(vm *VendingMachine)
	SelectProduct(vm *VendingMachine)
	DispenseProduct(vm *VendingMachine)
}

// VendingMachine 自动售货机
type VendingMachine struct {
	state       State
	hasProduct  bool
	productName string
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		hasProduct:  true,
		productName: "Coke",
	}
	vm.state = &IdleState{}
	return vm
}

func (vm *VendingMachine) SetState(state State) {
	vm.state = state
}

func (vm *VendingMachine) InsertCoin() {
	vm.state.InsertCoin(vm)
}

func (vm *VendingMachine) SelectProduct() {
	vm.state.SelectProduct(vm)
}

func (vm *VendingMachine) DispenseProduct() {
	vm.state.DispenseProduct(vm)
}

func (vm *VendingMachine) GetProductName() string {
	return vm.productName
}

func (vm *VendingMachine) HasProduct() bool {
	return vm.hasProduct
}

func (vm *VendingMachine) ReleaseProduct() {
	if vm.hasProduct {
		fmt.Printf("Dispensing %s\n", vm.productName)
		vm.hasProduct = false
	}
}

// IdleState 空闲状态
type IdleState struct{}

func (is *IdleState) InsertCoin(vm *VendingMachine) {
	fmt.Println("Coin inserted")
	vm.SetState(&HasCoinState{})
}

func (is *IdleState) SelectProduct(vm *VendingMachine) {
	fmt.Println("Please insert coin first")
}

func (is *IdleState) DispenseProduct(vm *VendingMachine) {
	fmt.Println("Please insert coin first")
}

// HasCoinState 有硬币状态
type HasCoinState struct{}

func (hcs *HasCoinState) InsertCoin(vm *VendingMachine) {
	fmt.Println("Coin already inserted")
}

func (hcs *HasCoinState) SelectProduct(vm *VendingMachine) {
	fmt.Println("Product selected")
	if vm.HasProduct() {
		vm.SetState(&DispensingState{})
	} else {
		fmt.Println("Product out of stock")
		vm.SetState(&IdleState{})
	}
}

func (hcs *HasCoinState) DispenseProduct(vm *VendingMachine) {
	fmt.Println("Please select product first")
}

// DispensingState 出货状态
type DispensingState struct{}

func (ds *DispensingState) InsertCoin(vm *VendingMachine) {
	fmt.Println("Please wait, dispensing product")
}

func (ds *DispensingState) SelectProduct(vm *VendingMachine) {
	fmt.Println("Please wait, dispensing product")
}

func (ds *DispensingState) DispenseProduct(vm *VendingMachine) {
	vm.ReleaseProduct()
	vm.SetState(&IdleState{})
}