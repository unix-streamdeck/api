package api

import (
	"image"
	"testing"
)

// MockIconHandler implements the IconHandler interface for testing
type MockIconHandler struct {
	running bool
	started bool
	stopped bool
}

func (h *MockIconHandler) Start(key KeyConfigV3, info StreamDeckInfoV1, callback func(image image.Image)) {
	h.started = true
	h.running = true

	// Create a simple image and call the callback
	img := image.NewRGBA(image.Rect(0, 0, info.IconSize, info.IconSize))
	callback(img)
}

func (h *MockIconHandler) IsRunning() bool {
	return h.running
}

func (h *MockIconHandler) SetRunning(running bool) {
	h.running = running
}

func (h *MockIconHandler) Stop() {
	h.stopped = true
	h.running = false
}

// MockKeyHandler implements the KeyHandler interface for testing
type MockKeyHandler struct {
	keyPressed bool
	lastKey    KeyConfigV3
	lastInfo   StreamDeckInfoV1
}

func (h *MockKeyHandler) Key(key KeyConfigV3, info StreamDeckInfoV1) {
	h.keyPressed = true
	h.lastKey = key
	h.lastInfo = info
}

// MockLcdHandler implements the LcdHandler interface for testing
type MockLcdHandler struct {
	running bool
	started bool
	stopped bool
}

func (h *MockLcdHandler) Start(key KnobConfigV3, info StreamDeckInfoV1, callback func(image image.Image)) {
	h.started = true
	h.running = true

	// Create a simple image and call the callback
	img := image.NewRGBA(image.Rect(0, 0, info.LcdWidth, info.LcdHeight))
	callback(img)
}

func (h *MockLcdHandler) IsRunning() bool {
	return h.running
}

func (h *MockLcdHandler) SetRunning(running bool) {
	h.running = running
}

func (h *MockLcdHandler) Stop() {
	h.stopped = true
	h.running = false
}

// MockKnobOrTouchHandler implements the KnobOrTouchHandler interface for testing
type MockKnobOrTouchHandler struct {
	inputReceived bool
	lastKey       KnobConfigV3
	lastInfo      StreamDeckInfoV1
	lastEvent     InputEvent
}

func (h *MockKnobOrTouchHandler) Input(key KnobConfigV3, info StreamDeckInfoV1, event InputEvent) {
	h.inputReceived = true
	h.lastKey = key
	h.lastInfo = info
	h.lastEvent = event
}

func TestIconHandler(t *testing.T) {
	// Create a mock icon handler
	handler := &MockIconHandler{}

	// Create a key config and stream deck info
	key := KeyConfigV3{
		Icon:    "test.png",
		Text:    "Test",
		Command: "echo test",
	}

	info := StreamDeckInfoV1{
		Cols:     8,
		Rows:     4,
		IconSize: 96,
		Serial:   "ABCD1234",
	}

	// Create a callback function
	var callbackCalled bool
	var callbackImage image.Image
	callback := func(img image.Image) {
		callbackCalled = true
		callbackImage = img
	}

	// Start the handler
	handler.Start(key, info, callback)

	// Check that the handler was started and is running
	if !handler.started {
		t.Error("Handler should be started")
	}

	if !handler.IsRunning() {
		t.Error("Handler should be running")
	}

	// Check that the callback was called with an image
	if !callbackCalled {
		t.Error("Callback should have been called")
	}

	if callbackImage == nil {
		t.Error("Callback should have been called with an image")
	}

	// Stop the handler
	handler.Stop()

	// Check that the handler was stopped and is not running
	if !handler.stopped {
		t.Error("Handler should be stopped")
	}

	if handler.IsRunning() {
		t.Error("Handler should not be running")
	}

	// Test SetRunning
	handler.SetRunning(true)
	if !handler.IsRunning() {
		t.Error("Handler should be running after SetRunning(true)")
	}

	handler.SetRunning(false)
	if handler.IsRunning() {
		t.Error("Handler should not be running after SetRunning(false)")
	}
}

