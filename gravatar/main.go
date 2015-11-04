package main

import (
"fmt"
"strings"
	"crypto/md5"
	"io"
	"encoding/hex"
	"os"
)

func getGravataHash(email string)string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)
	finalByte := h.Sum(nil)
	finalString := hex.EncodeToString(finalByte)
	return finalString
}

//*TODO	add in feature to get user name and email address and redisplay them to the page

func main() {

	fmt.Fprintln(os.Stderr,"Hello and Welcome. \nWhat is your name")
	var name string
	fmt.Scan(&name)

	fmt.Fprintln(os.Stderr, "What is your email address")
	var email string
	fmt.Scan(&email)
	gravatarHas := getGravataHash(email)

	fmt.Println(`<!DOCTYPE html>
	<html>
	<head>
	</head>
	<body>
		<h1>Welcome ` + name + ` </h1>

		<img src="http://www.gravatar.com/avatar/` + gravatarHas + `?d=identicon">
	</body>
	</hmtl>`)
}
