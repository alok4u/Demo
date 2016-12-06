package main

import (
	"encoding/json"
	"flag"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/cli"
	"github.wdf.sap.corp/Ariba-cobalt/COBALT.git/cobalt/manage"
	"github.wdf.sap.corp/Ariba-cobalt/COBALT.git/cobalt/util"
)

// DestroyCommand implements the UI interfaces.
type DestroyCommand struct {
	UI cli.Ui
}

// Run is really the entry point to all the work for cobalt deploy command.
func (c *DestroyCommand) Run(args []string) int {
	log.Info("Destroy Command[command/destroy.go]: Starting run().")
	defer log.Info("Destroy Command[command/destroy.go]:  run() is done.")

	// Instantiate a new Manager object (manager.go), this is require in order to execute it func.
	manager := new(manage.Manager)

	// create new flag object to storage different arguments data
	// A FlagSet represents a set of defined flags.  The zero value of a FlagSet
	// has no name and has ContinueOnError error handling.
	cmdFlags := flag.NewFlagSet("destroy", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }

	// parse cobaltID argument
	cmdFlags.StringVar(&manager.CobaltID, "cobaltid", "", "cobaltid is used to identify a job")

	// parse storage argument
	cmdFlags.StringVar(&manager.ConsulURL, "storage", "", "storage is used to connect to storage system")

	// parse --vv argument
	var robust bool
	cmdFlags.BoolVar(&robust, "vv", false, "if vv turn on debug mode")

	// this command parses the value from commandline
	// if there is an error, break
	if err := cmdFlags.Parse(args); err != nil {
		log.Errorf("Destroy Command[command/destroy.go]: Parse rrror found:%v\n", err)
		c.UI.Error("")
		c.UI.Error(c.Help())
		return 1
	}
if err := cmdFlags.Parse(arguments); err != nil {
	log.Errorf ("Destroy Command[command/dsetory.go]: Parse rrror found:%v\n", err)
	
}
	if robust {
		config := &util.LogConfig{
			Level:  "DEBUG",
			Format: "text",
			Output: "stdout",
		}
		config.Init()
	} else {
		config := &util.LogConfig{
			Level:  "INFO",
			Format: "text",
			Output: "stdout",
		}
		config.Init()
	}

	/* throw error if cobaltid is empty */
	if manager.CobaltID == "" {
		log.Error("Destroy Command[command/destroy.go]: Please check arg cobaltid")
		c.UI.Error("Please check arg cobaltid")
		c.UI.Error(c.Help())
		return 1
	}

	res, e := manager.DestroyManager()
	resByte, _ := json.MarshalIndent(res, "", "	")
	fmt.Println(string(resByte))
	if e != nil {
		fmt.Println(e)
		return 1
	}

	return 0
}

// Help functions print out the helo command
func (c *DestroyCommand) Help() string {
	helpText := `
Usage: cobalt destroy [options] <cobaltid, storage>
This command is used to destroy a service.
Options:
	--cobaltid[required]         a unique id to represent unique instantiation of a service
	--storage[required]          storage system to store services info (consul as default)

Example
========
1. No environment specify
cobalt destroy --cobaltid="1234eddff677837" --storage=consul:10.9.43.159:8500
`
	return helpText
}

// Synopsis display the status of exec
func (c *DestroyCommand) Synopsis() string {
	return "Destroy a micro service from a cluster."
}
