package x11

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	envKey = "DISPLAY"
)

var (
	displayRegex = regexp.MustCompilePOSIX("^([^:]*):([0-9]+)(.([0-9]*))?$")
)

// DisplayNetAddr retrieves the $DISPLAY environment variable, and returns the corresponding ‘network’ and ‘address’,
// that could be used in a call to the "net" package's net.Dial() func, for the X11 client to connect to the X11 server.
//
//
// X11 Clients Versus X11 Servers
//
// With X11, there are clients and servers.
//
// Most people (who work with X11) create X11 clients, as part of creating their application.
//
// (Most people have an X11 server already installed, and already running on their computer. The X11 server probably
// came pre-installed on their computer. Or was installed when they installed their operating system.)
//
// An X11 server is what displays images to the screen or monitor, and is communicating with the graphics hardware.
// There is usually only one X11 server runnig on a computer. And often the X11 server is automatically started when
// the computer is started.
//
// It is probably easiest to think of an X11 client as an application. An X11 client communicates with the X11 server,
// and asks it (for example) to create a window, display an image, draw various graphics, etc. The X11 client tells
// the X11 server what to display and draw.
//
// With X11, before an X11 client connects to an X11 server, it looks at the $DISPLAY environment variable, to determine
// where the server is.
//
// DisplayNetAddr would potentially be used by those creating an X11 client. (It is unlikely to be useful to an X11 server.)
// Although most people using this package will probably not call DisplayNetAddr directly, and instead use x11.Dial().
// (More on x11.Dial() shortly.)
//
//
// Example Usage
//
// This is how one might use DisplayNetAddr:
//
//	network, address, err := x11.DisplayNetAddr()
//	if nil != err {
//		return err
//	}
//	
//	conn, err := net.Dial(network, address)
//
// However, note that most people will probably not call DisplayNetAddr directly themselves, but instead will use this package's xll.Dial() func:
//
//	conn, err := x11.Dial()
//
//
// $DISPLAY Environment Variable
//
// Although not necessary to know to use DisplayNetAddr, to see what the value of the $DISPLAY environment variable is on your system, go to the shell (ex: bash), and type the following command:
//
//	echo $DISPLAY
//
// You might see something such as:
//
//	:0
//
// Or:
//
//	:0.0
//
// Or (if you connecting into a machine via “ssh -X ...”) something such as:
//
//	93.184.216.34:12.0
//
//
// $DISPLAY Environment Variable Format
//
// Although not necessary to know to use DisplayNetAddr, the $DISPLAY environment variable can appear in these format:
//
// • :display
//
// • :display.screen
//
// • host:display
//
// • host:display.screen
//
// So, effectively, ‘display’ is mandatory, and ‘host’ and ‘screen’ are optional.
//
// With these formats:....
//
// • ‘display’ is an integer that can range from 0 (zero) to 59,535 (= the maximum TCP port number (i.e., 65,535) minus 6,000);
//
// • ‘host’ can be (left out, or) “unix”, or an IP address, or a domain name;
//
// • ‘screen’ can be (left out, or) 0 (zero) or any positive integer; (the author is not actually sure if or what the maximum value for ‘screen’ is); and
//
// Some examples of the :display format are:
//
//	:0
//
//	:1
//
//	:2
//
//	:13
//
//	:32884
//
// Some examples of the :display.screen format are:
//
//	:0.0
//
//	:1.0
//
//	:2.0
//
//	:13.0
//
//	:32884.0
//
// And also:
//
//	:0.1
//
//	:1.1
//
//	:2.1
//
//	:13.1
//
//	:32884.1
//
// And also:
//
//	:0.5
//
//	:1.5
//
//	:2.5
//
//	:13.5
//
//	:32884.5
//
// And also:
//
//	:0.4528
//
//	:1.4528
//
//	:2.4528
//
//	:13.4528
//
//	:32884.4528
//
// Some examples of the host:display format, where the ‘host’ is “unix” are:
//
//	unix:0
//
//	unix:1
//
//	unix:2
//
//	unix:13
//
//	unix:32884
//
// Some examples of the host:display format, where the ‘host’ is an IP address are:
//
//	127.0.0.1:0
//
//	127.0.0.1:1
//
//	127.0.0.1:2
//
//	127.0.0.1:13
//
//	127.0.0.1:32884
//
// And also:
//
//	93.184.216.34:0
//
//	93.184.216.34:1
//
//	93.184.216.34:2
//
//	93.184.216.34:13
//
//	93.184.216.34:32884
//
// Some examples of the host:display format, where the ‘host’ is a domain name are:
//
//	example.com:0
//
//	example.com:1
//
//	example.com:2
//
//	example.com:13
//
//	example.com:32884
//
// Some examples of the host:display.screen format, where the ‘host’ is “unix” are:
//
//	unix:0.0
//
//	unix:1.0
//
//	unix:2.0
//
//	unix:13.0
//
//	unix:32884.0
//
// And also:
//
//	unix:0.1
//
//	unix:1.1
//
//	unix:2.1
//
//	unix:13.1
//
//	unix:32884.1
//
// And also:
//
//	unix:0.5
//
//	unix:1.5
//
//	unix:2.5
//
//	unix:13.5
//
//	unix:32884.5
//
// And also:
//
//	unix:0.4528
//
//	unix:1.4528
//
//	unix:2.4528
//
//	unix:13.4528
//
//	unix:32884.4528
//
// Some examples of the host:display.screen format, where the ‘host’ is an IP address are:
//
//	127.0.0.1:0.0
//
//	127.0.0.1:1.0
//
//	127.0.0.1:2.0
//
//	127.0.0.1:13.0
//
//	127.0.0.1:32884.0
//
// And also:
//
//	127.0.0.1:0.1
//
//	127.0.0.1:1.1
//
//	127.0.0.1:2.1
//
//	127.0.0.1:13.1
//
//	127.0.0.1:32884.1
//
// And also:
//
//	127.0.0.1:0.5
//
//	127.0.0.1:1.5
//
//	127.0.0.1:2.5
//
//	127.0.0.1:13.5
//
//	127.0.0.1:32884.5
//
// And also:
//
//	127.0.0.1:0.4528
//
//	127.0.0.1:1.4528
//
//	127.0.0.1:2.4528
//
//	127.0.0.1:13.4528
//
//	127.0.0.1:32884.4528
//
// And also:
//
//	93.184.216.34:0.0
//
//	93.184.216.34:1.0
//
//	93.184.216.34:2.0
//
//	93.184.216.34:13.0
//
//	93.184.216.34:32884.0
//
// And also:
//
//	93.184.216.34:0.1
//
//	93.184.216.34:1.1
//
//	93.184.216.34:2.1
//
//	93.184.216.34:13.1
//
//	93.184.216.34:32884.1
//
// And also:
//
//	93.184.216.34:0.5
//
//	93.184.216.34:1.5
//
//	93.184.216.34:2.5
//
//	93.184.216.34:13.5
//
//	93.184.216.34:32884.5
//
// And also:
//
//	93.184.216.34:0.4528
//
//	93.184.216.34:1.4528
//
//	93.184.216.34:2.4528
//
//	93.184.216.34:13.4528
//
//	93.184.216.34:32884.4528
//
// Some examples of the host:display.screen format, where the ‘host’ is a domain name are:
//
//	example.com:0.0
//
//	example.com:1.0
//
//	example.com:2.0
//
//	example.com:13.0
//
//	example.com:32884.0
//
// And also:
//
//	example.com:0.1
//
//	example.com:1.1
//
//	example.com:2.1
//
//	example.com:13.1
//
//	example.com:32884.1
//
// And also:
//
//	example.com:0.5
//
//	example.com:1.5
//
//	example.com:2.5
//
//	example.com:13.5
//
//	example.com:32884.5
//
// And also:
//
//	example.com:0.4528
//
//	example.com:1.4528
//
//	example.com:2.4528
//
//	example.com:13.4528
//
//	example.com:32884.4528
//
//
// Interpretation
//
// If ‘host’ is left out, or it is “unix”, then then ‘network’ will be “unix”.
// (I.e., a Unix domain socket, which is also sometimes call an IPC socket.)
//
// Else, ‘host’ is interpreted as a domain name or an IP address.
//
// If ‘screen’ is left out, then it defaults to “0”.
//
// The value for ‘display’ is mandatory, and cannot be left out. (Leaving it out will cause DisplayNetAddr to return a error.)
//
//
// Unix Domain Socket
//
// If ‘host’ is left out or “unix”, such as with:
//
//	:0
//
//	:0.0
//
//	unix:0
//
//	unix:0.0
//
//	:1
//
//	:1.42
//
//	unix:1
//
//	unix:1.42
//
//	:2
//
//	:2.42
//
//	unix:2
//
//	unix:2.42
//
//	:13
//
//	:13.42
//
//	unix:13
//
//	unix:13.42
//
//	:4528
//
//	:4528.42
//
//	unix:4528
//
//	unix:4528.42
//
// … then the returned ‘address’ will be of the format:
//
//	/tmp/.X11-unix/X<display>
//
// So if ‘display’ is “0”, then the returns ‘address’ will be :
//
//	/tmp/.X11-unix/X0
//
// If ‘display’ is “1”, then the returns ‘address’ will be :
//
//	/tmp/.X11-unix/X1
//
// If ‘display’ is “2”, then the returns ‘address’ will be :
//
//	/tmp/.X11-unix/X2
//
// If ‘display’ is “13”, then the returns ‘address’ will be :
//
//	/tmp/.X11-unix/X13
//
// If ‘display’ is “4528”, then the returns ‘address’ will be :
//
//	/tmp/.X11-unix/X4528
//
// These are all paths to a Unix domain socket, in the local file system.
// (Remember, in Unix and Linux, “everything is a file”. Well… “everything is a file [or a process]” 🙂)
//
// And in these situations, the X11 client will connect to this Unix domain socket to communicate with the X11 server.
func DisplayNetAddr() (network string, address string, err error) {

	envValue := os.Getenv(envKey)

	return displaynetaddr(envValue)
}

func displaynetaddr(value string) (string, string, error) {

	if "" == value {
		return "", "", fmt.Errorf("x11: Problem with $%s environment variable: either not set or value is empty: %q", envKey, value)
	}

	matches := displayRegex.FindStringSubmatch(value)
	if nil == matches {
		return "", "", fmt.Errorf("x11: Problem with $%s environment variable: it is not in the expected format, it is: %q", envKey, value)

	}
	if 5 != len(matches) {
		return "", "", fmt.Errorf("x11: Problem parsing $%s environment variable: internal error.", envKey)
	}

	host    := matches[1]
	display := matches[2]
	screen  := matches[4]

	if "" == screen {
		screen = "0"
	}
//@TODO: We never actually do anything with screen!

	switch host {
	case "","unix":
		addr := fmt.Sprintf("/tmp/.X11-unix/X%s", display)

		return "unix", addr, nil
	default:
		displayUint64, err := strconv.ParseUint(display, 10, 64)
		if nil != err {
			return "", "", fmt.Errorf("x11: Problem parsing display number in $%s environment variable: %s", err)
		}

		var port uint64 = 6000 + displayUint64

		addr := fmt.Sprintf("%s:%d", host, port)

		return "tcp", addr, nil
	}
}
