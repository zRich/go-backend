package lab

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/db/models"
	"github.com/zRich/go-backend/internal/log"
	"github.com/zRich/go-backend/internal/mocha"
)

type Operator struct {
	DB    db.Database
	Chain *Chain
}

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
	cmd := exec.Command("npx", "hardhat", "node")
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

		logger.Info(fmt.Sprintf("chain process exit, pid: %d", chain_pid))
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
	cmd := exec.Command("npx", "hardhat", "test")
	cmd.Stdout = chain
	// out, err := cmd.Output()
	err := cmd.Run()
	if err != nil {
		logger.Error(fmt.Sprintf("deploy failed, output: %s", string(chain.output)))
	}

	logger.Info("deploy success")
}

// 将一个 mocha 测试报告存入数据库
func (o *Operator) SaveReport(studentNo string, report *mocha.Report) error {
	// 定义一个 TestReport 结构体
	var testReport models.TestReport
	// 将 report 转换为 json 格式字符串
	reportJson, err := json.Marshal(report)
	if err != nil {
		return err
	}
	db, err := o.DB.Connect()
	if err != nil {
		return err
	}

	// 给 testReport 赋值
	testReport.StudentNo = studentNo
	testReport.Report = string(reportJson)

	// 将 reportJson 存入数据库
	db.Create(&testReport)

	// (&testReport, models.TestReport{StudentNo: studentNo, Report: string(reportJson)})
	return nil
}

// 创建一个 Operator
func NewOperator(db db.Database, chain *Chain) *Operator {
	return &Operator{
		DB:    db,
		Chain: chain,
	}
}
