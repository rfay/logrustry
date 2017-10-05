package main

import (
	"os"

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

var logger = log.New()

func init() {

	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	//x := logger.(*log.MyTextFormatter).DisableTimestamp
	//x := log.MyTextFormatter(logger)
	//x := logger.(*log.MyTextFormatter)
}

func main() {
	log.Warn("Try this")
	log.Print("This is just a print")
	customFormatter := new(MyTextFormatter)
	customFormatter.DisableTimestamp = true
	customFormatter.isTerminal = false

	log.Info("Hello Walrus before DisableTimestamp=true")
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
	log.Info("Hello Walrus after FullTimestamp=true")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logger.WithFields(log.Fields{
		"whichlogger": "logger",
		"size":        10,
	}).Info("using logger")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Warnln("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Warn("I'll be logged with common and other field")
	contextLogger.WithField("x", 1).Warn("Me too")
}
