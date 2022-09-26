package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appId = "com.nayoso.example"
	app, _ := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	app.Connect("activate", func() {
		onActivate(app)
	})
	app.Run(os.Args)
}

// --- At the beginning we still use the familiar code
func onActivate(application *gtk.Application) {
	if builder, err := gtk.BuilderNewFromFile("builder.ui"); err != nil { //Create Builder from file
		log.Fatal(err)
	} else if winObj, err := builder.GetObject("window"); err != nil { //Read the window object from the file, which is actually a Gobject
		log.Fatal(err)
	} else {
		window := winObj.(*gtk.Window) //Since winObj is a Gobject, we use type assertion to get the Gtk.Window object
		application.AddWindow(window)  //Remember to add window to our application
		window.ShowAll()
	}
}
