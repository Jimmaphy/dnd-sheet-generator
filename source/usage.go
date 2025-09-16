package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Printf(`Usage:
		%s create -name CHARACTER_NAME -race RACE -class CLASS -level N -str N -dex N -con N -int N -wis N -cha N
		%s view -name CHARACTER_NAME
		%s list
		%s delete -name CHARACTER_NAME
		%s equip -name CHARACTER_NAME -weapon WEAPON_NAME -slot SLOT
		%s equip -name CHARACTER_NAME -armor ARMOR_NAME
		%s equip -name CHARACTER_NAME -shield SHIELD_NAME
		%s learn-spell -name CHARACTER_NAME -spell SPELL_NAME
		%s prepare-spell -name CHARACTER_NAME -spell SPELL_NAME 
		`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}
