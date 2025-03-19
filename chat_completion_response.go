package openroutergo

import (
	"encoding/json"

	"github.com/orsinium-labs/enum"
)

// chatCompletionFinishReason is an enum for the reason the model stopped generating tokens.
//
//   - https://openrouter.ai/docs/api-reference/overview#finish-reason
type chatCompletionFinishReason enum.Member[string]

// MarshalJSON implements the json.Marshaler interface for chatCompletionFinishReason.
func (cfr chatCompletionFinishReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(cfr.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for chatCompletionFinishReason.
func (cfr *chatCompletionFinishReason) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	*cfr = chatCompletionFinishReason{Value: value}
	return nil
}

var (
	// FinishReasonStop is when the model hit a natural stop point or a provided stop sequence.
	FinishReasonStop = chatCompletionFinishReason{"stop"}
	// FinishReasonLength is when the maximum number of tokens specified in the request was reached.
	FinishReasonLength = chatCompletionFinishReason{"length"}
	// FinishReasonContentFilter is when content was omitted due to a flag from our content filters.
	FinishReasonContentFilter = chatCompletionFinishReason{"content_filter"}
	// FinishReasonToolCalls is when the model called a tool.
	FinishReasonToolCalls = chatCompletionFinishReason{"tool_calls"}
	// FinishReasonError is when the model returned an error.
	FinishReasonError = chatCompletionFinishReason{"error"}
)

// ChatCompletionResponse is the response from the OpenRouter API for a chat completion request.
//
//   - https://openrouter.ai/docs/api-reference/overview#responses
//   - https://platform.openai.com/docs/api-reference/chat/object
type ChatCompletionResponse struct {
	// A unique identifier for the chat completion.
	ID string `json:"id"`
	// A list of chat completion choices (the responses from the model).
	Choices []ChatCompletionResponseChoice `json:"choices"`
	// Usage statistics for the completion request.
	Usage ChatCompletionResponseUsage `json:"usage"`
	// The Unix timestamp (in seconds) of when the chat completion was created.
	Created int `json:"created"`
	// The model used for the chat completion.
	Model string `json:"model"`
	// The object type, which is always "chat.completion"
	Object string `json:"object"`
}

type ChatCompletionResponseChoice struct {
	// The reason the model stopped generating tokens. This will be `stop` if the model hit a
	// natural stop point or a provided stop sequence, `length` if the maximum number of
	// tokens specified in the request was reached, `content_filter` if content was omitted
	// due to a flag from our content filters, `tool_calls` if the model called a tool, or
	// `error` if the model returned an error.
	FinishReason chatCompletionFinishReason `json:"finish_reason"`
	// A chat completion message generated by the model.
	Message ChatCompletionMessage `json:"message"`
}

type ChatCompletionResponseUsage struct {
	// The number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens"`
	// The number of tokens in the generated completion.
	CompletionTokens int `json:"completion_tokens"`
	// The total number of tokens used in the request (prompt + completion).
	TotalTokens int `json:"total_tokens"`
}
