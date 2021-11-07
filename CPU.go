package main

import (
	"fmt"
	"strconv"
)

func exec(rom []string, addr string) string {
	normalizedPointer := normalizePointer(addr)
	instr := ReadOne(rom, normalizedPointer)

	fmt.Println("Addr: " + normalizedPointer)

	return execInstr(rom, instr, normalizedPointer)
}

func execInstr(rom []string, instr string, normalizedPointer string) string {
	i, err2 := strconv.ParseInt(normalizedPointer, 0, 32)

	check(err2)
	addrInt := int(i)
	addrInt++

	fmt.Println("INST: " + instr)

	// TODO: How the fuck do I make this return the hexadecimal version of addrInt??????
	return string("-1")
}
