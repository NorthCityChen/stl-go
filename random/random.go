/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-04 16:11:37
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package random

import (
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStr(strLen int) string {
	str := strings.Builder{}
	str.Grow(strLen)
	for i, cache, remain := strLen-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			str.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return str.String()
}
