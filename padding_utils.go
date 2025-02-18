package main

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
