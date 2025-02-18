package utils

import (
	"encoding/binary"
	"fmt"
)

// IntelExpandOutputData
// Prende un byte (8 bit)
// lo espande in un uint32 con 24 zeri a destra (little endian o Intel)
// ritorna il uint32
func IntelExpandOutputData(currentByte byte) uint32 {
	// creo un uint32
	outUint32 := binary.LittleEndian.Uint32([]byte{currentByte, 0, 0, 0})
	return outUint32
}

// SanitizeUint32Data
// Prende in input:
// encodedData: un array di byte
// numBitsToWrite: il numero di bit da scrivere
// Converte encodedData in un uint32 (little endian o Intel)
// Shifta a destra di 32 - numBitsToWrite e a sinistra di 32 - numBitsToWrite per assicurarsi che a destra dell'ultimo bit da scrivere ci siano solo zeri
func SanitizeUint32Data(encodedData []byte, numBitsToWrite uint8) (uint32, error) {
	if len(encodedData) == 0 {
		return 0, fmt.Errorf("encodedData è vuoto")
	}
	// converto encodedData in un uint32(intel)
	startUint := binary.LittleEndian.Uint32(encodedData)

	// calcolo il numero di bit da shiftare
	shifts := uint32(32 - numBitsToWrite)
	// shifto a destra di shifts
	res := startUint >> shifts
	// shifto a sinistra di shifts
	// in questo modo a destra dell'ultimo bit da scrivere ci saranno solo zeri
	res <<= shifts
	return res, nil
}
