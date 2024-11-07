package rocket

import "fmt"

const (
	ENVIRONMENT string = "ENV"
	VERSION     string = "VER"
)

type Message struct {
	Header map[string]string
	Body   string
}

func (message *Message) ToString() string {
	return fmt.Sprintf("env:%s,ver:%s,msg:%s", message.Header["env"], message.Header["ver"], message.Body)
}
