package api

import "image"

type ConfigV3 struct {
    Modules           []string            `json:"modules,omitempty"`
    Decks             []DeckV3            `json:"decks"`
    ObsConnectionInfo ObsConnectionInfoV2 `json:"obs_connection_info,omitempty"`
}

type DeckV3 struct {
    Serial string   `json:"serial"`
    Pages  []PageV3 `json:"pages"`
}

type PageV3 struct {
    Keys  []KeyV3  `json:"keys"`
    Knobs []KnobV3 `json:"knobs"`
}

type KeyV3 struct {
    Application             map[string]*KeyConfigV3 `json:"application,omitempty"`
    ActiveBuff              image.Image             `json:"-"`
    ActiveIconHandlerStruct *IconHandler            `json:"-"`
    ActiveKeyHandlerStruct  *KeyHandler             `json:"-"`
    ActiveApplication       string                  `json:"-"`
}

type KnobV3 struct {
    Application       map[string]*KnobConfigV3 `json:"application,omitempty"`
    ActiveBuff        image.Image              `json:"-"`
    ActiveApplication string                   `json:"-"`
}

type KeyConfigV3 struct {
    Icon              string            `json:"icon,omitempty"`
    SwitchPage        int               `json:"switch_page,omitempty"`
    Text              string            `json:"text,omitempty"`
    TextSize          int               `json:"text_size,omitempty"`
    TextAlignment     string            `json:"text_alignment,omitempty"`
    Keybind           string            `json:"keybind,omitempty"`
    Command           string            `json:"command,omitempty"`
    Brightness        int               `json:"brightness,omitempty"`
    Url               string            `json:"url,omitempty"`
    ObsCommand        string            `json:"obs_command,omitempty"`
    ObsCommandParams  map[string]string `json:"obs_command_params,omitempty"`
    IconHandler       string            `json:"icon_handler,omitempty"`
    KeyHandler        string            `json:"key_handler,omitempty"`
    IconHandlerFields map[string]string `json:"icon_handler_fields,omitempty"`
    KeyHandlerFields  map[string]string `json:"key_handler_fields,omitempty"`
    Buff              image.Image       `json:"-"`
    IconHandlerStruct IconHandler       `json:"-"`
    KeyHandlerStruct  KeyHandler        `json:"-"`
}
type KnobConfigV3 struct {
    Icon               string      `json:"icon,omitempty"`
    Text               string      `json:"text,omitempty"`
    TextSize           int         `json:"text_size,omitempty"`
    TextAlignment      string      `json:"text_alignment,omitempty"`
    LcdHandler         string `json:"lcd_handler,omitempty"`
    KnobOrTouchHandler string `json:"knob_or_touch_handler,omitempty"`
    Buff               image.Image `json:"-"`
    LcdHandlerStruct LcdHandler `json:"-"`
    KnobOrTouchHandlerStruct KnobOrTouchHandler `json:"-"`
    LcdHandlerFields map[string]string `json:"lcd_handler_fields,omitempty"`
    KnobOrTouchHandlerFields  map[string]string `json:"knob_or_touch_handler_fields,omitempty"`
}

type KnobMoveActionV3 struct {
    Command string `json:"command,omitempty"`
}

type KnobPressActionV3 struct {
    SwitchPage       int               `json:"switch_page,omitempty"`
    Keybind          string            `json:"keybind,omitempty"`
    Command          string            `json:"command,omitempty"`
    Brightness       int               `json:"brightness,omitempty"`
    Url              string            `json:"url,omitempty"`
    ObsCommand       string            `json:"obs_command,omitempty"`
    ObsCommandParams map[string]string `json:"obs_command_params,omitempty"`
}
