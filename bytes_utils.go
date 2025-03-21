package utils

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

// HexToByteArray
// prende in input una stringa "esadecimale"
// se riesce a convertire la stringa in un byte array restituisce il byte array corrispondente
func HexToByteArray(hex string) ([]byte, error) {
	// controllo che la lunghezza della stringa sia pari
	if len(hex)%2 != 0 {
		return nil, fmt.Errorf("HexToByteArray() -> lunghezza stringa dispari")
	}
	// converto due caratteri esadecimali in un byte
	res := make([]byte, len(hex)/2)
	for i := 0; i < len(hex); i += 2 {
		_, err := fmt.Sscanf(hex[i:i+2], "%02x", &res[i/2])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// ByteArrayToHex
// prende in input un array di byte
// restituisce la stringa "esadecimale" corrispondente con spazi tra i byte
func ByteArrayToHex(arr []byte) string {
	str := ""
	for i := 0; i < len(arr); i++ {
		str += fmt.Sprintf("%02X ", arr[i])
	}
	return str
}

// BytesToHex
// prende in input un byte
// restituisce la stringa "esadecimale" corrispondente con spazi tra i byte
func BytesToHex(arr ...byte) string {
	str := ""
	for i := 0; i < len(arr); i++ {
		str += fmt.Sprintf("%02X ", arr[i])
	}
	return str
}

// BytesAsBinaryString
// Prende in input un array di byte
// Converte ogni byte in una stringa binaria di 8 bit con uno spazio tra un byte e l'altro
// Ritorna la stringa
func BytesAsBinaryString(data ...byte) string {
	str := ""
	for _, b := range data {
		str += fmt.Sprintf("%08b ", b)
	}
	str = strings.TrimSpace(str)
	return str
}

// XorBytes
// Prende in input due array di byte
// controlla che abbiano la stessa lunghezza
// Esegue l'operazione di XOR tra i due array
// Ritorna l'array risultante
func XorBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("a e b devono avere la stessa lunghezza")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res, nil
}

// ReplaceBytes
// prende il puntatore a un array di byte nel quale sostituire i valori
// l'indice di partenza
// i valori da sostituire
// restituisce:
// un intero che corrisponde al numero di byte scritti
// un errore se l'indice di partenza è minore di 0 o se l'indice di partenza + la lunghezza dei valori da sostituire è maggiore della lunghezza dell'array
func ReplaceBytes(data *[]byte, startIdx int, values ...byte) (int, error) {
	var str string
	if startIdx < 0 {
		str = fmt.Sprintf("strartIdx è minore di 0")
		return 0, fmt.Errorf(str)
	}
	if startIdx+len(values) > len(*data) {
		str = fmt.Sprintf("strartIdx + len(values) è maggiore della lunghezza di data")
		return 0, fmt.Errorf(str)
	}
	i := 0
	v := byte(0)
	for _, v = range values {
		(*data)[startIdx+i] = v
		i++
	}
	return i, nil
}

// func IsAllZero
// prende in ingresso un array di byte
// restituisce true se tutti i byte sono uguali a 0
func IsAllZero(data []byte) bool {
	for _, b := range data {
		if b != 0 {
			return false
		}
	}
	return true
}

// func ConcatBytes
// prende in ingresso "outData" un puntatore a un array di byte
// prende in ingresso "values" un puntatore array di array di byte
func ConcatArrBytes(outData *[]byte, values ...*[][]byte) {
	for _, arr := range values {
		for _, v := range *arr {
			*outData = append(*outData, v...)
		}
	}
}

// ReplaceArrayBytes
// prende in ingresso un puntatore a un array di byte nel quale sostituire i valori (di dimensione sufficiente, oltre startIdx, per contenere i valori da sostituire)
// prende in ingresso un array di array di byte da sostituire
// prende in ingresso la lunghezza totale dei byte da sostituire
// prende in ingresso l'indice di partenza
func ReplaceArrayBytes(output *[]byte, values [][]byte, lenBytesToReplace int, startIdx int) (int, error) {
	// variabile tenere traccia dei byte sostituiti
	replacedBytes := 0
	// ciclo per la lunghezza dei dati
	for i, bytes := range values {
		temp, err := ReplaceBytes(output, startIdx+replacedBytes, bytes...)
		if err != nil {
			str := fmt.Sprintf("ReplaceArrayBytes() -> %d: %s", i, err.Error())
			replacedBytes += temp
			return replacedBytes, fmt.Errorf(str)
		}
		replacedBytes += temp
	}
	// controllo che i byte sostituiti siano uguali alla lunghezza totale dei byte da sostituire
	if replacedBytes != lenBytesToReplace {
		str := fmt.Sprintf("ReplaceArrayBytes() -> replacedBytes != lenBytesToReplace: %d != %d", replacedBytes, lenBytesToReplace)
		return replacedBytes, fmt.Errorf(str)
	}
	return replacedBytes, nil
}

func PackLittleEndian(data []byte) uint64 {
	packed := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		packed[i] = data[len(data)-1-i]
	}
	return binary.LittleEndian.Uint64(packed)
}

