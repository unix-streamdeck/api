package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/draw"
	"image/png"
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/assert"
	"github.com/unix-streamdeck/api/v2/mocks/mock_api"
	"github.com/unix-streamdeck/api/v2/mocks/mock_dbus"
	"go.uber.org/mock/gomock"
)

func TestGetInfo(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	sdInfo := StreamDeckInfoV1{
		Cols:             8,
		Rows:             4,
		IconSize:         72,
		Page:             0,
		Serial:           "AABBCCDDE",
		Name:             "FakeDeck",
		Connected:        true,
		LastConnected:    time.Time{},
		LastDisconnected: time.Time{},
		LcdWidth:         0,
		LcdHeight:        0,
		LcdCols:          0,
		KnobCols:         0,
	}

	sdInfoString, _ := json.Marshal([]StreamDeckInfoV1{sdInfo})

	call := dbus.Call{
		Body: []any{string(sdInfoString)},
	}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.GetDeckInfo", dbus.Flags(0)).Times(1).Return(&call)

	actual, err := conn.GetInfo()

	assertions.Nil(err)

	assertions.Equal([]*StreamDeckInfoV1{&sdInfo}, actual)
}

func TestSetPage(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.SetPage", dbus.Flags(0), "AABBCCDD", 4).Times(1).Return(&dbus.Call{})

	err := conn.SetPage("AABBCCDD", 4)

	assertions.Nil(err)
}

func TestGetConfig(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	expectedConfig := ConfigV3{
		Modules: []string{
			"/test/module.so",
		},
		Decks: []DeckV3{
			{
				Serial: "AABBCCDD",
				Pages: []PageV3{
					{
						Keys: []KeyV3{
							{},
						},
					},
				},
			},
		},
	}

	expectedConfigString, _ := json.Marshal(expectedConfig)

	call := dbus.Call{
		Body: []any{string(expectedConfigString)},
	}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.GetConfig", dbus.Flags(0)).Times(1).Return(&call)

	actualConfig, err := conn.GetConfig()

	assertions.Nil(err)

	assertions.Equal(&expectedConfig, actualConfig)
}

func TestSetConfig(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	expectedConfig := ConfigV3{
		Modules: []string{
			"/test/module.so",
		},
		Decks: []DeckV3{
			{
				Serial: "AABBCCDD",
				Pages: []PageV3{
					{
						Keys: []KeyV3{
							{},
						},
					},
				},
			},
		},
	}

	expectedConfigString, _ := json.Marshal(expectedConfig)

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.SetConfig", dbus.Flags(0), string(expectedConfigString)).Times(1).Return(&dbus.Call{})

	err := conn.SetConfig(&expectedConfig)

	assertions.Nil(err)
}

func TestReloadConfig(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.ReloadConfig", dbus.Flags(0)).Times(1).Return(&dbus.Call{})

	err := conn.ReloadConfig()

	assertions.Nil(err)
}

func TestCommitConfig(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.CommitConfig", dbus.Flags(0)).Times(1).Return(&dbus.Call{})

	err := conn.CommitConfig()

	assertions.Nil(err)
}

func TestGetModules(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	modules := []*Module{
		{
			Name: "Example1",
			ForegroundFields: []Field{
				{
					Title:     "File",
					Name:      "file",
					Type:      File,
					FileTypes: []string{"*.gif"},
				},
			},
			InputFields: []Field{
				{
					Title: "Something",
					Name:  "something",
					Type:  Text,
				},
			},
			LinkedFields: nil,
			IsForeground: true,
			IsInput:      true,
		},
	}

	modulesString, _ := json.Marshal(modules)

	call := &dbus.Call{
		Body: []any{string(modulesString)},
	}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.GetModules", dbus.Flags(0)).Times(1).Return(call)

	actualModules, err := conn.GetModules()

	assertions.Nil(err)

	assertions.Equal(modules, actualModules)

}

