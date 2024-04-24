package main

import (
	"log"

	"github.com/justinbornais/gocheck/windows"
)

// Get-ItemProperty -Path "C:\Users\justin\Documents\key.txt" | Get-Member -MemberType Property
func main() {
	output, _ := windows.GetTarget("C:\\Users\\justin\\Documents\\key.txt")
	log.Println(output)

	output2, _ := GetFileBitType("C:\\Users\\justin\\Documents\\key.txt")
	log.Println(output2)

	log.Println(GetExtensionNoPeriod("C:\\Users\\justin\\Documents\\key.txt"))
	log.Println(GetFileType("C:\\Users\\justin\\Documents\\key.txt"))
}