func PackLittleEndianBlocks(data []byte) []uint64 {
	var resBlocks []uint64
	for i := 0; i < len(data); i += 8 {
		block := data[i:min(i+8, len(data))]
		resBlocks = append(resBlocks, PackLittleEndian(block))
	}
	return resBlocks
}

func PackBigEndian(data []byte) uint64 {
	return binary.BigEndian.Uint64(data)
}

func PackBigEndianBlocks(data []byte) []uint64 {
	var resBlocks []uint64
	for i := 0; i < len(data); i += 8 {
		block := data[i:min(i+8, len(data))]
		resBlocks = append(resBlocks, PackBigEndian(block))
	}
	return resBlocks
}

func UnpackLittleEndian(data uint64, size int) []byte {
	unpacked := make([]byte, size)
	binary.LittleEndian.PutUint64(unpacked, data)
	return unpacked
}

func UnpackLittleEndianBlocks(data []uint64, totalLen int) []byte {
	var res []byte
	for i := 0; i < len(data); i += 8 {
		block := UnpackLittleEndian(data[i], min(8, totalLen-i))
		res = append(res, block...)
	}
	return res
}

func UnpackBigEndian(data uint64, size int) []byte {
	unpacked := make([]byte, size)
	binary.BigEndian.PutUint64(unpacked, data)
	return unpacked
}

func UnpackBigEndianBlocks(data []uint64, totalLen int) []byte {
	var res []byte
	for i := 0; i < len(data); i += 8 {
		block := UnpackBigEndian(data[i], min(8, totalLen-i))
		res = append(res, block...)
	}
	return res
}

/*
func SanitizeUint64Array(dataArray []uint64, startIdx int, length int) []uint64 {
	if length == 0 {
		fmt.Println("length is 0")
		return []uint64{}
	}
	inputBytesLen := len(dataArray) * 8
	if inputBytesLen == 0 {
		fmt.Println("inputBytesLen is 0")
		return []uint64{}
	}

	// variabile con cui calcolare il numero di bit da shiftare
	leftShiftBits := 0

	res, err := CropUint64Array(dataArray, startIdx, length)
	if err != nil {
		fmt.Println(err)
		return []uint64{}
	}

	return res
}
*/

func ShiftUint64Array(inputArray []uint64, startIdx int) ([]uint64, error) {
	fmt.Println("ShiftUint64Array")
	var temp uint64
	var err error
	var overflow uint64
	var res []uint64

	// ciclo per shiftare di un 1 bit a sinistra tutti gli elementi dell'array
	// partendo dall'ultimo
	for i := len(inputArray) - 1; i >= 0; i-- {
		fmt.Println("\t i = ", i)
		temp, err = ShiftUint64(inputArray[i], startIdx, &overflow)
		if err != nil {
			fmt.Println(err)
		}
		// aggiungo il risultato all'inizio dell'array
		res = append([]uint64{temp}, res...)
	}
	return res, nil
}

func ExtractBitsFromUint64Array(dataArray []uint64, startIdx int, length int) ([]uint64, error) {
	fmt.Println("ExtractBitsFromUint64Array")
	if length == 0 {
		fmt.Println("length is 0")
		return []uint64{}, nil
	}
	inputBytesLen := len(dataArray) * 8
	if inputBytesLen == 0 {
		fmt.Println("inputBytesLen is 0")
		return []uint64{}, nil
	}
	if startIdx >= inputBytesLen*8 {
		fmt.Println("startIdx >= inputBytesLen * 8")
		return []uint64{}, nil
	}
	if startIdx+length > inputBytesLen*8 {
		fmt.Println("startIdx + length > inputBytesLen * 8")
		return []uint64{}, nil
	}
	leftShiftBits := 0
	res, err := CropUint64Array(dataArray, startIdx, length, &leftShiftBits)
	if err != nil {
		fmt.Println(err)
		return []uint64{}, err
	}

	finalLeftShift := startIdx - leftShiftBits
	fmt.Println("\tstartIdx: ", startIdx)
	fmt.Println("\tleftShiftBits: ", leftShiftBits)
	fmt.Println("\tfinalLeftShift: ", finalLeftShift)

	res, err = ShiftUint64Array(res, finalLeftShift)
	if err != nil {
		fmt.Println(err)
		return []uint64{}, err
	}

	return res, nil
}

func ConvertUint64ArrayToBytesArray(dataArray []uint64) []byte {
	res := make([]byte, len(dataArray)*8)
	for i := 0; i < len(dataArray); i++ {
		binary.LittleEndian.PutUint64(res[i*8:], dataArray[i])
	}
	return res
}

