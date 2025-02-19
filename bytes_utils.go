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
