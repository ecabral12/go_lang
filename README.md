# (TP1)

ELVIS
LHUKAS
DJAHNICK

### Actions

* **ajouter** : ajoute un contact

  ```bash
  --action ajouter --nom NOM --prenom PRENOM --tel TEL
  ```
* **rechercher** : affiche le numéro

  ```bash
  --action rechercher --nom NOM --prenom PRENOM
  ```
* **lister** : affiche tous les contacts

  ```bash
  --action lister
  ```
* **supprimer** : supprime un contact

  ```bash
  --action supprimer --nom NOM --prenom PRENOM
  ```

## Exemple

```bash
# ajouter
go run main.go --action ajouter --nom Dupont --prenom Alice --tel 0612345678
# lister
go run main.go --action lister
# rechercher
go run main.go --action rechercher --nom Dupont --prenom Alice
# supprimer
go run main.go --action supprimer --nom Dupont --prenom Alice
```

## Tests

Pour lancer les tests unitaires :

```bash
go test -v
```
