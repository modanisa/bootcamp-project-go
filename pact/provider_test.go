package pact

import (
	"bootcamp/server"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	svr := server.NewServer()
	go svr.StartServer(port)

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "Backend",
		Consumer:                 "Frontend",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs: []string{
			"/Users/abdulsamet.ileri/Desktop/bootcamp/bootcamp-project-vue/pact/pacts/frontend-backend.json",
		},
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}
