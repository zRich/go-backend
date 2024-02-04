# 使用说明

本系统是基于 hardhat 实现的自动评分系统。

## hardhat 基本命令

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```

## 目录结构

students目录存放学生提交的文件，每个学生在该目录下有一个子目录，目录名为学生的学号。
对学生提交的任务进行评分时，会将文件复制到项目目录下。例如 `students/2023000143/contracts`下到文件复制到 `contracts`。

scripts 目录用于存放评分的脚步。

## 评分过程

评分过程是基于 Mocha 测试框架实现的，结果是 mocha 的测试报告。后台程序根据测试报告打分。
