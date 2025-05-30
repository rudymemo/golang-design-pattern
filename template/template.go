package template

import "fmt"

type DataProcessor interface {
	ReadData() string
	ProcessData(data string) string
	WriteData(data string) string
	Execute() []string
}

type BaseDataProcessor struct {
	processor DataProcessor
}

func (bdp *BaseDataProcessor) Execute() []string {
	var results []string
	
	// Template method defining the algorithm structure
	data := bdp.processor.ReadData()
	results = append(results, data)
	
	processedData := bdp.processor.ProcessData(data)
	results = append(results, processedData)
	
	writeResult := bdp.processor.WriteData(processedData)
	results = append(results, writeResult)
	
	return results
}

type CSVProcessor struct {
	BaseDataProcessor
}

func NewCSVProcessor() *CSVProcessor {
	csv := &CSVProcessor{}
	csv.processor = csv
	return csv
}

func (csv *CSVProcessor) ReadData() string {
	return "Reading CSV data from file"
}

func (csv *CSVProcessor) ProcessData(data string) string {
	return fmt.Sprintf("Processing CSV: %s -> Parsed CSV data", data)
}

func (csv *CSVProcessor) WriteData(data string) string {
	return fmt.Sprintf("Writing to database: %s", data)
}

type JSONProcessor struct {
	BaseDataProcessor
}

func NewJSONProcessor() *JSONProcessor {
	json := &JSONProcessor{}
	json.processor = json
	return json
}

func (json *JSONProcessor) ReadData() string {
	return "Reading JSON data from API"
}

func (json *JSONProcessor) ProcessData(data string) string {
	return fmt.Sprintf("Processing JSON: %s -> Parsed JSON data", data)
}

func (json *JSONProcessor) WriteData(data string) string {
	return fmt.Sprintf("Writing to cache: %s", data)
}

type XMLProcessor struct {
	BaseDataProcessor
}

func NewXMLProcessor() *XMLProcessor {
	xml := &XMLProcessor{}
	xml.processor = xml
	return xml
}

func (xml *XMLProcessor) ReadData() string {
	return "Reading XML data from web service"
}

func (xml *XMLProcessor) ProcessData(data string) string {
	return fmt.Sprintf("Processing XML: %s -> Parsed XML data", data)
}

func (xml *XMLProcessor) WriteData(data string) string {
	return fmt.Sprintf("Writing to file system: %s", data)
}