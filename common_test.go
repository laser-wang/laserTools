package laserTools

import (
	"testing"
)

func Test_GetToken_1(t *testing.T) {

	token := GetToken()
	t.Log("Test_GetToken_1测试通过。" + token)

}
