package pawnshop_encryption

import (
	"fmt"
	"github.com/cryptography-research-lab/go-pawnshop-encryption/data"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Dictionary 笔画的映射字典
type Dictionary struct {

	// 加密使用的字典，记录笔画数到汉字的映射，是一个一对多的关系
	numberToChineseSetMap map[int]map[rune]struct{}

	// 解密使用的字典，记录函数到笔画数的映射，是一个一对一的关系
	chineseRuneToNumberMap map[rune]int
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		numberToChineseSetMap:  make(map[int]map[rune]struct{}),
		chineseRuneToNumberMap: make(map[rune]int),
	}
}

// Init 初始化字典，从原始字典中初始化加密解密需要的字典
func (x *Dictionary) Init(dictionary string) error {

	for _, line := range strings.Split(dictionary, "\n") {

		// 忽略空行
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 解析每一行，每行是一个汉字和它对应的笔画数，比如：
		// 鸵 10
		split := strings.Split(line, " ")
		if len(split) != 2 {
			return fmt.Errorf("init by dictionary error, line %s format error", line)
		}
		number, err := strconv.Atoi(split[1])
		if err != nil {
			return err
		}

		// 放入到加密字典
		chineseRune := ([]rune(split[0]))[0]
		x.chineseRuneToNumberMap[chineseRune] = number
		chineseRuneSet := x.numberToChineseSetMap[number]
		if chineseRuneSet == nil {
			chineseRuneSet = make(map[rune]struct{})
			x.numberToChineseSetMap[number] = chineseRuneSet
		}
		chineseRuneSet[chineseRune] = struct{}{}
	}
	return nil
}

// AddItem 把映射关系加到字典
func (x *Dictionary) AddItem(chineseRune rune, number int) {

	// 数字到中文的映射
	chineseSet := x.numberToChineseSetMap[number]
	if chineseSet == nil {
		chineseSet = make(map[rune]struct{})
		x.numberToChineseSetMap[number] = chineseSet
	}
	chineseSet[chineseRune] = struct{}{}

	// 中文到数字的映射
	x.chineseRuneToNumberMap[chineseRune] = number

}

// DeleteItem 从字典中删除某个字
func (x *Dictionary) DeleteItem(chineseRune rune) {

	// 删除数字到中文的映射
	for _, chineseSet := range x.numberToChineseSetMap {
		delete(chineseSet, chineseRune)
	}

	// 删除中文到数字的映射
	delete(x.chineseRuneToNumberMap, chineseRune)

}

func (x *Dictionary) GetChineseSetByNumber(number int) map[rune]struct{} {
	return x.numberToChineseSetMap[number]
}

func (x *Dictionary) GetNumberByChineseRune(chineseRune rune) int {
	return x.chineseRuneToNumberMap[chineseRune]
}

// Encrypt 对数字字符串进行当铺加密
func (x *Dictionary) Encrypt(numbers string) (string, error) {
	runeSlice := make([]rune, utf8.RuneCountInString(numbers))
	for index, numberRune := range numbers {
		// 要加密的明文的每一个字符都必须是位于区间[0, 9]的数字，否则无法加密
		if numberRune < '0' || numberRune > '9' {
			return "", ErrPlaintextUnavailable
		}
		number := x.numberRuneToInt(numberRune)
		chineseSet, exists := x.numberToChineseSetMap[number]
		if !exists {
			return "", fmt.Errorf("number %d do not have chinese ", number)
		}
		// TODO 尽量加密成通顺的句子
		for chineseRune := range chineseSet {
			runeSlice[index] = chineseRune
			break
		}
	}
	return string(runeSlice), nil
}

// Decrypt 对当铺密码加密的密文进行解密
func (x *Dictionary) Decrypt(chineseString string) (string, error) {
	chineseRunes := []rune(chineseString)
	numbers := make([]rune, len(chineseRunes))
	for index, chineseRune := range chineseRunes {
		number, exists := x.chineseRuneToNumberMap[chineseRune]
		if !exists {
			return "", fmt.Errorf("chinese %s do not have number", string(chineseString))
		}
		numbers[index] = x.numberIntToRune(number)
	}
	return string(numbers), nil
}

func (x *Dictionary) numberRuneToInt(number rune) int {
	return int(number-'0') + 1
}

func (x *Dictionary) numberIntToRune(number int) rune {
	return rune(number + '0' - 1)
}

// ---------------------------------------------------------------------------------------------------------------------

// DefaultDictionary 内置的字典，一般只使用这个字典就可以了
var DefaultDictionary *Dictionary

func init() {
	DefaultDictionary = NewDictionary()
	err := DefaultDictionary.Init(data.Dictionary)
	if err != nil {
		panic(err)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
