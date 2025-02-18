package utils

import (
	"fmt"
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
