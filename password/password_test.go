package password

import "testing"

func TestHashAndCompare(t *testing.T) {
	rawPassword := "secret123"

	hash, err := Hash(rawPassword)
	if err != nil {
		t.Fatal(err)
	}

	err = Compare(hash, rawPassword)

	if err != nil {
		t.Fatal(err)
	}
}
