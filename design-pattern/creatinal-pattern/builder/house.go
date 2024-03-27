/*
- Product: House
- Builder Interface:
- Concrete Builder 1: Normal House
- Concrete Builder 2: Igloo House
- Director:
*/
package builder

/*House is struct */
type House struct {
	windowType string
	doorType   string
	floor      int
}

/*GetWindowType is getter */
func (h *House) GetWindowType() string {
	return h.windowType
}

/*GetDoorType is getter */
func (h *House) GetDoorType() string {
	return h.doorType
}

/*GetNumFloor is getter */
func (h *House) GetNumFloor() int {
	return h.floor
}
