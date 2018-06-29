package slack

import (
	"testing"
)

func TestEvent(t *testing.T) {
	sc := setup(t)
	userID := "U11A2B8C1"
	userName := "testorizor1"
	selfID := "U11D00T0"
	channelName := "slirctest"
	channelID := "C11JBA78E"
	msg := "foo bar bäz!"
	// contains no IDs
	se1 := &Event{Username: userName, Channelname: channelName, Text: msg}
	// contains no name
	se2 := &Event{UserID: userID, ChannelID: channelID, Text: msg}
	// contains selfMSG
	se3 := &Event{UserID: selfID, ChannelID: channelID, Text: msg}

	// looksup only ChannelID
	sc.nameToID(se1)
	// looksup channel and username
	sc.idToName(se2)

	if se1.ChannelID != channelID {
		t.Logf("nameToID failed - expected: (%v) - got: (%v)", channelID, se1.ChannelID)
		t.Fail()
	}

	if se2.Chan() != channelName || se2.Usernick() != userName {
		t.Logf("idToName failed - expected: (%v) (%v) - got: (%v) (%v)", channelName, userName, se2.Chan(), se2.Usernick())
		t.Fail()
	}

	if se1.Chan() != channelName {
		t.Logf("Channel mismatch expected: (%v) - got: (%v)", channelName, se1.Chan())
		t.Fail()
	}

	if se1.Usernick() != userName {
		t.Logf("User mismatch expected: (%v) - got: (%v)", userName, se1.Usernick())
		t.Fail()
	}

	if se1.Msg() != msg {
		t.Logf("Msg mismatch expected: (%v) - got: (%v)", msg, se1.Msg())
		t.Fail()
	}

	if !sc.IsSelfMsg(se3) {
		t.Log("IsSelfMsg produced wrong result")
		t.Fail()
	}

}
