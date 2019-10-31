package idmanager

import "testing"

func Test_aa(t *testing.T)  {
	id := CreateMessageIdManager(7,12);

	idd := id.CreateId();

	getIDDetal(idd)
}