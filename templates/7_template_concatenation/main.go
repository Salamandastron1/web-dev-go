package main

import "fmt"

func main() {
	name := "Tim Garrity"

	tpl := `
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="UTF-8">
	<title>Hello World</title>
	<body>
	<h1>` + name + `</h1>
	</body>
	</head>
	</html>
	`
	fmt.Println(tpl)
}
