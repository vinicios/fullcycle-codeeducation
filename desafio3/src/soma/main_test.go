package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := soma(5, 5)
	if resultado != 10 {
		t.Errorf("Ocoreu um erro, valor esperado era 10, mas resultado foi %v", resultado)
	}
}
