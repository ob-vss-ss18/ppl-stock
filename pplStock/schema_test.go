package pplStock

import (
	"testing"
)

func TestSchemaSki(t *testing.T) {
	var testSki = ski{
		id:      10,
		useCase: "Langlauf",
		brand:   "Fischer",
	}

	if testSki.id != 10 {
		t.Error("Expected 10, got ", testSki.id)
	}
	if testSki.useCase != "Langlauf" {
		t.Error("Expected 'Langlauf', got ", testSki.useCase)
	}
	if testSki.brand != "Fischer" {
		t.Error("Expected 'Fischer', got ", testSki.brand)
	}
}