func TestPressButton(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.PressButton", dbus.Flags(0), "AABBCCDD", 12).Times(1).Return(&dbus.Call{})

	err := conn.PressButton("AABBCCDD", 12)

	assertions.Nil(err)
}

func TestGetObsFields(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	expectedFields := []*Field{
		{
			Title:     "File",
			Name:      "file",
			Type:      File,
			FileTypes: []string{"*.gif"},
		},
		{
			Title: "Something",
			Name:  "something",
			Type:  Text,
		},
	}

	fieldsString, _ := json.Marshal(expectedFields)

	call := &dbus.Call{
		Body: []any{string(fieldsString)},
	}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.GetObsFields", dbus.Flags(0)).Times(1).Return(call)

	actualFields, err := conn.GetObsFields()

	assertions.Nil(err)

	assertions.Equal(expectedFields, actualFields)
}

func TestGetHandlerExample(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	keyConfig := KeyConfigV3{
		Icon:       "/path/to/file.png",
		Command:    "ls",
		Brightness: 65,
	}

	keyConfigString, _ := json.Marshal(keyConfig)

	expectedImg := image.NewRGBA(image.Rect(0, 0, 72, 72))
	draw.Draw(expectedImg, expectedImg.Bounds(), image.Black, image.ZP, draw.Src)

	buf := new(bytes.Buffer)
	png.Encode(buf, expectedImg)
	imageBits := buf.Bytes()
	pngString := "data:image/png;base64," + base64.StdEncoding.EncodeToString(imageBits)

	call := &dbus.Call{
		Body: []any{pngString},
	}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.GetHandlerExample", dbus.Flags(0), "AABBCCDD", string(keyConfigString)).Times(1).Return(call)

	actualImg, err := conn.GetHandlerExample("AABBCCDD", keyConfig)

	assertions.Nil(err)

	assertions.Equal(expectedImg, actualImg)
}

func TestGetKnobHandlerExample(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	busObject := mock_dbus.NewMockBusObject(ctrl)

	conn := Connection{busobj: busObject}

	knobConfig := KnobConfigV3{
		Icon:     "/path/to/file.png",
		FontFace: "regular",
	}

	knobConfigString, _ := json.Marshal(knobConfig)

	expectedImg := image.NewRGBA(image.Rect(0, 0, 72, 72))
	draw.Draw(expectedImg, expectedImg.Bounds(), image.Black, image.ZP, draw.Src)

	buf := new(bytes.Buffer)
	png.Encode(buf, expectedImg)
	imageBits := buf.Bytes()
	pngString := "data:image/png;base64," + base64.StdEncoding.EncodeToString(imageBits)

	call := &dbus.Call{
		Body: []any{pngString},
	}

	busObject.EXPECT().Call("com.unixstreamdeck.streamdeckd.GetKnobHandlerExample", dbus.Flags(0), "AABBCCDD", string(knobConfigString)).Times(1).Return(call)

	actualImg, err := conn.GetKnobHandlerExample("AABBCCDD", knobConfig)

	assertions.Nil(err)

	assertions.Equal(expectedImg, actualImg)
}

func TestRegisterPageListener(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	c := mock_api.NewMockIConn(ctrl)

	conn := Connection{conn: c}

	c.EXPECT().AddMatchSignal(dbus.WithMatchObjectPath("/com/unixstreamdeck/streamdeckd"), dbus.WithMatchInterface("com.unixstreamdeck.streamdeckd"), dbus.WithMatchMember("Page")).Times(1).Return(nil)

	var ch chan<- *dbus.Signal

	c.EXPECT().Signal(gomock.Any()).Times(1).Do(func(channel chan<- *dbus.Signal) {
		ch = channel
		channel <- &dbus.Signal{
			Body: []any{"AABBCCDD", int32(12)},
		}
	})

	err := conn.RegisterPageListener(func(serial string, page int32) {
		assertions.Equal("AABBCCDD", serial)
		assertions.Equal(int32(12), page)
		close(ch)
	})

	assertions.Nil(err)
}
