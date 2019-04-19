package main

import (
	"fmt"

	"github.com/sqeven/go-ai/vision"
	"github.com/sqeven/go-ai/vision/face/v2"
)

func DetectAndAnalysis() {
	client := v2.NewFaceClient(APIKEY, APISECRET)
	options := map[string]interface{}{
		"max_face_num": 10,
		"face_fields":  "age,beauty,expression,faceshape,gender,glasses,landmark,race,qualities",
	}
	rs, err := client.DetectAndAnalysis(
		vision.MustFromFile("face.jpg"),
		options,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.ToString())
}
