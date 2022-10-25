package templates

var Main = `
	/*
	{{.Copyright}}
	{{ if .Legal.Header }}{{ .Legal.Header }}{{ end }}
	*/
	package main

	func main() {
		fmt.Println("hihihi")
	}
`
