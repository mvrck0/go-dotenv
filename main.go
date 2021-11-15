package main

import (
        // "fmt"
        // "github.com/thatisuday/commando"
        // "os"
        // "os/exec"
        // "path/filepath"

        "os"
        "strings"

        "github.com/fatih/color"
        "github.com/joho/godotenv"
)

const dotenvVersion = "0.0.1"

var verbose = true

type WalkerConfig struct {
	parseFiles  bool
	parseDirs   bool
	parsePrefix string
}

func main() {
        findUpstream()
}

func findUpstream() {
	path, _ := os.Getwd()
	paths := []string{}
	prevPath := ""

	for _, name := range strings.Split(path, "/") {
		if name == "" {
				continue
		}
		prevPath = prevPath + "/" + name
		paths = append(paths, prevPath)
	}

	for _, path := range paths {
		hasDotEnv(path)
	}
}

func hasDotEnv(dir string) {
	dotEnvPath := dir + "/" + ".env"
	if fileExists(dotEnvPath) {
		color.Green(dotEnvPath)
		loadFile(dotEnvPath, true)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func loadFile(filename string, overload bool) error {
	envMap, err := readFile(filename)
	if err != nil {
			return err
	}

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()

	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] || overload {
			os.Setenv(key, value)
			println(key, "=", value)
		}
	}

	return nil
}

func readFile(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	return godotenv.Parse(file)
}