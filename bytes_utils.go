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
// restituisce un errore se l'indice di partenza è minore di 0 o se l'indice di partenza + la lunghezza dei valori da sostituire è maggiore della lunghezza dell'array
func ReplaceBytes(data *[]byte, strartIdx int, values ...byte) error {
	var str string
	if strartIdx < 0 {
		str = fmt.Sprintf("strartIdx è minore di 0")
		return fmt.Errorf(str)
	}
	if strartIdx+len(values) > len(*data) {
		str = fmt.Sprintf("strartIdx + len(values) è maggiore della lunghezza di data")
		return fmt.Errorf(str)
	}
	for i, v := range values {
		(*data)[strartIdx+i] = v
	}
	return nil
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
