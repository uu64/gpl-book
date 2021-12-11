package ftp

const (
	status220 = "220 Service ready for new user.\n"
	status500 = "500 Syntax error, command unrecognized.\n"
	// USER
	status230 = "230 User loggedin, proceed.\n"
	status530 = "530 Not loggedin.\n"
	status331 = "331 User name okay, need password.\n"
	// QUIT
	status221 = "221 Service closing control connection.\n"
)
