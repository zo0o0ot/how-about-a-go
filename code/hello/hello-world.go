package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	that := "That Conference!"
	fmt.Println("Hello,", that)
	i18n := "هذا المؤتمر那次会议အဲဒီညီလာခံ"
	fmt.Println("Does", i18n, "work?") // in Arabic, Chinese, and Burmese?
}
