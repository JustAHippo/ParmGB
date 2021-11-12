package main

import (
	"encoding/hex"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("rom.gb")
	check(err)

	pointer := "0100"
	hexArr := Chunks(hex.EncodeToString(dat), 2)
	TITLEBYTES := Chunks(Read(hexArr, "134", "143"), 2)
	TITLE := ""
	for i := 0; i < len(TITLEBYTES); i++ {
		curr := TITLEBYTES[i]

		if curr != "00" {
			cb, _err := hex.DecodeString(curr)
			check(_err)
			TITLE += string(cb)
		}
	}
	TITLE = TITLE[1:] // Just don't question it
	fmt.Println(TITLE)

	scale := int32(4)

	rl.InitWindow(160*scale, 144*scale, TITLE)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("PC: 0"+string(pointer)+"\nFPS: "+fmt.Sprintf("%d", int(rl.GetFPS())), 4, 4, 18, rl.Black)
		pointer = exec(hexArr, pointer)

		// rl.DrawFPS(4, 4)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
