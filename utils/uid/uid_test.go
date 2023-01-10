package uid

import "testing"

func TestUID(t *testing.T) {
	uid := GenerateUID()
	t.Log(uid)
	println(uid)
}
