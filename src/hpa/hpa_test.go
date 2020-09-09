package main

import "testing"

func TestLooping(t *testing.T) {
	resultado := Looping()
	// resultado := "Code.education Rocks!a"

	if resultado != "Code.education Rocks!" {
		t.Errorf("Esperando: %s, obtido: %s", "Code.education Rocks!", resultado)
	}
}
