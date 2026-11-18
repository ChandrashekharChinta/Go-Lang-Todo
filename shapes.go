package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

type Shape interface {
	Area() float64
	Perimeter() float64
	fmt.Stringer
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func NewRectangle(width, height float64) (*Rectangle, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("Width and Height must be positive")
	}
	return &Rectangle{
		Width:  width,
		Height: height,
	}, nil

	// Rectangle{} creates a struct value.
	// &Rectangle{} creates a pointer to that struct.
	// Our constructor promises to return *Rectangle.
}

func NewCircle(radius float64) (*Circle, error) {
	if radius <= 0 {
		return nil, errors.New("Radius must be positive")
	}
	return &Circle{
		Radius: radius,
	}, nil
}

func NewTriangle(a, b, c float64) (*Triangle, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, errors.New("triangle sides must be positive")
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return nil, errors.New("invalid triangle sides")
	}
	return &Triangle{
		SideA: a,
		SideB: b,
		SideC: c,
	}, nil
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) String() string {
	return fmt.Sprintf(
		"Rectangle [Width: %.2f, Height: %.2f]",
		r.Width,
		r.Height,
	)
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) String() string {
	return fmt.Sprintf(
		"Circle [Radius: %.2f]",
		c.Radius,
	)
}

func (t *Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2

	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t *Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func (t *Triangle) String() string {
	return fmt.Sprintf(
		"Triangle [SideA: %.2f, SideB: %.2f, SideC: %.2f]",
		t.SideA, t.SideB, t.SideC,
	)
}

type ShapeCalculator struct{}

func NewShapeCalculator() *ShapeCalculator {
	return &ShapeCalculator{}
}

// ShapeCalculator is an empty struct because it doesn't need to store any data.
// NewShapeCalculator() creates and returns a pointer to a calculator.
// We'll add methods to this struct to calculate and display shape properties.

func (sc *ShapeCalculator) TotalArea(shapes []Shape) float64 {

	total := 0.0

	for _, shape := range shapes {
		if shape == nil {
			continue
		}
		total += shape.Area()
	}

	return total
}

func (sc *ShapeCalculator) LargestShape(shapes []Shape) Shape {
	var largest Shape

	for _, shape := range shapes {
		if shape == nil {
			continue
		}

		if largest == nil || shape.Area() > largest.Area() {
			largest = shape
		}
	}

	return largest
}

func (sc *ShapeCalculator) SortByArea(shapes []Shape, ascending bool) []Shape {
	result := make([]Shape, len(shapes))
	copy(result, shapes)

	sort.Slice(result, func(i, j int) bool {
		if result[i] == nil {
			return false
		}
		if result[j] == nil {
			return true
		}

		if ascending {
			return result[i].Area() < result[j].Area()
		}

		return result[i].Area() > result[j].Area()
	})

	return result
}

func main() {
	// Create shapes
	rectangle, err := NewRectangle(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	circle, err := NewCircle(7)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	triangle, err := NewTriangle(3, 4, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Store different shapes in one slice
	shapes := []Shape{
		rectangle,
		circle,
		triangle,
	}

	// Create calculator
	calculator := NewShapeCalculator()

	// Calculate total area
	fmt.Println("Total Area:", calculator.TotalArea(shapes))

	// Find largest shape
	fmt.Println("\n----- Largest Shape -----")
	largest := calculator.LargestShape(shapes)
	if largest != nil {
		fmt.Println(largest)
		fmt.Println("Area:", largest.Area())
	}

	// Sort in ascending order
	fmt.Println("\n----- Ascending Order -----")
	ascending := calculator.SortByArea(shapes, true)
	for _, shape := range ascending {
		fmt.Println(shape, "Area:", shape.Area())
	}

	// Sort in descending order
	fmt.Println("\n----- Descending Order -----")
	descending := calculator.SortByArea(shapes, false)
	for _, shape := range descending {
		fmt.Println(shape, "Area:", shape.Area())
	}
}
