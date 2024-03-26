package urlparsing

import (
	"fmt"
	"net"

	"net/url"
)

func Main() {
	dsn := "http://localhost:8080/user?name=an&&password=123123"

	u, err := url.Parse(dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host:: ", host)
	fmt.Println("port:: ", port)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("map query params:: ", m)
	fmt.Println("Value param[name]::", m["name"][0])

}
