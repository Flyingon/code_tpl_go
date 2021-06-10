package cos

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPutObject 需要填充配置
func TestPutObject(t *testing.T) {
	cosConf := &Config{
		AppID:     "XXXX",
		Bucket:    "XXXX",
		Region:    "ap-shanghai",
		SecretID:  "XXXX",
		SecretKey: "XXXX",
		CDNHost:   "https://XXXX.file.myqcloud.com",
	}
	url, err := PutObject(cosConf, "avatar/test.go", "./cos.go")
	fmt.Printf("cos test res: %s, err: %v\n", url, err)
	assert.NoError(t, err, "")
	assert.Greater(t, len(url), 0)
}
