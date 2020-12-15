package main

import "os"

const directory = "/home/botiyava/go_try/FellInLove/"
func main(){
os.Remove(directory + "main.tar")
os.Remove(directory + "main")
os.Remove(directory + "installer")

}
