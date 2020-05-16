package entity

const PersonSize = 3
type Person struct {
	*ModelS
}

func (p Person) GetSize() GridType {
	return p.ModelS.GetSize()
}

func (p Person) GetId() int {
	return p.ModelS.GetId()
}

func (p Person) GetPosition() (GridType, GridType) {
	modelPtr := p.ModelS
	return modelPtr.GetPosition()
}

func (p Person) SetPosition(newX GridType, newY GridType) {
	modelPtr := p.ModelS
	modelPtr.SetPosition(newX, newY)
}

func NewPerson(initX GridType, initY GridType) Person {
	return Person{
		ModelS: newModel(initX, initY, PersonSize),
	}
}
