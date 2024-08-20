package lambda_handler

type MessagePath struct {
	Path string `json:"path"`
	Body string `json:"body"`
}

const (
	PATH string = "test"
)
