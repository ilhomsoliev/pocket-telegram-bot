package storage

type Bucket string

const (
	AccessTokens  Bucket = "access_tokens"
	RequestTokens Bucket = "request_tokens"
)

type TokenStorage interface {
	Save(chatId int64, text string, bucket Bucket) error
	Get(chatId int64, bucket Bucket) (string, error)
}
