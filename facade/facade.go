package facade

// Subsystem components
type CPU struct{}

func (c *CPU) Start() string {
	return "CPU started"
}

func (c *CPU) Execute() string {
	return "CPU executing"
}

func (c *CPU) Stop() string {
	return "CPU stopped"
}

type Memory struct{}

func (m *Memory) Load() string {
	return "Memory loaded"
}

func (m *Memory) Free() string {
	return "Memory freed"
}

type HardDrive struct{}

func (hd *HardDrive) Read() string {
	return "Hard drive reading data"
}

func (hd *HardDrive) Write() string {
	return "Hard drive writing data"
}

// Facade
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (cf *ComputerFacade) StartComputer() []string {
	var results []string
	results = append(results, cf.cpu.Start())
	results = append(results, cf.memory.Load())
	results = append(results, cf.hardDrive.Read())
	results = append(results, cf.cpu.Execute())
	return results
}

func (cf *ComputerFacade) ShutdownComputer() []string {
	var results []string
	results = append(results, cf.cpu.Stop())
	results = append(results, cf.memory.Free())
	results = append(results, cf.hardDrive.Write())
	return results
}

func (cf *ComputerFacade) RestartComputer() []string {
	var results []string
	results = append(results, "Restarting computer...")
	shutdownResults := cf.ShutdownComputer()
	startupResults := cf.StartComputer()
	
	results = append(results, shutdownResults...)
	results = append(results, startupResults...)
	return results
}