package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invalido: Resulatdo %d. Esperado: %d.", total, 30)
	}
}

func TestSoma2(t *testing.T) {
	total := Soma(1, 1)

	if total != 2 {
		t.Errorf("Resultado da soma é invalido: Resulatdo %d. Esperado: %d.", total, 30)
	}
}
