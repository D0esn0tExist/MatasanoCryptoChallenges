package set2

import (
	"fmt"
	"strings"
)

// Profile type is a map of cookie attrs and values.
type Profile map[string]string

// Parser function is a k=v parsing routine.
// The routine should take:
// foo=bar&baz=qux&zap=zazzle
// ... and produce:
// {
//   foo: 'bar',
//   baz: 'qux',
//   zap: 'zazzle'
// }
func Parser(cookie string) Profile {
	fmt.Println("cook:", cookie)
	attr := strings.Split(cookie, "&")
	attrMap := make(map[string]string, len(attr))
	for i := 0; i < len(attr); i++ {
		kv := strings.Split(attr[i], "=")
		fmt.Println(kv)
		if _, ok := attrMap[kv[0]]; ok {
			continue
		}
		attrMap[kv[0]] = kv[1]
	}
	fmt.Println(attrMap)
	return attrMap

}

// ProfileFor function encodes and encrypts a user profile in that format, given an email address.
func ProfileFor(email string) string {
	// sanitize
	cleanedEmail := strings.Replace(strings.Replace(email, "&", "", -1), "=", "", -1)
	p := make(Profile)
	p["email"] = cleanedEmail
	p["uid"] = "10"
	p["role"] = "user"
	return p.CookieEncoder()
}

// CookieEncoder function encodes a Profile
func (p *Profile) CookieEncoder() string {
	var cookie []string

	handler := make(Profile)
	for k, v := range *p {
		handler[k] = v
	}

	assertOrder := func(order string) {
		if v, ok := handler[order]; ok {
			cookie = append(cookie, order+"="+v)
			delete(handler, order)
		}
	}

	assertOrder("email")
	assertOrder("uid")
	assertOrder("role")
	profileCookie := strings.Join(cookie, "&")

	fmt.Println(profileCookie)
	return profileCookie
}

// PrivEsc function performs a priviledge escalation hack to get back a profile with admin role.
func PrivEsc(cookieEncrypter func(string) []byte) []byte {
	begin := "email="

	createBlock := func(prefix string) string {
		msg := strings.Repeat("A", 16-len(begin)) + prefix
		return string(cookieEncrypter(msg))[16:32]
	}

	emailBlock := string(cookieEncrypter("me@foo.bar"))[:16]                       // email=me@foo.bar
	padEmailBlock := createBlock(strings.Repeat("A", 16-len("&uid=10&role=")))     // AAA&uid=10&role=
	adminBlock := createBlock("admin")                                             // admin&uid=10&rol
	trailingBlock := string(cookieEncrypter(strings.Repeat("A", 16-len(begin)-1))) // email=AAAAAAAAA&uid=10&role=user

	synthCookie := emailBlock + padEmailBlock + adminBlock + trailingBlock
	return []byte(synthCookie)
}
