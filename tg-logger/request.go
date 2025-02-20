package tg_logger

type MessageRequest struct {
	ServiceName string                 `json:"service_name"`
	TypeContent int                    `json:"type_content"`
	Content     map[string]interface{} `json:"content"`
}

func NewMessageRequest(serviceName string, typeContent int) *MessageRequest {
	return &MessageRequest{
		ServiceName: serviceName,
		TypeContent: typeContent,
		Content:     make(map[string]interface{}),
	}
}

func (r *MessageRequest) AddContent(key string, value interface{}) {
	r.Content[key] = value
}

func (r *MessageRequest) DeleteContent(key string) {
	delete(r.Content, key)
}

const (
	TYPE_FATAL = iota
	TYPE_ERROR
	TYPE_WARN
	TYPE_INFO

	TYPE_DEBUG
	TYPE_TRACE
)
