package utils

import (
	"encoding/binary"
	"fmt"
	"strings"
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

// IntelBytesToUint32
// Prende in input due array di byte
// controlla che abbiano la stessa lunghezza
// Esegue l'operazione di XOR tra i due array
// Ritorna l'array risultante
func XorIntelBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("a e b devono avere la stessa lunghezza")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res, nil
}

// XorIntelHex
// Prende in input due stringhe "esadecimali"
// controlla che abbiano la stessa lunghezza e che siano multipli di 2
// Converte le stringhe in array di byte
// Esegue l'operazione di XOR tra i due array
// Ritorna il risultato sotto forma di stringa "esadecimale"
func XorIntelHex(a, b string) (string, error) {
	// rimuovo gli spazi
	a = strings.ReplaceAll(a, " ", "")
	b = strings.ReplaceAll(b, " ", "")
	// mi assucuro che abbiano la stessa lunghezza e che siano multipli di 2
	if len(a) != len(b) || len(a)%2 != 0 || len(b)%2 != 0 {
		return "", fmt.Errorf("a e b devono avere la stessa lunghezza e devono essere multipli di 2")
	}
	// converto le stringhe in array di byte
	aBytes, err := HexToByteArray(a)
	if err != nil {
		return "", err
	}
	bBytes, err := HexToByteArray(b)
	if err != nil {
		return "", err
	}
	// eseguo l'operazione di XOR
	res, err := XorIntelBytes(aBytes, bBytes)
	if err != nil {
		return "", err
	}
	// converto il risultato in stringa
	return ByteArrayToHex(res), nil
}
