package main

import (
	"fmt"
)

// A Cylinder represents any gas cylinder, whether target or cascade
type Cylinder struct {
	PMax	float64	// Max pressure of cylinder
	V       float64 // Volume of cylinder
	P 	float64 // Current pressure of cylinder
}

// Transfer transfers gas between cylinders c1 and c2 until c1 reaches pressure p1.
// If the target pressure is unattainable, equalize the pressures (i.e get as close
// as possible).
func Transfer(c1, c2 *Cylinder, p1 float64) {
	// Calculate equalization pressures
	//pe := c2.P * (1 - (c2.P - c1.P) / c1.PMax / (c2.PMax/c1.PMax + c2.V/c1.V) )
	pe := (c1.P * c1.V / c1.PMax + c2.P * c2.V / c2.PMax) / (c1.V / c1.PMax + c2.V / c2.PMax)
	if (c1.P < c2.P && p1 > pe) || (c1.P > c2.P && p1 < pe)  {
		p1 = pe
	}
	c2.P = c2.P - (p1 - c1.P) * (c1.V / c1.PMax) / (c2.V / c2.PMax)
	c1.P = p1
}

func PrintPressures(c0 *Cylinder, cascade []*Cylinder) {
	fmt.Printf("%4.0f ", c0.P)

	for _, c := range cascade {
		fmt.Printf("%4.0f ", c.P)
	}
	fmt.Print("\n")
}

func main() {
	// Target cylinder
	c := &Cylinder{PMax: 2216, V: 48.2, P: 500}

	// Cascade
	n := 3 // Number of cylinders in cascade
	cascade := make([]*Cylinder, n)
	for i := 0; i < len(cascade); i++ {
		cascade[i] = &Cylinder{PMax: 2400, V: 300, P: 2400}
	}

	for n := 0; n < 10; n++ {
		c.P = 500
		PrintPressures(c, cascade)

		for _, cc := range cascade {
			Transfer(c, cc, c.PMax)
			PrintPressures(c, cascade)
			if c.P > c.PMax - 1 {
				break
			}
		}
		fmt.Println()
		if c.P < 1000 {
			break
		}
	}
}
