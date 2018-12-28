package gosweeper

import (
	"log"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

// Launcher creates a game board and launches the game
type Launcher struct {
	board *Board
	x, y  int
}

// NewLauncher creates game launcher, that able to create GTK game board
func NewLauncher(xLen, yLen, mines int) (*Launcher, error) {
	board, err := NewBoard(xLen, yLen, 10, 10, mines)
	if err != nil {
		return nil, err
	}
	return &Launcher{
		board: board,
		x:     xLen,
		y:     yLen,
	}, nil
}

// Start creates and displays a game goard in GTK
func (l *Launcher) Start() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("GoSweeper")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	for x := 0; x < l.x; x++ {
		for y := 0; y < l.y; y++ {
			var label string
			if x == l.x && y == l.y {
				label = "X"
			} else {
				label = ""
			}

			btn, err := gtk.ButtonNewWithLabel(label)
			// btn.SetSizeRequest(1, 1)
			btn.Connect("clicked", func(x, y int) func() {
				return func() {
					status, _ := l.board.GetStatus(x, y)
					switch status {
					case StatusMine:
						btn.SetLabel("X")

						md := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Boom!")
						md.Connect("response", func() { gtk.MainQuit() })
						md.ShowNow()

					default:
						btn.SetLabel(strconv.Itoa(status))
					}

				}
			}(x, y))
			if err != nil {
				log.Fatal("Unable to create button:", err)
			}
			grid.Attach(btn, x, y, 1, 1)
		}
	}

	win.Add(grid)
	win.ShowAll()

	gtk.Main()
}
