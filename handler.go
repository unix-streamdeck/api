package api

import (
	"github.com/unix-streamdeck/driver"
)

type IconHandler interface {
	Icon(page int, index int, key *Key, dev streamdeck.Device)
	Stop()
}

type KeyHandler interface {
	Key(page int, index int, key *Key, dev streamdeck.Device)
}
