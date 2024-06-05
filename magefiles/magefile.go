package main

import (
	"fmt"
	"log"
	"os/exec"
	// mg contains helpful utility functions, like Deps
)

// Build Protobuf files
func Pb() error {
	//mg.Deps(Clean)
	fmt.Println("Building protobuf files...")
	_, err := exec.Command("protoc", "--go_out=greet", "greet/greetpb/greet.proto").Output()
	if err != nil {
		log.Fatal(err)
	}
	return nil

}

// Clean Protobuf files
func Clean() error {
	//mg.Deps(setopt)
	fmt.Println("Removing generated protobuf files...")
	_, err := exec.Command("rm", "-f", "greet/greetpb/greet.pb.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// Removing files suppression confirmation
func setopt() error {
	_, err := exec.Command("setopt", "rmstarsilent").Output()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
