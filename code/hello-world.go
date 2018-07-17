package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	that := "That Conference!"
	i18n := "هذا المؤتمر那次会议အဲဒီညီလာခံ"
	fmt.Println("Hello,", that)
	fmt.Println("Does", i18n, "work?") // in Arabic, Chinese, and Burmese?
}
