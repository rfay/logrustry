package main

import (

	log "github.com/sirupsen/logrus"
)

// Overall strategy:
/**
* Add formatters for user interaction, normal logging, and full json interaction
* Use log for all output aimed at the user
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

}
