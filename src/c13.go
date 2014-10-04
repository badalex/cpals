package main

import (
	"bytes"
	"cpals"
	"fmt"
	"strings"
)

type user struct {
	email string
	uid   int
	role  string
}

func (u user) encode() string {
	return fmt.Sprintf("email=%s&uid=%d&role=%s", u.email, u.uid, u.role)
}

func profile_for(email string) user {
	u := user{}

	e := strings.Replace(strings.Replace(email, "&", "", -1), "=", "", -1)

	u.email = e
	u.uid = 10
	u.role = "user"

	return u
}

func main() {
	key := cpals.Fill(16, "A")

	// so the idea here is the server encrypts who you are and your role and stuffs it into a cookie
	// so we want to modify the encrypted data in the cookie so that we are admin
	// email has to be the right size so "user" of role=user ends up on its own block
	u := profile_for("aaaaaaa@b.com")
	u_enc := cpals.ECBEncrypt(key, []byte(u.encode()))

	// now we just need to make it so admin + padding has its own block
	// ok, so if we pad this out to a block size, then
	admin := profile_for("aaaa@b.com" + string(cpals.Pad7(16, []byte("admin"))))
	admin_enc := cpals.ECBEncrypt(key, []byte(admin.encode()))

	//cpals.PrintEnc(u.encode(), u_enc)
	//cpals.PrintEnc(admin.encode(), admin_enc)

	// then we copy the admin block + padding from admin_enc over the u_enc "user" block
	attack := u_enc
	copy(attack[32:48], admin_enc[16:32])

	//cpals.PrintEnc(u.encode(), attack)

	//bingo
	out := cpals.ECBDecrypt(key, attack)

	expected := cpals.Pad7(16, []byte("email=aaaaaaa@b.com&uid=10&role=admin"))
	if !bytes.Equal(out, expected) {
		fmt.Printf("%s\n%s %x\n", u.encode(), out, out)
		panic("failed")
	}
}
