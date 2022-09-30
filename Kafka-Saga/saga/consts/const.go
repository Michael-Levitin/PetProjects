package consts

var (
	Brokers = []string{"localhost:9095", "localhost:9096"}
)

type Order struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}
