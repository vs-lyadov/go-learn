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

	printIntegers()

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

	fmt.Printf("int8: %d\n", a)
	fmt.Printf("int16: %d\n", b)
	fmt.Printf("int32: %d\n", c)
	fmt.Printf("int64: %d\n", d)
	fmt.Printf("uint8: %d\n", e)
	fmt.Printf("uint16: %d\n", f)
	fmt.Printf("uint32: %d\n", g)
	fmt.Printf("uint64: %d\n", h)
	fmt.Printf("byte: %d\n", i)
	fmt.Printf("rune: %d\n", j)
	fmt.Printf("int: %d\n", k)
	fmt.Printf("uint: %d\n", l)
}
