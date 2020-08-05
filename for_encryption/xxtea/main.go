package main

import (
	"fmt"
)
const TEA_KEY_CLIENT = "502ccc35a0e27a28"
const TEA_KEY_CLIENT_NEW_201911 = "s1001ddd35a0e27a28"
const TEA_KEY_ADMIN = "300ccc35a0e27a28"

func main() {

	array1 := make([]string, 0, 0)
	array2 := []string{}
	array1 = append(array1, []string{"1", "2", "3"}...)
	array2 = append(array2, []string{"1", "2", "3"}...)
	fmt.Printf("array1: %+v", array1)
	fmt.Printf("array2: %+v", array2)

	//getParam := url.Values{}
	//getParam.Set("userId", "54523")
	//getParam.Set("token", "xxxxxx")
	//data := getParam.Encode()
	//encrypted := xxtea.Encrypt([]byte(data), []byte(TEA_KEY_CLIENT_NEW_201911))
	//fmt.Println(encrypted)
	////dataEncrypted := `[16 170 231 35 161 49 195 230 77 103 218 158 239 152 109 11 20 164 159 7 179 78 159 3 21 30 120 66 241 86 106 213]`
	//dataByte := []byte {102,8,188,152,115,187,225,204,17,246,215,224,128,175,5,126,243,41,99,70,243,186,2,15,235,170,103,112,177,44,33,132,133,44,61,10,225,234,245,13,216,26,91,159,214,204,154,147,245,233,61,73,73,143,23,67,247,21,28,27,120,215,247,84,16,103,54,24,226,98,95,81,67,81,76,55,211,51,130,201,108,171,110,104}
	//decoded := xxtea.Decrypt([]byte(dataByte), []byte(TEA_KEY_CLIENT_NEW_201911))
	//
	//fmt.Printf("%s", decoded)
}