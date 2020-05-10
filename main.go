package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {

reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your city1: ")
    city, _ := reader.ReadString('\n')
    fmt.Print("Enter your city2: ")
    city1, _ := reader.ReadString('\n')
    fmt.Print("Enter your city3: ")
    city2, _ := reader.ReadString('\n')
    fmt.Print("Enter your city4: ")
    city3, _ := reader.ReadString('\n')
    fmt.Print("Enter your city5: ")
    city4, _ := reader.ReadString('\n')
     fmt.Print("Enter your city6: ")
    city5, _ := reader.ReadString('\n')
     fmt.Print("Enter your city7: ")
    city6, _ := reader.ReadString('\n')
     fmt.Print("Enter your city8: ")
    city7, _ := reader.ReadString('\n')
     fmt.Print("Enter your city9: ")
    city8, _ := reader.ReadString('\n')
     fmt.Print("Enter your city10: ")
    city9, _ := reader.ReadString('\n')


	cities := []string{city, city1, city2,city3,city4,city5,city6,city7,city8, city9}

	afterSlice := cities[3:]
	afterMoreSlice := afterSlice[1:5]

	fmt.Println("cities = ", cities)
	fmt.Println("afterSlice = ", afterSlice)
	fmt.Println("afterMoreSlice = ", afterMoreSlice)
}