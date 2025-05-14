package annuaire

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Annuaire struct {
	Contacts []Contact `json:"contacts"`
}

func New() *Annuaire {
	return &Annuaire{
		Contacts: []Contact{},
	}
}

func (pb *Annuaire) AddContact(name, phone string) error {
	for _, contact := range pb.Contacts {
		if contact.Name == name {
			return errors.New("un contact avec ce nom existe déjà")
		}
	}
	pb.Contacts = append(pb.Contacts, Contact{
		Name:  name,
		Phone: phone,
	})

	return nil
}

func (pb *Annuaire) FindContact(name string) (Contact, bool) {
	for _, contact := range pb.Contacts {
		if contact.Name == name {
			return contact, true
		}
	}
	return Contact{}, false
}

func (pb *Annuaire) ListContacts() []Contact {
	return pb.Contacts
}

func (pb *Annuaire) RemoveContact(name string) bool {
	contactFound := false
	var newContactsList []Contact
	for _, contact := range pb.Contacts {
		if contact.Name != name {
			newContactsList = append(newContactsList, contact)
		} else {
			contactFound = true
		}
	}
	if contactFound {
		pb.Contacts = newContactsList
		return true
	}

	return false
}

func (pb *Annuaire) UpdateContact(name, newPhone string) bool {
	for i, contact := range pb.Contacts {
		if contact.Name == name {
			pb.Contacts[i].Phone = newPhone
			return true
		}
	}
	return false
}

func (pb *Annuaire) SaveToFile(filename string) error {
	data, err := json.Marshal(pb)
	if err != nil {
		return fmt.Errorf("erreur lors de la traduction de l'annuaire: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier: %w", err)
	}

	return nil
}

func (pb *Annuaire) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier: %w", err)
	}

	err = json.Unmarshal(data, pb)
	if err != nil {
		return fmt.Errorf("erreur lors de la traduction de l'annuaire: %w", err)
	}
	return nil
}
