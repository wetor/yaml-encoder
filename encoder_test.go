package yaml_encoder_test

import (
	"fmt"
	"testing"

	encoder "github.com/wetor/yaml-encoder"
)

func TestEncoder(t *testing.T) {
	type DBConfig struct {
		Username string `yaml:"username" comment:"this is the username of database"`
		Password string `yaml:"password" comment:"this is the password of database"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port" comment_key:"PORT"`
	}

	config := DBConfig{
		Username: "root",
		Password: "xxxxxx",
		Host:     "127.0.0.1",
		Port:     4444,
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
	// Output:
	// # this is the username of database
	// username: root
	// # this is the password of database
	// password: xxxxxx
	// # 主机名
	// host: 127.0.0.1
	// # 端口号
	// port: 4444
}
