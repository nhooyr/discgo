package discgo

import (
	"context"
	"testing"
)

var gID = "331307058660114433"

func TestClient_CreateGuild(t *testing.T) {
	params := &ParamsGuildsCreate{
		Name: "REKTERONIED",
	}
	g, err := client.Guilds().Create(context.Background(),params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g.ID)
}

func TestClient_DeleteGuild(t *testing.T) {
	g, err := client.Guild(gID).Delete(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g.ID)
}

func TestClient_GetChannels(t *testing.T) {
	channels, err := client.Guild(gID).Channels().Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(channels[3].Name)
}

func TestClient_GetGuildMember(t *testing.T) {
	gm, err := client.Guild(gID).Member(uID).Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if gm.Nick != nil {
		t.Log(*gm.Nick)
	} else {
		t.Log(gm.Nick)
	}
}

func TestClient_GetGuildMembers(t *testing.T) {
	guildMembers, err := client.Guild(gID).Members().Get(context.Background(),nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(guildMembers[0].User.Username)
}

func TestClient_ModifyGuildMember(t *testing.T) {
	params := &ParamsGuildMemberModify{
		Nick: "fdkg",
	}
	err := client.Guild(gID).Member(uID).Modify(context.Background(),params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_ModifyMyNick(t *testing.T) {
	nick, err := client.Guild(gID).Me().ModifyNick(context.Background(),"xd RssEKT")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(nick)
}
