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

	patternChan := make(chan string)
	go findMatches(patternChan)

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
	dialogBox.PackEnd(userEntry, false, false, 0)

	debounced := debounce.New(100 * time.Millisecond)
	userEntry.Connect("insert-text", func() {
		debounced(func() {
			s, _ := userEntry.GetText()
			patternChan <- s
		})
	})

	dialogWindow.ShowAll()
	dialogWindow.Run()
	text, _ := userEntry.GetText()
	dialogWindow.GetDestroyWithParent()
	fmt.Println(text)
}

// Reads Stdin and receives the pattern on a channel
// Prints matching results.
func findMatches(patternChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	lines := []string{}

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
			results := fuzzy.Find(pattern, lines)
			fmt.Println(len(results))
			// for _, r := range results {
			// 	fmt.Println(lines[r.Index])
			// }
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
