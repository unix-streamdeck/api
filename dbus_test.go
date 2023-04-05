package api

import (
    "github.com/godbus/dbus/v5"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/unix-streamdeck/api/mocks"
    "image"
    "testing"
    "time"
)

func TestConnection_SetPage(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.SetPage", dbus.Flags(0), "ABCD1234", 3).Return(&dbus.Call{})

    err := conn.SetPage("ABCD1234", 3)

    assert.Nil(t, err, "No error should be returned")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)
}

func TestConnection_GetInfo(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    body := "[{\"cols\":8,\"rows\":4,\"icon_size\":96,\"page\":0,\"serial\":\"ABCD1234\",\"name\":\"Elgato Stream Deck XL\",\"connected\":true,\"last_connected\":\"2022-11-11T19:33:50.447959984Z\",\"last_disconnected\":\"0001-01-01T00:00:00Z\"}]"

    y := make([]interface{}, 1)

    y[0] = body

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.GetDeckInfo", dbus.Flags(0)).Return(&dbus.Call{Body: y, Err: nil})

    infos, err := conn.GetInfo()

    assert.Nil(t, err, "No error should be returned")

    assert.Equal(t, 1, len(infos), "Info array should contain 1 item")

    actualInfo := infos[0]

    expectedInfo := &StreamDeckInfoV1{
        Rows:             4,
        Cols:             8,
        Serial:           "ABCD1234",
        Connected:        true,
        Name:             "Elgato Stream Deck XL",
        Page:             0,
        IconSize:         96,
        LastConnected:    time.Date(2022, 11, 11, 19, 33, 50, 447959984, time.UTC),
        LastDisconnected: time.Date(0001, 01, 01, 00, 00, 00, 0, time.UTC),
    }

    assert.Equal(t, expectedInfo, actualInfo, "StreamDeckInfo should be equal")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)

}

func TestConnection_GetConfig(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    body := "{\"decks\": [{\"serial\": \"ABCD1234\", \"pages\": [[{\"application\": {\"\": {\"icon\": \"example.png\", \"command\": \"notify-send 'Hello World'\"}}}]]}]}"

    y := make([]interface{}, 1)

    y[0] = body

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.GetConfig", dbus.Flags(0)).Return(&dbus.Call{Body: y, Err: nil})

    actualConfig, err := conn.GetConfig()

    assert.Nil(t, err, "No error should be returned")

    expectedConfig := &ConfigV3{
        Decks: []DeckV3{
            {
                Serial: "ABCD1234", Pages: []PageV3{
                []KeyV3{
                    {
                        Application: map[string]*KeyConfigV3{
                            "": {
                                Icon:    "example.png",
                                Command: "notify-send 'Hello World'",
                            },
                        },
                    },
                },
            }},
        },
    }

    assert.Equal(t, expectedConfig, actualConfig, "Configs should be equal")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)
}

func TestConnection_SetConfig(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    conf := &ConfigV3{
        Decks: []DeckV3{
            {
                Serial: "ABCD1234", Pages: []PageV3{
                []KeyV3{
                    {
                        Application: map[string]*KeyConfigV3{
                            "": {
                                Icon:    "example.png",
                                Command: "notify-send 'Hello World'",
                            },
                        },
                    },
                },
            }},
        },
    }

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.SetConfig", dbus.Flags(0), "{\"decks\":[{\"serial\":\"ABCD1234\",\"pages\":[[{\"application\":{\"\":{\"icon\":\"example.png\",\"command\":\"notify-send 'Hello World'\"}}}]]}],\"obs_connection_info\":{}}").Return(&dbus.Call{})

    err := conn.SetConfig(conf)

    assert.Nil(t, err, "No error should be returned")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)

}

func TestConnection_ReloadConfig(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.ReloadConfig", dbus.Flags(0)).Return(&dbus.Call{})

    err := conn.ReloadConfig()

    assert.Nil(t, err, "No error should be returned")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)
}

func TestConnection_CommitConfig(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.CommitConfig", dbus.Flags(0)).Return(&dbus.Call{})

    err := conn.CommitConfig()

    assert.Nil(t, err, "No error should be returned")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)

}

