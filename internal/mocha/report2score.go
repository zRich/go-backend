package mocha

import (
	"encoding/json"
	"os"

	"github.com/zRich/go-backend/internal/log"
)

// 定义报表的统计信息
// "stats": {
//     "suites": 6,
//     "tests": 9,
//     "passes": 8,
//     "pending": 0,
//     "failures": 1,
//     "start": "2024-01-29T14:33:44.034Z",
//     "end": "2024-01-29T14:33:45.612Z",
//     "duration": 1578
// 	  },

type Stats struct {
	Suites   int    `json:"suites"`
	Tests    int    `json:"tests"`
	Passes   int    `json:"passes"`
	Pending  int    `json:"pending"`
	Failures int    `json:"failures"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Duration int    `json:"duration"`
}

// 定义测试用例的错误信息，下面是一个例子
// "err": {
// 	"message": "expected 1738074823 to equal 1738074824. The numerical values of the given \"bigint\" and \"number\" inputs were compared, and they differed.",
// 	"showDiff": true,
// 	"actual": "1738074823",
// 	"expected": "1738074824",
// 	"operator": "strictEqual",
// 	"stack": "AssertionError: expected 1738074823 to equal 1738074824. The numerical values of the given \"bigint\" and \"number\" inputs were compared, and they differed.\n    at Context.<anonymous> (test/Lock.ts:33:42)\n    at processTicksAndRejections (node:internal/process/task_queues:95:5)"
//   }

type TestErr struct {
	Message  string `json:"message"`
	ShowDiff bool   `json:"showDiff"`
	Actual   string `json:"actual"`
	Expected string `json:"expected"`
	Operator string `json:"operator"`
	Stack    string `json:"stack"`
}

// 定义测试用例， 一个 json 文件包含多个测试用例， 下面是一个测试用例的例子

// {
// 	"title": "Should set the right unlockTime",
// 	"fullTitle": "Lock Deployment Should set the right unlockTime",
// 	"file": "/Users/richzhao/repos/zRich/blockchain-lab/verify-system/test/Lock.ts",
// 	"duration": 1461,
// 	"currentRetry": 0,
// 	"err": {
// 	  "message": "expected 1738074823 to equal 1738074824. The numerical values of the given \"bigint\" and \"number\" inputs were compared, and they differed.",
// 	  "showDiff": true,
// 	  "actual": "1738074823",
// 	  "expected": "1738074824",
// 	  "operator": "strictEqual",
// 	  "stack": "AssertionError: expected 1738074823 to equal 1738074824. The numerical values of the given \"bigint\" and \"number\" inputs were compared, and they differed.\n    at Context.<anonymous> (test/Lock.ts:33:42)\n    at processTicksAndRejections (node:internal/process/task_queues:95:5)"
// 	}
//   },

type TestCase struct {
	Title        string  `json:"title"`
	FullTitle    string  `json:"fullTitle"`
	File         string  `json:"file"`
	Duration     int     `json:"duration"`
	CurrentRetry int     `json:"currentRetry"`
	Err          TestErr `json:"err"`
}

//定义报表格式, 一个 json 文件 包含以下内容
// 1. stats 一个简要的统计，包括一共测试了多少用例通过了多少等。
// 2. tests 测试用例
// 3. pending
// 4. failures，失败的测试用例
// 5. passes，通过的测试用例

type Report struct {
	Stats    Stats      `json:"stats"`
	Tests    []TestCase `json:"tests"`
	Pending  []string   `json:"pending"`
	Failures []TestCase `json:"failures"`
	Passes   []TestCase `json:"passes"`
}

var logger = log.InitLogger()

// 将 json 格式的 mocha 测试报告转换为分数
func LoadMochaReport(reportFile string) (*Report, error) {
	// 1. 从reportFile读取 json 文件
	content, err := os.ReadFile(reportFile)
	if err != nil {
		logger.Error("read mocha report file failed")
		return nil, err
	}

	var report Report
	err = json.Unmarshal(content, &report)
	if err != nil {
		logger.Error("unmarshal mocha report failed")
		return nil, err
	}

	return &report, nil
}

// 根据title判断测试用例是否通过
func (r *Report) IsPass(title string) bool {
	for _, pass := range r.Passes {
		if pass.Title == title {
			return true
		}
	}
	return false
}
