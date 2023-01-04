package entity

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

var characterxiaogong character = character{
	id:          1,
	affiliation: enum.AffiliationInazuma,
	vision:      enum.ElementPyro,
	weapon:      enum.WeaponBow,
	maxHP:       10,
	currentHP:   10,
	status:      enum.CharacterStatusActive,
}
