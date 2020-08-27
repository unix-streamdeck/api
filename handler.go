package api

import "image"

type IconHandler interface {
	Start(key Key, info StreamDeckInfo, callback func(image image.Image))
	IsRunning() bool
	SetRunning(running bool)
	Stop()
}

type KeyHandler interface {
	Key(key Key, info StreamDeckInfo)
}
