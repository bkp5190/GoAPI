package main
import "fmt"

func main()  {
  printString("Hello World from outer function!")
}

func printString(target string){
  fmt.Println(target)
}
