package builder

import "fmt"

func Main() {
	normalBuilder := GetBuilder("normal")
	// igloolBuilder := GetBuilder("igloo")

	diector := NewDirector(normalBuilder)
	normalHouse := diector.BuildHouse()

	fmt.Printf("Normal House Window Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.GetNumFloor())
}
