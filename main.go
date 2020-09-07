package main

import (
	"Security-Benchmarking-Tool/constants"
	"Security-Benchmarking-Tool/files"
	"Security-Benchmarking-Tool/store"
	"Security-Benchmarking-Tool/utils"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
)

func init() {
	fmt.Println("This will get called on main initialization")
	HOME, err := utils.GetUserHome()
	if err != nil {
		log.Println(err)
	}
	log.Println(HOME)
	if _, err := os.Stat(HOME + constants.DESKTOP + constants.AuditDirectory); err != nil {
		if os.IsNotExist(err) {
			if err := utils.CreateAuditsDir(HOME); err != nil {
				log.Println("Error creating my directory")
			}
		} else {
			log.Println("File exists!!!")
		}
	}
}

func main() {

	HOME, err := utils.GetUserHome()
	auditPath := HOME + constants.DESKTOP + constants.AuditDirectory
	policyFileName, err := utils.GenerateSavedFileName(auditPath+constants.SavedFileDIRECTORY, constants.AuditFormat, constants.Policy)


	gtk.Init(nil)


	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}


	err = b.AddFromFile("resources/about_dialog_with_buttons.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	obj, err := b.GetObject("about_dialog")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}



	win := obj.( *gtk.AboutDialog)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = b.GetObject("download_button")
	downloadButton := obj.(*gtk.Button)

	obj, _ = b.GetObject("parse_button")
	parseButton := obj.(*gtk.Button)

	obj, _ = b.GetObject("message_or_error")
	label1 := obj.(*gtk.Label)



	downloadButton.Connect("clicked", func() {
		if err != nil {
			log.Println(err)
		}
		if err := store.DownloadFileToExpectedLocation(policyFileName); err != nil {
			log.Println("ERROR IN DOWNLOADING: ", err)
			label1.SetText(fmt.Sprintf("ERROR IN DOWNLOADING: %s",err))
		}else{
			label1.SetText(fmt.Sprintf("Your file was downloaded successfully in: %s",policyFileName))
		}
	})

	parseButton.Connect("clicked", func() {
		arrayData := files.ParseFile(policyFileName)
		info := store.CreateMapForMultipleItems(arrayData)
		jsonFileName, err := utils.GenerateSavedFileName(auditPath+constants.ParsedDataDirectory, constants.ParsedFileFormat, constants.ParsedPolicy)
		if err = store.CreateJsonResponse(jsonFileName, info); err != nil {
			log.Println("ERROR CREATING JSON", err)
			label1.SetText(fmt.Sprintf("ERROR CREATING JSON: %s",err))
		}else{
			label1.SetText(fmt.Sprintf("Your file was parsed successfully in: %s",jsonFileName))
		}
	})

	win.ShowAll()

	gtk.Main()
}