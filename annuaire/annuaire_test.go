package annuaire

import (
	"os"
	"testing"
)

func TestAddContact(t *testing.T) {
	pb := New()

	err := pb.AddContact("Seb", "0123456789")
	if err != nil {
		t.Errorf("Erreur lors de l'ajout d'un contact: %v", err)
	}
	contact, found := pb.FindContact("Seb")
	if !found {
		t.Error("Le contact ajouté n'a pas été trouvé")
	}
	if contact.Phone != "0123456789" {
		t.Errorf("Le numéro de téléphone attendu était 0123456789, mais on a obtenu %s", contact.Phone)
	}
	err = pb.AddContact("Seb", "9876543210")
	if err == nil {
		t.Error("L'ajout d'un contact avec un nom déjà existant devrait échouer")
	}
}

func TestFindContact(t *testing.T) {
	pb := New()
	pb.AddContact("Seb", "0987654321")
	contact, found := pb.FindContact("Seb")
	if !found {
		t.Error("Contact existant non trouvé")
	}
	if contact.Name != "Seb" || contact.Phone != "0987654321" {
		t.Errorf("Contact incorrect trouvé: %v", contact)
	}
	_, found = pb.FindContact("Nico")
	if found {
		t.Error("Un contact inexistant a été trouvé")
	}
}

func TestListContacts(t *testing.T) {
	pb := New()
	contacts := pb.ListContacts()
	if len(contacts) != 0 {
		t.Errorf("L'annuaire vide devrait retourner 0 contacts, mais a retourné %d", len(contacts))
	}
	pb.AddContact("Seb", "0123456789")
	pb.AddContact("Yanis", "0987654321")
	contacts = pb.ListContacts()
	if len(contacts) != 2 {
		t.Errorf("L'annuaire devrait contenir 2 contacts, mais en contient %d", len(contacts))
	}
}

func TestRemoveContact(t *testing.T) {
	pb := New()
	pb.AddContact("Yanis", "0567891234")
	removed := pb.RemoveContact("Yanis")
	if !removed {
		t.Error("La suppression d'un contact existant a échoué")
	}
	_, found := pb.FindContact("Yanis")
	if found {
		t.Error("Le contact supprimé a été trouvé")
	}
	removed = pb.RemoveContact("Nico")
	if removed {
		t.Error("La suppression d'un contact inexistant a réussi")
	}
}

func TestUpdateContact(t *testing.T) {
	pb := New()
	pb.AddContact("Seb", "0654321098")
	updated := pb.UpdateContact("Seb", "0123456789")
	if !updated {
		t.Error("La modification d'un contact existant a échoué")
	}
	contact, found := pb.FindContact("Seb")
	if !found {
		t.Error("Le contact modifié n'a pas été trouvé")
	}
	if contact.Phone != "0123456789" {
		t.Errorf("Le numéro mis à jour devrait être 0123456789, mais est %s", contact.Phone)
	}
	updated = pb.UpdateContact("Nico", "1111111111")
	if updated {
		t.Error("La modification d'un contact inexistant a réussi")
	}
}

func TestSaveAndLoadFromFile(t *testing.T) {
	filename := "test_contacts.json"
	os.Remove(filename)
	pb1 := New()
	pb1.AddContact("Seb", "0123456789")
	pb1.AddContact("Yanis", "0987654321")
	err := pb1.SaveToFile(filename)
	if err != nil {
		t.Errorf("Erreur lors de la sauvegarde de l'annuaire: %v", err)
	}
	pb2 := New()
	err = pb2.LoadFromFile(filename)
	if err != nil {
		t.Errorf("Erreur lors du chargement de l'annuaire: %v", err)
	}
	contacts := pb2.ListContacts()
	if len(contacts) != 2 {
		t.Errorf("L'annuaire chargé devrait contenir 2 contacts, mais en contient %d", len(contacts))
	}
	os.Remove(filename)
}
