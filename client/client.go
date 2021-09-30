package client

import "sync"

type MoySklad struct {
	Login string
	Password string
	WG  *sync.WaitGroup
}