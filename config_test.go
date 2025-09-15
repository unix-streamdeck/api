package api

import (
	"encoding/json"
	"testing"
	"time"
)

func TestConfigV1Serialization(t *testing.T) {
	// Create a ConfigV1 instance
	config := ConfigV1{
		Modules: []string{"module1", "module2"},
		Pages: []PageV1{
			{
				{
					Icon:       "icon1.png",
					SwitchPage: 1,
					Text:       "Button 1",
					TextSize:   12,
					Command:    "echo hello",
				},
				{
					Icon:       "icon2.png",
					SwitchPage: 2,
					Text:       "Button 2",
					TextSize:   14,
					Command:    "echo world",
				},
			},
			{
				{
					Icon:       "icon3.png",
					SwitchPage: 0,
					Text:       "Button 3",
					TextSize:   16,
					Command:    "echo test",
				},
			},
		},
	}

	// Serialize to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Errorf("Failed to marshal ConfigV1: %v", err)
		t.FailNow()
	}

	// Deserialize from JSON
	var newConfig ConfigV1
	err = json.Unmarshal(data, &newConfig)
	if err != nil {
		t.Errorf("Failed to unmarshal ConfigV1: %v", err)
		t.FailNow()
	}

	// Check that the deserialized config matches the original
	if len(newConfig.Modules) != len(config.Modules) {
		t.Errorf("Expected %d modules, got %d", len(config.Modules), len(newConfig.Modules))
	}

	for i, module := range config.Modules {
		if newConfig.Modules[i] != module {
			t.Errorf("Expected module %s, got %s", module, newConfig.Modules[i])
		}
	}

	if len(newConfig.Pages) != len(config.Pages) {
		t.Errorf("Expected %d pages, got %d", len(config.Pages), len(newConfig.Pages))
	}

	for i, page := range config.Pages {
		if len(newConfig.Pages[i]) != len(page) {
			t.Errorf("Expected page %d to have %d keys, got %d", i, len(page), len(newConfig.Pages[i]))
		}

		for j, key := range page {
			if newConfig.Pages[i][j].Icon != key.Icon {
				t.Errorf("Expected key %d on page %d to have icon %s, got %s", j, i, key.Icon, newConfig.Pages[i][j].Icon)
			}
			if newConfig.Pages[i][j].SwitchPage != key.SwitchPage {
				t.Errorf("Expected key %d on page %d to have switch page %d, got %d", j, i, key.SwitchPage, newConfig.Pages[i][j].SwitchPage)
			}
			if newConfig.Pages[i][j].Text != key.Text {
				t.Errorf("Expected key %d on page %d to have text %s, got %s", j, i, key.Text, newConfig.Pages[i][j].Text)
			}
			if newConfig.Pages[i][j].TextSize != key.TextSize {
				t.Errorf("Expected key %d on page %d to have text size %d, got %d", j, i, key.TextSize, newConfig.Pages[i][j].TextSize)
			}
			if newConfig.Pages[i][j].Command != key.Command {
				t.Errorf("Expected key %d on page %d to have command %s, got %s", j, i, key.Command, newConfig.Pages[i][j].Command)
			}
		}
	}
}

func TestConfigV2Serialization(t *testing.T) {
	// Create a ConfigV2 instance
	config := ConfigV2{
		Modules: []string{"module1", "module2"},
		Decks: []DeckV2{
			{
				Serial: "ABCD1234",
				Pages: []PageV1{
					{
						{
							Icon:       "icon1.png",
							SwitchPage: 1,
							Text:       "Button 1",
							TextSize:   12,
							Command:    "echo hello",
						},
						{
							Icon:       "icon2.png",
							SwitchPage: 2,
							Text:       "Button 2",
							TextSize:   14,
							Command:    "echo world",
						},
					},
				},
			},
		},
		ObsConnectionInfo: ObsConnectionInfoV2{
			Host: "localhost",
			Port: 4444,
		},
	}

	// Serialize to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Errorf("Failed to marshal ConfigV2: %v", err)
		t.FailNow()
	}

	// Deserialize from JSON
	var newConfig ConfigV2
	err = json.Unmarshal(data, &newConfig)
	if err != nil {
		t.Errorf("Failed to unmarshal ConfigV2: %v", err)
		t.FailNow()
	}

	// Check that the deserialized config matches the original
	if len(newConfig.Modules) != len(config.Modules) {
		t.Errorf("Expected %d modules, got %d", len(config.Modules), len(newConfig.Modules))
	}

	if len(newConfig.Decks) != len(config.Decks) {
		t.Errorf("Expected %d decks, got %d", len(config.Decks), len(newConfig.Decks))
	}

	if newConfig.Decks[0].Serial != config.Decks[0].Serial {
		t.Errorf("Expected deck serial %s, got %s", config.Decks[0].Serial, newConfig.Decks[0].Serial)
	}

	if newConfig.ObsConnectionInfo.Host != config.ObsConnectionInfo.Host {
		t.Errorf("Expected OBS host %s, got %s", config.ObsConnectionInfo.Host, newConfig.ObsConnectionInfo.Host)
	}

	if newConfig.ObsConnectionInfo.Port != config.ObsConnectionInfo.Port {
		t.Errorf("Expected OBS port %d, got %d", config.ObsConnectionInfo.Port, newConfig.ObsConnectionInfo.Port)
	}
}

