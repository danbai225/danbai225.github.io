---
title: 在github自动构建go程序
date: "2023-03-08 18:43:04"
updated: "2023-03-08 18:43:04"
url: https://p00q.cn/?p=835
categories:
    - Go
tags:
    - Go
    - github
summary: 在GitHub上自动构建Go程序可以通过GitHub Actions完成。首先，在GitHub上创建一个新的仓库，并使用Git将代码推送到该仓库。然后，在该仓库的主页上，点击"Actions"选项卡，然后点击"New Workflow"按钮。选择要构建的Go版本，并创建一个包含构建和测试代码的Workflow文件。将Workflow文件提交到GitHub，并确保仓库中存在可构建的Go项目。打开仓库的Actions页面，可以看到一个新的构建正在运行。构建完成后，可以在构建日志中查看构建结果。注意事项包括使用正确的Go版本和命令，添加其他步骤和正确配置环境变量和GitHub Secrets。
id: "835"
---

# 在GitHub上自动构建Go程序可以通过GitHub Actions完成

1.  在GitHub上创建一个新的仓库，并使用Git将代码推送到该仓库。

2.  在该仓库的主页上，单击"Actions"选项卡，然后单击"New Workflow"按钮。

3. 选择要构建的Go版本，可以使用以下示例Workflow文件（.github/workflows/go.yml）：
```
name: Go

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        go-version: [ '1.16.x' ]
        os: [ ubuntu-latest, windows-latest, macOS-latest ]
    steps:
    - uses: actions/checkout@v2
    - name: Set up Go ${{ matrix.go-version }}
      uses: actions/setup-go@v2
      with:
        go-version: ${{ matrix.go-version }}
    - name: Build
      run: go build -v ./...
    - name: Test
      run: go test -v ./...

```
`该文件定义了一个包含一个job的Workflow，该job用于构建和测试代码。`

4. 将Workflow文件提交到GitHub，并确保该仓库的main分支中存在一个可构建的Go项目。

5. 打开仓库的Actions页面，可以看到一个新的构建正在运行。构建完成后，可以在构建日志中查看构建结果。

注意事项：
- 在Workflow文件中，需要使用正确的Go版本，以便构建和测试过程可以成功完成。
- 在构建和测试过程中，需要使用正确的命令，例如"go build"和"go test"。
- 可以添加其他步骤，例如在构建过程中编译并打包二进制文件。
- 如果您需要使用第三方依赖项，请确保在构建和测试过程中正确安装这些依赖项。
- 如果您需要使用环境变量或其他GitHub Secrets，请确保在Workflow文件中正确配置这些变量。

