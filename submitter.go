package main

import (
	client "github.com/astriaorg/go-sequencer-client/client"
)

type submitter struct {
	queuedTxs []string

	signer *client.Signer
	c      client.Client
}

func (s *submitter) startSubmitter() {
	signer, err := client.GenerateSigner()
	if err != nil {
		panic(err)
	}

	// default tendermint RPC endpoint
	c, err := client.NewClient("http://localhost:26657")
	if err != nil {
		panic(err)
	}
}

func submitCollectedTxs() {

}
