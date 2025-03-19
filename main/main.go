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
	EsempioRightShiftUint64()
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

func EsempioExtractBitsFromUint64Array() {
	// Creo un array di Uint64 vuoto
	fmt.Println("\t\tTest EmptyArray, startBit = 0, length = 0")
	var emptyInputArray []uint64
	utils.PrintUint64ArrayAsBinary(emptyInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array vuoto
	res, err := utils.ExtractBitsFromUint64Array(emptyInputArray, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest EmptyArray, startBit = 1, length = 0")
	utils.PrintUint64ArrayAsBinary(emptyInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array vuoto
	res, err = utils.ExtractBitsFromUint64Array(emptyInputArray, 1, 0)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest EmptyArray, startBit = 0, length = 1")
	utils.PrintUint64ArrayAsBinary(emptyInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array vuoto
	res, err = utils.ExtractBitsFromUint64Array(emptyInputArray, 0, 1)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultatoF
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest EmptyArray, startBit = 1, length = 1")
	utils.PrintUint64ArrayAsBinary(emptyInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array vuoto
	res, err = utils.ExtractBitsFromUint64Array(emptyInputArray, 1, 1)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest SingleArray, startBit = 0, length = 0")
	var singleInputArray = []uint64{0x0000000000000001}
	utils.PrintUint64ArrayAsBinary(singleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con un solo elemento e startBit = 0 e length = 0
	res, err = utils.ExtractBitsFromUint64Array(singleInputArray, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest SingleArray, startBit = 0, length = 8")
	utils.PrintUint64ArrayAsBinary(singleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con un solo elemento e startBit = 0 e length = 8
	res, err = utils.ExtractBitsFromUint64Array(singleInputArray, 0, 8)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest SingleArray, startBit = 0, length = 4")
	utils.PrintUint64ArrayAsBinary(singleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con un solo elemento e startBit = 0 e length = 4
	res, err = utils.ExtractBitsFromUint64Array(singleInputArray, 0, 4)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest SingleArray, startBit = 0, length = 12")
	utils.PrintUint64ArrayAsBinary(singleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con un solo elemento e startBit = 0 e length = 12
	res, err = utils.ExtractBitsFromUint64Array(singleInputArray, 0, 12)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest SingleArray, startBit = 1, length = 8")
	utils.PrintUint64ArrayAsBinary(singleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con un solo elemento e startBit = 1 e length = 8
	res, err = utils.ExtractBitsFromUint64Array(singleInputArray, 1, 8)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest SingleArray, startBit = 58, length = 8")
	utils.PrintUint64ArrayAsBinary(singleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con un solo elemento e startBit = 58 e length = 8
	res, err = utils.ExtractBitsFromUint64Array(singleInputArray, 58, 8)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	var doubleInputArray = []uint64{0x0000000000000001, 0x0000000000000002}
	fmt.Println("\t\tTest DoubleArray, startBit = 0, length = 0")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 0 e length = 0
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 0, length = 8")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 0 e length = 8
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 0, 8)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 0, length = 64")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 0 e length = 64
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 0, 64)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 0, length = 65")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 0 e length = 65
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 0, 65)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 1, length = 64")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 1 e length = 64
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 1, 64)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 1, length = 63")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 1 e length = 63
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 1, 63)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 64, length = 1")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 64 e length = 1
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 64, 1)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 64, length = 16")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 64 e length = 16
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 64, 16)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 65, length = 16")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 65 e length = 16
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 65, 16)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 60, length = 16")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 60 e length = 16
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 60, 16)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()

	fmt.Println("\t\tTest DoubleArray, startBit = 55, length = 16")
	utils.PrintUint64ArrayAsBinary(doubleInputArray...)
	// Provo a chiamare ExtractBitsFromUint64Array con un array con due elementi e startBit = 55 e length = 16
	res, err = utils.ExtractBitsFromUint64Array(doubleInputArray, 55, 16)
	if err != nil {
		fmt.Println(err)
	}
	// Stampo il risultato
	utils.PrintUint64ArrayAsBinary(res...)
	fmt.Println()
}
