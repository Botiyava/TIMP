	package main

	import (
		"fmt"
		"gopkg.in/gookit/color.v1"
		"io/ioutil"
		"log"
		"os"
		"strconv"
		"strings"
		"time"
	)
	const directory = "/home/botiyava/go_try/FellInLove/"
	type User struct{
		name string
		patronymic string
		surname string

		numberOfLaunch int
		superUser int
	}
	func findUser(list []User, len int) {

		fmt.Print("Введите имя пользователя:\n")
		var name string
		fmt.Fscan(os.Stdin, &name)

		fmt.Print("Введите отчество пользователя:\n")
		var patronymic string
		fmt.Fscan(os.Stdin, &patronymic)

		fmt.Print("Введите фамилию пользователя:\n")
		var surname string
		fmt.Fscan(os.Stdin, &surname)

		var great = User{name,patronymic,surname, 1, 0}
		res := false
		var flag int
		for i := 0; i < len; i++{
			if  great.name == list[i].name && great.surname == list[i].surname && great.patronymic == list[i].patronymic{
				flag = i
				res = true
			}
		}
		if res == true{
			list[flag].numberOfLaunch += 1
			if list[flag].numberOfLaunch >5 && list[flag].superUser == 0{
				color.Red.Print("Вы исчерпали лимит на использование бесплатной версии этой программы.\n" +
					"Чтобы продолжить работу приобретите полную версию за $300.\n" +
					"Вы хотите приобрести полную версию программы? (д/н)\n")
				var full string
				fmt.Fscan(os.Stdin, &full)
				if full == "д"{
					list[flag].superUser = 1

				}else{
					for i := 0; i <3; i++ {
						fmt.Printf("Программа будет удалена через: %d\n",3 - i)
						time.Sleep(1 * time.Second)

					}
					os.Remove(directory + "Lab2/main.tar")
					os.Remove(directory + "Lab2/main")
					os.Remove(directory + "Lab2/installer")
					os.Remove(directory + "Lab2/instruction.txt")
					os.Remove(directory +"Lab2")
					os.Remove(directory + "Lab2.tar" )
					os.Exit(0)
				}


			}
			mydata := []byte("")
			ioutil.WriteFile(directory + ".userlist.txt", mydata, 0777 )
			for _,val := range list{

				f, err := os.OpenFile(directory + ".userlist.txt", os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					panic(err)
				}
				defer f.Close()

				if _, err = f.WriteString(val.name + " " + val.patronymic + " " + val.surname + " " + strconv.Itoa(val.numberOfLaunch) +" "+ strconv.Itoa( val.superUser) + "\n"); err != nil {
					panic(err)
				}
			}

				color.Cyan.Printf("Добро пожаловать, %v %v!\n", list[flag].name, list[flag].patronymic)
			if list[flag].numberOfLaunch <6 && list[flag].numberOfLaunch >0 {
				color.Cyan.Printf("\nУ вас осталось %v бесплатных запусков.\n", 5-(list[flag].numberOfLaunch))
			}

			if list[flag].superUser != 1 {
				time.Sleep(30 * time.Second)
			} else{
				//time.Sleep(100000 * time.Hour)
				a:= 0
				fmt.Scan(&a)
				if a == 1{
				}
			}
				color.Cyan.Printf("Завершение программы...\n")

		}else{

			list = append(list,great)
			f, err := os.OpenFile(directory + ".userlist.txt", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			if _, err = f.WriteString(great.name + " " + great.patronymic + " " + great.surname + " " + strconv.Itoa(great.numberOfLaunch) +" "+ strconv.Itoa(great.superUser) +"\n"); err != nil {
				panic(err)
			}
				color.Cyan.Printf("\nУ вас осталось 4 бесплатных запуска.\n\n")
			time.Sleep(30 * time.Second)
				color.Cyan.Printf("Завершение программы...\n")
			}



		return
	}
	func main() {

		users := make([]User, 0)
		userFile, err := ioutil.ReadFile(directory + ".userlist.txt")
		if err != nil {
			log.Fatal(err)
		}
		userLines := strings.Split(string(userFile), "\n")
		for i := 0; i < len(userLines); i++ {
			if userLines[i] != "" {
				userLine := strings.Split(userLines[i], " ")

				number, err := strconv.Atoi(userLine[3])
				if err != nil {
					log.Fatal(err)
				}
				super,err :=strconv.Atoi(userLine[4])
				if err != nil{
					log.Fatal(err)
				}

				newUser := User{userLine[0], userLine[1], userLine[2],
					number,super }
				users = append(users, newUser)
			}
		}
		findUser(users, len(users))



	}
