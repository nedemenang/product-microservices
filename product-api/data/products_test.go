package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "nn",
		Price: 1.00,
		SKU: "abs-absc-abdi",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}