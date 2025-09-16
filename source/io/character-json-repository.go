package io

import (
	"github.com/jimmaphy/dnd-sheet-generator/character"
)

type JsonCharacterRepository struct {
	handler *JsonHandler
}

func NewJsonCharacterRepository() *JsonCharacterRepository {
	handler, _ := NewJsonHandler("./storage/characters/")

	return &JsonCharacterRepository{
		handler: handler,
	}
}

func (repository *JsonCharacterRepository) Add(character character.Character) error {
	return repository.handler.Store(character, character.Name)
}

func (repository *JsonCharacterRepository) Get(name string) (character.Character, error) {
	var character character.Character
	err := repository.handler.Read(name, &character)
	return character, err
}

func (repository *JsonCharacterRepository) List() ([]character.Character, error) {
	names, err := repository.handler.List()
	if err != nil {
		return nil, err
	}

	var characters []character.Character
	for _, name := range names {
		character, err := repository.Get(name)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return characters, nil
}
