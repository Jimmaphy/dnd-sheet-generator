# !/bin/bash


# Prepare for the tests
DND_TESTS_SUCCEEDED=0
DND_TESTS_FAILED=0



# Test 1: Compile the program
go build -o dcg main.go

if [ $? -ne 0 ]; then
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 1 Failed: Compilation error."
else
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 1 Passed: Compilation successful."
fi



# Test 2: Run the program without arguments and check for usage message
output=$(./dcg 2>&1)

if [[ $output == *"Usage:"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 2 Passed: Defaults to usage message."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 2 Failed: Does not default to usage message without arguments."
fi



# Test 3: Run the usage command and check for usage message
output=$(./dcg usage 2>&1)

if [[ $output == *"Usage:"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 3 Passed: Usage command works."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 3 Failed: Usage command does not work."
fi



# Test 4: Create a character with just a name
output=$(./dcg create -name "TestCharacter" 2>&1)

if [[ $output == *"TestCharacter"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 4 Passed: Character created with name."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 4 Failed: Character not created with name."
fi



# Test 5: Name should be required when creating a character
output=$(./dcg create 2>&1)

if [[ $output == *"name is required"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 5 Passed: Name is required for create."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 5 Failed: Name requirement not enforced."
fi



# Test 6: List characters should find the created character
output=$(./dcg list 2>&1)

if [[ $output == *"TestCharacter"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 6 Passed: character found in list."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 6 Failed: character not found in list."
fi



# Test 7: The character can be shown
output=$(./dcg view -name "TestCharacter" 2>&1)

if [[ $output == *"TestCharacter"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 7 Passed: character can be viewed."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 7 Failed: character cannot be viewed."
fi



# Test 8: The character can be deleted
./dcg delete -name "TestCharacter" >/dev/null 2>&1
output=$(./dcg view -name "TestCharacter" 2>&1)

if [[ $output == *"not found"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 8 Passed: character was deleted."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 8 Failed: character not deleted."
fi



# Test 9: Create a new character and add a basic weapon
./dcg create -name "Merry Brandybuck" -race "lightfoot halfling" -class rogue -str 8 -dex 15 -con 14 -int 10 -wis 12 -cha 13 >/dev/null 2>&1
output=$(./dcg equip -name "Merry Brandybuck" -weapon shortsword -slot "main hand" 2>&1)

if [[ $output == *"Equipped"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 9 Passed: Weapon equipped successfully."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 9 Failed: Weapon not equipped."
fi



# Test 10: Verify weapon is listed in character view
output=$(./dcg view -name "Merry Brandybuck" 2>&1)

if [[ $output == *"shortsword"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 10 Passed: Weapon listed in character view."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 10 Failed: Weapon not listed in character view."
fi



# Test 11: Equip armor to the character
output=$(./dcg equip -name "Merry Brandybuck" -armor "chain shirt" 2>&1)

if [[ $output == *"Equipped"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 11 Passed: Armor equipped successfully."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 11 Failed: Armor not equipped."
fi



# Test 12: Verify armor is listed in character view
output=$(./dcg view -name "Merry Brandybuck" 2>&1)

if [[ $output == *"chain shirt"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 12 Passed: Armor listed in character view."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 12 Failed: Armor not listed in character view."
fi



# Test 13: Equip a shield to the character
output=$(./dcg equip -name "Merry Brandybuck" -shield "shield" 2>&1)

if [[ $output == *"Equipped"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 13 Passed: Shield equipped successfully."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 13 Failed: Shield not equipped."
fi



# Test 14: Verify shield is listed in character view
output=$(./dcg view -name "Merry Brandybuck" 2>&1)

if [[ $output == *"shield"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 14 Passed: Shield listed in character view."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 14 Failed: Shield not listed in character view."
fi



# Test 15: Equip to off-hand 
output=$(./dcg equip -name "Merry Brandybuck" -weapon dagger -slot "off hand" 2>&1)

if [[ $output == *"Equipped"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 15 Passed: Off-hand weapon equipped successfully."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 15 Failed: Off-hand weapon not equipped."
fi



# Test 16: Verify off-hand weapon is listed in character view
output=$(./dcg view -name "Merry Brandybuck" 2>&1)

if [[ $output == *"dagger"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 16 Passed: Off-hand weapon listed in character view."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 16 Failed: Off-hand weapon not listed in character view."
fi



# Test 17: Cannot equip weapon to occupied slot
output=$(./dcg equip -name "Merry Brandybuck" -weapon longsword -slot "main hand" 2>&1)

if [[ $output == *"already occupied"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 17 Passed: Cannot equip to occupied slot."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 17 Failed: Equipped to occupied slot."
fi



# Remove the test character and make new ones for spell tests
./dcg delete -name "Merry Brandybuck" >/dev/null 2>&1
./dcg create -name "Aragorn" -race "human" -class fighter -level 20 -str 10 -dex 13 -con 14 -int 10 -wis 12 -cha 15 >/dev/null 2>&1
./dcg create -name "Gandalf" -race "human" -class wizard -level 20 -str 8 -dex 10 -con 12 -int 14 -wis 15 -cha 13 >/dev/null 2>&1
./dcg create -name "Galadriel" -race "high elf" -class sorcerer -level 20 -str 8 -dex 10 -con 12 -int 15 -wis 14 -cha 13 >/dev/null 2>&1



# Test 18: Prepare a spell
output=$(./dcg prepare-spell -name "Gandalf" -spell "burning hands" 2>&1)

if [[ $output == *"Prepared"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 18 Passed: Spell prepared successfully."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 18 Failed: Spell not prepared."
fi



# Test 19: Learn a spell
output=$(./dcg learn-spell -name "Galadriel" -spell "false life" 2>&1)

if [[ $output == *"Learned"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 19 Passed: Spell learned successfully."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 19 Failed: Spell not learned."
fi



# Test 20: Can't prepare a spell if you're a learner
output=$(./dcg prepare-spell -name "Galadriel" -spell "false life" 2>&1)

if [[ $output == *"learns"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 20 Passed: Learner cannot prepare spell."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 20 Failed: Learner prepared spell."
fi



# Test 21: Can't learn a spell if you're a preparer
output=$(./dcg learn-spell -name "Gandalf" -spell "burning hands" 2>&1)

if [[ $output == *"prepares"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 21 Passed: Preparer cannot learn spell."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 21 Failed: Preparer learned spell."
fi



# Test 22: Can't learn spell if you're not a spellcaster
output=$(./dcg learn-spell -name "Aragorn" -spell "fireball" 2>&1)

if [[ $output == *"can't cast"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 22 Passed: Non-spellcaster cannot learn spell."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 22 Failed: Non-spellcaster learned spell."
fi



# Test 23: Can't prepare spell if you're not a spellcaster
output=$(./dcg prepare-spell -name "Aragorn" -spell "fireball" 2>&1)

if [[ $output == *"can't cast"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 23 Passed: Non-spellcaster cannot prepare spell."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 23 Failed: Non-spellcaster prepared spell."
fi



# Clearning up the compiled binary
echo ""
echo "Tests completed. Succeeded: $DND_TESTS_SUCCEEDED. Failed: $DND_TESTS_FAILED"

./dcg delete -name "Aragorn" >/dev/null 2>&1
./dcg delete -name "Gandalf" >/dev/null 2>&1
./dcg delete -name "Galadriel" >/dev/null 2>&1
rm dcg
