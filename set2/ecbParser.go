package set2

import (
	"encoding/json"
	"fmt"
	"strings"
)

type profile map[string]string

// Parser function is a k=v parsing routine.
// The routine should take:
// foo=bar&baz=qux&zap=zazzle
// ... and produce:
// {
//   foo: 'bar',
//   baz: 'qux',
//   zap: 'zazzle'
// }
func Parser(cookie string) string {
	attr := strings.Split(cookie, "&")
	fmt.Println(attr)
	attrMap := make(map[string]string, len(attr))
	for i := 0; i < len(attr); i++ {
		kv := strings.Split(attr[i], "=")
		fmt.Println(kv)
		attrMap[kv[0]] = kv[1]
	}
	attrJSON, _ := json.Marshal(attrMap)
	return string(attrJSON)

}

// ProfileFor function encodes a user profile in that format, given an email address.
func ProfileFor(email string) string {
	cleanedEmail := strings.Replace(strings.Replace(email, "&", "", -1), "=", "", -1)
	p := make(profile)
	p["email"] = cleanedEmail
	p["uid"] = "10"
	p["role"] = "user"

	return p.CookieEncoder()
}

func (p *profile) CookieEncoder() string {
	var cookie []string
	// handler := make(profile)
	for k, v := range *p {
		// handler[k] = v
		cookie = append(cookie, k+"="+v)
	}
	profileCookie := strings.Join(cookie, "&")
	fmt.Println(profileCookie)
	return profileCookie
}
