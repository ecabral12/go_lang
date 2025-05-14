// main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Personne struct {
	Nom    string
	Prenom string
	Tel    string
}

func main() {
	// 1) Déclaration des flags
	action := flag.String("action", "", "ajouter/rechercher/lister/supprimer")
	nom := flag.String("nom", "", "Nom")
	prenom := flag.String("prenom", "", "Prénom")
	tel := flag.String("tel", "", "Téléphone")
	flag.Parse()

	const fileName = "contacts.txt"
	contacts := make(map[string]Personne)

	// 2) Chargement depuis le fichier (s'il existe)
	if f, err := os.Open(fileName); err == nil {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			parts := strings.Fields(scanner.Text())
			if len(parts) == 3 {
				key := parts[0] + " " + parts[1]
				contacts[key] = Personne{parts[0], parts[1], parts[2]}
			}
		}
		f.Close()
	}

	key := *nom + " " + *prenom
	modified := false

	// 3) Traitement de l'action
	switch *action {
	case "ajouter":
		if *nom == "" || *prenom == "" || *tel == "" {
			fmt.Println("Usage: --action ajouter --nom NOM --prenom PRENOM --tel TEL")
			return
		}
		if _, exists := contacts[key]; exists {
			fmt.Println("Ce contact existe déjà")
			return
		}
		contacts[key] = Personne{*nom, *prenom, *tel}
		modified = true
		fmt.Println("Contact ajouté")

	case "rechercher":
		if *nom == "" || *prenom == "" {
			fmt.Println("Usage: --action rechercher --nom NOM --prenom PRENOM")
			return
		}
		if p, found := contacts[key]; found {
			fmt.Printf("%s %s : %s\n", p.Nom, p.Prenom, p.Tel)
		} else {
			fmt.Println("Aucun contact trouvé")
		}

	case "lister":
		if len(contacts) == 0 {
			fmt.Println("Annuaire vide")
			return
		}
		for _, p := range contacts {
			fmt.Printf("- %s %s : %s\n", p.Nom, p.Prenom, p.Tel)
		}

	case "supprimer":
		if *nom == "" || *prenom == "" {
			fmt.Println("Usage: --action supprimer --nom NOM --prenom PRENOM")
			return
		}
		if _, exists := contacts[key]; !exists {
			fmt.Println("Contact non trouvé")
			return
		}
		delete(contacts, key)
		modified = true
		fmt.Println("Contact supprimé")

	default:
		fmt.Println("Action invalide : ajouter, rechercher, lister, supprimer")
		return
	}

	// 4) Si on a ajouté ou supprimé, on réécrit le fichier
	if modified {
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Erreur écriture :", err)
			return
		}
		for _, p := range contacts {
			fmt.Fprintln(f, p.Nom, p.Prenom, p.Tel)
		}
		f.Close()
	}
}
