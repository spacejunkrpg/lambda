package main

import (
	"encoding/json"
	"testing"
)

func TestRollDicec(t *testing.T) {
	roll := rollDice(1)

	if roll < 1 || roll > 6 {
		t.Errorf("rollDice(1) = %d; want >=1 and <= 6", roll)
	}
}

func TestChooseSpecies(t *testing.T) {
	species := chooseSpecies(1)

	if species != "Kaa" {
		t.Errorf("chooseSpecies(1) = %s; want 'Kaa'", species)
	}
}

func TestChooseName(t *testing.T) {
	var name string
	name = chooseName("Terran", 1, 1)

	if name != "Amari Devi" {
		t.Errorf("chooseName('Terran', 1, 1) = %s; want 'Amari Devi'", name)
	}

	name = chooseName("Kaa", 2, 3)

	if name != "Ekene Nwakaeme" {
		t.Errorf("chooseName('Kaa', 2, 3) = %s; want 'Ekene Nwakaeme'", name)
	}

	name = chooseName("Myxos", 6, 1)
	if name != "Xenophon of Agrinio" {
		t.Errorf("chooseName('Myxos', 6, 1) = %s; want 'Xenophon of Agrinio'", name)
	}
}

func TestChoosePersonality(t *testing.T) {
	personality := choosePersonality(2, 3, 4, 5, 6)

	if personality != "Close-mminded, Sloppy, Reserved, Selfish, Optimistic" {
		t.Errorf("choosePersonality(2, 3, 4, 5, 6) = %s; want 'Close-mminded, Sloppy, Reserved, Selfish, Optimistic'", personality)
	}
}

func TestChooseProfession(t *testing.T) {
	profession := chooseProfession(3)

	if profession != "Scientist" {
		t.Errorf("chooseProfession(3) = %s; want 'Scientist'", profession)
	}
}

func TestChooseMotivation(t *testing.T) {
	motivation := chooseMotivation("Scientist", 4)

	if motivation != "Translate a new language" {
		t.Errorf("chooseMotivation('Scientist', 4) = %s; want 'Translate a new language'", motivation)
	}
}

func TestChooseAttitude(t *testing.T) {
	attitude := chooseAttitude(2)

	if attitude != "Neutral" {
		t.Errorf("chooseAttitude(2) = %s; want 'Neutral'", attitude)
	}
}

func TestGenerateCharacter(t *testing.T) {
	j := GenerateCharacter(1642211868225597000)

	var c Character
	err := json.Unmarshal(j, &c)

	if err != nil {
		t.Errorf("unable to parse JSON response")
	}

}
