package template

import "testing"

func TestTemplateMethod(t *testing.T) {
	// Test CSV Processor
	csvProcessor := NewCSVProcessor()
	csvResults := csvProcessor.Execute()
	
	expectedCSV := []string{
		"Reading CSV data from file",
		"Processing CSV: Reading CSV data from file -> Parsed CSV data",
		"Writing to database: Processing CSV: Reading CSV data from file -> Parsed CSV data",
	}
	
	if len(csvResults) != len(expectedCSV) {
		t.Errorf("Expected %d CSV results, got %d", len(expectedCSV), len(csvResults))
	}
	
	for i, expected := range expectedCSV {
		if csvResults[i] != expected {
			t.Errorf("CSV result %d: expected '%s', got '%s'", i, expected, csvResults[i])
		}
	}
	
	// Test JSON Processor
	jsonProcessor := NewJSONProcessor()
	jsonResults := jsonProcessor.Execute()
	
	expectedJSON := []string{
		"Reading JSON data from API",
		"Processing JSON: Reading JSON data from API -> Parsed JSON data",
		"Writing to cache: Processing JSON: Reading JSON data from API -> Parsed JSON data",
	}
	
	if len(jsonResults) != len(expectedJSON) {
		t.Errorf("Expected %d JSON results, got %d", len(expectedJSON), len(jsonResults))
	}
	
	for i, expected := range expectedJSON {
		if jsonResults[i] != expected {
			t.Errorf("JSON result %d: expected '%s', got '%s'", i, expected, jsonResults[i])
		}
	}
	
	// Test XML Processor
	xmlProcessor := NewXMLProcessor()
	xmlResults := xmlProcessor.Execute()
	
	expectedXML := []string{
		"Reading XML data from web service",
		"Processing XML: Reading XML data from web service -> Parsed XML data",
		"Writing to file system: Processing XML: Reading XML data from web service -> Parsed XML data",
	}
	
	if len(xmlResults) != len(expectedXML) {
		t.Errorf("Expected %d XML results, got %d", len(expectedXML), len(xmlResults))
	}
	
	for i, expected := range expectedXML {
		if xmlResults[i] != expected {
			t.Errorf("XML result %d: expected '%s', got '%s'", i, expected, xmlResults[i])
		}
	}
}