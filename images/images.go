package images

type RequestBodyImage struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
}

type RequestBodyImageEdit struct {
	Image          string `json:"image"`
	Mask           string `json:"mask"`
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
}

type RequestBodyImageVariation struct {
	Image          string `json:"image"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
}

type ResponseBodyImage struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

func NewRequestBodyImage() RequestBodyImage {
	return RequestBodyImage{
		Prompt:         "A cute baby sea otter",
		N:              2,
		Size:           "1024x1024",
		ResponseFormat: "url",
		User:           "test",
	}
}

func NewRequestBodyImageEdit() RequestBodyImageEdit {
	return RequestBodyImageEdit{
		Image:          "@otter.png",
		Mask:           "@mask.png",
		Prompt:         "A cute baby sea otter",
		N:              2,
		Size:           "1024x1024",
		ResponseFormat: "url",
		User:           "test",
	}
}

func NewRequestBodyImageVriation() RequestBodyImageVariation {
	return RequestBodyImageVariation{
		Image:          "@otter.png",
		N:              2,
		Size:           "1024x1024",
		ResponseFormat: "url",
		User:           "test",
	}
}
