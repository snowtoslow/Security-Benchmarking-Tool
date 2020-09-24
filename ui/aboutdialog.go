package ui

import (
	"Security-Benchmarking-Tool/constants"
	"Security-Benchmarking-Tool/store"
	"Security-Benchmarking-Tool/utils"
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

var (
	policyFileNameNew string
	label1            *gtk.Label
)

func SetupAboutDialogWindow() (err error) {

	gtk.Init(nil)

	builder, err := createBuilder("resources/about_dialog_with_buttons.glade")
	if err != nil {
		log.Println("err builder: ", err)
		return err
	}

	object, err := createObject(builder, "about_dialog")
	if err != nil {
		log.Println("error about dialog: ", err)
		return err
	}

	window := object.(*gtk.AboutDialog)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	object, err = createObject(builder, "download_button")
	if err != nil {
		log.Println("error getting download: ", err)
		return err
	}

	if _, err = createButtons(object, setupDownloadButton); err != nil {
		log.Println("HERE:", err)
		return err
	}

	object, err = createObject(builder, "parse_button")
	if err != nil {
		log.Println("error getting download: ", err)
		return err
	}

	if _, err = createButtons(object, setupParseButton); err != nil {
		log.Println("error parsing file:", err)
		return err
	}

	object, err = createObject(builder, "custom_policy_button")
	if err != nil {
		log.Println("error opening tree view: ", err)
		return err
	}

	if _, err = createButtons(object, openSelection); err != nil {
		log.Println("HERE1:", err)
		return err
	}

	object, _ = builder.GetObject("message_or_error")
	label1 = object.(*gtk.Label)

	window.ShowAll()

	gtk.Main()

	return nil
}

// function to create a builder;
func createBuilder(pathToGladeFile string) (myBuilder *gtk.Builder, err error) {
	myBuilder, err = gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	if err = myBuilder.AddFromFile(pathToGladeFile); err != nil {
		return nil, err
	}

	return
}

// create objects
func createObject(myBuilder *gtk.Builder, objId string) (myObj glib.IObject, err error) {
	myObj, err = myBuilder.GetObject(objId)
	if err != nil {
		return nil, err
	}
	return
}

func createButtons(myObj glib.IObject, myFunc func()) (button *gtk.Button, err error) {
	button = myObj.(*gtk.Button)
	button.Connect("clicked", myFunc)

	return button, nil
}

func setupDownloadButton() {
	HOME, err := utils.GetUserHome()
	policyFileName, err := utils.GenerateSavedFileName(HOME+constants.DESKTOP+constants.AuditDirectory+constants.SavedFileDIRECTORY, constants.AuditFormat, constants.Policy)
	policyFileNameNew = policyFileName
	if err = store.DownloadFileToExpectedLocation(policyFileName); err != nil {
		log.Println("ERROR IN DOWNLOADING: ", err)
		label1.SetText(fmt.Sprintf("ERROR IN DOWNLOADING: %s", err))
	} else {
		label1.SetText(fmt.Sprintf("Your file was downloaded successfully in: %s", policyFileName))
	}
}

func setupParseButton() {
	arrayData := utils.ParseFile(policyFileNameNew)
	HOME, err := utils.GetUserHome()
	auditPath := HOME + constants.DESKTOP + constants.AuditDirectory
	info := store.CreateMapForMultipleItems(arrayData)
	jsonFileName, err := utils.GenerateSavedFileName(auditPath+constants.ParsedDataDirectory, constants.ParsedFileFormat, constants.ParsedPolicy)
	if err = store.CreateJsonResponse(jsonFileName, info); err != nil {
		log.Println("ERROR CREATING JSON", err)
		label1.SetText(fmt.Sprintf("ERROR CREATING JSON: %s", err))
	} else {
		label1.SetText(fmt.Sprintf("Your file was parsed successfully in: %s", jsonFileName))
	}
}

// open tree view with selection;
func openSelection() {
	DisplayOptionsToCreateCustomShit()
}
