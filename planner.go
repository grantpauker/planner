//usr/bin/env go run "$0" "$@"; exit "$?"
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func getDate() string {
	return time.Now().Local().Format("01-02-06")
}

func regexReplace(text string, regex string, replace_with string) string {
	rp := regexp.MustCompile(regex)
	text = rp.ReplaceAllString(text, replace_with) // "def abc ghi"
	return text
}

func initPlanner(file string, subjects []string) {
	if !fileExists(file) {
		fmt.Print("Already plannerized, create a new one (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)
		if text != "y" {
			return
		}
	}
	f, _ := os.Create(file)
	w := bufio.NewWriter(f)
	for _, class := range subjects {
		fmt.Print("Homework for " + class + ": ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		w.WriteString(class + " " + text)
	}
	w.Flush()
}

func createSubjectMap(filename string, n map[string]string) {
	tmp := []string{}
	b, _ := os.Open(filename)
	scanner := bufio.NewScanner(b)

	for scanner.Scan() {
		tmp = strings.Split(scanner.Text(), " ")
		n[tmp[0]] = tmp[1]
	}

}

func map2File(file string, m map[string]string) {
	f, _ := os.Create(file)
	w := bufio.NewWriter(f)
	for subject, work := range m {
		w.WriteString(subject + " " + work + "\n")
	}
	w.Flush()
}
func editSubject(subject string, m map[string]string) {
	fmt.Print("Homework for " + subject + ": ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	if text == "" {
		text = "none"
	}
	m[subject] = text

}
func doneSubject(subject string, m map[string]string) {
	m[subject] = regexReplace(m[subject], "(.*)", "done ($1)")
}

func notDoneSubject(subject string, m map[string]string) {
	m[subject] = regexReplace(m[subject], "done \\((.*)\\)", "$1")
}

func printSubjectMap(subjects []string, m map[string]string) {
	for _, subject := range subjects {
		if strings.Contains(m[subject], "done (") || strings.Contains(m[subject], "none") {
			printRGB(subject+": "+m[subject]+"\n", 50, 255, 100)

		} else {
			printRGB(subject+": "+m[subject]+"\n", 255, 50, 50)

		}
	}
}
func printRGB(text string, r int, g int, b int) {
	fmt.Print("\x1b[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m" + text + "\x1b[0m")
}

func main() {
	subjects := []string{"biology", "english", "history", "math", "spanish"}
	the_args := os.Args[1:]
	path := "/home/god/Documents/school/planner/"

	file := path + getDate()
	m := make(map[string]string)
	createSubjectMap(file, m)
	length := len(the_args)
	if length > 0 {
		switch the_args[0] {
		case "-l":
			printSubjectMap(subjects, m)
		case "-e":
			if length < 2 {
				fmt.Println("Please specify a subject")
			} else {
				editSubject(the_args[1], m)
			}
		case "-d":
			if length < 2 {
				fmt.Println("Please specify a subject")
			} else {
				doneSubject(the_args[1], m)
			}
		case "+d":
			if length < 2 {
				fmt.Println("Please specify a subject")
			} else {
				notDoneSubject(the_args[1], m)
			}

		}

	} else {
		initPlanner(file, subjects)
	}
	map2File(file, m)
	// editSubject("math", n)

	// doneSubject("history", n)
	// notDoneSubject("history", n)
}
