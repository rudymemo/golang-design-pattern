package singleton

import "testing"

func TestSingleton(t *testing.T) {
	db1 := GetInstance()
	db2 := GetInstance()

	if db1 != db2 {
		t.Error("Expected same instance, got different instances")
	}

	result := db1.Query("SELECT * FROM users")
	expected := "Executing: SELECT * FROM users on database connection established"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}