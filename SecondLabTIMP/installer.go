	package main

	import (
		"log"
		"os"
		"os/exec"
	)
	const directory = "/home/botiyava/go_try/FellInLove/"
	func main() {
	cmd := exec.Command("tar","-xf",directory + "Lab2/main.tar", "-C", directory + "Lab2/")
	error := cmd.Run()
	if error != nil{
		log.Fatal("Архива программы не существует или он не повреждён")
		os.Exit(1000)
	}


		_, err := os.Stat(directory + ".userlist.txt")
		if err != nil {
			if os.IsNotExist(err) {
				os.Create(directory + ".userlist.txt")
			} else {
				log.Fatal(err)
			}
		}
		os.Remove("installer")
	}
