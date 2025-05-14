package main

import (
	"flag"
	"fmt"
	"github.com/Nebsu/tp1/annuaire"
	"os"
)

func main() {
	action := flag.String("action", "", "Action à réaliser (ajouter, rechercher, lister, supprimer, modifier)")
	name := flag.String("nom", "", "Nom de la personne")
	phone := flag.String("tel", "", "Numéro de téléphone")
	newPhone := flag.String("nouveau-tel", "", "Nouveau numéro de téléphone")

	flag.Parse()

	pb := annuaire.New()
	err := pb.LoadFromFile("contacts.json")
	if err != nil {
		os.Exit(1)
	}

	switch *action {
	case "ajouter":
		if *name == "" || *phone == "" {
			fmt.Println("Erreur: nom et numéro de téléphone sont requis pour ajouter un contact")
			os.Exit(1)
		}
		err := pb.AddContact(*name, *phone)
		if err != nil {
			fmt.Println("Erreur:", err)
			os.Exit(1)
		}
		fmt.Printf("Contact ajouté: %s - %s\n", *name, *phone)

	case "rechercher":
		if *name == "" {
			fmt.Println("Erreur: nom est requis pour rechercher un contact")
			os.Exit(1)
		}
		contact, found := pb.FindContact(*name)
		if !found {
			fmt.Printf("Aucun contact trouvé pour le nom: %s\n", *name)
			return
		}
		fmt.Printf("Contact trouvé: %s - %s\n", contact.Name, contact.Phone)

	case "lister":
		contacts := pb.ListContacts()
		if len(contacts) == 0 {
			fmt.Println("L'annuaire est vide")
			return
		}
		fmt.Println("Liste des contacts:")
		for _, contact := range contacts {
			fmt.Printf("- %s: %s\n", contact.Name, contact.Phone)
		}

	case "supprimer":
		if *name == "" {
			fmt.Println("Erreur: nom est requis pour supprimer un contact")
			os.Exit(1)
		}
		removed := pb.RemoveContact(*name)
		if !removed {
			fmt.Printf("Aucun contact trouvé pour le nom: %s\n", *name)
			return
		}
		fmt.Printf("Contact supprimé: %s\n", *name)

	case "modifier":
		if *name == "" || *newPhone == "" {
			fmt.Println("Erreur: nom et nouveau numéro sont requis pour modifier un contact")
			os.Exit(1)
		}
		modified := pb.UpdateContact(*name, *newPhone)
		if !modified {
			fmt.Printf("Aucun contact trouvé pour le nom: %s\n", *name)
			return
		}
		fmt.Printf("Contact modifié: %s - %s\n", *name, *newPhone)

	default:
		fmt.Println("Action non reconnue. Utilisez: ajouter, rechercher, lister, supprimer ou modifier")
		os.Exit(1)
	}
	err = pb.SaveToFile("contacts.json")
	if err != nil {
		return
	}
}
