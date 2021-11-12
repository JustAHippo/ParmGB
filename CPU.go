package main

import (
	"fmt"
	"math/big"
)

func exec(rom []string, addr string) string {
	normalizedPointer := normalizePointer(addr)
	instr := ReadOne(rom, normalizedPointer)

	fmt.Println("; Addr: " + normalizedPointer + " ")

	return execInstr(rom, instr, normalizedPointer, addr)
}

//goland:noinspection GoUnusedParameter
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
		fmt.Println("NOP ;00")
		break

	case "06": // TODO: Implement
		a1 := ReadOne(rom, addrNext)

		n1 := big.NewInt(0)
		n1.SetString(a1, 16)

		fmt.Println("ld b," + a1 + " ;" + a1)

		return addrNextNext

	case "0e": // TODO: Implement
		a1 := ReadOne(rom, addrNext)

		n1 := big.NewInt(0)
		n1.SetString(a1, 16)

		fmt.Println("ld c," + a1 + " ;0e " + a1)

		return addrNextNext

	case "10": // STOP(Whatever is listed in IE Register)
		//This would likely terminate whatever is specified in the IE Register
		// 1. We have no implementation yet
		// 2. We can't read sprite data or context that might make this not execute, but rather be sprite data or other usage.
		// So for the time being, we'll simply print the fact that it should have happened
		fmt.Println("STOP ;10 Something should possibly have been stopped.")
		break

	case "21": // TODO: Implement
		a1 := ReadOne(rom, addrNext)
		a2 := ReadOne(rom, addrNextNext)

		n1 := big.NewInt(0)
		n2 := big.NewInt(0)
		n1.SetString(a1, 16)
		n2.SetString(a2, 16)

		fmt.Println("ld " + a1 + ", " + a2 + " ;21 " + a2 + " " + a1)

		return addrNextNextNext

	case "3e": // TODO: Implement
		a1 := ReadOne(rom, addrNext)

		n1 := big.NewInt(0)
		n1.SetString(a1, 16)

		fmt.Println("ld a, " + a1 + " ;" + a1)

		return addrNextNext

	case "20": // TODO: Implement
		a1 := ReadOne(rom, addrNext)

		n1 := big.NewInt(0)
		n1.SetString(a1, 16)

		fmt.Println("jr nz, " + a1 + " ;20 " + a1)

		return addrNextNext

	case "af": // TODO: Implement

		fmt.Println(";xor a")
		return addrNext

	case "c3":
		// we read it in reverse bcs of little endian (I think)
		a2 := ReadOne(rom, addrNext)
		a1 := ReadOne(rom, addrNextNext)
		fmt.Println("jp " + a1 + a2 + " ;C3" + " " + a2 + " " + a1)
		return a1 + a2

	default:
		fmt.Println(";Unknown inst: " + instr)

	}

	return addrNext
}