func TestConfigV3Serialization(t *testing.T) {
	// Create a ConfigV3 instance
	config := ConfigV3{
		Modules: []string{"module1", "module2"},
		Decks: []DeckV3{
			{
				Serial: "ABCD1234",
				Pages: []PageV3{
					{
						Keys: []KeyV3{
							{
								Application: map[string]*KeyConfigV3{
									"app1": {
										Icon:       "icon1.png",
										SwitchPage: 1,
										Text:       "Button 1",
										TextSize:   12,
										Command:    "echo hello",
									},
								},
							},
							{
								Application: map[string]*KeyConfigV3{
									"app2": {
										Icon:       "icon2.png",
										SwitchPage: 2,
										Text:       "Button 2",
										TextSize:   14,
										Command:    "echo world",
									},
								},
							},
						},
						Knobs: []KnobV3{
							{
								Application: map[string]*KnobConfigV3{
									"app1": {
										Icon:     "knob1.png",
										Text:     "Knob 1",
										TextSize: 12,
										KnobPressAction: KnobActionV3{
											Command: "echo press",
										},
										KnobTurnUpAction: KnobActionV3{
											Command: "echo up",
										},
										KnobTurnDownAction: KnobActionV3{
											Command: "echo down",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		ObsConnectionInfo: ObsConnectionInfoV2{
			Host: "localhost",
			Port: 4444,
		},
	}

	// Serialize to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Errorf("Failed to marshal ConfigV3: %v", err)
		t.FailNow()
	}

	// Deserialize from JSON
	var newConfig ConfigV3
	err = json.Unmarshal(data, &newConfig)
	if err != nil {
		t.Errorf("Failed to unmarshal ConfigV3: %v", err)
		t.FailNow()
	}

	// Check that the deserialized config matches the original
	if len(newConfig.Modules) != len(config.Modules) {
		t.Errorf("Expected %d modules, got %d", len(config.Modules), len(newConfig.Modules))
	}

	if len(newConfig.Decks) != len(config.Decks) {
		t.Errorf("Expected %d decks, got %d", len(config.Decks), len(newConfig.Decks))
	}

	if newConfig.Decks[0].Serial != config.Decks[0].Serial {
		t.Errorf("Expected deck serial %s, got %s", config.Decks[0].Serial, newConfig.Decks[0].Serial)
	}

	if len(newConfig.Decks[0].Pages) != len(config.Decks[0].Pages) {
		t.Errorf("Expected %d pages, got %d", len(config.Decks[0].Pages), len(newConfig.Decks[0].Pages))
	}

	if len(newConfig.Decks[0].Pages[0].Keys) != len(config.Decks[0].Pages[0].Keys) {
		t.Errorf("Expected %d keys, got %d", len(config.Decks[0].Pages[0].Keys), len(newConfig.Decks[0].Pages[0].Keys))
	}

	if len(newConfig.Decks[0].Pages[0].Knobs) != len(config.Decks[0].Pages[0].Knobs) {
		t.Errorf("Expected %d knobs, got %d", len(config.Decks[0].Pages[0].Knobs), len(newConfig.Decks[0].Pages[0].Knobs))
	}
}

func TestStreamDeckInfoV1Serialization(t *testing.T) {
	// Create a StreamDeckInfoV1 instance
	now := time.Now()
	info := StreamDeckInfoV1{
		Cols:             8,
		Rows:             4,
		IconSize:         96,
		Page:             0,
		Serial:           "ABCD1234",
		Name:             "Elgato Stream Deck XL",
		Connected:        true,
		LastConnected:    now,
		LastDisconnected: time.Time{},
	}

	// Serialize to JSON
	data, err := json.Marshal(info)
	if err != nil {
		t.Errorf("Failed to marshal StreamDeckInfoV1: %v", err)
		t.FailNow()
	}

	// Deserialize from JSON
	var newInfo StreamDeckInfoV1
	err = json.Unmarshal(data, &newInfo)
	if err != nil {
		t.Errorf("Failed to unmarshal StreamDeckInfoV1: %v", err)
		t.FailNow()
	}

	// Check that the deserialized info matches the original
	if newInfo.Cols != info.Cols {
		t.Errorf("Expected cols %d, got %d", info.Cols, newInfo.Cols)
	}
	if newInfo.Rows != info.Rows {
		t.Errorf("Expected rows %d, got %d", info.Rows, newInfo.Rows)
	}
	if newInfo.IconSize != info.IconSize {
		t.Errorf("Expected icon size %d, got %d", info.IconSize, newInfo.IconSize)
	}
	if newInfo.Page != info.Page {
		t.Errorf("Expected page %d, got %d", info.Page, newInfo.Page)
	}
	if newInfo.Serial != info.Serial {
		t.Errorf("Expected serial %s, got %s", info.Serial, newInfo.Serial)
	}
	if newInfo.Name != info.Name {
		t.Errorf("Expected name %s, got %s", info.Name, newInfo.Name)
	}
	if newInfo.Connected != info.Connected {
		t.Errorf("Expected connected %t, got %t", info.Connected, newInfo.Connected)
	}
}

func TestModuleSerialization(t *testing.T) {
	// Create a Module instance
	module := Module{
		Name:   "TestModule",
		IsIcon: true,
		IsKey:  false,
		IconFields: []Field{
			{
				Title:     "Icon",
				Name:      "icon",
				Type:      "File",
				FileTypes: []string{".png", ".jpg"},
			},
			{
				Title:     "Text",
				Name:      "text",
				Type:      "Text",
				FileTypes: nil,
			},
		},
	}

	// Serialize to JSON
	data, err := json.Marshal(module)
	if err != nil {
		t.Errorf("Failed to marshal Module: %v", err)
		t.FailNow()
	}

	// Deserialize from JSON
	var newModule Module
	err = json.Unmarshal(data, &newModule)
	if err != nil {
		t.Errorf("Failed to unmarshal Module: %v", err)
		t.FailNow()
	}

	// Check that the deserialized module matches the original
	if newModule.Name != module.Name {
		t.Errorf("Expected name %s, got %s", module.Name, newModule.Name)
	}
	if newModule.IsIcon != module.IsIcon {
		t.Errorf("Expected IsIcon %t, got %t", module.IsIcon, newModule.IsIcon)
	}
	if newModule.IsKey != module.IsKey {
		t.Errorf("Expected IsKey %t, got %t", module.IsKey, newModule.IsKey)
	}
	if len(newModule.IconFields) != len(module.IconFields) {
		t.Errorf("Expected %d icon fields, got %d", len(module.IconFields), len(newModule.IconFields))
	}

	for i, field := range module.IconFields {
		if newModule.IconFields[i].Title != field.Title {
			t.Errorf("Expected field title %s, got %s", field.Title, newModule.IconFields[i].Title)
		}
		if newModule.IconFields[i].Name != field.Name {
			t.Errorf("Expected field name %s, got %s", field.Name, newModule.IconFields[i].Name)
		}
		if newModule.IconFields[i].Type != field.Type {
			t.Errorf("Expected field type %s, got %s", field.Type, newModule.IconFields[i].Type)
		}
		if len(field.FileTypes) > 0 {
			if len(newModule.IconFields[i].FileTypes) != len(field.FileTypes) {
				t.Errorf("Expected %d file types, got %d", len(field.FileTypes), len(newModule.IconFields[i].FileTypes))
			}
			for j, fileType := range field.FileTypes {
				if newModule.IconFields[i].FileTypes[j] != fileType {
					t.Errorf("Expected file type %s, got %s", fileType, newModule.IconFields[i].FileTypes[j])
				}
			}
		}
	}
}
