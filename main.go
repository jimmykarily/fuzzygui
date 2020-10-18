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

const (
	PatternEntryID   = "pattern_entry"
	MatchesListboxID = "matches_list_box"
	MainWindowID     = "main_window"
	MaxResults       = 100
)

var (
	lines        []string
	patternEntry *gtk.Entry
	listBox      *gtk.ListBox
	mainWindow   *gtk.Window
	desiredRow   int
	matches      []Match
)

type Match struct {
	Str   string
	Label *gtk.Label
}

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

	const appID = "gr.brainbytes.fuzzygui"
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatalln("Couldn't create app:", err)
	}
	app.Connect("activate", func() {
		initilizeWidgets()

		mainWindow.ShowAll()
		app.AddWindow(mainWindow)

		debounced := debounce.New(100 * time.Millisecond)
		patternEntry.Connect("changed", func() {
			debounced(func() {
				pattern, _ := patternEntry.GetText()
				matches = findMatches(pattern, &lines, MaxResults)
				// Cleanup
				_, err = glib.IdleAdd(CleanList, listBox)

				_, err = glib.IdleAdd(func() {
					for i, match := range matches {
						listBox.Insert(match.Label, i)
					}
				})
				_, err = glib.IdleAdd(SelectClosestRow)
				_, err = glib.IdleAdd(listBox.ShowAll)
			})
		})

		mainWindow.Connect("key-press-event", func(entry *gtk.Window, event *gdk.Event) bool {
			keyval := gdk.EventKeyNewFromEvent(event).KeyVal()
			for _, key := range []string{"Up", "Down", "Return"} {
				if keyval == gdk.KeyvalFromName(key) {
					_, err = glib.IdleAdd(HandleKey, key)
					return true
				}
			}
			return false
		})

		patternEntry.Emit("changed") // Emit once to start with all results

		go readLines()
	})

	app.Run(os.Args)
}

func initilizeWidgets() {
	builder, err := gtk.BuilderNew()
	if err != nil {
		log.Fatalln("Couldn't make builder:", err)
	}

	guiGlade, err := Asset("gui.glade")
	if err != nil {
		fmt.Println("gui.glade asset was not found. Run go-bindata and re-compile")
		os.Exit(1)
	}

	err = builder.AddFromString(string(guiGlade))
	if err != nil {
		log.Fatalln("Couldn't add UI XML to builder:", err)
	}

	obj, _ := builder.GetObject(MainWindowID)
	mainWindow = obj.(*gtk.Window)
	obj, _ = builder.GetObject(PatternEntryID)
	patternEntry = obj.(*gtk.Entry)
	obj, _ = builder.GetObject(MatchesListboxID)
	listBox = obj.(*gtk.ListBox)
}

func HandleKey(key string) {
	switch key {
	case "Return":
		PrintSelectionAndExit()
	case "Up":
		desiredRow -= 1
		SelectClosestRow()
		patternEntry.GrabFocus()
	case "Down":
		desiredRow += 1
		SelectClosestRow()
		patternEntry.GrabFocus()
	}
}

// Select the row above, the row below or the current line respecting
// the current limits. Won't move below bottom row, above top row and
// will adjust to closer row if the step is zero but there is no such row
// (because we may have less results available now)
func SelectClosestRow() {
	totalMatches := len(matches)
	if totalMatches == 0 {
		desiredRow = -1
		return
	}

	if desiredRow > totalMatches-1 {
		desiredRow = totalMatches - 1
	} else if desiredRow < 0 {
		desiredRow = 0
	}

	listBox.SelectRow(listBox.GetRowAtIndex(desiredRow))
}

func CleanList(listBox *gtk.ListBox) {
	listBox.GetChildren().Foreach(func(child interface{}) {
		listBox.Remove(child.(gtk.IWidget))
	})
}

func PrintSelectionAndExit() {
	if desiredRow >= 0 {
		fmt.Print(matches[desiredRow].Str)
	}
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
// and retuns a sclice of the first maxResults Matches.
// If pattern is empty, it returns unfiltered results.
func findMatches(pattern string, lines *[]string, maxResults int) []Match {
	results := []Match{}
	if pattern == "" {
		for i, match := range *lines {
			label, _ := gtk.LabelNew(match)
			label.SetXAlign(0)
			results = append(results, Match{Str: match, Label: label})
			if i >= maxResults {
				break
			}
		}
	} else {
		for i, match := range fuzzy.Find(pattern, *lines) {
			label, _ := gtk.LabelNew(match.Str)
			label.SetXAlign(0)
			results = append(results, Match{Str: match.Str, Label: label})
			if i >= maxResults {
				break
			}
		}
	}

	return results
}
