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
	listBox      *gtk.ListBox
	dialogWindow *gtk.MessageDialog
	currentMatch string
)

type employee struct {
	name string
	age  int
}

type employees []employee

func (e employees) String(i int) string {
	return e[i].name
}

func (e employees) Len() int {
	return len(e)
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

	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_POPUP)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	dialogWindow = gtk.MessageDialogNew(win, gtk.DIALOG_MODAL|gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_QUESTION, 0, "Fuzzy search")
	dialogWindow.SetTitle("Fuzzy search")

	dialogBox, _ := dialogWindow.GetContentArea()
	userEntry, _ := gtk.EntryNew()
	userEntry.SetSizeRequest(250, 0)
	listBox, _ = gtk.ListBoxNew()
	dialogBox.PackEnd(userEntry, false, false, 0)
	dialogBox.PackEnd(listBox, false, false, 0)

	debounced := debounce.New(100 * time.Millisecond)
	userEntry.Connect("changed", func() {
		debounced(func() {
			pattern, _ := userEntry.GetText()
			matches := findMatches(pattern, &lines)
			// Cleanup
			_, err = glib.IdleAdd(CleanList)

			numOfResults := 0
			matchesLen := len(matches)
			if matchesLen > 10 {
				numOfResults = 10
			} else {
				numOfResults = matchesLen
			}
			for _, r := range matches[:numOfResults] {
				label, _ := gtk.LabelNew(r)
				_, err = glib.IdleAdd(listBox.Prepend, label)
				_, err = glib.IdleAdd(listBox.ShowAll)
			}
			match := ""
			if matchesLen > 0 {
				match = matches[0]
			}
			_, _ = glib.IdleAdd(SelectFirstRow, match)
		})
	})

	userEntry.Connect("key-press-event", func(entry *gtk.Entry, event *gdk.Event) {
		if gdk.KeyvalFromName("Return") == gdk.EventKeyNewFromEvent(event).KeyVal() {
			_, err = glib.IdleAdd(PrintSelectionAndExit)
		}
	})

	go readLines()

	dialogWindow.ShowAll()
	dialogWindow.Run()
	dialogWindow.GetDestroyWithParent()
}

func CleanList() {
	listBox.GetChildren().Foreach(func(child interface{}) {
		listBox.Remove(child.(gtk.IWidget))
	})
}

func PrintSelectionAndExit() {
	fmt.Print(currentMatch)
	os.Exit(0)
}

func SelectFirstRow(match string) {
	r := listBox.GetRowAtIndex(int(listBox.GetChildren().Length()) - 1)
	listBox.SelectRow(r)
	currentMatch = match
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
