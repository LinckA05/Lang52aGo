package main
import "fmt"

func main() {
  var length, width, area, parameter int
  fmt.Println("Enter the length: ")
   fmt.Scanln(&length)
  fmt.Println("Enter the width: ")
   fmt.Scanln(&width)

  area = length * width
  parameter = (length * 2) + (width * 2)

  fmt.Println("The area is:", area)

}
