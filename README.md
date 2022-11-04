# yaml-encoder

This is a yaml encoder which implements `yaml.Marshaler` for marshal with comments.

Inspired by [The doc's encoder of talos](https://github.com/siderolabs/talos/blob/main/pkg/machinery/config/encoder/encoder.go)

## example

```go
package main

import (
	"fmt"
	encoder "github.com/wetor/yaml-encoder"
)

type DBConfig struct {
	Username string `yaml:"username" comment:"this is the username of database"`
	Password string `yaml:"password" comment:"this is the password of database"`
	Host string `yaml:"host"`
	Port int `yaml:"port" comment_key:"PORT"`
}

func main(){
	config := DBConfig{
		Username: "root",
		Password: "xxxxxx",
		Host: "127.0.0.1",
		Port: 4444,
	}

	encoder := encoder.NewEncoder(config,
		encoder.WithComments(encoder.CommentsOnHead),
		encoder.WithCommentsMap(map[string]string{
			"host": "主机名",
			"PORT": "端口号",
		}),
	)
	content, _ := encoder.Encode()
	fmt.Printf("%s", content)
}
```

output:

```yaml
# this is the username of database
username: root
# this is the password of database
password: xxxxxx
# 主机名
host: 127.0.0.1
# 端口号
port: 4444
want:
# this is the username of database
username: root
# this is the password of database
password: xxxxxx
```
