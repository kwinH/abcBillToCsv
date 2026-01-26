# abcBillToCsv

这是一个将农业银行账单转换为 CSV 格式的 Go 应用程序。

## 功能特性

- 将农业银行账单数据转换为标准 CSV 格式
- 支持多平台（Linux、macOS、Windows）
- 自动化构建和发布流程

## 使用方法

1. 下载适合您操作系统的预编译二进制文件：

   - Linux: [abcBillToCsv-linux-amd64](releases/download/vX.X.X/abcBillToCsv-linux-amd64)
   - macOS: [abcBillToCsv-darwin-amd64](releases/download/vX.X.X/abcBillToCsv-darwin-amd64)
   - Windows: [abcBillToCsv-windows-amd64.exe](releases/download/vX.X.X/abcBillToCsv-windows-amd64.exe)
2. 给予执行权限（Linux/macOS）：

   ```bash
   chmod +x abcBillToCsv-[platform]
   ```
3. 运行应用程序并传入您的 ABC 银行账单 PDF 文件路径：

   ```bash
   # 基本用法：将 PDF 文件路径作为第一个参数
   ./abcBillToCsv-[platform] /path/to/your/bank_statement.pdf

   # 指定输出文件名：将 PDF 文件路径作为第一个参数，输出文件名作为第二个参数
   ./abcBillToCsv-[platform] /path/to/your/bank_statement.pdf my_output.csv
   ```

   - 如果不指定输出文件名，程序将自动生成一个带时间戳的文件名，格式为 `output_YYYYMMDDHHMMSS.csv`
   - 程序会自动解析农业银行 PDF 账单，并提取交易日期、交易时间、交易摘要、交易金额、本次余额、对手信息、日志号、交易渠道和交易附言等信息

## 开发

### 构建

```bash
go build -o abcBillToCsv .
```

### 测试

```bash
go test ./...
```

## 发布流程

本项目使用 GitHub Actions 进行自动化构建和发布，包含以下功能：

### 构建与测试

- 自动拉取代码
- 安装 Go 依赖
- 编译应用程序
- 运行单元测试

### 版本管理

- 支持手动指定版本号
- 支持自动版本递增（major、minor、patch）
- 基于 Git 标签的版本跟踪

#### 版本递增类型

- `major`: 主版本号递增（如 v1.2.3 → v2.0.0）
- `minor`: 次版本号递增（如 v1.2.3 → v1.3.0）
- `patch`: 修订号递增（如 v1.2.3 → v1.2.4）

### 多平台构建

- Linux (amd64)
- macOS (amd64)
- Windows (amd64)

### 自动发布

- 自动创建 Git 标签
- 生成 GitHub Release
- 上传多平台二进制文件

## 发布操作

您可以手动触发发布流程：

1. 在 Actions 选项卡中选择 "Go Build and Release" 工作流
2. 选择是否手动指定版本号或使用自动版本递增
3. 如果选择自动递增，请指定递增类型（major、minor 或 patch）
4. 工作流将自动构建、测试并发布新版本

## 许可证

[MIT License](LICENSE)
