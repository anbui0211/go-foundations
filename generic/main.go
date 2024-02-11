package generic

import "fmt"

func generic() {
	usersData := GetUserBSON()

	//list := GetListUser(usersData)
	//fmt.Println(ToString(list))

	//list := GetListShort(usersData)
	//fmt.Println(ToString(list))

	//responses := GetResponse(usersData, Short)
	//fmt.Println(ToString(responses))

	responses := GetResponse(usersData, Brief)
	fmt.Println(ToString(responses))
	go func() {

	}()
}
