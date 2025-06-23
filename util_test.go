package api

import (
	"testing"
)

func TestCompareKeyConfigs(t *testing.T) {
	// Test case 1: Identical configs
	config1 := KeyConfigV3{
		Icon:          "icon.png",
		SwitchPage:    1,
		Text:          "Button",
		TextSize:      12,
		TextAlignment: "center",
		Keybind:       "ctrl+a",
		Command:       "echo hello",
		Brightness:    100,
		Url:           "https://example.com",
		IconHandler:   "handler1",
		KeyHandler:    "handler2",
	}
	
	config2 := KeyConfigV3{
		Icon:          "icon.png",
		SwitchPage:    1,
		Text:          "Button",
		TextSize:      12,
		TextAlignment: "center",
		Keybind:       "ctrl+a",
		Command:       "echo hello",
		Brightness:    100,
		Url:           "https://example.com",
		IconHandler:   "handler1",
		KeyHandler:    "handler2",
	}
	
	if !CompareKeyConfigs(config1, config2) {
		t.Error("Identical configs should be equal")
	}
	
	// Test case 2: Different icon
	config2.Icon = "different.png"
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different icons should not be equal")
	}
	
	// Test case 3: Different switch page
	config2.Icon = config1.Icon
	config2.SwitchPage = 2
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different switch pages should not be equal")
	}
	
	// Test case 4: Different text
	config2.SwitchPage = config1.SwitchPage
	config2.Text = "Different"
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different text should not be equal")
	}
	
	// Test case 5: Different text size
	config2.Text = config1.Text
	config2.TextSize = 14
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different text sizes should not be equal")
	}
	
	// Test case 6: Different text alignment
	config2.TextSize = config1.TextSize
	config2.TextAlignment = "left"
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different text alignments should not be equal")
	}
	
	// Test case 7: Different keybind
	config2.TextAlignment = config1.TextAlignment
	config2.Keybind = "ctrl+b"
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different keybinds should not be equal")
	}
	
	// Test case 8: Different command
	config2.Keybind = config1.Keybind
	config2.Command = "echo world"
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different commands should not be equal")
	}
	
	// Test case 9: Different brightness
	config2.Command = config1.Command
	config2.Brightness = 50
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different brightness should not be equal")
	}
	
	// Test case 10: Different URL
	config2.Brightness = config1.Brightness
	config2.Url = "https://different.com"
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different URLs should not be equal")
	}
	
	// Test case 11: Different icon handler
	config2.Url = config1.Url
	config2.IconHandler = "handler3"
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different icon handlers should not be equal")
	}
	
	// Test case 12: Different key handler
	config2.IconHandler = config1.IconHandler
	config2.KeyHandler = "handler4"
	if CompareKeyConfigs(config1, config2) {
		t.Error("Configs with different key handlers should not be equal")
	}
}

func TestCompareKeys(t *testing.T) {
	// Test case 1: Identical keys
	key1 := KeyV3{
		Application: map[string]*KeyConfigV3{
			"app1": {
				Icon:    "icon1.png",
				Command: "echo app1",
			},
			"app2": {
				Icon:    "icon2.png",
				Command: "echo app2",
			},
		},
	}
	
	key2 := KeyV3{
		Application: map[string]*KeyConfigV3{
			"app1": {
				Icon:    "icon1.png",
				Command: "echo app1",
			},
			"app2": {
				Icon:    "icon2.png",
				Command: "echo app2",
			},
		},
	}
	
	if !CompareKeys(key1, key2) {
		t.Error("Identical keys should be equal")
	}
	
	// Test case 2: Different number of applications
	key2.Application["app3"] = &KeyConfigV3{
		Icon:    "icon3.png",
		Command: "echo app3",
	}
	
	if CompareKeys(key1, key2) {
		t.Error("Keys with different number of applications should not be equal")
	}
	
	// Test case 3: Different application configs
	delete(key2.Application, "app3")
	key2.Application["app1"].Icon = "different.png"
	
	if CompareKeys(key1, key2) {
		t.Error("Keys with different application configs should not be equal")
	}
	
	// Test case 4: Missing application in key2
	key2 = KeyV3{
		Application: map[string]*KeyConfigV3{
			"app1": {
				Icon:    "icon1.png",
				Command: "echo app1",
			},
		},
	}
	
	if CompareKeys(key1, key2) {
		t.Error("Keys with missing applications should not be equal")
	}
	
	// Test case 5: Missing application in key1
	key1 = KeyV3{
		Application: map[string]*KeyConfigV3{
			"app1": {
				Icon:    "icon1.png",
				Command: "echo app1",
			},
		},
	}
	
	if !CompareKeys(key1, key2) {
		t.Error("Keys with same applications should be equal")
	}
}