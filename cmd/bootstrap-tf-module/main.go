package main

import (
	"fmt"
	"os"
)

var moduleDirPath = os.Args[1]
var moduleName = os.Args[2]
var path = "/home/users/lbrewerton/Documents/test.txt"
var moduleFullPath = moduleDirPath + moduleName

var variableFileName = "variables.tf"
var moduleFileName = moduleName + ".tf"
var outputsFileName = "outputs.tf"

func main() {
	fmt.Println("Bootstrapping Terraform module...")

	createModuleDirectory()
	createVariablesFile()
	createModuleFile()
	createOutputsFile()
}

func createModuleDirectory() {
	var _, err = os.Stat(moduleFullPath)

	if os.IsNotExist(err) {
		os.Mkdir(moduleFullPath, 0755)
		if isError(err) {
			return
		}
	}

	fmt.Println("===> done creating variables file")
}

func createVariablesFile() {
	var _, err = os.Stat(moduleFullPath + variableFileName)

	if os.IsNotExist(err) {
		var file, err = os.Create(moduleFullPath + variableFileName)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("===> done creating variables file")
}

func writeVariablesFile() {
	var file, err = os.OpenFile(moduleFullPath+variableFileName, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("# Variables file")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("===> done bootstrapping variables file")
}

func createModuleFile() {
	var _, err = os.Stat(moduleFullPath + moduleFileName)

	if os.IsNotExist(err) {
		var file, err = os.Create(moduleFullPath + moduleFileName)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("===> done creating module file")
}

func writeModuleFile() {
	var file, err = os.OpenFile(moduleFullPath+moduleFileName, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("# Module file")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("===> done bootstrapping module file")
}

func createOutputsFile() {
	var _, err = os.Stat(moduleFullPath + outputsFileName)

	if os.IsNotExist(err) {
		var file, err = os.Create(moduleFullPath + outputsFileName)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("===> done creating outputs file")
}

func writeOutputsFile() {
	var file, err = os.OpenFile(moduleFullPath+outputsFileName, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("# Outputs file")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("===> done bootstrapping outputs file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
