package main

import (
	"fmt"
	"math/big"
)

func exec(rom []string, addr string) string {
	normalizedPointer := normalizePointer(addr)
	instr := ReadOne(rom, normalizedPointer)

	fmt.Println("Addr: " + normalizedPointer + " ")

	return execInstr(rom, instr, normalizedPointer, addr)
}

func execInstr(rom []string, instr string, normalizedPointer string, orig string) string {

	p1 := big.NewInt(0)
	p1.SetString(orig, 16)
	p2 := big.NewInt(1)
	p3 := big.NewInt(2)
	addrNext := fmt.Sprintf("%0x", p2.Add(p1, p2))
	addrNextNext := fmt.Sprintf("%0x", p2.Add(p1, p3))

	switch instr {
	case "00": // NOP
		break

	case "c3": // JP
		a1 := ReadOne(rom, addrNext)
		a2 := ReadOne(rom, addrNextNext)
		// we read it in reverse bcs of little endian (I think)
		return a2 + a1
	case "10": // STOP(display)
		//This would likely terminate the display until requsted to return, however:
		// 1. We have no implementation for display yet
		// 2. We can't read sprite data or context that might make this not execute, but rather be sprite data.
		// So for the time being, we'll simply print the fact that it should have happened
		fmt.Println("Display should possibly have been stopped.")

	default:
		fmt.Println("Unknown inst: " + instr)

	}

	return addrNext
}
