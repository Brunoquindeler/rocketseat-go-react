package main

import (
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command(
			"cmd",
			"/c",
			"tern",
			"migrate",
			"--migrations",
			"./internal/store/pgstore/migrations",
			"--config",
			"./internal/store/pgstore/migrations/tern.conf",
		)
	default:
		cmd = exec.Command(
			"tern",
			"migrate",
			"--migrations",
			"./internal/store/pgstore/migrations",
			"--config",
			"./internal/store/pgstore/migrations/tern.conf",
		)
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
