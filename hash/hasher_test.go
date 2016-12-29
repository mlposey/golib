package hash

import "testing"

func TestStringHasher_Hash(t *testing.T) {
	var hasher Hasher = NewStringHasher()

	abc := hasher.Hash("abc")
	bac := hasher.Hash("bac")

	if abc == 0 || bac == 0 {
		t.Error("0 as hash value in StringHasher")
	}

	if abc == bac {
		t.Error("collisions in StringHasher")
	}
}
