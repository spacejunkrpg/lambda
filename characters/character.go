package characters

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Character struct {
	Seed        int    `json:"seed"`
	Name        string `json:"name"`
	Species     string `json:"species"`
	Profession  string `json:"profession"`
	Personality string `json:"personality"`
	Motivation  string `json:"motivation"`
	Attitude    string `json:"attitude"`
}

func rollDice(x int) int {
	sum := 0

	for i := 1; i <= x; i++ {
		sum += rand.Int()%6 + 1
	}

	return sum
}

func chooseSpecies(roll int) string {
	m := map[int]string{
		1: "Kaa",
		2: "Kaa",
		3: "Terran",
		4: "Terran",
		5: "Myxos",
		6: "Myxos",
	}

	return m[roll]
}

func chooseName(species string, givenRoll int, familyRoll int) string {
	var g map[int]string
	var f map[int]string

	switch species {
	case "Terran":
		g = map[int]string{
			1: "Amari",
			2: "Darra",
			3: "Fenix",
			4: "Nuru",
			5: "Omid",
			6: "Valo",
		}

		f = map[int]string{
			1: "Devi",
			2: "Hernandez",
			3: "Ivanova",
			4: "Mwangi",
			5: "Smith",
			6: "Wang",
		}
	case "Kaa":
		g = map[int]string{
			1: "Chike",
			2: "Ekene",
			3: "Ndidi",
			4: "Ngozi",
			5: "Oluchi",
			6: "Uzoma",
		}

		f = map[int]string{
			1: "Abagun",
			2: "Bakare",
			3: "Nwakaeme",
			4: "Ogunode",
			5: "Osakwe",
			6: "Umelo",
		}
	case "Myxos":
		g = map[int]string{
			1: "Aegus",
			2: "Clio",
			3: "Despina",
			4: "Ianthe",
			5: "Selene",
			6: "Xenophon",
		}

		f = map[int]string{
			1: "Agrinio",
			2: "Ilio",
			3: "Kavala",
			4: "Smyrna",
			5: "Volos",
			6: "Zografou",
		}
	}

	given := g[givenRoll]
	family := f[familyRoll]
	separator := " "
	if species == "Myxos" {
		separator = " of "
	}
	return given + separator + family
}

func choosePersonality(rollOpen int, rollCon int, rollExtra int, rollAgree int, rollStability int) string {

	o := map[int]string{
		2:  "Close-mminded",
		3:  "Literal",
		4:  "Rigid",
		5:  "Analytical",
		6:  "Knowledgable",
		7:  "Smart",
		8:  "Curious",
		9:  "Imaginative",
		10: "Artistic",
		11: "Creative",
		12: "Adventurous",
	}

	open := o[rollOpen]

	c := map[int]string{
		2:  "Irresponsible",
		3:  "Sloppy",
		4:  "Disorganized",
		5:  "Spontaneous",
		6:  "Effective",
		7:  "Practical",
		8:  "Reliable",
		9:  "Responsible",
		10: "Careful",
		11: "Cautious",
		12: "Serious",
	}

	con := c[rollCon]

	e := map[int]string{
		2:  "Solitary",
		3:  "Shy",
		4:  "Reserved",
		5:  "Quiet",
		6:  "Sociable",
		7:  "Friendly",
		8:  "Energetic",
		9:  "Outgoing",
		10: "Assertive",
		11: "Noisy",
		12: "Gregarious",
	}

	extra := e[rollExtra]

	a := map[int]string{
		2:  "Manipulative",
		3:  "Aggressive",
		4:  "Callous",
		5:  "Selfish",
		6:  "Gruff",
		7:  "Hospitable",
		8:  "Polite",
		9:  "Kind",
		10: "Thoughtful",
		11: "Helpful",
		12: "Generous",
	}

	agree := a[rollAgree]

	s := map[int]string{
		2:  "Carefree",
		3:  "Calm",
		4:  "Resilient",
		5:  "Centered",
		6:  "Optimistic",
		7:  "Realistic",
		8:  "Pessimistic",
		9:  "Nervous",
		10: "Jumpy",
		11: "Emotional",
		12: "Volatile",
	}

	stability := s[rollStability]

	return open + ", " + con + ", " + extra + ", " + agree + ", " + stability

}

