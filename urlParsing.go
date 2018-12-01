package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	pt := fmt.Println

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	eCheck(err)

	pt(u.Scheme)
	pt(u.User)
	pt(u.User.Username())
	pt(u.User.Password())
	pt(u.User.String())
	pt(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	pt(host)
	pt(port)

	pt(u.RawQuery)

}

func eCheck(e error) {
	if e != nil {
		panic(e)
	}
}
