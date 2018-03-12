package utils

import "strings"

type hostAndPort struct {
	host string
	port string
	uri  string
}

func (h *hostAndPort) initialize() {
	strList := strings.Split(h.uri, ":")
	h.host = strList[0]
	h.port = strList[1]
}

// Base connection
type Conn struct {
	hapList []hostAndPort
	uri     string
}

func (c *Conn) setHostList(hapList []string) {
	for i := 0; i < len(hapList); i++ {
		hap := hostAndPort{uri: hapList[i]}
		hap.initialize()
		c.hapList = append(c.hapList, hap)
	}
}

func (c *Conn) generateURI() {
	for i := 0; i < len(c.hapList); i++ {
		hapURI := c.hapList[i].uri
		if i == 0 {
			c.uri += hapURI
		} else {
			c.uri = c.uri + "," + hapURI
		}
	}
}

// Connection with auth setting
type authConn struct {
	conn     Conn
	username string
	password string
	uri      string
}

func (ac *authConn) generateURI() {
	ac.uri = ac.authURI() + ac.conn.uri
}

func (ac authConn) authURI() string {
	return ac.username + ":" + ac.password + "@"
}
