package tester

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/pk910/blob-sender/txbuilder"
)

func (tester *Tester) PrepareClients() error {
	tester.clients = make([]*txbuilder.Client, 0)
	wg := &sync.WaitGroup{}
	mtx := sync.Mutex{}

	var chainId *big.Int
	for _, rpcHost := range tester.config.RpcHosts {
		wg.Add(1)

		go func(rpcHost string) {
			defer wg.Done()

			client, err := txbuilder.NewClient(rpcHost)
			if err != nil {
				tester.logger.Errorf("failed creating client for '%v': %v", rpcHost, err.Error())
				return
			}
			client.Timeout = 5 * time.Second
			cliChainId, err := client.GetChainId()
			if err != nil {
				tester.logger.Errorf("failed getting chainid from '%v': %v", rpcHost, err.Error())
				return
			}
			if chainId == nil {
				chainId = cliChainId
			} else if cliChainId.Cmp(chainId) != 0 {
				tester.logger.Errorf("chainid missmatch from %v (chain ids: %v, %v)", rpcHost, cliChainId, chainId)
				return
			}
			client.Timeout = 30 * time.Second
			mtx.Lock()
			tester.clients = append(tester.clients, client)
			mtx.Unlock()
		}(rpcHost)
	}

	wg.Wait()
	tester.chainId = chainId
	if len(tester.clients) == 0 {
		return fmt.Errorf("no useable clients")
	}
	return nil
}
