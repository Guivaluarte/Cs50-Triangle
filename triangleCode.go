package main

import "fmt"

func main() {

	var side1, side2, side3 int

	fmt.Println("Welcome to the triangle validation.")
	fmt.Println("Type the side of the triangles")

	fmt.Scan(&side1, &side2, &side3)

	var is_triangle_possible bool = triangle_calc(side1, side2, side3)

	if is_triangle_possible == true {
		fmt.Println("Your Triangle is possible")
	} else {
		fmt.Println("your triangle is not possible")
	}
}

func triangle_calc(side1, side2, side3 int) bool {

	var is_triangle_possible bool = true

	if side1+side2 <= side3 || side2+side3 <= side1 {
		is_triangle_possible = false
	}

	return is_triangle_possible

}
