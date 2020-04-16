package entity

const PersonSize = 3
type Person struct {
	ModelS
}

func (p Person) GetSize() GridType {
	return p.ModelS.GetSize()
}

func (p Person) GetId() int {
	return p.ModelS.GetId()
}

func (p Person) GetPosition() (GridType, GridType) {
	return p.GetPosition()
}

func NewPerson(initX GridType, initY GridType) Person {
	return Person{
		ModelS: newModel(initX, initY, PersonSize),
	}
}

func (person *Person) Move(newX GridType, newY GridType) {
	person.x = newX
	person.y = newY
}

