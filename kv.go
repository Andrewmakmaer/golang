package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	k := parser.String("k", "key", &argparse.Options{Required: true, Help: "Name of key"})
	v := parser.String("v", "value", &argparse.Options{Required: false, Help: "Value", Default: ""})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	path, _ := os.Getwd()
	jsonFile, err := os.OpenFile(path+"/file1.json", os.O_RDONLY|os.O_CREATE, 0666)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := make(map[string]string)
	json.Unmarshal([]byte(byteValue), &result)
	if *v == "" {
		fmt.Println(result[*k])
	} else {
		result[*k] = *v
		rawDataOut, _ := json.MarshalIndent(&result, "", "  ")
		err = ioutil.WriteFile(path+"/file1.json", rawDataOut, 0)
	}
}
