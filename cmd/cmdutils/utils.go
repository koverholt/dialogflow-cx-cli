package cmdutils

import (
	log "github.com/sirupsen/logrus"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/utils"
)

func CheckUpdate(output bool) {

	if output {
		global.Log.Infof("Current version: %s", global.VersionString)
	}

	latestVersion, _ := utils.CheckAvailableUpdate(global.VersionString, output)

	if latestVersion != "" {
		global.Log.Warnf("A new version is available: %s. Please execute \"brew update && brew upgrade cxcli\"", latestVersion)
	} else {
		if output && latestVersion == "" {
			global.Log.Infof("You have installed the latest version!")
		}
	}
}

func PreRun(command string) {
	if global.VersionString == "" {
		global.VersionString = "development"
	}

	global.Log = *log.New()

	if global.Output == "json" {
		global.Log.Formatter = new(log.JSONFormatter)

	} else {
		global.Log.Formatter = new(log.TextFormatter)
		global.Log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
	}

	if global.Verbose {
		//global.Log.SetReportCaller(true)
		global.Log.SetLevel(log.TraceLevel)
	} else {
		global.Log.SetLevel(log.InfoLevel)
	}

	if !global.SkipUpdate && command != "version" {
		CheckUpdate(false)
	}
}
