package devframework

import (
	"os"
	"time"

	"github.com/0xkumi/incongito-dev-framework/account"
	"github.com/0xkumi/incongito-dev-framework/rpcclient"
)

func NewRPCClient(endpoint string) *rpcclient.RPCClient {
	return rpcclient.NewRPCClient(&RemoteRPCClient{endpoint: endpoint})
}

func NewStandaloneSimulation(name string, config Config, disableLog bool) *SimulationEngine {
	os.RemoveAll(name)
	sim := &SimulationEngine{
		config:            config,
		simName:           name,
		timer:             NewTimeEngine(),
		accountSeed:       "master_account",
		accountGenHistory: make(map[int]int),
		committeeAccount:  make(map[int][]account.Account),
		listennerRegister: make(map[int][]func(msg interface{})),
	}
	sim.DisableChainLog(disableLog)
	sim.init()
	time.Sleep(1 * time.Second)
	return sim
}
