package game

import "fmt"

func New()  {
	var x, y int
	fmt.Println("please specify the size of the field in length and height")
	fmt.Scanf("%s\n", &x, &y)

	fmt.Println(x, y)
}
