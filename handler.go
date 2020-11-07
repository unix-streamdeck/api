package api

import "image"

type Handler interface {

}

type IconHandler interface {
	Handler
	Start(key Key, info StreamDeckInfo, callback func(image image.Image))
	IsRunning() bool
	SetRunning(running bool)
	Stop()
}

type KeyHandler interface {
	Handler
	Key(key Key, info StreamDeckInfo)
}
