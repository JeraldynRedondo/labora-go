package model

import (
	"interfacesGame/fighters"
)

var Police fighters.Police = fighters.Police{
	BaseFighter: fighters.BaseFighter{
		Life: 10, InitialLife: 10,
	},
	Armour: 5,
}

var Criminal fighters.Criminal = fighters.Criminal{
	BaseFighter: fighters.BaseFighter{
		Life: 10, InitialLife: 10,
	},
}

var Paladin fighters.Paladin = fighters.Paladin{
	BaseFighter: fighters.BaseFighter{
		Life: 10, InitialLife: 10,
	},
	Armour: 6,
}