func chooseProfession(roll int) string {
	m := map[int]string{
		2:  "Scientist",
		3:  "Scientist",
		4:  "Merchant",
		5:  "Merchant",
		6:  "Criminal",
		7:  "Junker",
		8:  "Mercenary",
		9:  "Surveyor",
		10: "Surveyor",
		11: "Diplomat",
		12: "Diplomat",
	}

	return m[roll]
}

func chooseMotivation(profession string, roll int) string {
	var m map[int]string

	switch profession {
	case "Scientist":
		m = map[int]string{
			1: "Research a new species",
			2: "Lead an engineering project",
			3: "Develop a biological enhancement",
			4: "Translate a new language",
			5: "Test a nanotech device",
			6: "Investigate a gravitational anomaly",
		}
	case "Merchant":
		m = map[int]string{
			1: "Buy a swarm of nanobots that are used for terraforming",
			2: "Sell a pair of overly chatty, refurbished construction robots",
			3: "Deliver a load of fuel to a stranded spacecraft",
			4: "Sell food and water to a recently established colony",
			5: "Deliver medicine to treat a deadly pandemic",
			6: "Buy a crate of weapons for an army fighting a civil war",
		}
	case "Criminal":
		m = map[int]string{
			1: "Deliver a shipment of stolen goods",
			2: "Collect protection payments",
			3: "Run a gambling ring",
			4: "Pick up a shipment of illegal drugs",
			5: "Train new recruits",
			6: "Break some kneecaps",
		}
	case "Junker":
		m = map[int]string{
			1: "Biological",
			2: "Construction",
			3: "Medical",
			4: "Military",
			5: "Mineral",
			6: "Technology",
		}
	case "Mercenary":
		m = map[int]string{
			1: "Quell a rebellion on a recently colonized moon",
			2: "Provide security for a very important creature",
			3: "Train the troops of a developing society",
			4: "Protect a mining colony from dangerous native wildlife",
			5: "Provide humanitarian assistance after a natural disaster",
			6: "Enforce a cease-fire between two hostile forces",
		}
	case "Surveyor":
		m = map[int]string{
			1: "Assess the repairability of a salvage site",
			2: "Measure the boundaries of a construction site",
			3: "Assess the mining potential of an asteroid belt",
			4: "Adjudicate a dispute on a recently colonized moon",
			5: "Map a trade route between distant planetary systems",
			6: "Scout ideal locations for terraforming equipment",
		}
	case "Diplomat":
		m = map[int]string{
			1: "Negotiate a treaty between regional powers",
			2: "Review the terms of a trade agreement",
			3: "Deliver a gift to an allied power",
			4: "Exchange prisoners with a rival power",
			5: "Speak at a galactic summit",
			6: "Visit an ambassador",
		}
	}

	return m[roll]
}

func chooseAttitude(roll int) string {
	m := map[int]string{
		1: "Friendly",
		2: "Neutral",
		3: "Neutral",
		4: "Neutral",
		5: "Neutral",
		6: "Hostile",
	}

	return m[roll]
}

func GenerateCharacter(seed int64) {
	rand.Seed(seed)
	c := Character{}
	c.Seed = int(seed)

	c.Species = chooseSpecies(rollDice(1))
	c.Name = chooseName(c.Species, rollDice(1), rollDice(1))
	c.Personality = choosePersonality(rollDice(2), rollDice(2), rollDice(2), rollDice(2), rollDice(2))

	c.Profession = chooseProfession(rollDice(2))
	c.Motivation = chooseMotivation(c.Profession, rollDice(1))

	c.Attitude = chooseAttitude(rollDice(1))

	j, _ := json.Marshal(c)
	fmt.Println(string(j))
}
