package components

// Icon values
const (
	MBOk                = 0x00000000
	MBOkCancel          = 0x00000001
	MBAbortRetryIgnore  = 0x00000002
	MBYesNoCancel       = 0x00000003
	MBYesNo             = 0x00000004
	MBRetryCancel       = 0x00000005
	MBCancelTryContinue = 0x00000006
	MBIconHand          = 0x00000010
	MBIconQuestion      = 0x00000020
	MBIconExclamation   = 0x00000030
	MBIconAsterisk      = 0x00000040
	MBUserIcon          = 0x00000080
	MBIconWarning       = MBIconExclamation
	MBIconError         = MBIconHand
	MBIconInformation   = MBIconAsterisk
	MBIconStop          = MBIconHand

	MBDefaultButton1 = 0x00000000
	MBDefaultButton2 = 0x00000100
	MBDefaultButton3 = 0x00000200
	MBDefaultButton4 = 0x00000300

	ErrorAlreadyExists = 183

	MEMCommit  = 0x1000
	MEMReserve = 0x2000

	PageExecuteReadWrite = 0x40
)

const (
	// Virtual-Key Codes
	vkBACK      = 0x08
	vkTAB       = 0x09
	vkCLEAR     = 0x0C
	vkRETURN    = 0x0D
	vkSHIFT     = 0x10
	vkCONTROL   = 0x11
	vkMENU      = 0x12
	vkPAUSE     = 0x13
	vkCAPITAL   = 0x14
	vkESCAPE    = 0x1B
	vkSPACE     = 0x20
	vkPRIOR     = 0x21
	vkNEXT      = 0x22
	vkEND       = 0x23
	vkHOME      = 0x24
	vkLEFT      = 0x25
	vkUP        = 0x26
	vkRIGHT     = 0x27
	vkDOWN      = 0x28
	vkSELECT    = 0x29
	vkPRINT     = 0x2A
	vkEXECUTE   = 0x2B
	vkSNAPSHOT  = 0x2C
	vkINSERT    = 0x2D
	vkDELETE    = 0x2E
	vkLWIN      = 0x5B
	vkRWIN      = 0x5C
	vkAPPS      = 0x5D
	vkSLEEP     = 0x5F
	vkNUMPAD0   = 0x60
	vkNUMPAD1   = 0x61
	vkNUMPAD2   = 0x62
	vkNUMPAD3   = 0x63
	vkNUMPAD4   = 0x64
	vkNUMPAD5   = 0x65
	vkNUMPAD6   = 0x66
	vkNUMPAD7   = 0x67
	vkNUMPAD8   = 0x68
	vkNUMPAD9   = 0x69
	vkMULTIPLY  = 0x6A
	vkADD       = 0x6B
	vkSEPARATOR = 0x6C
	vkSUBTRACT  = 0x6D
	vkDECIMAL   = 0x6E
	vkDIVIDE    = 0x6F
	vkF1        = 0x70
	vkF2        = 0x71
	vkF3        = 0x72
	vkF4        = 0x73
	vkF5        = 0x74
	vkF6        = 0x75
	vkF7        = 0x76
	vkF8        = 0x77
	vkF9        = 0x78
	vkF10       = 0x79
	vkF11       = 0x7A
	vkF12       = 0x7B
	vkNUMLOCK   = 0x90
	vkSCROLL    = 0x91
	vkLSHIFT    = 0xA0
	vkRSHIFT    = 0xA1
	vkLCONTROL  = 0xA2
	vkRCONTROL  = 0xA3
	vkLMENU     = 0xA4
	vkRMENU     = 0xA5
	vkOEM1      = 0xBA
	vkOEMPLUS   = 0xBB
	vkOEMCOMMA  = 0xBC
	vkOEMMINUS  = 0xBD
	vkOEMPERIOD = 0xBE
	vkOEM2      = 0xBF
	vkOEM3      = 0xC0
	vkOEM4      = 0xDB
	vkOEM5      = 0xDC
	vkOEM6      = 0xDD
	vkOEM7      = 0xDE
	vkOEM8      = 0xDF
)
