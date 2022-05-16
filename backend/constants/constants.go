package constants

const (
    USER_INDEX = "user"
    POST_INDEX = "post"
    ES_URL = "http://35.192.173.21:9200"
    ES_USERNAME = "bettsongy"
    ES_PASSWORD = "671115Sjm"
    GCS_BUCKET = "kevinsong_bucket"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Age      int64  `json:"age"`
    Gender   string `json:"gender"`
}