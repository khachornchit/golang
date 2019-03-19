package main

import ("fmt"
	"math"
	"math/rand"
	"time")

func main() {
    // runMath()
    // num1, num2 := 5.6, 9.5
    // fmt.Println(add(num1, num2))
    multiple("welcome, ", "GO !")
}

func multiple(a, b string) (string, string) {
    return a, b
}

func sqrt () {
	fmt.Println("Sqrt(4) = ", math.Sqrt(9));
}

func random() {
	fmt.Println("Random number 1-100 is ", rand.Intn(100));
}

func add(x float64, y float64) float64 {
	return x+y;
}

func runMath() {
	sqrt()
	random()
	fmt.Println("add(5, 9) = ", add(5, 9))

	fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    fmt.Println(rand.Float64())

    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
}