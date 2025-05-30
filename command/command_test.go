package command

import "testing"

func TestCommand(t *testing.T) {
	light := NewLight()
	remote := NewRemoteControl()
	
	// Initially light should be off
	if light.IsOn() {
		t.Error("Light should be initially off")
	}
	
	// Test turning light on
	lightOnCmd := NewLightOnCommand(light)
	remote.SetCommand(lightOnCmd)
	result := remote.PressButton()
	
	if result != "Light is ON" {
		t.Errorf("Expected 'Light is ON', got %s", result)
	}
	
	if !light.IsOn() {
		t.Error("Light should be on after executing turn on command")
	}
	
	// Test turning light off
	lightOffCmd := NewLightOffCommand(light)
	remote.SetCommand(lightOffCmd)
	result = remote.PressButton()
	
	if result != "Light is OFF" {
		t.Errorf("Expected 'Light is OFF', got %s", result)
	}
	
	if light.IsOn() {
		t.Error("Light should be off after executing turn off command")
	}
	
	// Test undo functionality
	result = remote.PressUndo()
	if result != "Light is ON" {
		t.Errorf("Expected 'Light is ON' after undo, got %s", result)
	}
	
	if !light.IsOn() {
		t.Error("Light should be on after undoing turn off command")
	}
	
	// Test undo again
	result = remote.PressUndo()
	if result != "Light is OFF" {
		t.Errorf("Expected 'Light is OFF' after second undo, got %s", result)
	}
	
	if light.IsOn() {
		t.Error("Light should be off after undoing turn on command")
	}
	
	// Test undo when no commands in history
	result = remote.PressUndo()
	if result != "No command to undo" {
		t.Errorf("Expected 'No command to undo', got %s", result)
	}
}