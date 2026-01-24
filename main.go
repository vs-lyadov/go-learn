package main

import "fmt"

func main() {
	// Chapter 1
	fmt.Println("# Chapter 1")
	fmt.Println("Hello, vs-lyadov!")

	// Chapter 2
	fmt.Println("# Chapter 2")
	var name string = "vs-lyadov"
	var age int = 25

	var firstname, lastname string = "Vyacheslav", "Lyadov"

	var (
		country string = "Russia"
		city    string = "Bryansk"
	)

	fmt.Printf("My name is %s %s. I'm %d years old.\n", firstname, lastname, age)
	fmt.Printf("Username: %s\n", name)
	fmt.Printf("I live in %s, %s.\n", city, country)

	printAllTypes()
	printConstants()
	printIota()

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

func printAllTypes() {
	printIntegers()
	printFloats()
	printComplex()
	printBool()
	printString()
}

func printString() {
	fmt.Println("\nString:")
	var a string = "Hello, World!"

	fmt.Printf("string: %s (type: %T)\n", a, a)
}

func printBool() {
	fmt.Println("\nBool:")
	var a bool = true
	var b bool = false

	fmt.Printf("bool: %t (type: %T)\n", a, a)
	fmt.Printf("bool: %t (type: %T)\n", b, b)
}

func printComplex() {
	fmt.Println("\nComplex:")
	var a complex64 = 1 + 2i
	var b complex128 = 3 + 4i

	fmt.Printf("complex64: %v (type: %T)\n", a, a)
	fmt.Printf("complex128: %v (type: %T)\n", b, b)
}

func printFloats() {
	fmt.Println("\nFloats:")
	var a float32 = 3.14
	var b float64 = 2.71828

	fmt.Printf("\nfloat32: %.2f (type: %T)\n", a, a)
	fmt.Printf("float64: %.5f (type: %T)\n", b, b)
}

func printIntegers() {
	fmt.Println("Integers:")
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

	var k int = 2147483647
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
