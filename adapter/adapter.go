package adapter

import "fmt"

type OldPrinter struct{}

func (p *OldPrinter) Print(text string) string {
	return fmt.Sprintf("old printer: %s", text)
}

type NewPrinter interface {
	Print(text string) string
}

type PrinterAdapter struct {
	oldPrinter *OldPrinter
}

func (p *PrinterAdapter) Print(text string) string {
	return p.oldPrinter.Print(text)
}
