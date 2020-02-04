package main

import (
	"fmt"

	"os"

	"github.com/BurntSushi/toml"
	"github.com/kaepa3/rotation"
)

const RotationFile = "./rotation.toml"

func main() {
	rotation, err := ReadStruct()
	if err != nil {
		fmt.Println("no Read")
		os.Exit(-1)
	}
	result := rotation.Rotate()
	fmt.Println(result)

	WriteStruct(rotation)
}

func ReadStruct() (*rotation.Rotation, error) {
	var rotation rotation.Rotation
	_, err := toml.DecodeFile(RotationFile, &rotation)
	if err != nil {
		return nil, err
	}
	return &rotation, nil
}

func WriteStruct(rotation *rotation.Rotation) error {
	file, err := os.Create(RotationFile)
	if err != nil {
		return err
	}
	defer file.Close()
	encorder := toml.NewEncoder(file)
	err = encorder.Encode(rotation)
	if err != nil {
		return err
	}
	return nil
}
