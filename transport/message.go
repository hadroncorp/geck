package transport

type Data struct {
	Data any `json:"data"`
}

type ErrorDetail struct {
	Type     string            `json:"@type"`
	Reason   string            `json:"reason"`
	Metadata map[string]string `json:"metadata"`
}

type Error struct {
	Code     int           `json:"code"`
	Message  string        `json:"message"`
	Status   string        `json:"status"`
	Details  []ErrorDetail `json:"details"`
	Internal error         `json:"-"`
}

type Errors struct {
	Code   int     `json:"code"`
	Errors []Error `json:"errors"`
}
