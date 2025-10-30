Usage:
  create -name CHARACTER_NAME -race RACE -class CLASS -level N -str N -dex N -con N -int N -wis N -cha N
  view -name CHARACTER_NAME
  list
  delete -name CHARACTER_NAME
  equip -name CHARACTER_NAME -weapon WEAPON_NAME -slot SLOT
  equip -name CHARACTER_NAME -armor ARMOR_NAME
  equip -name CHARACTER_NAME -shield SHIELD_NAME
  learn-spell -name CHARACTER_NAME -spell SPELL_NAME
  prepare-spell -name CHARACTER_NAME -spell SPELL_NAME
  export -name CHARACTER_NAME
  update-data



CREATE
---------------------------------------------------------------------
Create a new character with specified attributes.
Save the character to a JSON file named CHARACTER_NAME.json.
		 
-name CHARACTER_NAME   Name of the character (string)
-race RACE             Race of the character (string)
-class CLASS           Class of the character (string)
-level N [optional]    Level of the character (integer, default: 1)
-str N [optional]      Strength attribute (integer, default: 10)
-dex N [optional]      Dexterity attribute (integer, default: 10)
-con N [optional]      Constitution attribute (integer, default: 10)
-int N [optional]      Intelligence attribute (integer, default: 10)
-wis N [optional]      Wisdom attribute (integer, default: 10)
-cha N [optional]      Charisma attribute (integer, default: 10)


VIEW
---------------------------------------------------------------------
View the details of a character in a readable format.
The output is directed to the console.
		 
-name CHARACTER_NAME   Name of the character to view (string)


LIST
---------------------------------------------------------------------
List all existing characters by their names.
The output is directed to the console.


DELETE
---------------------------------------------------------------------
Delete a character by removing its JSON file.
		 
-name CHARACTER_NAME   Name of the character to delete (string)


EQUIP
---------------------------------------------------------------------
Equip a character with a weapon, armor, or shield.
Updates the character's JSON file accordingly.
There are three patterns for this command:


1. Equip a weapon to a specified slot.
   - Name CHARACTER_NAME   Name of the character (string)
   - Weapon WEAPON_NAME   Name of the weapon to equip (string)
   - Slot SLOT            Slot to equip the weapon to (string)
2. Equip armor.
   - Name CHARACTER_NAME   Name of the character (string)
   - Armor ARMOR_NAME     Name of the armor to equip (string)
3. Equip a shield.
   - Name CHARACTER_NAME   Name of the character (string)
   - Shield SHIELD_NAME   Name of the shield to equip (string)


LEARN-SPELL
---------------------------------------------------------------------
Add a spell to the character's list of known spells by learning it.
Updates the character's JSON file accordingly.
		 
-name CHARACTER_NAME   Name of the character (string)
-spell SPELL_NAME      Name of the spell to learn (string)


PREPARE-SPELL
---------------------------------------------------------------------
Add a spell to the character's list of prepared spells.
Updates the character's JSON file accordingly.
		 
-name CHARACTER_NAME   Name of the character (string)
-spell SPELL_NAME      Name of the spell to prepare (string)


EXPORT
---------------------------------------------------------------------
Export the character sheet to an HTML file named CHARACTER_NAME.html.

-name CHARACTER_NAME   Name of the character to export (string)


UPDATE-DATA
---------------------------------------------------------------------
Update local data files (e.g., spells, items) from a remote source.
This command does not require any additional arguments.
It only updates local files that are not already present.
