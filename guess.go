package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	fmt.Println("Я загадал число от 1 до 100. Угадаешь его?")
	fmt.Println(target)

	fmt.Print("Твоя попытка: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	if guess > target {
		fmt.Println("Задуманное число меньше!")
		fmt.Print("Попробуй еще раз:")
	} else if guess < target {
		fmt.Println("Задуманное число больше!")
		fmt.Print("Попробуй еще раз:")
	} else {
		fmt.Println("Угадал!")
	}

}