func TestConnection_GetModules(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    body := "[{\"name\":\"Gif\",\"icon_fields\":[{\"title\":\"Icon\",\"name\":\"icon\",\"type\":\"File\",\"file_types\":[\".gif\"]},{\"title\":\"Text\",\"name\":\"text\",\"type\":\"Text\"},{\"title\":\"Text Size\",\"name\":\"text_size\",\"type\":\"Number\"},{\"title\":\"Text Alignment\",\"name\":\"text_alignment\",\"type\":\"TextAlignment\"}],\"is_icon\":true},{\"name\":\"Time\",\"is_icon\":true},{\"name\":\"Counter\",\"is_icon\":true,\"is_key\":true},{\"name\":\"Spotify\",\"is_icon\":true}]"

    y := make([]interface{}, 1)

    y[0] = body

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.GetModules", dbus.Flags(0)).Return(&dbus.Call{Body: y})

    actualModules, err := conn.GetModules()

    assert.Nil(t, err, "No error should be returned")

    expectedModules := []*Module{
        {
            Name:   "Gif",
            IsIcon: true,
            IconFields: []Field{
                {
                    Title:     "Icon",
                    Name:      "icon",
                    Type:      "File",
                    FileTypes: []string{".gif"}},
                {
                    Title: "Text",
                    Name:  "text",
                    Type:  "Text",
                },
                {
                    Title: "Text Size",
                    Name:  "text_size",
                    Type:  "Number",
                },
                {
                    Title: "Text Alignment",
                    Name:  "text_alignment",
                    Type:  "TextAlignment",
                },
            },
        },
        {
            Name:   "Time",
            IsIcon: true,
        },
        {
            Name:   "Counter",
            IsIcon: true,
            IsKey:  true,
        },
        {
            Name:   "Spotify",
            IsIcon: true,
        },
    }

    assert.Equal(t, expectedModules, actualModules, "Modules should be equal")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)
}

func TestConnection_PressButton(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.PressButton", dbus.Flags(0), "ABCD1234", 18).Return(&dbus.Call{})

    err := conn.PressButton("ABCD1234", 18)

    assert.Nil(t, err, "No error should be returned")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)
}

func TestConnection_GetObsFields(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    body := "[{\"title\":\"Icon\",\"name\":\"icon\",\"type\":\"File\",\"file_types\":[\".gif\"]},{\"title\":\"Text\",\"name\":\"text\",\"type\":\"Text\"},{\"title\":\"Text Size\",\"name\":\"text_size\",\"type\":\"Number\"},{\"title\":\"Text Alignment\",\"name\":\"text_alignment\",\"type\":\"TextAlignment\"}]"

    y := make([]interface{}, 1)

    y[0] = body

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.GetObsFields", dbus.Flags(0)).Return(&dbus.Call{Body: y})

    actualFields, err := conn.GetObsFields()

    assert.Nil(t, err, "No error should be returned")

    expectedFields := []*Field{
        {
            Title:     "Icon",
            Name:      "icon",
            Type:      "File",
            FileTypes: []string{".gif"}},
        {
            Title: "Text",
            Name:  "text",
            Type:  "Text",
        },
        {
            Title: "Text Size",
            Name:  "text_size",
            Type:  "Number",
        },
        {
            Title: "Text Alignment",
            Name:  "text_alignment",
            Type:  "TextAlignment",
        },
    }

    assert.Equal(t, expectedFields, actualFields, "Fields should be equal")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)

}

func TestConnection_GetHandlerExample(t *testing.T) {

    mockBusObj := mocks.NewBusObject(t)

    conn := Connection{busobj: mockBusObj}

    body := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGAAAABgCAYAAADimHc4AAAA3ElEQVR4nOzRQREAAAgEIcf+oS/GfqACf6QExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEBMQExATEFgAA//9H+QDBW18vywAAAABJRU5ErkJggg=="

    y := make([]interface{}, 1)

    y[0] = body

    mockBusObj.On("Call", "com.unixstreamdeck.streamdeckd.GetHandlerExample", dbus.Flags(0), "ABCD1234", "{\"icon_handler\":\"NoOp\"}").Return(&dbus.Call{Body: y})

    actualImg, err := conn.GetHandlerExample("ABCD1234", KeyConfigV3{IconHandler: "NoOp"})

    assert.Nil(t, err, "No error should be returned")

    expectedImg := image.NewNRGBA(image.Rect(0, 0, 96, 96)).SubImage(image.Rect(0, 0, 96, 96))

    assert.Equal(t, expectedImg, actualImg, "Images should match")

    mockBusObj.AssertNumberOfCalls(t, "Call", 1)

}

func TestConnection_RegisterPageListener(t *testing.T) {
    mockConnObj := mocks.NewIConn(t)

    conn := Connection{conn: mockConnObj}

    mockConnObj.On("AddMatchSignal", dbus.WithMatchObjectPath("/com/unixstreamdeck/streamdeckd"), dbus.WithMatchInterface("com.unixstreamdeck.streamdeckd"), dbus.WithMatchMember("Page")).Return(nil)

    mockConnObj.On("Signal", mock.Anything).Return(nil)

    go conn.RegisterPageListener(func(serial string, keyIndex int32) {
        assert.Equal(t, "ABCD1234", serial)
        assert.Equal(t, int32(16), keyIndex)
    })

    y := make([]interface{}, 2)

    y[0] = "ABCD1234"

    y[1] = int32(16)

    time.Sleep(100 * time.Millisecond)

    assert.NotNil(t, mockConnObj.Sig)

    mockConnObj.Sig <- &dbus.Signal{Body: y}

}