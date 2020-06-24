package codec

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"reflect"

	"github.com/go-playground/form"
)

func init() {
	RegisterSerializer(SerializationTypeForm, &FormSerialization{})
}

// FormSerialization 打包http get请求的kv结构
type FormSerialization struct {
	encoder *form.Encoder
	decoder *form.Decoder
}

// NewFormSerialization 初始化from序列化对象
func NewFormSerialization(tag string) Serializer {
	encoder := form.NewEncoder()
	decoder := form.NewDecoder()
	encoder.SetTagName(tag)
	decoder.SetTagName(tag)
	return &FormSerialization{
		encoder: encoder,
		decoder: decoder,
	}
}

// Unmarshal 解包kv结构
func (j *FormSerialization) UnmarshalOld(in []byte, body interface{}) error {
	values, _ := url.ParseQuery(string(in))
	params := map[string]interface{}{}
	for k, v := range values {
		if len(v) == 1 {
			params[k] = v[0]
		} else {
			params[k] = v
		}
	}
	fmt.Printf("in: %s\n params: %+v\n", in, params)
	config := &mapstructure.DecoderConfig{TagName: "json", Result: body, WeaklyTypedInput: true, Metadata: nil}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(params)
}

// Unmarshal 解包kv结构
func (j *FormSerialization) Unmarshal(in []byte, body interface{}) error {

	values, err := url.ParseQuery(string(in))
	if err != nil {
		return err
	}
	fmt.Println("values:", values)
	fmt.Println("body type: ", reflect.TypeOf(body))
	decoder := form.NewDecoder()
	decoder.SetTagName("json")
	err = decoder.Decode(&body, values)
	return err
}

// Marshal 打包kv结构
func (j *FormSerialization) Marshal(body interface{}) ([]byte, error) {
	if req, ok := body.(url.Values); ok { // 用于向后端post发送form urlencode请求
		return []byte(req.Encode()), nil
	}

	jsonSerializer := GetSerializer(SerializationTypeJSON) // 用于收到Get请求给前端回json包
	if jsonSerializer == nil {
		return nil, errors.New("empty json serializer")
	}
	return jsonSerializer.Marshal(body)
}
