package service

import (
	"Security-Benchmarking-Tool/constants"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

// usage example in locale test2
func applyChecks(stringToSwitch string, myMap map[string]string) (ok bool, err error) {
	switch stringToSwitch {
	case "BANNER_CHECK":
		if ok, err := fileBannerCheck(myMap); err != nil {
			return false, err
		} else {
			return ok, nil
		}
	case "CMD_EXEC":
		if ok, err := cmdExecOperation(myMap); err != nil {
			return false, err
		} else {
			return ok, nil
		}
	case "FILE_CONTENT_CHECK":
		if ok, err := fileContentCheckOperation(myMap); err != nil {
			return false, err
		} else {
			return ok, nil
		}
	case "FILE_CHECK_NOT":
		if ok, err := fileCheckNotOperation(myMap); err != nil {
			return false, err
		} else {
			return ok, nil
		}
	case "FILE_CONTENT_CHECK_NOT":
		if ok, err := fileContentCheckNotOperation(myMap); err != nil {
			return false, err
		} else {
			return ok, nil
		}
	case "FILE_CHECK":
		if ok, err := fileCheckOperation(myMap); err != nil {
			return false, err
		} else {
			return ok, nil
		}
	}

	return true, nil
}

func fileContentCheckOperation(customPolicyMap map[string]string) (ok bool, err error) {
	fileData, err := ioutil.ReadFile(customPolicyMap["file"])
	if err != nil {
		return false, err
	}

	ok, err = regexp.Match(customPolicyMap["regex"], fileData)
	if err != nil {
		return
	}

	return
}

func fileBannerCheck(customPolicyMap map[string]string) (ok bool, err error) {

	pathOfFileToEditOrCreate := customPolicyMap["file"]

	if _, err = os.Stat(pathOfFileToEditOrCreate); err == nil {
		ok = false
	}
	ok = true

	return
}

func fileCheckOperation(customPolicyMap map[string]string) (ok bool, err error) {
	log.Println("WILL BE IMPLEMENTED IN NEXT LAAAB!")

	fmt.Println(customPolicyMap["info"])

	//need to check description and to parse commands

	return true, nil
}

func fileContentCheckNotOperation(customPolicy map[string]string) (ok bool, err error) {
	fileData, err := ioutil.ReadFile(customPolicy["file"])
	if err != nil {
		return false, err
	}

	ok, err = regexp.Match(customPolicy["regex"], fileData)
	if err != nil {
		return
	}

	return
}

func fileCheckNotOperation(customPolicy map[string]string) (ok bool, err error) {

	pathOfFileToEditOrCreate := customPolicy["file"]

	if _, err = os.Stat(pathOfFileToEditOrCreate); err == nil {
		ok = false
	}
	ok = true
	return true, nil
}

func cmdExecOperation(customPolicyMap map[string]string) (ok bool, err error) {

	err, stdout, stderr := Shellout(customPolicyMap["cmd"], "-c")
	if err != nil {
		return false, err
	}

	if ok, err = regexp.Match(customPolicyMap["expect"], []byte(stdout)); err != nil {
		log.Println("OK IN MY CHECK: ", ok)
		return
	}

	log.Println("stderr: ", stderr)

	return
}

// function for cmd
func Shellout(command string, args string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(constants.ShellToUse, args, command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
