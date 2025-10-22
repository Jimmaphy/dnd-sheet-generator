# !/bin/bash


# Prepare for the tests
DND_TESTS_SUCCEEDED=0
DND_TESTS_FAILED=0



# Test 1: Compile the program
go build -o dnd-character-generator main.go

if [ $? -ne 0 ]; then
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 1 Failed: Compilation error."
else
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 1 Passed: Compilation successful."
fi



# Test 2: Run the program without arguments and check for usage message
output=$(./dnd-character-generator 2>&1)

if [[ $output == *"Usage:"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 2 Passed: Defaults to usage message."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 2 Failed: Does not default to usage message without arguments."
fi



# Test 3: Run the usage command and check for usage message
output=$(./dnd-character-generator usage 2>&1)

if [[ $output == *"Usage:"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 3 Passed: Usage command works."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 3 Failed: Usage command does not work."
fi



# Test 4: Create a character with just a name
output=$(./dnd-character-generator create -name "TestCharacter" 2>&1)

if [[ $output == *"TestCharacter"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 4 Passed: Character created with name."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 4 Failed: Character not created with name."
fi



# Test 5: Name should be required when creating a character
output=$(./dnd-character-generator create 2>&1)

if [[ $output == *"name is required"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 5 Passed: Name is required for create."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 5 Failed: Name requirement not enforced."
fi



# Test 6: List characters should find the created character
output=$(./dnd-character-generator list 2>&1)

if [[ $output == *"TestCharacter"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 6 Passed: character found in list."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 6 Failed: character not found in list."
fi



# Test 7: The character can be shown
output=$(./dnd-character-generator view -name "TestCharacter" 2>&1)

if [[ $output == *"TestCharacter"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 7 Passed: character can be viewed."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 7 Failed: character cannot be viewed."
fi



# Test 8: The character can be deleted
./dnd-character-generator delete -name "TestCharacter" >/dev/null 2>&1
output=$(./dnd-character-generator view -name "TestCharacter" 2>&1)

if [[ $output == *"not found"* ]]; then
    DND_TESTS_SUCCEEDED=$((DND_TESTS_SUCCEEDED + 1))
    echo "Test 8 Passed: character was deleted."
else
    DND_TESTS_FAILED=$((DND_TESTS_FAILED + 1))
    echo "Test 8 Failed: character not deleted."
fi


# Clearning up the compiled binary
echo ""
echo "Tests completed. Succeeded: $DND_TESTS_SUCCEEDED. Failed: $DND_TESTS_FAILED"
rm dnd-character-generator