func TestKeyHandler(t *testing.T) {
	// Create a mock key handler
	handler := &MockKeyHandler{}

	// Create a key config and stream deck info
	key := KeyConfigV3{
		Icon:    "test.png",
		Text:    "Test",
		Command: "echo test",
	}

	info := StreamDeckInfoV1{
		Cols:     8,
		Rows:     4,
		IconSize: 96,
		Serial:   "ABCD1234",
	}

	// Call the Key method
	handler.Key(key, info)

	// Check that the key was pressed
	if !handler.keyPressed {
		t.Error("Key should have been pressed")
	}

	// Check that the key and info were stored
	if handler.lastKey.Icon != key.Icon || handler.lastKey.Text != key.Text || handler.lastKey.Command != key.Command {
		t.Error("Last key does not match the key passed to Key()")
	}

	if handler.lastInfo.Serial != info.Serial || handler.lastInfo.IconSize != info.IconSize {
		t.Error("Last info does not match the info passed to Key()")
	}
}

func TestLcdHandler(t *testing.T) {
	// Create a mock LCD handler
	handler := &MockLcdHandler{}

	// Create a knob config and stream deck info
	key := KnobConfigV3{
		Icon: "test.png",
		Text: "Test",
		KnobPressAction: KnobActionV3{
			Command: "echo press",
		},
	}

	info := StreamDeckInfoV1{
		Cols:      8,
		Rows:      4,
		IconSize:  96,
		Serial:    "ABCD1234",
		LcdWidth:  200,
		LcdHeight: 100,
	}

	// Create a callback function
	var callbackCalled bool
	var callbackImage image.Image
	callback := func(img image.Image) {
		callbackCalled = true
		callbackImage = img
	}

	// Start the handler
	handler.Start(key, info, callback)

	// Check that the handler was started and is running
	if !handler.started {
		t.Error("Handler should be started")
	}

	if !handler.IsRunning() {
		t.Error("Handler should be running")
	}

	// Check that the callback was called with an image
	if !callbackCalled {
		t.Error("Callback should have been called")
	}

	if callbackImage == nil {
		t.Error("Callback should have been called with an image")
	}

	// Stop the handler
	handler.Stop()

	// Check that the handler was stopped and is not running
	if !handler.stopped {
		t.Error("Handler should be stopped")
	}

	if handler.IsRunning() {
		t.Error("Handler should not be running")
	}
}

func TestKnobOrTouchHandler(t *testing.T) {
	// Create a mock knob or touch handler
	handler := &MockKnobOrTouchHandler{}

	// Create a knob config and stream deck info
	key := KnobConfigV3{
		Icon: "test.png",
		Text: "Test",
		KnobPressAction: KnobActionV3{
			Command: "echo press",
		},
	}

	info := StreamDeckInfoV1{
		Cols:     8,
		Rows:     4,
		IconSize: 96,
		Serial:   "ABCD1234",
	}

	// Create an input event
	event := InputEvent{
		EventType:     KNOB_PRESS,
		RotateNotches: 0,
	}

	// Call the Input method
	handler.Input(key, info, event)

	// Check that the input was received
	if !handler.inputReceived {
		t.Error("Input should have been received")
	}

	// Check that the key, info, and event were stored
	if handler.lastKey.Icon != key.Icon || handler.lastKey.Text != key.Text {
		t.Error("Last key does not match the key passed to Input()")
	}

	if handler.lastInfo.Serial != info.Serial || handler.lastInfo.IconSize != info.IconSize {
		t.Error("Last info does not match the info passed to Input()")
	}

	if handler.lastEvent.EventType != event.EventType || handler.lastEvent.RotateNotches != event.RotateNotches {
		t.Error("Last event does not match the event passed to Input()")
	}

	// Test with a different event type
	event = InputEvent{
		EventType:     KNOB_CW,
		RotateNotches: 2,
	}

	handler.Input(key, info, event)

	if handler.lastEvent.EventType != event.EventType || handler.lastEvent.RotateNotches != event.RotateNotches {
		t.Error("Last event does not match the event passed to Input()")
	}
}

func TestInputEventTypes(t *testing.T) {
	// Test that the input event types are defined correctly
	if KNOB_CCW != 0 {
		t.Errorf("KNOB_CCW should be 0, got %d", KNOB_CCW)
	}

	if KNOB_CW != 1 {
		t.Errorf("KNOB_CW should be 1, got %d", KNOB_CW)
	}

	if KNOB_PRESS != 2 {
		t.Errorf("KNOB_PRESS should be 2, got %d", KNOB_PRESS)
	}

	if SCREEN_SHORT_TAP != 3 {
		t.Errorf("SCREEN_SHORT_TAP should be 3, got %d", SCREEN_SHORT_TAP)
	}

	if SCREEN_LONG_TAP != 4 {
		t.Errorf("SCREEN_LONG_TAP should be 4, got %d", SCREEN_LONG_TAP)
	}
}
