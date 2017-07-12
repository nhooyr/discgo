package discgo_test

import "testing"

var inviteCode = "NP9NQ8v"

func TestClient_GetInvite(t *testing.T) {
	inv, err := c.GetInvite(inviteCode)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(inv.Code)
}