package lab

import (
	"fmt"
	"os/exec"

	"github.com/zRich/go-backend/internal/log"
)

type Chain struct {
	Host    string
	ChainId int
	Port    int
	output  []byte
}

// Write implements io.Writer.
func (chain *Chain) Write(p []byte) (n int, err error) {
	chain.output = p
	// logger.Info("chain output: %s", string(p), "")
	return len(chain.output), nil
}

var logger = log.InitLogger()

func (chain *Chain) Start() {
	// 执行 yarn chain 命令
	cmd := exec.Command("yarn", "chain")
	cmd.Stdout = chain
	// out, err := cmd.Output()
	err := cmd.Start()

	if err != nil {
		log.Log.Error("start chain failed")
	}

	chain_pid := cmd.Process.Pid

	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Log.Error("chain process exit")
		}

		//print pid
		fmt.Printf("chain pid: %d\n", chain_pid)
	}()
}

func (chain *Chain) Stop() {
	// 根据端口号杀掉进程
	// kill -9 $(lsof -i:8545 -t)
	cmd := exec.Command("lsof", "-t", "-i", fmt.Sprintf("tcp:%d", chain.Port))
	out, err := cmd.Output()
	if err != nil {
		log.Log.Error("lsof chain failed")
	}
	//去除换行符
	pid := string(out[:len(out)-1])
	// pid := string(out)

	cmd = exec.Command("kill", "-9", pid)
	err = cmd.Run()

	if err != nil {
		log.Log.Error("kill chain failed")
	}
}

func (chain *Chain) Deploy() {
	// 执行 yarn deploy 命令
	cmd := exec.Command("yarn", "deploy")
	cmd.Stdout = chain
	// out, err := cmd.Output()
	err := cmd.Run()
	if err != nil {
		log.Log.Error("deploy failed")
	}

	fmt.Printf("yarn deploy output\n%s\n", chain.output)
}
