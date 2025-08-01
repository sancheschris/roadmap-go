package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500
	expected := 5.0

	result := CalculateTax(float64(amount))

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// go tool cover -html=coverage.out or go test -coverprofile=coverag.out
// go test -v
// go test -bench=. or go test -bench=. -run=^#, 
// go test -bench=. -run=^# -benchmem
// go test -bench=. -run=^# -count=3 