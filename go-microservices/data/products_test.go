package data

import "testing"

func TestChecksValidaion(t *testing.T) {
	p := &Product{Name: "nics", Price: 1.00, SKU: "abs-das-das-vd"}
	if err := p.Validate(); err != nil {
		t.Fatal(err)
	}
}
