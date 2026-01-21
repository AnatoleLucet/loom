package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AnatoleLucet/loom"
	. "github.com/AnatoleLucet/loom-term/components"
	. "github.com/AnatoleLucet/loom/components"

	. "github.com/AnatoleLucet/loom/signals"
)

func MyApp() loom.Node {
	frame, setFrame := Signal(0)
	childs, setChilds := Signal([]int{})

	go func() {
		for {
			time.Sleep(time.Second / 60)
			setFrame(frame() + 1)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second / 60)
			setChilds(append(childs(), frame()))
		}
	}()

	// for range 100 {
	// 	setChilds(append(childs(), frame()))
	// }

	return Box(
		Console(),

		BindText(func() string { return fmt.Sprintf("Length: %d", len(childs())) }),

		Box(
			For(childs, func(child Accessor[int], index Accessor[int]) loom.Node {
				return Box(
					BindText(func() string { return fmt.Sprintf("%d", frame()) }),
					Apply(Style{
						Width:           4,
						JustifyContent:  "center",
						BackgroundColor: RGBA(0, 0, 150, 1),
					}),
				)
			}),

			Apply(Style{
				GapAll:       "2pt",
				FlexWrap:     "wrap",
				AlignContent: "start",
			}),
		),

		Apply(Style{
			Width:  "100%",
			Height: "100%",

			FlexDirection: "column",
		}),
	)
}

func main() {
	app := NewApp()

	for err := range app.Run(RenderFullscreen, MyApp) {
		// for err := range app.Run(RenderInline, MyApp) {
		app.Close()
		log.Fatal(err)
		os.Exit(1)
	}
}
