package builder

type igloolBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIgloolBuilder() *igloolBuilder {
	return &igloolBuilder{}
}

func (b *igloolBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *igloolBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *igloolBuilder) setNumFloor() {
	b.floor = 1
}

func (b *igloolBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
