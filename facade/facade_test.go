package facade

import "testing"

func TestFacade(t *testing.T) {
	computer := NewComputerFacade()
	
	// Test starting computer
	startResults := computer.StartComputer()
	expectedStart := []string{
		"CPU started",
		"Memory loaded",
		"Hard drive reading data",
		"CPU executing",
	}
	
	if len(startResults) != len(expectedStart) {
		t.Errorf("Expected %d start results, got %d", len(expectedStart), len(startResults))
	}
	
	for i, expected := range expectedStart {
		if startResults[i] != expected {
			t.Errorf("Expected start result %d to be '%s', got '%s'", i, expected, startResults[i])
		}
	}
	
	// Test shutting down computer
	shutdownResults := computer.ShutdownComputer()
	expectedShutdown := []string{
		"CPU stopped",
		"Memory freed",
		"Hard drive writing data",
	}
	
	if len(shutdownResults) != len(expectedShutdown) {
		t.Errorf("Expected %d shutdown results, got %d", len(expectedShutdown), len(shutdownResults))
	}
	
	for i, expected := range expectedShutdown {
		if shutdownResults[i] != expected {
			t.Errorf("Expected shutdown result %d to be '%s', got '%s'", i, expected, shutdownResults[i])
		}
	}
	
	// Test restarting computer
	restartResults := computer.RestartComputer()
	expectedRestartLength := 1 + len(expectedShutdown) + len(expectedStart) // "Restarting..." + shutdown + startup
	
	if len(restartResults) != expectedRestartLength {
		t.Errorf("Expected %d restart results, got %d", expectedRestartLength, len(restartResults))
	}
	
	if restartResults[0] != "Restarting computer..." {
		t.Errorf("Expected first restart result to be 'Restarting computer...', got '%s'", restartResults[0])
	}
}