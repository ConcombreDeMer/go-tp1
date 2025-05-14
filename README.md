Membres du groupe :
- Yanis Rozier
- Sebastien Zhou

Liste des actions:
`go run main.go --action="<action>"`
- `ajouter` : Ajouter un utilisateur
- `supprimer` : Supprimer un utilisateur
- `lister` : Lister les utilisateurs
- `modifier` : Mettre à jour un utilisateur
- `rechercher` : Obtenir un utilisateur

Ajouter un utilisateur:
- `go run main.go --action="ajouter" --nom="<nom>" --tel="<numero de tel>"`

Supprimer un utilisateur:
- `go run main.go --action="supprimer" --nom="<nom>"`

Lister les utilisateurs:
- `go run main.go --action="lister"`

Modifier un utilisateur:
- `go run main.go --action="modifier" --nom="<nom>" --tel="<nouveau numero de tel>"`

Rechercher un utilisateur:
- `go run main.go --action="rechercher" --nom="<nom>"`