package ui

import (
	"Security-Benchmarking-Tool/constants"
	"Security-Benchmarking-Tool/store"
	"Security-Benchmarking-Tool/utils"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

var (
	myStore *gtk.ListStore
	paths   []string
	RootBox *gtk.Box
)

func DisplayOptionsToCreateCustomShit() {
	gtk.Init(nil)

	win := setupWindow("Security Benchmarking Tool")
	arrayData := utils.ParseFile("/home/snowtoslow/Desktop/audit/new-audits/policy092320200.audit") // change here the path in future;
	info := store.CreateMapForMultipleItems(arrayData)
	//
	treeView, myStore, positionWithKeys := setupTreeView(getMapsWithMaxNumberOfKey(info))
	width, height := 600, 300

	for i := 0; i < len(info); i++ {
		addRow(myStore, createInterface(test(info[i], positionWithKeys)))
	}

	vAdj, err := gtk.AdjustmentNew(0, 0, float64(width), 1, 10, float64(height))
	if err != nil {
		log.Println("vadjerr:", err)
	}

	hAdj, err := gtk.AdjustmentNew(0, 0, float64(width), 1, 10, float64(height))
	if err != nil {
		log.Println("hadj:", err)
	}

	scrolledWindow, err := gtk.ScrolledWindowNew(hAdj, vAdj)
	if err != nil {
		log.Println("scrolled window error:", err)
	}

	createCustomPolicyButton, err := gtk.ButtonNew()
	if err != nil {
		log.Println("custom create policy button error: ", err)
	}

	createCustomPolicyButton.SetLabel("Save")

	HOME, err := utils.GetUserHome()
	auditPath := HOME + constants.DESKTOP + constants.AuditDirectory

	intArray, err := utils.ConvertArrayToInt(paths)
	if err != nil {
		log.Println("convert to array of ints error: ", err)
	}

	customAuditsMap := utils.CreateMapOfAuditsFromIndexArray(intArray, info)

	createCustomPolicyButton.Connect("clicked", func() {

		customPolicyFileName, err := utils.GenerateSavedFileName(auditPath+constants.CustomAuditDirectory, constants.AuditFormat, constants.CustomAudit)
		if err != nil {
			log.Println("Custom policy file name generator: ", err)
		}
		if err := store.CreateCustomPolicy(customPolicyFileName, customAuditsMap); err != nil {
			log.Println("KAKOITO NOVYI EBANYI ERROR: ", err)
		}

	})

	RootBox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)

	scrolledWindow.Add(treeView)
	RootBox.PackStart(scrolledWindow, true, true, 0)
	RootBox.PackStart(createCustomPolicyButton, false, false, 0)
	scrolledWindow.SetHExpand(true)
	scrolledWindow.SetVExpand(true)

	win.SetPosition(gtk.WIN_POS_CENTER)

	win.SetDefaultSize(width, height)

	selection, err := treeView.GetSelection()
	if err != nil {
		log.Fatal("Could not get tree selection object.")
	}

	selection.SetMode(gtk.SELECTION_MULTIPLE)
	selection.Connect("changed", SelectionChanged)

	win.Add(RootBox)
	win.ShowAll()
	gtk.Main()
}

// working with single selection
func SelectionChanged(s *gtk.TreeSelection) {
	// Returns glib.List of gtk.TreePath pointers
	rows := s.GetSelectedRows(myStore)
	paths = make([]string, 0, rows.Length())

	for l := rows; l != nil; l = l.Next() {
		path := l.Data().(*gtk.TreePath)
		paths = append(paths, path.String())
	}
	log.Println(paths)
}

// Add a column to the tree view (during the initialization of the tree view)
func createColumn(title string, id int) *gtk.TreeViewColumn {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal("Unable to create text cell renderer:", err)
	}

	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "text", id)
	if err != nil {
		log.Fatal("Unable to create cell column:", err)
	}

	return column
}

// Creates a tree view and the list store that holds its data
func setupTreeView(maxSizeMap map[string]string) (*gtk.TreeView, *gtk.ListStore, map[int]string) {
	positionWithKeys := make(map[int]string)
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal("Unable to create tree view:", err)
	}
	var counter int

	for k := range maxSizeMap {
		counter++
		treeView.AppendColumn(createColumn(k, counter))
		positionWithKeys[counter] = k
	}

	// Creating a list store. This is what holds the data that will be shown on our tree view.
	listStore, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING,
		glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("Unable to create list store:", err)
	}
	myStore = listStore
	treeView.SetModel(listStore)

	return treeView, listStore, positionWithKeys
}

// add just one row
func addRow(listStore *gtk.ListStore, myInterface []interface{}) {
	// Get an iterator for a new row at the end of the list store
	iter := listStore.Append()

	err := listStore.Set(iter,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		myInterface)

	if err != nil {
		log.Fatal("Unable to add row:", err)
	}
}

// Create and initialize the window
func setupWindow(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	return win
}

func test(map1 map[string]string, mapWithInts map[int]string) (arrayWithLengthEleven []string) {
	arrayWithLengthEleven = make([]string, len(mapWithInts))
	for k, v := range mapWithInts {
		for key, value := range map1 {
			if v == key {
				arrayWithLengthEleven[k-1] = value
			}
		}
	}
	return
}

func createInterface(stringArr []string) (myInterface []interface{}) {
	myInterface = make([]interface{}, len(stringArr))
	for i, s := range stringArr {
		myInterface[i] = s
	}

	return
}

func getMapsWithMaxNumberOfKey(myMap []map[string]string) (maxMap map[string]string) {
	maxMap = myMap[0]
	for _, v := range myMap {
		if len(maxMap) < len(v) {
			maxMap = v
		}
	}
	return
}
