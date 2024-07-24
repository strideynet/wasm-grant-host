package types

type Request struct {
	Target []string `json:"target"`
}

type Response struct {
	Allow bool `json:"allow"`
}
