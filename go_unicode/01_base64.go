package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var str string = "luffy"
	strByte := []byte(str)
	fmt.Println(strByte)
	// base64编码
	ss := base64.StdEncoding.EncodeToString(strByte)
	fmt.Println(ss)

	// base64解码
	ss64, _ := base64.StdEncoding.DecodeString(ss)
	fmt.Println(ss64)
	fmt.Println(string(ss64))
}

/*
108 117 102 102 121                                       //每个字符转为对应的Ascii编码表中的值.
01101100 01110101 01100110 01100110 01111001              //Ascii 值分别转为对应的8位二进制格式
011011 000111 010101 100110 011001 100111 1001            //重新分组，base32是6位一组。
011011 000111 010101 100110 011001 100111 100100 000000   //因为最后一组不满6位，所以补8个0，至于补多少0，下文会有说明
27 7 21 38 25 39 36 =                                     //每组都转换成十进制,最后一组因为是补的，而且都是0，使用=号代替
b H V m Z n k =                                           //对应base64编码表转换，得到最终的值为bHVmZnk=
*/

//测试base64的用法
/*
编码流程
（1）首先将字符串中的每个字符转为对应的Ascii编码表中的值.
（2）将第一步中的 Ascii 值分别转为对应的8位二进制格式（不足8比特位高位补0）
（3）根据不同编码规则，对二进制串重新分组，比如base16为4位一组，base32为5位一组，base64为6位一组.
（4）将每组二进制重新转换为十进制，然后去对应编码表再转换为对应的字符。
*/
