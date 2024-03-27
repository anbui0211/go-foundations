package builder

type Director struct {
	builder IBuilder
}

/*NewDirector is function*/
func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

/*SetBuilder is function */
func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
