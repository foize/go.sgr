package main

import (
	"fmt"
	"go.tcolor"
)

func main() {
	fmt.Println("normal text")

	r, err := tcolor.Format("Dit word rood: ", tcolor.FgRed, "%s\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	yc, err := tcolor.Format("Dit word geel op cyan: ", tcolor.FgYellow, tcolor.BgCyan, "%s\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Because of the reset that's being placed at the end of the formatstring, we need the text to be Printf'd
	fmt.Printf(r, "rood rood rood")
	fmt.Printf(yc, "geel op cyan")
}
