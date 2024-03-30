package hsv

import (
	"fmt"
	"image/color"
	"testing"
)

func TestConvertToRGB(t *testing.T) {
	tests := []convertTest{
		convertTest{HSV: Color{85, 255, 255}, RGB: color.RGBA{255, 140, 0, 0}},
	}
	for _, v := range tests {
		test := v
		expected := v.RGB
		t.Run(fmt.Sprintf("%x:%x:%x", test.HSV.H, test.HSV.S, test.HSV.V), func(t *testing.T) {
			r, g, b := test.HSV.ConvertToRGB()
			converted := color.RGBA{r, g, b, 0}
			if converted != expected {
				println("converted: ", converted.R, converted.G, converted.B, converted.A)
				println("expected:  ", expected.R, expected.G, expected.B, expected.A)
				t.Fail()
			}
		})
	}
}

type convertTest struct {
	HSV Color
	RGB color.RGBA
}
