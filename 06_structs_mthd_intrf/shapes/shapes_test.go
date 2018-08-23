package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %2.f want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, v := range areaTests {
		t.Run(v.name, func(t *testing.T) {
			got := v.shape.Area()
			if got != v.want {
				t.Errorf("got %.2f want %.2f", got, v.want)
			}
		})
	}
}
