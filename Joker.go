package main

type joker_config struct {
	_Auth string
	_Self string
	_Os string
	_Addr string
	_history []string
	_command string
	_Active int64
}

type joker struct {
	_connected string
}