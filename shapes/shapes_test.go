package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectange{10, 10}
	got := r.Perimeter()
	want := 40.0
	assertFloats(t, got, want)
}

func TestArea(t *testing.T) {
	testArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertFloats(t, got, want)
	}
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{
			Rectange{12, 15},
			180.0,
		},
		{
			Circle{12},
			452.3893421169302087,
		},
		{
			Triangle{12, 6},
			36.0,
		},
	}
	for _, testCase := range areaTests{
		testArea(t, testCase.shape, testCase.want)
	}
}

func assertFloats(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.16f want %.16f", got, want)
	}
}
