package main

import (
	"fmt"
	"math/big"
)

func exec(rom []string, addr string) string {
	normalizedPointer := normalizePointer(addr)
	instr := ReadOne(rom, normalizedPointer)

	fmt.Print("Addr: " + normalizedPointer + " ")

	return execInstr(rom, instr, normalizedPointer, addr)
}

func execInstr(rom []string, instr string, normalizedPointer string, orig string) string {

	fmt.Println("INST: " + instr)

	p1 := big.NewInt(0)
	p1.SetString(orig, 16)
	p2 := big.NewInt(1)

	return fmt.Sprintf("%0x", p2.Add(p1, p2))
}
