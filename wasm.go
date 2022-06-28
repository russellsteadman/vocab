//go:build js
// +build js

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/russellsteadman/vocab/internal/vocab"
)

func sendData(data interface{}) string {
	jsonText, err := json.Marshal(struct {
		Error bool        `json:"error"`
		Data  interface{} `json:"data"`
	}{false, data})
	if err != nil {
		return sendError(err.Error())
	}
	return string(jsonText)
}

func sendError(text string) string {
	jsonText, err := json.Marshal(struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}{true, text})
	if err != nil {
		return "{\"error\":true,\"message\":\"Unable to encode error\"}"
	}
	return string(jsonText)
}

func main() {
	js.Global().Set("createVocab", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return sendError("Invalid number of arguments")
		}

		rawInput := args[0].String()

		input := &vocab.VocabInput{}
		err := json.Unmarshal([]byte(rawInput), input)
		if err != nil {
			return sendError("Unable to decode input")
		}

		vocab, err := vocab.CreateSet(input)
		if err != nil {
			return sendError(err.Error())
		}

		return sendData(vocab)
	}))

	<-make(chan bool)
}
