package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	testFunc()
	buffer()
	reader()
}

func testFunc() {
	for i, v := range []struct {
		a []byte
		b []byte
	}{
		{[]byte{'a'}, []byte{'b'}},
		{[]byte{'a'}, []byte{'a'}},
		{[]byte{'b'}, []byte{'a'}},
		{[]byte{'a'}, []byte{'A'}},
		{[]byte{'a', 'b'}, []byte{'a', 'b'}},
	} {
		fmt.Printf("=====  case.%d  =====\n", i)

		// 返回 -1/0/1
		ri := bytes.Compare(v.a, v.b)
		log.Println(ri)

		// 返回 true/false
		rb := bytes.Contains(v.a, v.b)
		log.Println(rb)

		rb = bytes.ContainsAny(v.a, string(v.b))
		log.Println(rb)

		rb = bytes.ContainsRune(v.a, rune(v.b[0]))
		log.Println(rb)

		ri = bytes.Count(v.a, v.b)
		log.Println(ri)

		rb = bytes.Equal(v.a, v.b)
		log.Println(rb)

		rb = bytes.EqualFold(v.a, v.b) // 不区分大小写
		log.Println(rb)

		rbs := bytes.Fields(v.a) // 以空格为分隔符
		log.Println(rbs)

		rbs = bytes.FieldsFunc(v.a, func(r rune) bool {
			if r == 'a' { // 以 a 为分割符
				return true
			}
			return false
		})
		log.Println(rbs)

		rb = bytes.HasPrefix(v.a, v.b)
		log.Println(rb)

		rb = bytes.HasSuffix(v.a, v.b)
		log.Println(rb)

		ri = bytes.Index(v.a, v.b)
		log.Println(ri)

		ri = bytes.IndexAny(v.a, string(v.b[0]))
		log.Println(ri)

		ri = bytes.IndexByte(v.a, v.b[0])
		log.Println(ri)

		ri = bytes.IndexFunc(v.a, func(r rune) bool {
			if r == 'a' {
				return true
			}
			return false
		})
		log.Println(ri)

		ri = bytes.IndexRune(v.a, rune(v.b[0]))
		log.Println(ri)

		rbyte := bytes.Join([][]byte{v.a, v.b}, []byte(","))
		log.Println(rbyte)

		ri = bytes.LastIndex(v.a, v.b)
		log.Println(ri)

		ri = bytes.LastIndexAny(v.a, string(v.b[0]))
		log.Println(ri)

		ri = bytes.LastIndexByte(v.a, v.b[0])
		log.Println(ri)

		// 可以用于寻找最后一个非XX条件的字符
		ri = bytes.LastIndexFunc(v.a, func(r rune) bool {
			if r == 'a' {
				return true
			}
			return false
		})
		log.Println(ri)

		// 修改指定范围字符
		rbyte = bytes.Map(func(r rune) rune {
			if r == 'a' {
				return 'b'
			}
			return 'a'
		}, v.a)
		log.Println(rbyte)

		rbyte = bytes.Repeat(v.a, 2) // 0 则空；1 则相等；2 则翻倍；以此类推
		log.Println(rbyte)

		rbyte = bytes.Replace(v.a, v.b, v.a, -1)
		log.Println(rbyte)

		rr := bytes.Runes(v.a)
		log.Println(rr)

		rbs = bytes.Split(v.a, v.a[1:])
		log.Println(rbs)

		rbs = bytes.SplitAfter(v.a, v.a[1:])
		log.Println(rbs)

		rbs = bytes.SplitAfterN(v.a, v.a[1:], 2)
		log.Println(rbs)

		rbs = bytes.SplitN(v.a, v.a[1:], 2)
		log.Println(rbs)

		rbyte = bytes.Title(v.a)
		log.Println(rbyte)

		rbyte = bytes.ToLower(v.a)
		log.Println(rbyte)

		rbyte = bytes.ToLowerSpecial(unicode.AzeriCase, v.a)
		log.Println(rbyte)

		rbyte = bytes.ToTitle(v.a)
		log.Println(rbyte)

		rbyte = bytes.ToTitleSpecial(unicode.AzeriCase, v.a)
		log.Println(rbyte)

		rbyte = bytes.ToUpper(v.a)
		log.Println(rbyte)

		rbyte = bytes.ToUpperSpecial(unicode.AzeriCase, v.a)
		log.Println(rbyte)

		rbyte = bytes.Trim(v.a, string(v.a[0]))
		log.Println(rbyte)

		rbyte = bytes.TrimFunc(v.a, func(r rune) bool {
			if r == rune(v.a[0]) {
				return true
			}
			return false
		})
		log.Println(rbyte)

		rbyte = bytes.TrimLeft(v.a, string(v.a[0]))
		log.Println(rbyte)

		rbyte = bytes.TrimLeftFunc(v.a, func(r rune) bool {
			if r == rune(v.a[0]) {
				return true
			}
			return false
		})
		log.Println(rbyte)

		rbyte = bytes.TrimPrefix(v.a, v.a)
		log.Println(rbyte)

		rbyte = bytes.TrimRight(v.a, string(v.b[0]))
		log.Println(rbyte)

		rbyte = bytes.TrimRightFunc(v.a, func(r rune) bool {
			if r == rune(v.b[0]) {
				return true
			}
			return false
		})
		log.Println(rbyte)

		rbyte = bytes.TrimSpace(v.a)
		log.Println(rbyte)

		rbyte = bytes.TrimSuffix(v.a, v.b)
		log.Println(rbyte)
	}
}

func buffer() {
	var buf []byte
	bb := bytes.NewBuffer(buf)

	// 取值
	lpln := func() {
		log.Println(bb.Bytes(), bb.Len(), bb.Cap(), bb.String())
	}
	lpln()

	bb.Grow(10) // Cap 会变大，最小扩大 10，实际上是 64
	lpln()

	lpFa := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
		lpln()
	}
	_, err := bb.Write([]byte("hello"))
	lpFa(err)

	// buf = make([]byte, bb.Len())
	buf2 := make([]byte, bb.Len())
	n, err := bb.Read(buf2) // 读出来多少，就少掉多少
	lpFa(err)
	log.Println(buf2, n)

	err = bb.UnreadByte()
	// err = bb.UnreadRune() // 对应的是 ReadRune
	lpFa(err) // 如果在 Read 操作后，执行 UnreadByte ，这里会出现一个 o ，也就是上次读到的最后一个字符

	bb.Reset()

	// 从文件读取内容
	filename := "about.md"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, os.ModeAppend)
	lpFa(err)
	defer f.Close()
	_, err = bb.ReadFrom(f)
	lpFa(err)

	// 写到文件
	_, err = bb.WriteTo(f)
	lpFa(err)
}

func reader() {
	var buf = []byte{'a', 'b', 'c'}
	br := bytes.NewReader(buf)

	lpln := func() {
		log.Println(br.Len(), br.Size())
	}
	lpln()

	lpFa := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
		lpln()
	}

	var buf2 = make([]byte, br.Len())
	n, err := br.Read(buf2)
	lpFa(err)
	log.Println(buf2, n)
}
