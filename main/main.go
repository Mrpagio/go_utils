package main

import (
	"fmt"
	utils "github.com/Mrpagio/go_utils"
)

// Voglio che nel main provare il funzionamento di ExtractBitsFromUint64Array

func main() {

	// TEST ExtractBitsFromUint64Array
	//EsempioExtractBitsFromUint64Array()

	// TEST EsempioLeftShiftUint64
	//EsempioLeftShiftUint64()

	// TEST EsempioRightShiftUint64
	//EsempioRightShiftUint64()

	// TEST EsempioShiftJoinUint64
	//EsempioShiftUint64Array()

	// TEST EsempioReplaceBits
	EsempioReplaceBits()
}

func EsempioReplaceBits() {
	fmt.Println("EsempioReplaceBits()")
	// Creo un array di byte
	input := []byte{0xFF, 0xFF, 0xFF, 0xFF}
	fmt.Print("\tinput: ")
	for _, b := range input {
		fmt.Printf("%08b ", b)
	}
	fmt.Println("")

	pack := utils.PackBigEndian(input)
	fmt.Println("\tpack: ", pack)

	newBytes := []byte{0x7E}
	fmt.Print("\tnewBytes: ")
	for _, b := range newBytes {
		fmt.Printf("%08b ", b)
	}
	fmt.Println("")

	err := utils.ReplaceBits(&input, newBytes, 8, 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("\toutput: ")
	for _, b := range input {
		fmt.Printf("%08b ", b)
	}
	fmt.Println()
}

func EsempioShiftUint64Array() {
	// Creo un []uint64 con due elementi 0x0 e 0xFF00000000000000
	var input = []uint64{0x0, 0xFFFFFFFFFFFFFFFF, 0xFF00000000000000}
	fmt.Print("\tinput: ")
	utils.PrintUint64ArrayAsBinary(input...)
	var err error
	var res []uint64

	res, err = utils.ShiftUint64Array(input, 7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("\toutput: ")
	utils.PrintUint64ArrayAsBinary(res...)
}

func EsempioRightShiftUint64() {

	fmt.Println("\t\tTest RightShiftUint64, shift = 1")
	// Creo un uint64 con un 1 in posizione 12-13-14-15-16-17-18-19
	//0000000000001111111100000000000000000000000000000000000000000000
	input := uint64(0xFF00000000000)

	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(input)
	res, err := utils.RightShiftUint64(input, 1)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(res)
	fmt.Println()

	fmt.Println("\t\tTest RightShiftUint64, shift = 8")
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(input)
	res, err = utils.RightShiftUint64(input, 8)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(res)
	fmt.Println()

	fmt.Println("\t\tTest RightShiftUint64, shift = 16")
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(input)
	res, err = utils.RightShiftUint64(input, 16)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(res)
	fmt.Println()

	fmt.Println("\t\tTest RightShiftUint64, shift = 32")
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(input)
	res, err = utils.RightShiftUint64(input, 32)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(res)
	fmt.Println()

	fmt.Println("\t\tTest RightShiftUint64, shift = 56")
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(input)
	res, err = utils.RightShiftUint64(input, 56)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(res)
	fmt.Println()
}

func EsempioLeftShiftUint64() {

	fmt.Println("\t\tTest LeftShiftUint64, shift = 1")
	// Creo un uint64 con un 1 in posizione 12-13-14-15-16-17-18-19
	//0000000000001111111100000000000000000000000000000000000000000000
	input := uint64(0xFF00000000000)

	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(input)
	res, err := utils.LeftShiftUint64(input, 1)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(res)
	fmt.Println()

	fmt.Println("\t\tTest LeftShiftUint64, shift = 8")
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(input)
	res, err = utils.LeftShiftUint64(input, 8)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(res)
	fmt.Println()

	fmt.Println("\t\tTest LeftShiftUint64, shift = 16")
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(input)
	res, err = utils.LeftShiftUint64(input, 16)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il numero in binario
	utils.PrintUint64ArrayAsBinary(res)

}
