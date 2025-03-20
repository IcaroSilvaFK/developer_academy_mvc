package utils

import "github.com/matthewhartstonge/argon2"

func MakeHash(plan string) string {
	argon := argon2.DefaultConfig()

	enc, err := argon.HashEncoded([]byte(plan))

	if err != nil {
		panic(err)
	}

	return string(enc)
}

func VerifyHash(plan, hash string) bool {
	ok, err := argon2.VerifyEncoded([]byte(plan), []byte(hash))

	if err != nil {
		return false
	}

	return ok
}
