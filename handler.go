package api

import "image"

type IconHandler interface {
	Icon(key Key, info StreamDeckInfo, callback func(image image.Image))
	Stop()
}

type KeyHandler interface {
	Key(key Key, info StreamDeckInfo)
}
