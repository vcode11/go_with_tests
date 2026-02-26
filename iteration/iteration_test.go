package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T){
	repeated := Repeat("a", 4)
	expected := "aaaa"
	if repeated != expected {
		t.Errorf("Expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 1000)
	}
}

func BenchmarkRepeatWithConcact(b *testing.B){
	concat := func (input string, numTimes int) string{
		var output string
		for range numTimes {
			output += input
		}
		return output
	}
	for b.Loop() {
		concat("a", 1000)
	}
}

func ExampleRepeat(){
	repeated := Repeat("ab", 4)
	fmt.Println(repeated)
	// Output: abababab
}