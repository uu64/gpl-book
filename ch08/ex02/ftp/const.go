package ftp

const (
	protocolIp4 = "ip4"
	protocolIp6 = "ip6"
)

const (
	typeAscii = "A"
	typeImage = "I"
	// typeLocal  = "L"
	// typeEbcdic = "E"
)

const (
	modeStream = "S"
	// modeBlock      = "B"
	// modeCompressed = "C"
)

const (
	struFile = "F"
	// struRecord = "R"
	// struPage   = "P"
)

const (
	// status125 = "125 Data connection already open; transfer starting."
	status150 = "150 File status okay; about to open data connection."

	status200 = "200 Command okay."
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection."
	status230 = "230 User loggedin, proceed."
	status250 = "250 Requested file action okay, completed."
	status257 = "257 \"%s\" is the current directory."

	// status331 = "331 User name okay, need password."

	// status421 = "421 Service not available, closing control connection."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status450 = "450 Requested file action not taken."
	status451 = "451 Requested action aborted. Local error in processing."
	status452 = "452 Requested action not taken."

	status500 = "500 Syntax error, command unrecognized."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status503 = "503 Bad sequence of commands."
	status504 = "504 Command not implemented for that parameter."
	status530 = "530 Not loggedin."
	status550 = "550 Requested action not taken."
	status551 = "551 Requested action aborted. Page type unknown."
	status552 = "552 Requested file action aborted."
	status553 = "553 Requested action not taken."
)
