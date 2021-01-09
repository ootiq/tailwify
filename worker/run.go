package worker

import (
	"log"
	"os/exec"

	"github.com/leaanthony/spinner"
)

// check for the main installer
func (w *Worker) checkNodePackager() (string, []string) {
	// check command
	var nodePkger string
	var nodePkgerCommand []string

	// just loop in the package managers
	for _, p := range w.appConfig.installer {
		if p.pkgManager == w.jsPkger {
			nodePkger = p.pkgInstaller
			nodePkgerCommand = p.pkgInstArgs

			break
		}
	}

	// if the package manager is not `npm` or `yarn`
	// this will error if the code is modified
	if nodePkger == "" {
		log.Fatal("Error NODE Package Manager: " + w.jsPkger)
	}

	// append the project name to the args
	newArgs := append(nodePkgerCommand, w.ProjectName)

	return nodePkger, newArgs
}

// MAIN HANDLER FOR EVERYTHING
func (w *Worker) run() {
	// start install
	w.installSpinner = spinner.New("Installing " + w.appConfig.name)
	w.installSpinner.Start()

	// run the app installer //
	cmdCommand, cmdArg := w.checkNodePackager()
	cmd := exec.Command(cmdCommand, cmdArg...)
	if err := cmd.Run(); err != nil {
		w.installSpinner.Error("There was a problem while trying to install " + w.appConfig.name)
		log.Fatal(err)
	}

	// after create-
	// if needed
	if w.appConfig.afterCreateInstall {
		w.afterInstall()
	}

	// show success message
	w.installSpinner.Success("Succesfully installed " + w.appConfig.name)

	// install tailwind
	w.installTailwind()

	// configure and modify files
	w.fileModifier()
}

// after install function
// after the create- something, if there is
// some of the frameworks, do not automatically install it
// like, .. having `npm install`  after
func (w *Worker) afterInstall() {
	cmd := exec.Command(w.jsPkger, "install")
	cmd.Dir = w.projectDir

	// install it
	if err := cmd.Run(); err != nil {
		w.installSpinner.Error("There was a problem while trying to install " + w.appConfig.name)
		log.Fatal(err)
	}
}
