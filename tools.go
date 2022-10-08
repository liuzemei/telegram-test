package main

import (
	"bytes"
	"encoding/json"
	"log"
	"runtime"
)

func PrintJson(d interface{}) {

	s, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, s, "", "\t")
	if err != nil {
		log.Println("JSON parse error: ", err)
		return
	}
	log.Println(prettyJSON.String())
}

func Log(v ...interface{}) {
	caller := 1
	for {
		if _, file, line, ok := runtime.Caller(caller); ok {
			log.Println("warning...", v, file, line)
			caller++
		} else {
			return
		}
	}
}
