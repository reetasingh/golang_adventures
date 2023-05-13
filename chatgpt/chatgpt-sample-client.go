package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func chat() {
	data, err := os.ReadFile("api.key")
	if err != nil {
		panic(err)
	}

	client := openai.NewClient(string(data))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "rephrase in context of badminton - Smash receiving technique Your left leg is in front Your hand up , like in lift position",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func imageGeneration() {
	data, err := os.ReadFile("/Users/reeta/secrets/chatgpt-api.key")
	client := openai.NewClient(string(data))
	// reqUrl := openai.ImageRequest{
	// 	Prompt:         "Golang gopher cartoon for medium article sitting on a laptop and writing unit tests in english language",
	// 	Size:           openai.CreateImageSize256x256,
	// 	ResponseFormat: openai.CreateImageResponseFormatURL,
	// 	N:              1,
	// }

	// respUrl, err := client.CreateImage(context.Background(), reqUrl)
	// if err != nil {
	// 	fmt.Printf("Image creation error: %v\n", err)
	// 	return
	// }
	//fmt.Println(respUrl.Data[0].URL)
	imageFile, err := os.Open("golang.png")
	if err != nil {
		panic(err)
	}
	reqUrl := openai.ImageEditRequest{
		Image:          imageFile,
		Prompt:         "Golang gopher cartoon dancing in rain",
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}
	respUrl, err := client.CreateEditImage(context.Background(), reqUrl)
	if err != nil {
		// failing as need to buy credit for this
		panic(err)
	}
	fmt.Println(respUrl.Data[0].URL)
}

func main() {
	//chat()
	imageGeneration()
}
