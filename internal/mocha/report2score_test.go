package mocha

import (
	"os"
	"testing"
)

const (
	// ReportFile = "./test/data/test-results.json"
	//定义报表文件是 test/data/test-results.json
	ReportFile = "../../test/data/test-results.json"
)

func TestLoadMochaReport(t *testing.T) {
	// 打印当前路径
	curPath, _ := os.Getwd()
	t.Logf("current path: %s", curPath)

	report, err := LoadMochaReport(ReportFile)
	if err != nil {
		t.Errorf("LoadMochaReport failed: %v", err)
	}

	// 打印报告的统计信息
	t.Logf("report stats: %+v", report.Stats)
}
