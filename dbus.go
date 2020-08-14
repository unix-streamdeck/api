package streamdeckd_dbus_lib

import (
	"encoding/json"
	dbus "github.com/godbus/dbus/v5"
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
		busobj: conn.Object("com.thejonsey.streamdeckd", "/com/thejonsey/streamdeckd"),
	}, nil
}

func (c *Connection) Close()  {
	c.conn.Close()
}

func (c *Connection) GetInfo() (*StreamDeckInfo, error) {
	var s string
	err := c.busobj.Call("com.thejonsey.streamdeckd.GetDeckInfo", 0).Store(&s)
	if err != nil {
		return nil, err
	}
	var info *StreamDeckInfo
	err = json.Unmarshal([]byte(s), &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c *Connection) SetPage(page int) error {
	call := c.busobj.Call("com.thejonsey.streamdeckd.SetPage", 0, page)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func (c *Connection) GetConfig() (*Config, error) {
	var s string
	err := c.busobj.Call("com.thejonsey.streamdeckd.GetConfig", 0).Store(&s)
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
	call := c.busobj.Call("com.thejonsey.streamdeckd.SetConfig", 0, string(configString))
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func (c *Connection) ReloadConfig() error {
	call := c.busobj.Call("com.thejonsey.streamdeckd.ReloadConfig", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func (c *Connection) CommitConfig() error {
	call := c.busobj.Call("com.thejonsey.streamdeckd.CommitConfig", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func (c *Connection) RegisterPageListener(cback func(int32)) error {
	err := c.conn.AddMatchSignal(dbus.WithMatchObjectPath("/com/thejonsey/streamdeckd"), dbus.WithMatchInterface("com.thejonsey.streamdeckd"), dbus.WithMatchMember("Page"))
	if err != nil {
		return err
	}
	ch := make(chan *dbus.Signal, 10)
	c.conn.Signal(ch)
	for v := range ch {
		cback(v.Body[0].(int32))
	}
	return nil
}