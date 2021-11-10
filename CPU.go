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
	p4 := big.NewInt(3)
	addrNext := fmt.Sprintf("%0x", p2.Add(p1, p2))
	addrNextNext := fmt.Sprintf("%0x", p2.Add(p1, p3))
	addrNextNextNext := fmt.Sprintf("%0x", p2.Add(p1, p4))

	switch instr {
	case "00": // NOP
		break

	case "af": // XOR ** **
		a2 := ReadOne(rom, addrNext)
		a1 := ReadOne(rom, addrNextNext)

		n1 := big.NewInt(0)
		n2 := big.NewInt(0)
		n1.SetString(a1, 16)
		n2.SetString(a2, 16)

		res := n1.Int64() ^ n2.Int64()

		fmt.Println("Unimplemented XOR: " + n1.String() + " " + n2.String() + " -> " + fmt.Sprintf("%d", res))
		return addrNextNextNext

	case "c3": // JP ** **
		// we read it in reverse bcs of little endian (I think)
		a2 := ReadOne(rom, addrNext)
		a1 := ReadOne(rom, addrNextNext)
		fmt.Println("Jumping to: $" + a1 + a2)
		return a1 + a2
	case "10": // STOP(Whatever is listed in IE Register)
		//This would likely terminate whatever is specified in the IE Register
		// 1. We have no implementation yet
		// 2. We can't read sprite data or context that might make this not execute, but rather be sprite data or other usage.
		// So for the time being, we'll simply print the fact that it should have happened
		fmt.Println("Something should possibly have been stopped.")
		break

	default:
		fmt.Println("Unknown inst: " + instr)

	}

	return addrNext
}
