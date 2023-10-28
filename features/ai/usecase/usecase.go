package usecase

import (
	"context"
	"errors"
	"recycle/features/ai/entity"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type RubbishUseCase struct {
	userInput entity.UseCaseInterface
	openAIKey string
}

func NewAIUsecase(userInput entity.UseCaseInterface, openAIKey string) entity.UseCaseInterface {
	return &RubbishUseCase{
		userInput: userInput,
		openAIKey: openAIKey,
	}
}

// RecommendRecyclable implements entity.UseCaseInterface.
func (uc *RubbishUseCase) RecommendRecyclable(itemName string) (string, error) {
	chatSystem := "Saya adalah sebuah sistem yang dapat menjelaskan jenis sampah yang bisa didaur ulang dan memberikan penjelasan contoh-contoh sampahnya secara rinci berdasarkan inputan type nya. Kemudian saya dapat menjelaskan fungsi dan kegunaan daur ulang dari sampah bersadarkan type tersebut dan terakhir memberikan func fact terhadap jenis sampah yang diinputkan"
	itemName = strings.ToLower(itemName)
	validItems := []string{"plastik", "kertas", "logam", "kaca", "karton"}
	isValidItem := false
	for _, validItem := range validItems {
		if itemName == validItem {
			isValidItem = true
			break
		}
	}

	if !isValidItem {
		return "", errors.New("Jenis sampah yang dimasukkan tidak valid. Mohon pastikan Anda memberikan jenis sampah yang benar. Contoh: 'Kertas', 'Plastik', 'Logam', 'Kaca', atau 'Karton'.")
	}

	ctx := context.Background()
	client := openai.NewClient(uc.openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: chatSystem,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: itemName,
		},
	}

	request := entity.ChatRequest{
		Context:  ctx,
		Client:   client,
		Model:    model,
		Messages: messages,
	}

	resp, err := uc.GetCompletionFromMessages(request)
	if err != nil {
		return "", err
	}

	answer := resp.Choices[0].Message.Content
	return answer, nil
}

// GetCompletionFromMessages implements entity.UseCaseInterface.
func (uc *RubbishUseCase) GetCompletionFromMessages(request entity.ChatRequest) (openai.ChatCompletionResponse, error) {
	if request.Model == "" {
		request.Model = openai.GPT3Dot5Turbo
	}

	resp, err := request.Client.CreateChatCompletion(
		request.Context,
		openai.ChatCompletionRequest{
			Model:    request.Model,
			Messages: request.Messages,
		},
	)
	return resp, err
}
