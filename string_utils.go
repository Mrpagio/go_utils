package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// LeftPad
// prende in input una stringa
// la lunghezza della stringa finale
// il padding da aggiungere
// restituisce la stringa finale
// se la stringa è già più lunga della lunghezza richiesta restituisce un errore
// se la differenza tra lunghezza richiesta e lunghezza stringa non è divisibile per la lunghezza del padding restituisce un errore
func LeftPad(str string, length int, pad string) (string, error) {
	// Controllo se la stringa è già più lunga della lunghezza richiesta
	if len(str) > length {
		return str, fmt.Errorf("Errore: la stringa è già più lunga di %d caratteri", length)
	}
	diffLen := length - len(str)
	// Controllo che la differenza tra lunghezza richiesta e lunghezza stringa sia divisibile per la lunghezza del padding
	if diffLen%len(pad) != 0 {
		return str, fmt.Errorf("Errore: la differenza tra lunghezza richiesta e lunghezza stringa non è divisibile per la lunghezza del padding")
	}
	padding := strings.Repeat(pad, diffLen/len(pad))
	res := padding + str
	return res, nil
}

// IntLeftPad
// prende in input un intero
// la lunghezza della stringa finale
// il padding da aggiungere
// restituisce la stringa finale
func IntLeftPad(val int, length int, pad string) (string, error) {
	// Converto l'intero in stringa
	t := strconv.Itoa(val)
	// Chiamo la funzione LeftPad
	res, err := LeftPad(t, length, pad)
	if err != nil {
		return "", err
	}
	return res, nil
}

// XorHex
// Prende in input due stringhe "esadecimali"
// controlla che abbiano la stessa lunghezza e che siano multipli di 2
// Converte le stringhe in array di byte
// Esegue l'operazione di XOR tra i due array
// Ritorna il risultato sotto forma di stringa "esadecimale"
func XorHex(a, b string) (string, error) {
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
	res, err := XorBytes(aBytes, bBytes)
	if err != nil {
		return "", err
	}
	// converto il risultato in stringa
	return BytesToHex(res), nil
}
