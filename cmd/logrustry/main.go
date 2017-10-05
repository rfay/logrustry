package main

import (

	log "github.com/sirupsen/logrus"
)

// Overall strategy:
/**
* Add formatters for user interaction, normal logging, and full json interaction
* Use logrus (UserOut) for all output aimed at the user
* Remove all imports of the stdlib log
* Remove all usages of fmt.Print* and the like

Consider one logger for user output (what we *expect* them to see) and another for error/warn output.
One could be called userout? The other errorout?
*/

var UserOut = log.New()
var UserOutFormatter = new(UserTextFormatter)


func init() {
	UserOut.Formatter = UserOutFormatter
	UserOutFormatter.DisableTimestamp = true
	UserOut.Level = log.DebugLevel

	log.SetLevel(log.InfoLevel)
}

func main() {

	// UserOut has a custom formatter that doesn't add timestamp etc.
	UserOut.Info("UserOut info")
	UserOut.Print("UserOut print")
	UserOut.Debug("UserOut debug should show")

	// log behaves same as always; level is set in init() to InfoLevel
	log.Debug("log.Debug() - should not show")
	log.Info("log.Info() - should show with timestamp, and with extra data if no terminal")

	jsonF := new(log.JSONFormatter)
	UserOut.Formatter = jsonF
	fields := make(log.Fields)
	fields["tag1"] = "data1"
	type AppRow struct {
		X      int;
		Y      string;
		Name   string;
		Status string;
		Timeup string;
	}
	type vertex struct {
		X, Y, Z int;
		A string
	}
	randommap := make(map[string]string)
	randommap["r1key"] = "r1data"
	randommap["r2key"] = "r2data"
	fields["complexrow1"] = AppRow{0, "1", "myapp", "running", "sometime"}
	fields["row2"] = vertex{1, 2, 3, "4"}
	fields["randommap"] = randommap
	UserOut.WithFields(fields).Print("And this is the message")
}
