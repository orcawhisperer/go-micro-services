package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 1.13,
		SKU:   "abc-abcd-abcd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
