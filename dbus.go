package api

import (
	"bytes"
	"encoding/json"
	"github.com/godbus/dbus/v5"
	"image"
	"image/png"
	"encoding/base64"
	"strings"
)
type Connection struct {
	busobj dbus.BusObject
	conn *dbus.Conn
}

func Connect() (*Connection, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return nil, err
	}
	return &Connection{
		conn: conn,
		busobj: conn.Object("com.unixstreamdeck.streamdeckd", "/com/unixstreamdeck/streamdeckd"),
	}, nil
}

func (c *Connection) Close()  {
	c.conn.Close()
}

func (c *Connection) GetInfo() ([]*StreamDeckInfo, error) {
	var s string
	err := c.busobj.Call("com.unixstreamdeck.streamdeckd.GetDeckInfo", 0).Store(&s)
	if err != nil {
		return nil, err
	}
	var info []*StreamDeckInfo
	err = json.Unmarshal([]byte(s), &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c *Connection) SetPage(serial string, page int) error {
	call := c.busobj.Call("com.unixstreamdeck.streamdeckd.SetPage", 0, serial, page)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func (c *Connection) GetConfig() (*Config, error) {
	var s string
	err := c.busobj.Call("com.unixstreamdeck.streamdeckd.GetConfig", 0).Store(&s)
	if err != nil {
		return nil, err
	}
	var config *Config
	err = json.Unmarshal([]byte(s), &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Connection) SetConfig(config *Config) error {
	configString, err := json.Marshal(config)
	if err != nil {
		return err
	}
	call := c.busobj.Call("com.unixstreamdeck.streamdeckd.SetConfig", 0, string(configString))
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func (c *Connection) ReloadConfig() error {
	call := c.busobj.Call("com.unixstreamdeck.streamdeckd.ReloadConfig", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func (c *Connection) CommitConfig() error {
	call := c.busobj.Call("com.unixstreamdeck.streamdeckd.CommitConfig", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func (c *Connection) GetModules() ([]*Module, error) {
	var s string
	err := c.busobj.Call("com.unixstreamdeck.streamdeckd.GetModules", 0).Store(&s)
	if err != nil {
		return nil, err
	}
	var modules []*Module
	err = json.Unmarshal([]byte(s), &modules)
	if err != nil {
		return nil, err
	}
	return modules, nil
}

func (c *Connection) PressButton(serial string, keyIndex int) error  {
	return c.busobj.Call("com.unixstreamdeck.streamdeckd.PressButton", 0, serial, keyIndex).Err
}

func (c *Connection) GetHandlerExample(serial string, key Key) (image.Image, error) {
	keyString, err := json.Marshal(key)
	if err != nil {
		return nil, err
	}
	var s string
	err = c.busobj.Call("com.unixstreamdeck.streamdeckd.GetHandlerExample", 0, serial, string(keyString)).Store(&s)
	if err != nil {
		return nil, err
	}
	imgData := s[strings.IndexByte(s, ',')+1:]
	decodedImgData, err := base64.StdEncoding.DecodeString(imgData)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(decodedImgData)
	return png.Decode(reader)
}

func (c *Connection) RegisterPageListener(cback func(string, int32)) error {
	err := c.conn.AddMatchSignal(dbus.WithMatchObjectPath("/com/unixstreamdeck/streamdeckd"), dbus.WithMatchInterface("com.unixstreamdeck.streamdeckd"), dbus.WithMatchMember("Page"))
	if err != nil {
		return err
	}
	ch := make(chan *dbus.Signal, 10)
	c.conn.Signal(ch)
	for v := range ch {
		cback(v.Body[0].(string), v.Body[1].(int32))
	}
	return nil
}