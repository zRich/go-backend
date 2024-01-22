package lab

import "testing"

func TestChain_Start(t *testing.T) {
	chain := &Chain{
		Port: 8545,
	}
	chain.Start()
}

func TestChain_Deploy(t *testing.T) {
	chain := &Chain{}
	chain.Start()
	chain.Deploy()
}