func CropUint64Array(dataArray []uint64, startIdx int, length int, bitLeftCrop *int) ([]uint64, error) {
	fmt.Println("CropUint64Array")

	if length == 0 {
		fmt.Println("length is 0")
		return []uint64{}, nil
	}
	inputBytesLen := len(dataArray) * 8
	if inputBytesLen == 0 {
		fmt.Println("inputBytesLen is 0")
		return []uint64{}, nil
	}

	startLen := len(dataArray)
	res := fmt.Sprintf("len(dataArray): %d", startLen)
	fmt.Println(res)

	StartUint64Idx, EndUint64Idx := CalcBounds(startIdx, length)
	if EndUint64Idx > startLen {
		res = fmt.Sprintf("EndUint64Idx > startLen")
		fmt.Println(res)
		return nil, fmt.Errorf(res)
	}
	crop := dataArray[StartUint64Idx:EndUint64Idx]
	res = fmt.Sprintf("len(crop): %d", len(crop))
	fmt.Println(res)

	return dataArray[StartUint64Idx:EndUint64Idx], nil
}

func CalcBounds(startIdx int, length int) (int, int) {
	StartUint64Idx := CalcUint64Idx(startIdx)
	sum := startIdx + length
	str := fmt.Sprintf("startIdx: %d, length: %d, sum: %d", startIdx, length, sum)
	fmt.Println(str)
	EndUint64Idx := CalcUint64Idx(sum-1) + 1

	res := fmt.Sprintf("StartUint64Idx: %d, EndUint64Idx: %d", StartUint64Idx, EndUint64Idx)
	fmt.Println(res)
	return StartUint64Idx, EndUint64Idx
}

func CalcUint64Idx(bitIdx int) int {
	res := int(math.Floor(float64(bitIdx) / 64))
	return res
}

func PrintUint64ArrayAsBinary(dataArray ...uint64) {
	for i := 0; i < len(dataArray); i++ {
		fmt.Printf("%064b\t", dataArray[i])
	}
	fmt.Println()
}

// Func SanitizeUint64
// Prende in input:
// startUint64: un uint64
// numBitsToWrite: il numero di bit da scrivere
// Shifta a destra di 64 - numBitsToWrite e a sinistra di 64 - numBitsToWrite per assicurarsi che a destra dell'ultimo bit da scrivere ci siano solo zeri
func SanitizeUint64(startUint64 uint64, numBitsToWrite uint8) (uint64, error) {
	// calcolo il numero di bit da shiftare
	shifts := uint64(64 - numBitsToWrite)
	// shifto a destra di shifts
	res := startUint64 >> shifts
	// shifto a sinistra di shifts
	// in questo modo a destra dell'ultimo bit da scrivere ci saranno solo zeri
	res <<= shifts
	return res, nil
}

// ShiftUint64
// Prende in input:
// startUint (uint64)
// startIdx (int) l'indice di partenza
// overflow (un *uint64) il puntatore a un uint64 che conterrà nei primi bit alla parte shiftata di startUint
// Ritorna:
// un uint64 che contiene startUint shiftata a sinistra di startIdx unito a ai primi bit di overflow
// un errore se startIdx è minore di 0 o se startIdx + length è maggiore di 64
func ShiftUint64(startUint uint64, startIdx int, overflow *uint64) (uint64, error) {
	var err error
	if startIdx < 0 {
		return 0, fmt.Errorf("startIdx è minore di 0")
	}

	// assegno a tempCopy la parte che verrà shiftata
	tempCopy, err := RightShiftUint64(startUint, startIdx)
	if err != nil {
		return 0, err
	}

	// assegno a res la parte dopo lo shift sinistro
	tempRes, err := LeftShiftUint64(startUint, startIdx)
	if err != nil {
		return 0, err
	}

	// eseguo il join tra tempRes e overflow
	res := tempRes | *overflow

	// assegno a overflow tempCopy
	*overflow = tempCopy

	return res, nil
}

func RightShiftUint64(startUint64 uint64, lastIdx int) (uint64, error) {
	//fmt.Println("RightShiftUint64")
	// controllo che lastIdx sia minore di 64
	if lastIdx > 64 {
		return 0, fmt.Errorf("lastIdx > 64")
	}
	// calcolo il numero di shift da effetuare
	shifts := 64 - lastIdx
	// ritorno il risultato dello shift
	return startUint64 >> shifts, nil
}

func LeftShiftUint64(startUint64 uint64, startIdx int) (uint64, error) {
	//fmt.Println("LeftShiftUint64")
	// controllo che startIdx sia minore di 64
	if startIdx > 64 {
		return 0, fmt.Errorf("startIdx > 64")
	}
	// ritorno il risultato dello shift
	return startUint64 << startIdx, nil
}
