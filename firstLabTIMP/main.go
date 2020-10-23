package main

import (
	"bufio"
	"fmt"
	"gopkg.in/gookit/color.v1"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)
const directory = "/home/botiyava/go_try/"
func password(list []string){
	var choice string
	fmt.Print("Для завершения работы программы нажмите клавишу \"s\"\n")
	fmt.Fscan(os.Stdin, &choice)
	if choice == "s" {
		color.Println("<lightRed>Обращаем ваше внимание на то, что при вводе неверного пароля " +
			"программа \n экстренно завершится, файлы будут заблокированы и будет вызвана охрана\n" +
			"Введите пароль:\n</>")
		var pwd string
		fmt.Fscan(os.Stdin, &pwd)
		if pwd == list[0] {
			for _, val := range list {
				if val != list[0] {
					cmd := exec.Command("sudo", "chattr", "-i", "/home/botiyava/go_try/"+val)

					cmd.Run()
					os.Chmod(directory+val, 0700)

				}
			}
			os.Exit(1)
		}
	}
}

func Difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}


func main() {

	file, err := os.Open(directory + "list0.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		os.Chmod(directory + scanner.Text(),0000)
	}

	for _,val := range lines{
		if val != "qw" {
			cmd := exec.Command("sudo", "chattr", "+i", "/home/botiyava/go_try/"+val)

			cmd.Run()
		}
	}
	go password(lines)
	files, err := ioutil.ReadDir("/home/botiyava/go_try/")
	if err != nil {
		log.Fatal(err)
	}
	var startFile []string
	for _, file := range files {
		startFile = append(startFile, file.Name())
	}
	var secureList []string
	for i := 0; i < len(startFile); i++ {
		for j := 0; j < len(lines); j++ {
			if startFile[i] == lines[j] {
				secureList = append(secureList, startFile[i])
			}
		}
	}

	for {

		files, err := ioutil.ReadDir("/home/botiyava/go_try/")
		if err != nil {
			log.Fatal(err)
		}
		var currentFile []string
		for _, file := range files {
			currentFile = append(currentFile, file.Name())
		}
	targets := Difference(lines,secureList)

		for i := 0;i < len(currentFile); i++{
			for j:= 0; j < len(targets); j++{

				if currentFile[i] == targets[j]{
					os.Remove(directory + currentFile[i])
				}

			}
		}
	time.Sleep(100 * time.Microsecond)
        }
}
