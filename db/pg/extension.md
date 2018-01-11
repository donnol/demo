# 插件

1 pt_trgm

    A trigram is a group of three consecutive characters taken from a string. We can measure the similarity of two strings by counting the number of trigrams they share. This simple idea turns out to be very effective for measuring the similarity of words in many natural languages.
    -- 一个 trigram 是从字符串中取到的一组三个连续字符。我们通过计算两个字符串共享的 trigram 数量来度量它们的相似度。这个简单的想法在许多自然语言度量单词相似度方面非常有效。

```go
func similarity(leftStr , rightStr string) int {
    const charLen = 3

    if len(leftStr) == 0 ||
        len(rightStr) == 0 {
        return 0
    }
    if leftStr == rightStr {
        return 1
    }

    // 分别取出左右字符串的三字符组--所有不大于3字符的、连续的字符组合
    trigramList := func(str string) []string {
        trigrams := []string{}

        strRune := []rune(Str)
        strRuneLen := len(strRune)
        if strRuneLen <= charLen {
            for i := range strRune {
                trigram := []rune{}
                if i == 0 {
                    trigram = append(trigram, strRune[0])
                }else if i == 1 {
                    trigram = append(trigram, strRune[0], strRune[1])
                }else if i == 2 {
                    trigram = append(trigram, strRune[0], strRune[1], strRune[2])
                }
                trigrams = append(trigrams, string(trigram))
            }
        }else{
            for i := range strRune {
                if i > strRuneLen - charLen {
                    break
                }
                trigram := []rune{}
                trigram = append(trigram, strRune[i], strRune[i+1], strRune[i+2])
                leftTrigrams = append(leftTrigrams, string(trigram))
            }
        }
        return trigrams
    }
    leftTrigrams := trigramList(leftStr)
    rightTrigrams := trigramList(rightStr)

    // 比较三字符组，统计相同的数量
    simCnt := 0
    for _, l := range leftTrigrams {
        for _, r := range rightTrigrams {
            if l == r {
                simCnt ++
            }
        }
    }

    // 相同数量占总数量的比例
    maxTrigram := max(len(leftTrigrams), len(rightTrigrams))

    return simCnt / maxTrigram
}
```
