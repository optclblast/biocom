package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/gogo/protobuf/proto"
	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
)

func TestConnection(t *testing.T) {
	login := os.Getenv("T_OPENERP_LOGIN")
	password := os.Getenv("T_OPENERP_PASSWORD")
	org := os.Getenv("T_OPENERP_ORG")
	addr := os.Getenv("T_OPENERT_ADDR")

	login = "123"
	password = "123"
	org = "123"
	addr = "127.0.0.1:8080"

	c := NewTestClient(login, password, org, addr)

	r := &apiv1.Request{
		Id: 1,
	}

	data, err := proto.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	data, err = c.Do(data)
	if err != nil {
		t.Fatal(err)
	}

	var resp *apiv1.Response = new(apiv1.Response)

	err = proto.Unmarshal(data, resp)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)
}
