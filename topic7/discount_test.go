package main

import (
	"math"
	"testing"
)

func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		name       string
		price      float64
		percentage float64
		want       float64
		wantErr    bool
	}{
		{name: "10 percent discount", price: 100, percentage: 10, want: 90},
		{name: "zero discount", price: 250, percentage: 0, want: 250},
		{name: "full discount", price: 80, percentage: 100, want: 0},
		{name: "negative price", price: -10, percentage: 10, wantErr: true},
		{name: "negative percentage", price: 100, percentage: -1, wantErr: true},
		{name: "percentage above 100", price: 100, percentage: 101, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CalculateDiscount(test.price, test.percentage)
			if (err != nil) != test.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, test.wantErr)
			}
			if !test.wantErr && math.Abs(got-test.want) > 1e-9 {
				t.Errorf("got %.2f, want %.2f", got, test.want)
			}
		})
	}
}
