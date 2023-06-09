package utils

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

type Link struct {
	Url string
	Action map[string]string
}

type Menu struct {
	Url string
	Show string
	Create string
	Put string
	Delete string
	Detail string
	Change string
}

func GetMenu() []Menu {
	m := []Menu{
		{Url: "clients", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "modules", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "operations", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "roles", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "users", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail", Change: "change"},
		{Url: "dashboard", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "profiles", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	Change: "change"},
	}
	return m
}

func BirthdayTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:04:05"
	t, _ := time.Parse(Format, timeStr)
	return t
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func Execute(cmd string, args string) {
	out, err := exec.Command(cmd, args).Output()

	if err != nil {
			fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

func Chdir(newdir string) {
  
  // Getting the current working directory
	CurrentWD, _ := syscall.Getwd()
	fmt.Println("CurrentWD:", CurrentWD)

	// Changing the working directory
	syscall.Chdir("/home/tech")

	// Again,
	// getting the current working directory
	CurrentWD, _ = syscall.Getwd()
	fmt.Println("CurrentWD:", CurrentWD)
}
