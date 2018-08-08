package main

import (
	"encoding/json"
	"log"
)

const conf = `
{
	"expire_seconds": 123,
	"action_name":"QR_STR_SCENE",
	"action_info":{
	"scene": {
		"scene_str": "hello"
	}}
}`

func main() {
	request := struct {
		ExpireSeconds int    `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Scene struct {
				SceneStr string `json:"scene_str"`
			} `json:"scene"`
		} `json:"action_info"`
	}{
		1234,
		"QR_STR_SCENE_2",
		struct {
			Scene struct {
				SceneStr string `json:"scene_str"`
			} `json:"scene"`
		}{
			struct {
				SceneStr string `json:"scene_str"`
			}{
				"world",
			},
		},
	}

	log.Printf("request = %v", request)

	if err := json.Unmarshal([]byte(conf), &request); err != nil {
		panic(err)
	}

	log.Printf("request = %v", request)
}
