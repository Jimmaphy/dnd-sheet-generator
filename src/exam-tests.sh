#!/bin/bash

# This script tests the functionality of the exam feature.
# The base application has been tested through CodeGrade.


# Prepare the environment
TESTS_SUCCEEDED=0
TESTS_FAILED=0


# Test 1: Compile the application
echo "[SETUP] Compiling the application"
go build -o dndcsg > /dev/null 2>&1

if [ $? -ne 0 ]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Compilation failed"
    exit 1
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Compilation succeeded"
fi



# Test 2: Create two characters according to the exam specifications
echo "[SETUP] Creating two characters"

output=$(./dndcsg create -name "Pip Quickstep" -race "halfling" -class "rogue" -level 2 -str 10 -dex 16 -con 12 -int 13 -wis 11 -cha 14 2>&1)
if [[ $? -ne 0 || $output != *"saved character"* ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Failed to create character Pip Quickstep"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Created character Pip Quickstep"
fi

output=$(./dndcsg create -name "Thrain Stonefist" -race "hill dwarf" -class "barbarian" -level 6 -str 17 -dex 13 -con 12 -int 9 -wis 10 -cha 8 2>&1)
if [[ $? -ne 0 || $output != *"saved character"* ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Failed to create character Thrain Stonefist"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Created character Thrain Stonefist"
fi



# Test 3: Make sure the required data is stored in the JSON file
# Only the hit die is stored in the json. Everything else is calculated on the fly.
echo " [TEST] Verifying hit points data in JSON file"

output=$(grep -a '"HitDie": 8' './storage/characters/Pip Quickstep.json' 2>&1)
if [[ $output ]]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Pip Quickstep has correct hit die stored"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Pip Quickstep has incorrect hit die stored"
fi

output=$(grep -a '"HitDie": 12' './storage/characters/Thrain Stonefist.json' 2>&1)
if [[ !($output) ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Thrain Stonefist has incorrect hit die stored"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Thrain Stonefist has correct hit die stored"
fi



# Test 4: Make sure the calculated hit points are correct
echo " [TEST] Verifying calculated hit points"

output=$(./dndcsg view -name "Pip Quickstep" 2>&1)
if [[ $? -ne 0 || $output != *"Total hit points: 15"* ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Pip Quickstep has incorrect hit points"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Pip Quickstep has correct hit points"
fi

output=$(./dndcsg view -name "Thrain Stonefist" 2>&1)
if [[ $? -ne 0 || $output != *"Total hit points: 59"* ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Thrain Stonefist has incorrect hit points"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Thrain Stonefist has correct hit points"
fi



# Test 5: Export the character sheets
# Remember, this does not print anything to the console when successful
# Only errors are printed
echo "[SETUP] Exporting character sheets"

output=$(./dndcsg export -name "Pip Quickstep" 2>&1)
if [[ $? -ne 0 || $output == *" "* ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Failed to export character sheet for Pip Quickstep"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Exported character sheet for Pip Quickstep"
fi

output=$(./dndcsg export -name "Thrain Stonefist" 2>&1)
if [[ $? -ne 0 || $output == *" "* ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Failed to export character sheet for Thrain Stonefist"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Exported character sheet for Thrain Stonefist"
fi



# Test 6: Make sure the calculated hit points are included in the exported character sheet
echo " [TEST] Verifying hit points in exported character sheets"
./dndcsg export -name "Pip Quickstep" > /dev/null 2>&1
./dndcsg export -name "Thrain Stonefist" > /dev/null 2>&1

output=$(grep -a '<input name="maxhp" placeholder="10" type="text" value="15" />' './export/Pip Quickstep.html' 2>&1)
if [[ $output ]]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Pip Quickstep export has correct max hit points"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Pip Quickstep export has incorrect max hit points"
fi

output=$(grep -a '<input name="maxhp" placeholder="10" type="text" value="59" />' './export/Thrain Stonefist.html' 2>&1)
if [[ $output ]]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Thrain Stonefist export has correct max hit points"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Thrain Stonefist export has incorrect max hit points"
fi

output=$(grep -a '<input name="currenthp" type="text" value="15" />' './export/Pip Quickstep.html' 2>&1)
if [[ $output ]]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Pip Quickstep export has correct current hit points"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Pip Quickstep export has incorrect current hit points"
fi

output=$(grep -a '<input name="currenthp" type="text" value="59" />' './export/Thrain Stonefist.html' 2>&1)
if [[ $output ]]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Thrain Stonefist export has correct current hit points"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Thrain Stonefist export has incorrect current hit points"
fi

output=$(grep -a '<input name="totalhd" placeholder="2d10" type="text" value="2d8"/>' './export/Pip Quickstep.html' 2>&1)
if [[ $output ]]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Pip Quickstep export has correct hit die"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Pip Quickstep export has incorrect hit die"
fi

output=$(grep -a '<input name="totalhd" placeholder="2d10" type="text" value="6d12"/>' './export/Thrain Stonefist.html' 2>&1)
if [[ $output ]]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Thrain Stonefist export has correct hit die"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Thrain Stonefist export has incorrect hit die"
fi



# Test 7: Delete a class and make sure the class gets downloaded again
echo "[SETUP] Deleting a class and verifying re-download"
rm -f ./storage/classes/fighter.json >> /dev/null 2>&1
output=$(./dndcsg update-data 2>&1)
check=$(echo $output | grep -o "classes downloaded: 1")

if [[ $? -ne 0 || !($check) ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Failed to re-download fighter class"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Re-downloaded fighter class"
fi



# Test 8: Make sure the downloaded class contains the hit die information
echo " [TEST] Verifying hit die information in re-downloaded class"
if grep -q '"HitDie": 10' './storage/classes/fighter.json'; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Fighter class has correct hit die information"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Fighter class has incorrect hit die information"
fi



# Test 9: Delete created characters
# Exported files should not be deleted. They are the user's responsibility.
echo "[CLEAN] Deleting created characters"

output=$(./dndcsg delete -name "Pip Quickstep" 2>&1)
if [[ $? -ne 0 || $output != *"deleted"* ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Failed to delete character Pip Quickstep"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Deleted character Pip Quickstep"
fi

output=$(./dndcsg delete -name "Thrain Stonefist" 2>&1)
if [[ $? -ne 0 || $output != *"deleted"* ]]; then
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Failed to delete character Thrain Stonefist"
else
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Deleted character Thrain Stonefist"
fi

if [ ! -f './storage/characters/Pip Quickstep.json' ]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Pip Quickstep JSON file deleted"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Pip Quickstep JSON file not deleted"
fi

if [ ! -f './storage/characters/Thrain Stonefist.json' ]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Thrain Stonefist JSON file deleted"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Thrain Stonefist JSON file not deleted"
fi

if [ -f './export/Pip Quickstep.html' ]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Pip Quickstep export file still exists"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Pip Quickstep export file missing"
fi

if [ -f './export/Thrain Stonefist.html' ]; then
    TESTS_SUCCEEDED=$((TESTS_SUCCEEDED + 1))
    echo "        [SUCCESS] Thrain Stonefist export file still exists"
else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo "        [ERROR] Thrain Stonefist export file missing"
fi



# Print the summary
echo ""
echo "===================="
echo "TOTAL TESTS RUN: $((TESTS_SUCCEEDED + TESTS_FAILED))"
echo "TESTS SUCCEEDED: $TESTS_SUCCEEDED"
echo "TESTS FAILED:    $TESTS_FAILED"
echo "===================="

# Cleaup the executable
rm -f ./dndcsg >> /dev/null 2>&1
