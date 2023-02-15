package main

import (
	"fmt"
	"regexp"
	"time"
)

func display_sessions() {
	var timenow int64 = time.Now().UTC().Unix()

	var longest_username int
	for i:=0;i<len(sessions);i++ {
		if len(sessions[i]._Self) > longest_username {
			longest_username = len(sessions[i]._Self)
		}
	}

	if longest_username < 4 {
		longest_username = 4
	}

	fmt.Println("SessionID            IPAddress       Os      User",gchar(" ", longest_username - 4)+"Status")
	fmt.Println("-------------------- --------------- -------",gchar("-", longest_username)+" --------")
	for i:=0;i<len(sessions);i++ {
		var status string = Fore["RED"]+"Inactive"+Fore["RESET"]
		if sessions[i]._Active > timenow - 5 {
			status = Fore["GREEN"]+"Active"+Fore["RESET"]
		}
		fmt.Println(sessions[i]._Auth,sessions[i]._Addr+gchar(" ", 15 - len(sessions[i]._Addr)),sessions[i]._Os+gchar(" ", 7 - len(sessions[i]._Os)),sessions[i]._Self+gchar(" ", longest_username - len(sessions[i]._Self)),status)
	}

	if len(sessions) == 0 {
		fmt.Println("None                 None            None    None",gchar(" ", longest_username - 4)+"None")
	}
}


func VerifySessionID(sessionID string) bool {
    pattern := "^[a-f0-9]{6}-[a-f0-9]{6}-[a-f0-9]{6}$"

    regex := regexp.MustCompile(pattern)

    return regex.MatchString(sessionID)
}