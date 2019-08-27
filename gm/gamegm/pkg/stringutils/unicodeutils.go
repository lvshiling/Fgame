package stringutils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
)

func StringToUnicode(p_string string) string {
	rst := ""
	for _, r := range p_string {
		rint := int(r)
		sig := strconv.FormatInt(int64(rint), 16)
		if len(sig) < 4 {
			sig = "0000" + sig
			sig = sig[len(sig)-4:]
		}
		rst += "\\u" + sig // json
	}
	return rst
}

// func UnicodeToString(from string) string {
// 	sUnicodev := strings.Split(from, "\\u")
// 	var context string
// 	for _, v := range sUnicodev {
// 		if len(v) < 1 {
// 			continue
// 		}
// 		temp, err := strconv.ParseInt(v, 16, 32)
// 		if err != nil {
// 			panic(err)
// 		}

// 		context += fmt.Sprintf("%c", temp)
// 	}
// 	return context
// }

func UnicodeToString(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return

}
