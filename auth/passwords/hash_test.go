package passwords

import "testing"

func TestHash(t *testing.T) {
	testString := []byte("p@sSword#")
	testSecret := []byte("1q2w3e4r!")

	t.Run("hash_and_validate", func(t *testing.T) {
		hash := CreateHash(testString, testSecret)
		if !Validate(testString, hash, testSecret) {
			t.Errorf("Failed to validate password")
		}
	})
}
