package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Chapter 1
	fmt.Println("# Chapter 1")
	fmt.Println("Hello, vs-lyadov!")

	// Chapter 2
	fmt.Println("# Chapter 2")
	var name = "vs-lyadov"
	var age = 25

	var firstname, lastname = "Vyacheslav", "Lyadov"

	var (
		country = "Russia"
		city    = "Bryansk"
	)

	fmt.Printf("My name is %s %s. I'm %d years old.\n", firstname, lastname, age)
	fmt.Printf("Username: %s\n", name)
	fmt.Printf("I live in %s, %s.\n", city, country)

	var skills = []string{"Go", "PHP", "JavaScript"}
	fmt.Printf("Skills: %v\n", skills)

	if len(skills) > 0 {
		fmt.Printf("Best skill: %s\n", skills[0])
	}

	switch len(skills) {
	case 0:
		fmt.Println("No skills")
		fallthrough
	case 1:
		fmt.Println("One skill")
	default:
		fmt.Printf("%d skills\n", len(skills))
	}

	printAllTypes()
	printLoops()

	// Chapter 3
	fmt.Println("# Chapter 3")
	printFunctions()

	// Chapter 4
	fmt.Println("# Chapter 4")
	printPointers()
}

func printPointers() {
	var x = 10
	var p *int

	p = &x
	fmt.Printf("x: %d (type: %T)\n", x, x)
	fmt.Printf("p: %v (type: %T)\n", p, p)
	fmt.Printf("*p: %d (type: %T)\n", *p, *p)

	*p = 20
	fmt.Printf("x: %d (type: %T)\n", x, x)

	var pointer *float64

	if pointer == nil {
		fmt.Println("pointer is nil")
		pointer = new(float64)
		*pointer = 3.14
	} else {
		fmt.Println("pointer is not nil")
	}

	fmt.Printf("pointer: %v (type: %T)\n", pointer, pointer)
	fmt.Printf("*pointer: %f (type: %T)\n", *pointer, *pointer)

	var a, b, d = 1, 2, 4
	var nums [4]*int
	nums[0] = &a
	nums[1] = &b
	nums[3] = &d
	fmt.Printf("nums: %v (type: %T)\n", nums, nums)

	var ptr = sumAndCreatePointers(&a, &b)
	fmt.Printf("a: %v (type: %T)\n", a, a)
	fmt.Printf("*ptr: %d (type: %T)\n", *ptr, *ptr)
}

func sumAndCreatePointers(a, b *int) *int {
	*a = *a + *b
	var ptr = new(int)
	return ptr
}

func printFunctions() {
	fmt.Println("\nFunctions:")

	var nums = []int{1, 2, 3, 4, 5}
	var b, randInt = sum(nums...)
	fmt.Printf("sum(nums...): %d (type: %T)\n", b, b)
	fmt.Printf("rand: %d (type: %T)\n", randInt, randInt)

	add := func(a, b int) int { return a + b }
	fmt.Printf("add(1, 2): %d (type: %T)\n", add(1, 2), add(1, 2))

	f := closure()
	f()
	f()

	fmt.Printf("factorial(10): %d \n", factorial(10))
	fmt.Printf("fibonacci(10): %d \n", fibonacci(10))

}

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func closure() func() {
	var x = 10
	return func() {
		x++
		fmt.Println(x)
	}
}

func sum(numbers ...int) (result int, randStr int32) {
	for _, num := range numbers {
		result += num
	}

	randStr = rand.New(rand.NewSource(time.Now().UnixNano())).Int31()
	return
}

func printLoops() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Loop i = %d\n", i)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	str := "Hello, World!"
	for _, val := range str {
		fmt.Printf("%c", val)
	}
	fmt.Println()

	var users = []string{"vs-lyadov", "john_doe", "jane_smith"}
	for i, user := range users {
		fmt.Printf("Index: %d, User: %s", i, user)

		if user == "vs-lyadov" {
			fmt.Print(" — это я!")
		}

		fmt.Println()
	}

	for i := 0; i < len(users); i++ {
		if users[i] == "john_doe" {
			continue
		}

		if users[i] == "jane_smith" {
			break
		}

		fmt.Printf("User: %s\n", users[i])

	}

}

func printAllTypes() {
	printIntegers()
	printFloats()
	printComplex()
	printBool()
	printString()

	printConstants()
	printIota()
	printBitwiseShift()
	printArray()
}

func printArray() {
	numbers := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Printf("Numbers: %v\n", numbers)
}

func printBitwiseShift() {
	var i = 8
	fmt.Printf("%d (type: %T)\n", i<<1, i)
	fmt.Printf("%d (type: %T)\n", i>>2, i)
}

func printIota() {
	const (
		a = iota
		b
		c
	)

	fmt.Println("\niota:")
	fmt.Printf("a: %d (type: %T)\n", a, a)
	fmt.Printf("b: %d (type: %T)\n", b, b)
	fmt.Printf("c: %d (type: %T)\n", c, c)
}

func printConstants() {
	const pi float64 = 3.14
	const e float64 = 2.71828

	fmt.Println("\nConstants:")
	fmt.Printf("pi: %.2f (type: %T)\n", pi, pi)
	fmt.Printf("e: %.5f (type: %T)\n", e, e)
}

func printString() {
	fmt.Println("\nString:")
	var a = "Hello, World!"

	fmt.Printf("string: %s (type: %T)\n", a, a)
}

func printBool() {
	fmt.Println("\nBool:")
	var a = true
	var b = false

	fmt.Printf("bool: %t (type: %T)\n", a, a)
	fmt.Printf("bool: %t (type: %T)\n", b, b)
}

func printComplex() {
	fmt.Println("\nComplex:")
	var a complex64 = 1 + 2i
	var b = 3 + 4i

	fmt.Printf("complex64: %v (type: %T)\n", a, a)
	fmt.Printf("complex128: %v (type: %T)\n", b, b)
}

func printFloats() {
	fmt.Println("\nFloats:")
	var a float32 = 3.14
	var b = 2.71828

	fmt.Printf("\nfloat32: %.2f (type: %T)\n", a, a)
	fmt.Printf("float64: %.5f (type: %T)\n", b, b)
}

func printIntegers() {
	fmt.Println("\nIntegers:")
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	var e uint8 = 255
	var f uint16 = 65535
	var g uint32 = 4294967295
	var h uint64 = 18446744073709551615

	var i byte = 255
	var j rune = 2147483647

	var k = 2147483647
	var l uint = 4294967295

	fmt.Printf("int8: %d (type: %T)\n", a, a)
	fmt.Printf("int16: %d (type: %T)\n", b, b)
	fmt.Printf("int32: %d (type: %T)\n", c, c)
	fmt.Printf("int64: %d (type: %T)\n", d, d)
	fmt.Printf("uint8: %d (type: %T)\n", e, e)
	fmt.Printf("uint16: %d (type: %T)\n", f, f)
	fmt.Printf("uint32: %d (type: %T)\n", g, g)
	fmt.Printf("uint64: %d (type: %T)\n", h, h)
	fmt.Printf("byte: %d (type: %T)\n", i, i)
	fmt.Printf("rune: %d (type: %T)\n", j, j)
	fmt.Printf("int: %d (type: %T)\n", k, k)
	fmt.Printf("uint: %d (type: %T)\n", l, l)
}
