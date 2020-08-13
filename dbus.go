package streamdeckd_dbus_lib

import (
	"encoding/json"
	"errors"
	"github.com/godbus/dbus/v5"
)

var busobj dbus.BusObject
var conn *dbus.Conn

func InitDBUS() error {
	var err error
	conn, err = dbus.ConnectSessionBus()
	if err != nil {
		return err
	}
	busobj = conn.Object("com.thejonsey.streamdeckd", "/com/thejonsey/streamdeckd")
	return nil
}

func Close()  {
	conn.Close()
}

func GetInfo() (*StreamDeckInfo, error) {
	if busobj == nil {
		return nil, errors.New("DBus not connected")
	}
	var s string
	err := busobj.Call("com.thejonsey.streamdeckd.GetDeckInfo", 0).Store(&s)
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

func SetPage(page int) error {
	if busobj == nil {
		return errors.New("DBus not connected")
	}
	call := busobj.Call("com.thejonsey.streamdeckd.SetPage", 0, page)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func GetConfig() (*Config, error) {
	if busobj == nil {
		return nil, errors.New("DBus not connected")
	}
	var s string
	err := busobj.Call("com.thejonsey.streamdeckd.GetConfig", 0).Store(&s)
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

func SetConfig(config *Config) error {
	if busobj == nil {
		return errors.New("DBus not connected")
	}
	configString, err := json.Marshal(config)
	if err != nil {
		return err
	}
	call := busobj.Call("com.thejonsey.streamdeckd.SetConfig", 0, string(configString))
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func ReloadConfig() error {
	if busobj == nil {
		return errors.New("DBus not connected")
	}
	call := busobj.Call("com.thejonsey.streamdeckd.ReloadConfig", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func CommitConfig() error {
	if busobj == nil {
		return errors.New("DBus not connected")
	}
	call := busobj.Call("com.thejonsey.streamdeckd.CommitConfig", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func RegisterPageListener(cback func(int32)) error {
	err := conn.AddMatchSignal(dbus.WithMatchObjectPath("/com/thejonsey/streamdeckd"), dbus.WithMatchInterface("com.thejonsey.streamdeckd"), dbus.WithMatchMember("Page"))
	if err != nil {
		return err
	}
	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)
	for v := range c {
		cback(v.Body[0].(int32))
	}
	return nil
}