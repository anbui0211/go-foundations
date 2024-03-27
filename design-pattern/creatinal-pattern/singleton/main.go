package singleton

import (
	"fmt"
	"sync"
	"time"
)

/*Singleton interface*/
type Singleton interface {
	AddOne() int
}

/*
- Không cho các package khác truy cập
- Chỉ cho phép khởi tạo 1 đối tượng duy nhất trong xuyên suốt chương trình chạy
*/
type singleton struct {
	count int
}

var (
	instance *singleton
	once     sync.Once
)

/* GetInstance function return object */
/*
- Hàm GetInstance() kiểm tra xem instance có nil hay không. Nếu nil, nó sẽ tạo một instance mới và gán cho instance.
- Có thể xảy ra trường hợp nhiều goroutines cùng lúc kiểm tra instance thấy nil, và tất cả đều tiến hành tạo instance mới.
- Điều này dẫn đến việc có nhiều hơn một instance được tạo, vi phạm nguyên tắc singleton.
*/
func GetInstance() Singleton {
	once.Do(func() {
		time.Sleep(time.Second * 2)
		instance = &singleton{count: 100}
	})
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count

}

func Main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", GetInstance())
		}()
	}

	time.Sleep(time.Second * 10)
}
