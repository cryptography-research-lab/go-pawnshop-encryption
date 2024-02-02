package data

import _ "embed"

// Dictionary 内置的汉字笔画映射字典
//
//go:embed dictionary.txt
var Dictionary string
