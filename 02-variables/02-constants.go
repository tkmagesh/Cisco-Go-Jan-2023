package main

import "fmt"

func main() {
	/*
		const pi float32 = 3.14
		fmt.Println(pi)
	*/

	const (
		app_name    = "constants.go"
		app_version = "1.0"
	)
	fmt.Println(app_name, app_version)

	//iota
	/*
		const red int = 0
		const green int = 1
		const blue int = 2
	*/

	/*
		const (
			red   int = 0
			green int = 1
			blue  int = 2
		)
	*/

	/*
		const (
			red   = 0
			green = 1
			blue  = 2
		)
	*/

	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 1
			green
			blue
		)
	*/

	/*
		const (
			red = iota * 2
			green
			blue
		)
	*/

	/*
		const (
			red = (iota * 2) + 1
			green
			blue
		)
	*/

	const (
		red = (iota * 2) + 1
		green
		_
		blue
	)
	fmt.Printf("red = %d, green = %d and blue = %d\n", red, green, blue)

	fmt.Println("iota applied")
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}
