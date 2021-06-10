package cos

import (
	"context"
	"fmt"
	"github.com/siddontang/go-log/log"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Config cos 配置
type Config struct {
	AppID     string `json:"appid"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	CDNHost   string `json:"cdn_host"`
}

// PutObject 东西上传
func PutObject(cosConf *Config, cosPath, filePath string) (string, error) {
	cosUrl := fmt.Sprintf("https://%s.cos.%s.myqcloud.com",
		cosConf.Bucket, cosConf.Region)
	u, _ := url.Parse(cosUrl)
	b := &cos.BaseURL{
		BucketURL: u,
	}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosConf.SecretID,
			SecretKey: cosConf.SecretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
		Timeout: 5 * time.Second, // HTTP超时时间
	})

	res, err := c.Object.PutFromFile(context.Background(), cosPath, filePath, nil)
	if err != nil {
		log.Errorf(err.Error())
		return "", err
	}
	retUrl := path.Join(cosUrl, cosPath)
	log.Infof("cos res: %+v", res.StatusCode)
	if len(cosConf.CDNHost) > 0 {
		retUrl = path.Join(cosConf.CDNHost, cosPath)
	}
	return retUrl, nil
}
