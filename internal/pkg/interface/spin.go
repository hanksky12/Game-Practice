package _interface

import "sync"

type ISpin interface {
	Spin(wg *sync.WaitGroup)
}
