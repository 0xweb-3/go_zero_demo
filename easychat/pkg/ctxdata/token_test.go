package ctxdata

import (
	"testing"
	"time"
)

func Test_GetJwtToken(t *testing.T) {
	token, err := GetJwtToken("xin", time.Now().Unix(), 3600, "1")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
