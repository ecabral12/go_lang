// main_test.go
package main

import "testing"

func TestAjoutContact(t *testing.T) {
	m := make(map[string]string)
	key := "Dupont Alice"
	m[key] = "0612345678"

	if got := m[key]; got != "0612345678" {
		t.Errorf("Téléphone attendu «0612345678», obtenu «%s»", got)
	}
}

func TestSuppressionContact(t *testing.T) {
	key := "Dupont Alice"
	m := map[string]string{
		key: "0612345678",
	}

	delete(m, key)
	if _, ok := m[key]; ok {
		t.Errorf("Après delete, la clé %q devrait avoir disparu", key)
	}
}

func TestTailleMap(t *testing.T) {
	m := map[string]string{
		"Dupont Alice": "0612345678",
		"Martin Bob":   "0622334455",
	}
	if want, got := 2, len(m); got != want {
		t.Errorf("Taille de la map attendue = %d, obtenue = %d", want, got)
	}
}
