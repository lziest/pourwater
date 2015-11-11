package main

import (
	"flag"
	"fmt"
)

var (
	X = flag.Int("x", 5, "X value")
	Y = flag.Int("y", 3, "Y value")
	N = flag.Int("n", 4, "N value")
)

// gcd(x,y) returns d in the form of ax + by = d , where
// d is the greatest common divisor.
func gcd(x, y int) (d, a, b int) {
	if x%y == 0 {
		return y, 0, 1
	} else {
		// a' * x + b' * (x % y) = d
		// a' * x + b' * (x - x/y * y) = d
		// (a' - b'*(x/y)) - b'*x = d
		d, aprime, bprime := gcd(y, x%y)
		return d, bprime, aprime - bprime*(x/y)
	}
}

// lcm(x,y) return least common multiple m, and a and b such that
// a*x = m, b*y= m
func lcm(x, y int) (m, a, b int) {
	d, _, _ := gcd(x, y)
	return x / d * y, y / d, x / d
}

func findDecantationNumber(x, y, a, b int) int {
	if a <= 0 {
		a, b = b, a
		x, y = y, x
	}
	// fill x-glass a times
	// empty x-glass a-1 times if a > 0
	// fill y-glass |b| times
	// empty y-glass |b|-1 times if |b| > 0
	s := Abs(a) + Abs(b)
	if a != 0 {
		s += Abs(a) - 1
	}
	if b != 0 {
		s += Abs(b) - 1
	}
	// extra decantation:
	// if emptying x-glass (|a|-1) times exceeded filling  y-glass |b| times
	if a > 0 && Abs(b) > 0 && (a-1)*x+b*y > 0 {
		s = s + 1
	}

	return s
}

// return s, u, x, v, y such that s is the minimal steps to
// get t liter of water from using x-liter and y-liter cup
// and u*x + v*y = t
func solve(t, x, y int) (s, u, v, newX, newY int) {
	d, a, b := gcd(x, y)
	// make sure a > 0, b < 0
	if a < 0 {
		a, b = b, a
		x, y = y, x
	}

	if t%d != 0 {
		// unachievable
		return 0, 0, 0, x, y
	}
	multiplier := t / d
	am := a * multiplier
	bm := b * multiplier

	// find the factor fx and fy such that fx * x = fy * y
	_, fx, fy := lcm(x, y)
	// reduce am to the minimal postive number
	// and also the maximal negative number
	// am * x + bm * y = t
	// (am - fx)*x + (bm + fy)*y = t ...
	// (am - r*fx)*x + (bm +r*fy)*y = t ...
	r := am / fx
	ar := am % fx
	br := bm + r*fy

	s1 := findDecantationNumber(x, y, ar, br)

	ar = ar - fx
	br = br + fy
	s2 := findDecantationNumber(x, y, ar, br)

	if s1 < s2 {
		return s1, ar + fx, br - fy, x, y
	}
	return s2, ar, br, x, y

}

func simulate(t, x, y, a, b int) {
	if a <= 0 {
		x, y = y, x
		a, b = b, a
	}

	gx, gy := 0, 0 //initially glasses are empty
	r := x + y
	b = -b
	c := 0
	for a > 0 || b > 0 {
		c += 1
		fmt.Println("Step", c, ":")
		switch {
		case gx == 0: // fill the first glass
			gx = x
			r = r - x
			fmt.Printf("Pouring water to glass of %d-liter\n", x)
			fmt.Printf("Situation: %d, %d, %d\n", gx, gy, r)
			a = a - 1
			continue
		case gy < y: // fill the second glass with the first glass
			dy := y - gy
			dx := gx
			d := 0
			if dx > dy {
				d = dy
			} else {
				d = dx
			}

			gy = gy + d
			gx = gx - d
			if gy == y {
				b = b - 1
			}
			fmt.Printf("Pouring water from glass of %d-liter to glass of %d-liter\n", x, y)
			fmt.Printf("Situation: %d, %d, %d\n", gx, gy, r)
			continue
		case gy == y: // empty the second glass
			gy = 0
			r = r + y
			fmt.Printf("Pouring water from glass of %d-liter to reservior\n", y)
			fmt.Printf("Situation: %d, %d, %d\n", gx, gy, r)
			continue

		default:
			fmt.Println("impossible")
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	flag.Parse()

	x := *X
	y := *Y
	n := *N

	// target 1: get n directly
	s1, a1, b1, x, y := solve(n, x, y)
	// target 2: get (x+y-n) directly
	s2, a2, b2, x, y := solve(x+y-n, x, y)

	if s1 > 0 && s1 < s2 {
		fmt.Println("Minimal number of pouring", s1)
		simulate(n, x, y, a1, b1)
		return
	} else if s2 > 0 {
		fmt.Println("Minimal number of pouring", s2)
		simulate(x+y-n, x, y, a2, b2)
		return
	}

	fmt.Println("It is impossible")
}
