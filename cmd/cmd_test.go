package main

import (
	"context"
	"testing"
	"time"

	"github.com/g1eng/savac/pkg/vps"
	fakevps "github.com/g1eng/savac/testutil/fake_vps"
	"github.com/g1eng/savac/testutil/test_parameters"
)

func TestGenerate(t *testing.T) {
	g := Generate(&vps.SavaClient{})
	if g == nil {
		t.Errorf("Generate() returned nil")
	}
}

func TestServerListWithGlobalFlag(t *testing.T) {
	go func() {
		err := fakevps.StartFakeServer(test_parameters.FakeServerEndpoint["cmd0"])
		if err != nil {
			panic("failed to start fake server: %v" + err.Error())
		}
	}()
	g := Generate(vps.NewTestClient(test_parameters.FakeServerEndpoint["cmd0"]))
	if g == nil {
		t.Fatalf("Generate() returned nil")
	}
	time.Sleep(1 * time.Second)
	err := g.Run(context.Background(), []string{"--debug", "--json", "server", "list"})
	if err != nil {
		t.Errorf("Run() 1 cannot show the list of servers: returned %v", err)
	}
	err = g.Run(context.Background(), []string{"--yaml", "server", "list"})
	if err != nil {
		t.Errorf("Run() 2 cannot show the list of servers: returned %v", err)
	}
	err = g.Run(context.Background(), []string{"--output-format=table", "server", "list"})
	if err != nil {
		t.Errorf("Run() 3 cannot show the list of servers: returned %v", err)
	}
	err = g.Run(context.Background(), []string{"--output-format=json", "apikey", "list"})
	if err != nil {
		t.Errorf("Run() 3 cannot show the list of servers: returned %v", err)
	}
	err = g.Run(context.Background(), []string{"--output-format=json", "perm", "list"})
	if err != nil {
		t.Errorf("Run() 3 cannot show the list of servers: returned %v", err)
	}
}
