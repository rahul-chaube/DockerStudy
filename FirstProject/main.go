package main

import "fmt"

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"Message": "Pong",
	// 	})

	// })
	// r.Run(":8080")

	slice := make([]int, 10)
	slice1 := []int{}

	fmt.Println(" Slice  ", len(slice), slice)
	fmt.Println(" Slice 1 ", len(slice1), slice1)

	slice = append(slice, 10)

	fmt.Println(" Slice  ", len(slice), slice)
	fmt.Println(" Slice 1 ", len(slice1), slice1)

	temp := slice[10:10]

	temp = append(temp, 20)
	slice = append(slice, 8)
	temp = append(temp, 30)

	temp = append(temp, 40)

	fmt.Println(" Slice  ", len(slice), slice, temp)
	fmt.Println(" Slice 1 ", len(slice1), slice1)
	temp[0] = 5

	fmt.Println(" Slice  ", len(slice), slice, temp)
	fmt.Println(" Slice 1 ", len(slice1), slice1)
}
