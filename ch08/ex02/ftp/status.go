package ftp

const (
	status200 = "200 Command okay."
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status230 = "230 User loggedin, proceed."

	// status331 = "331 User name okay, need password."

	// status421 = "421 Service not available, closing control connection."

	status500 = "500 Syntax error, command unrecognized."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status503 = "503 Bad sequence of commands."
	status504 = "504 Command not implemented for that parameter."
	status530 = "530 Not loggedin."
)
