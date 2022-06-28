package encodeanddecodetinyurlmid

import (
	"math/rand"
	"testing"
)

type Codec struct {
	Pre      string
	DataBase map[string]string
}

func Constructor() Codec {
	return Codec{
		Pre:      "http://tinyurl.com/",
		DataBase: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (p *Codec) encode(longUrl string) string {
	k := 6
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	short := make([]byte, k)
	for {
		for i := 0; i < k; i++ {
			short[i] = str[rand.Intn(len(str))]
		}
		if _, has := p.DataBase[string(short)]; has {
			continue
		} else {
			break
		}
	}
	p.DataBase[string(short)] = longUrl
	return p.Pre + string(short)
}

// Decodes a shortened URL to its original URL.
func (p *Codec) decode(shortUrl string) string {
	shortUrl = shortUrl[len(p.Pre):]
	return p.DataBase[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
func Test_encode_and_decode_tinyurl(t *testing.T) {
	urls := make([]string, 10)
	urls[0] = "https://leetcode.com/problems/design-tinyurl"
	urls[1] = "https://leetcode.cn/problems/encode-and-decode-tinyurl"
	urls[2] = "https://bbs.nga.cn/"
	urls[3] = "http://wow.178.com/"
	urls[4] = "https://www.gcores.com/"
	urls[5] = "https://www.bilibili.com/"
	urls[6] = "https://zh.wikipedia.org/"
	urls[7] = "http://www.google.com/"
	urls[8] = "https://time.geekbang.org/"
	urls[9] = "https://www.runoob.com/"
	obj := Constructor()
	for _, url := range urls {
		obj.encode(url)
	}
	for k, v := range obj.DataBase {
		t.Log(k, v)
	}
}
