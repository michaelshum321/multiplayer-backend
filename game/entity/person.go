package entity

import "multiplayer-backend/game"

const PersonSize = 3
type Person struct {
	ModelS
}

func (p Person) GetSize() game.GridType {
	return p.ModelS.GetSize()
}

func (p Person) GetId() int {
	return p.ModelS.GetId()
}

func (p Person) GetPosition() (game.GridType, game.GridType) {
	return p.GetPosition()
}

func NewPerson(initX game.GridType, initY game.GridType) Person {
	return Person{
		ModelS: newModel(initX, initY, PersonSize),
	}
}

func (person *Person) Move(newX game.GridType, newY game.GridType) {
	person.x = newX
	person.y = newY
}

