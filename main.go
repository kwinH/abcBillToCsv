package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
	"time"
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		fmt.Println("用法: abcBillToCsv-linux-amd64 <pdf文件路径> [输出文件名]")
		os.Exit(1)
	}

	// 从命令行参数获取PDF文件路径
	pdfPath := os.Args[1]

	content, err := readPdfGroup(pdfPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取PDF文件失败: %v\n", err)
		os.Exit(1)
	}

	var outputFilename string
	if len(os.Args) >= 3 {
		// 如果提供了第二个参数，则用作输出文件名
		outputFilename = os.Args[2]
	} else {
		// 生成带时间戳的文件名
		timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS 格式
		outputFilename = fmt.Sprintf("output_%s.csv", timestamp)
	}

	// 将CSV内容写入文件
	err = writeCSVToFile(content, outputFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入CSV文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("账单内容已成功提取并保存到%s文件中\n", outputFilename)
}

func readPdfGroup(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	w := make(map[float64]int)

	hanldFrist := false
	var csvBuilder strings.Builder
	totalPage := r.NumPage()

	csvBuilder.Grow(1024 * 100)

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			if len(row.Content) < 7 {
				continue
			}

			// 检查是否是表头行
			if row.Content[0].S == "交易日期" && row.Content[1].S == "交易时间" {
				if !hanldFrist {
					// 只有第一次出现时才建立列索引映射
					w = make(map[float64]int)
					for k, word := range row.Content {
						w[word.X] = k
					}
					hanldFrist = true
					csvBuilder.WriteString(fmt.Sprintf("%s\n", strings.Join([]string{"交易日期", "交易时间", "交易摘要", "交易金额", "本次余额", "对手信息", "日 志 号", "交易渠道", "交易附言"}, ",")))
				}
				continue // 跳过所有表头行，不再重复添加
			}

			// 正常数据行处理
			rowStr := make([]string, 9)
			for _, word := range row.Content {
				if colIdx, exists := w[word.X]; exists {
					rowStr[colIdx] = word.S
				}
			}
			csvBuilder.WriteString(fmt.Sprintf("%s\n", strings.Join(rowStr, ",")))
		}
	}

	return csvBuilder.String(), nil
}

func writeCSVToFile(content, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
