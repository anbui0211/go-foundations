package generic

import (
	"encoding/json"
	"sync"
)

//
// PUBLIC METHOD
//

// GetListUser ...
func GetListUser(users []User) (res []UserBrief) {
	if len(users) == 0 {
		return
	}

	var (
		total = len(users)
		wg    sync.WaitGroup
	)

	res = make([]UserBrief, total)
	wg.Add(total)
	for index, user := range users {
		go func(i int, u User) {
			defer wg.Done()
			res[i] = Brief(u)
		}(index, user)
	}

	wg.Wait()
	return
}

// GetListShort ...
func GetListShort(users []User) (res []UserShort) {
	if len(users) == 0 {
		return
	}

	var (
		total = len(users)
		wg    sync.WaitGroup
	)

	res = make([]UserShort, total)
	wg.Add(total)
	for index, user := range users {
		go func(i int, u User) {
			defer wg.Done()
			res[i] = Short(u)
		}(index, user)
	}

	wg.Wait()
	return
}

// ToString ...
func ToString(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

// GetResponse ...
func GetResponse[K, V any](data []K, covertResponse func(K) V) (res []V) {
	if len(data) == 0 {
		return
	}

	var (
		total = len(data)
		wg    sync.WaitGroup
	)

	res = make([]V, total)
	wg.Add(total)
	for index, user := range data {
		go func(i int, u K) {
			defer wg.Done()
			res[i] = covertResponse(u)
		}(index, user)
	}

	wg.Wait()
	return
}

//
// PRIVATE METHOD
//

// Brief ...
func Brief(m User) UserBrief {
	return UserBrief{
		ID:        m.ID,
		Name:      m.Name,
		Age:       m.Age,
		Gender:    m.Gender,
		Phone:     m.Phone,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	}
}

// Short ...
func Short(m User) UserShort {
	return UserShort{
		ID:   m.ID,
		Name: m.Name,
	}
}
