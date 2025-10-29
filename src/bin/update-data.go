package bin

import (
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/jimmaphy/dnd-sheet-generator/adapters"
	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/infrastructure"
	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type UpdateDataCommand struct {
	ApiService infrastructure.APIService
}

// Create a new instance of UpdateDataCommand.
// All data will always be updated when this command is executed.
// So, no arguments are necessary.
func NewUpdateDataCommand() Command {
	return &UpdateDataCommand{
		ApiService: *infrastructure.NewAPIService("https://www.dnd5eapi.co/api/2014"),
	}
}

// ParseArguments for UpdateDataCommand does not require any arguments and always returns nil.
func (command *UpdateDataCommand) ParseArguments(args []string) error {
	return nil
}

// Execute reads and prints the usage instructions from the usage.txt template file.
// If the file cannot be read, an error message is printed and the program exits with status 1.
// Subroutines for updating armor and weapons are called concurrently.
// This function will never error out, errors messages are directed to standard output.
func (command *UpdateDataCommand) Execute() error {
	var group sync.WaitGroup

	group.Go(command.updateArmor)
	group.Go(command.updateWeapons)
	group.Go(command.updateClasses)

	group.Wait()
	return nil
}

// To update armor, the equiment categories endpoint is queried for armor listings.
// The local armor repository is checked for existing armor names.
// If an armor name from the API listing is not found locally, a message is printed indicating it should be fetched.
func (command *UpdateDataCommand) updateArmor() {
	var armorListing *adapters.ListingResponse
	command.ApiService.GetData(&armorListing, "equipment-categories", "armor")

	armorRepository := repository.NewArmorJSONRepository()
	localArmorNames, err := armorRepository.List()
	if err != nil {
		fmt.Println("failed to list local armor names: " + err.Error())
	}

	updatedArmorCount := 0
	var armorGroup sync.WaitGroup

	for _, armor := range armorListing.List() {
		if !slices.Contains(localArmorNames, strings.ToLower(armor.Name)) {
			updatedArmorCount += 1
			armorGroup.Go(func() {
				command.downloadArmor(armor.Index, armorRepository)
			})
		}
	}

	armorGroup.Wait()
	fmt.Println("New armor items downloaded: ", updatedArmorCount)
}

// Download armor downloads a specific armor by index from the API.
// It is then adapted to the domain model and stored in the local armor repository.
func (command *UpdateDataCommand) downloadArmor(index string, armorRepository *repository.ArmorJSONRepository) {
	var fetchedArmor *adapters.ApiArmor
	command.ApiService.GetData(&fetchedArmor, "equipment", index)
	armorRepository.Add(fetchedArmor.ToDomainModel())
}

// To update weapons, the equiment categories endpoint is queried for weapon listings.
// The local weapon repository is checked for existing weapon names.
// If a weapon name from the API listing is not found locally, a message is printed indicating it should be fetched.
func (command *UpdateDataCommand) updateWeapons() {
	var weaponListing *adapters.ListingResponse
	command.ApiService.GetData(&weaponListing, "equipment-categories", "weapon")

	weaponRepository := repository.NewWeaponJSONRepository()
	localWeaponNames, err := weaponRepository.List()
	if err != nil {
		fmt.Println("failed to list local weapon names: " + err.Error())
	}

	updatedWeaponCount := 0
	var weaponGroup sync.WaitGroup

	for _, weapon := range weaponListing.List() {
		if !slices.Contains(localWeaponNames, strings.ToLower(weapon.Name)) {
			updatedWeaponCount += 1
			weaponGroup.Go(func() {
				command.downloadWeapon(weapon.Index, weaponRepository)
			})
		}
	}

	weaponGroup.Wait()
	fmt.Println("New weapon items downloaded: ", updatedWeaponCount)
}

// Download weapon downloads a specific weapon by index from the API.
// It is then adapted to the domain model and stored in the local weapon repository.
func (command *UpdateDataCommand) downloadWeapon(index string, weaponRepository *repository.WeaponJSONRepository) {
	var fetchedWeapon *adapters.ApiWeapon
	command.ApiService.GetData(&fetchedWeapon, "equipment", index)
	weaponRepository.Add(fetchedWeapon.ToDomainModel())
}

// To update classes, the classs endpoint is queried for spell listings.
// The local class repository is checked for existing spell names.
// If a class name from the API listing is not found locally, a message is printed indicating it should be fetched.
func (command *UpdateDataCommand) updateClasses() {
	var classListing *adapters.ClassResponse
	command.ApiService.GetData(&classListing, "classes")

	classRepository := repository.NewClassJSONRepository()
	localClassNames, err := classRepository.List()
	if err != nil {
		fmt.Println("failed to list local class names: " + err.Error())
	}

	updatedClassCount := 0
	var classGroup sync.WaitGroup

	for _, class := range classListing.List() {
		if !slices.Contains(localClassNames, strings.ToLower(class.Name)) {
			updatedClassCount += 1
			classGroup.Go(func() {
				command.downloadClass(class.Index, classRepository)
			})
		}
	}

	classGroup.Wait()
	fmt.Println("New classes downloaded: ", updatedClassCount)
}

// Download class downloads a specific class by name from the API.
// It is then adapted to the domain model and stored in the local class repository.
func (command *UpdateDataCommand) downloadClass(name string, classRepository *repository.ClassJSONRepository) {
	var fetchedClass *adapters.ApiClass
	var fetchedLevels []adapters.ApiLevel
	var fetchedSpells *adapters.ApiSpellList
	var levelDomainModels []*domain.SpellLevel
	var spellDomainModels []*domain.Spell

	command.ApiService.GetData(&fetchedClass, "classes", name)
	command.ApiService.GetData(&fetchedLevels, "classes", name, "levels")
	command.ApiService.GetData(&fetchedSpells, "classes", name, "spells")

	for _, level := range fetchedLevels {
		levelDomainModels = append(levelDomainModels, level.ToDomainModel())
	}

	if fetchedSpells.Count > 0 {
		for _, spell := range fetchedSpells.Results {
			spellDomainModels = append(spellDomainModels, spell.ToDomainModel())
		}
	}

	domainClass := fetchedClass.ToDomainModel(spellDomainModels, levelDomainModels)
	classRepository.Add(domainClass)
}
