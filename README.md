# 当铺密码（Pawnshop Encryption）

# 一、当铺密码概述

当铺密码是用来把数字加密为中文的，加密的时候把数字映射到对应笔画的汉字，解密时按照汉字的笔画数映射回数字。

适用场景，比如用于，比如：

- 用来加密手机号
- QQ号
- 身份证号码等等

两种形式的当铺密码：

- 按笔画数（隐蔽性较好）
- 按笔画出头数（比较简单）

# 二、安装

```bash
go get -u github.com/cryptography-research-lab/go-pawnshop-encryption
```

# 三、API代码示例

```go
package main

import (
	"fmt"
	pawnshop_encryption "github.com/cryptography-research-lab/go-pawnshop-encryption"
)

func main() {

	// 对笔者的QQ号进行加密
	plaintext := "1451546085"
	encrypt, err := pawnshop_encryption.Encrypt(plaintext)
	if err != nil {
		fmt.Println("加密失败： " + err.Error())
		return
	}
	fmt.Println("加密结果： " + encrypt) // Output: 加密结果： 力训朴八向汀钉一姣亦

	decrypt, err := pawnshop_encryption.Decrypt(encrypt)
	if err != nil {
		fmt.Println("解密失败： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decrypt) // Output: 解密结果： 1451546085

}
```


# 四、参考资料
- https://www.cnblogs.com/cc11001100/p/9357263.html

# 五、TODO

- 提供笔画出头数加密方式的选择 
