package main

import (
	"fmt"
	irc "github.com/fluffle/goirc/client"
)

func main() {
	fmt.Println("IRC bridge initialized")

	// Creating a simple IRC client is simple.
	c := irc.SimpleClient("clusterioBridge")

	// Add handlers to do things here!
	// e.g. join a channel on connect.
	c.HandleFunc(irc.CONNECTED,
		func(conn *irc.Conn, line *irc.Line) { conn.Join("#factorio") })
	// And a signal on disconnect
	quit := make(chan bool)
	c.HandleFunc(irc.DISCONNECTED,
		func(conn *irc.Conn, line *irc.Line) { quit <- true })

	// handle people typing things
	c.HandleFunc("privmsg",
		func(conn *irc.Conn, line *irc.Line) {
			str := "/c game.print( '" + line.Nick + " | " + line.Args[1] + "' )"
			//str = strings.Replace(str, "'", "\\'", -1)
			fmt.Println(str)
		})

	// ... or, use ConnectTo instead.
	if err := c.ConnectTo("ipo.esper.net"); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}

	// Wait for disconnect
	<-quit
}
