package golang

import (
	"errors"
	"fmt"
)

func echo(s fmt.Stringer) {
	print(s.String())
}

func throw(e error) {
	print(e.Error())
}

type iface interface {
	Name() string
	Age() int
}

func needIF(i iface) {
	fmt.Println(i.Name(), i.Age())
}

type myStringer interface {
	fmt.Stringer
	error
}

type messages struct {
	message string
}

type err struct {
	er error
}

func (e err) get() error {
	return e.er
}

func newThrow(messages string) err {
	e := err{errors.New(messages)}
	return e
}

func (m messages) String() string {
	return fmt.Sprintf("%v", m.message)
}

func newStringer(m string) fmt.Stringer {
	mm := messages{m}
	return mm
}
func strErr(s myStringer) {
	fmt.Println(s.Error(), s.String())
}
