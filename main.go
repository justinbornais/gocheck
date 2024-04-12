package main

import (
	"log"

	"github.com/justinbornais/gocheck/windows"
)

// Get-ItemProperty -Path "C:\Users\justin\Documents\key.txt" | Get-Member -MemberType Property
func main() {
	output, err := windows.GetTarget("C:\\Users\\justin\\Documents\\key.txt")
	if err != nil {
		// log.Fatalln(err)
	}
	log.Println(output)
}
