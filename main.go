package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bep/debounce"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sahilm/fuzzy"
)

var (
	lines        []string
	currentMatch string
)

// https://coderwall.com/p/zyxyeg/golang-having-fun-with-os-stdin-and-shell-pipes
func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: find * | fuzzygui")
		return
	}

	// Initialize GTK without parsing any command line arguments.
	//gtk.Init(nil)

	guiGlade, err := Asset("gui.glade")
	if err != nil {
		fmt.Println("gui.glade asset was not found. Run go-bindata and re-compile")
		os.Exit(1)
	}

	const appID = "com.retc3.mytest"
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatalln("Couldn't create app:", err)
	}
	app.Connect("activate", func() {
		builder, err := gtk.BuilderNew()
		if err != nil {
			log.Fatalln("Couldn't make builder:", err)
		}

		err = builder.AddFromString(string(guiGlade))
		if err != nil {
			log.Fatalln("Couldn't add UI XML to builder:", err)
		}

		obj, _ := builder.GetObject("main_window")
		wnd := obj.(*gtk.Window)
		wnd.ShowAll()
		app.AddWindow(wnd)

		obj, _ = builder.GetObject("pattern_entry")
		patternEntry := obj.(*gtk.Entry)
		obj, _ = builder.GetObject("matches_list_box")
		listBox := obj.(*gtk.ListBox)

		debounced := debounce.New(100 * time.Millisecond)
		patternEntry.Connect("changed", func() {
			debounced(func() {
				pattern, _ := patternEntry.GetText()
				matches := findMatches(pattern, &lines)
				// Cleanup
				_, err = glib.IdleAdd(CleanList, listBox)

				numOfResults := 0
				matchesLen := len(matches)
				if matchesLen > 10 {
					numOfResults = 10
				} else {
					numOfResults = matchesLen
				}
				for i, r := range matches[:numOfResults] {
					label, _ := gtk.LabelNew(r)
					label.SetXAlign(0)
					_, err = glib.IdleAdd(func() {
						listBox.Insert(label, i)
					})
				}

				match := ""
				if matchesLen > 0 {
					match = matches[0]
					_, err = glib.IdleAdd(func() {
						SelectFirstRow(listBox)
					})
				}
				currentMatch = match

				_, err = glib.IdleAdd(listBox.ShowAll)
			})
		})

		patternEntry.Connect("key-press-event", func(entry *gtk.Entry, event *gdk.Event) {
			if gdk.KeyvalFromName("Return") == gdk.EventKeyNewFromEvent(event).KeyVal() {
				_, err = glib.IdleAdd(PrintSelectionAndExit)
			}
		})

		go readLines()
	})

	app.Run(os.Args)
}

func CleanList(listBox *gtk.ListBox) {
	listBox.GetChildren().Foreach(func(child interface{}) {
		listBox.Remove(child.(gtk.IWidget))
	})
}

func SelectFirstRow(listBox *gtk.ListBox) {
	listBox.SelectRow(listBox.GetRowAtIndex(0))
}

func PrintSelectionAndExit() {
	fmt.Print(currentMatch)
	os.Exit(0)
}

func readLines() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		lines = append(lines, string(input))
		time.Sleep(time.Duration(1) * time.Microsecond)
	}
}

// Looks for fuzzy matches in lines using the pattern
func findMatches(pattern string, lines *[]string) []string {
	matches := fuzzy.Find(pattern, *lines)
	results := []string{}
	for _, r := range matches {
		results = append(results, (*lines)[r.Index])
	}

	return results
}
