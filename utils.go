package main

func gchar(c string, l int) string {
	var result string

	for i:=0;i<l;i++ {
		result+=c
	}

	return result
}