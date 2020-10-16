package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bep/debounce"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sahilm/fuzzy"
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

	dialogWindow := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL|gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_QUESTION, 0, "Fuzzy search")
	dialogWindow.SetTitle("Fuzzy search")

	dialogBox, _ := dialogWindow.GetContentArea()
	userEntry, _ := gtk.EntryNew()
	userEntry.SetSizeRequest(250, 0)
	listBox, _ := gtk.ListBoxNew()
	dialogBox.PackEnd(userEntry, false, false, 0)
	dialogBox.PackEnd(listBox, false, false, 0)

	patternChan := make(chan string)
	matchesChan := make(chan []string)
	go findMatches(patternChan, matchesChan)

	debounced := debounce.New(100 * time.Millisecond)
	userEntry.Connect("insert-text", func() {
		debounced(func() {
			s, _ := userEntry.GetText()
			patternChan <- s
		})
	})

	dialogWindow.ShowAll()
	dialogWindow.Run()
	dialogWindow.GetDestroyWithParent()

	for {
		select {
		case matches := <-matchesChan:
			// Cleanup
			listBox.GetChildren().Foreach(func(child interface{}) {
				listBox.Remove(child.(gtk.IWidget))
			})
			for i, r := range matches {
				label, _ := gtk.LabelNew(r)
				listBox.Prepend(label)
				if i == 0 {
					listBox.SelectRow(listBox.GetRowAtIndex(i))
				}
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Reads Stdin and receives the pattern on a channel
// Prints matching results.
func findMatches(patternChan chan string, matchesChan chan []string) {
	reader := bufio.NewReader(os.Stdin)
	lines := []string{}
	results := []string{}

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(input))
	}

	for {
		select {
		case pattern := <-patternChan:
			fmt.Println("new pattern:", pattern)
			matches := fuzzy.Find(pattern, lines)
			for _, r := range matches {
				results = append(results, lines[r.Index])
			}
			fmt.Println("Will send: ", results)
			matchesChan <- results
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
