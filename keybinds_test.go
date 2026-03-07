package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseXDoToolKeybindString_EmptyString(t *testing.T) {
	assertions := assert.New(t)
	keycodes, err := ParseXDoToolKeybindString("")
	assertions.Nil(keycodes)
	assertions.EqualError(err, "empty keybind")
}

func TestParseXDoToolKeybindString_UnknownKey(t *testing.T) {
	assertions := assert.New(t)
	keycodes, err := ParseXDoToolKeybindString("test")
	assertions.Nil(keycodes)
	assertions.EqualError(err, "unknown key: test")
}

func TestParseXDoToolKeybindString_NoValidBinds(t *testing.T) {
	assertions := assert.New(t)
	keycodes, err := ParseXDoToolKeybindString(" +  ")
	assertions.Nil(keycodes)
	assertions.EqualError(err, "no valid keys in keybind")
}

func TestParseXDoToolKeybindString_ValidBind(t *testing.T) {
	assertions := assert.New(t)
	keycodes, err := ParseXDoToolKeybindString("ctrl+esc")
	assertions.Equal(keycodes, []int{29, 1})
	assertions.Nil(err)
}

func TestFindXDoToolKeybindString_InvalidKeycode(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal("", FindXDoToolKeybindString(999))
}

func TestFindXDoToolKeybindString_ValidKeycode(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal("backspace", FindXDoToolKeybindString(14))
}
