package uid

import "testing"

func TestUUID(t *testing.T) {
	uid := GenerateUUID()
	t.Log(uid)
	println(uid)
}
