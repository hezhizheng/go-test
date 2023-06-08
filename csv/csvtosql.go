package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// 打包 go build -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" csv/csvtosql.go

	//table := "users"
	//columns := []string{"id", "name", "email"}
	//values := [][]string{
	//	{"1", "Alice", "alice@example.com"},
	//	{"2", "Bob", "bob@example.com"},
	//	{"3", "Charlie", "charlie@example.com"},
	//}
	//updateColumns := []string{"name", "email"}
	//
	//query, err := GenerateInsertOrUpdateQuery(table, columns, values, updateColumns)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("Generated SQL query:", query)
	//}

	file := `C:\Users\DexterHo\Desktop\jd_copy1.csv`

	args := os.Args

	if len(args) > 1 && args[1] != "" {
		file = args[1]
	}

	table := getPathInfo(file)

	x, y, _ := readCSV(file)

	sql, eee_ := GenerateInsertOrUpdateQuery(table["basename"], x, y, x)

	filePath := "./" + table["basename"] + ".sql"
	fileFormat := "%s"
	saveStringToFile(sql, filePath, fileFormat)

	log.Println(x, eee_)
}

func readCSV(file string) ([]string, [][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
		return nil, nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)

	// 读取第一行数据
	headers, err := reader.Read()
	if err != nil {
		panic(err)
		return nil, nil, err
	}

	// 读取剩余所有行数据
	var rows [][]string

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}

		rows = append(rows, row)
	}

	return headers, rows, nil
}

func GenerateInsertOrUpdateQuery2(table string, columns []string, values [][]string, updateColumns []string) (string, error) {
	if len(columns) != len(values[0]) {
		return "", fmt.Errorf("number of columns does not match number of values")
	}

	var builder strings.Builder
	builder.WriteString("INSERT INTO `")
	builder.WriteString(table)
	builder.WriteString("` (`")
	builder.WriteString(strings.Join(columns, "`, `"))
	builder.WriteString("`) VALUES ")

	valueStrings := make([]string, len(values))
	for i, row := range values {
		formattedValues := make([]string, len(row))
		for j, v := range row {
			formattedValues[j] = fmt.Sprintf(`"%s"`, strings.Replace(v, `"`, `\"`, -1))
		}
		valueStrings[i] = fmt.Sprintf("(%s)", strings.Join(formattedValues, ","))
	}
	builder.WriteString(strings.Join(valueStrings, ","))

	if len(updateColumns) > 0 {
		setStatements := make([]string, len(updateColumns))
		for j, col := range updateColumns {
			idx := indexString(columns, col)
			if idx == -1 {
				return "", fmt.Errorf("update column %s not found", col)
			}
			setStatements[j] = fmt.Sprintf("`%s`=VALUES(`%s`)", col, col)
		}

		builder.WriteString(" ON DUPLICATE KEY UPDATE ")
		builder.WriteString(strings.Join(setStatements, ", "))
	}

	return builder.String(), nil
}

func GenerateInsertOrUpdateQuery(table string, columns []string, values [][]string, updateColumns []string) (string, error) {
	if len(columns) != len(values[0]) {
		panic("number of columns does not match number of values")
		return "", fmt.Errorf("number of columns does not match number of values")
	}

	var builder strings.Builder
	builder.WriteString("INSERT INTO `")
	builder.WriteString(table)
	builder.WriteString("` (`")
	builder.WriteString(strings.Join(columns, "`, `"))
	builder.WriteString("`) VALUES ")

	valueStrings := make([]string, len(values))
	for i, row := range values {
		formattedValues := make([]string, len(row))
		for j, v := range row {
			//formattedValues[j] = fmt.Sprintf(`"%s"`, strings.ReplaceAll(v, `"`, `\"`))
			// 处理反斜杠
			formattedValues[j] = fmt.Sprintf(`"%s"`, strings.ReplaceAll(strings.ReplaceAll(v, `"`, `\"`), "\\", "\\\\"))
		}
		valueStrings[i] = fmt.Sprintf("(%s)", strings.Join(formattedValues, ","))
	}
	builder.WriteString(strings.Join(valueStrings, ","))

	if len(updateColumns) > 0 {
		setStatements := make([]string, len(updateColumns))
		for j, col := range updateColumns {
			idx := indexString(columns, col)
			if idx == -1 {
				return "", fmt.Errorf("update column %s not found", col)
			}
			setStatements[j] = fmt.Sprintf("`%s`=VALUES(`%s`)", col, col)
		}

		builder.WriteString(" ON DUPLICATE KEY UPDATE ")
		builder.WriteString(strings.Join(setStatements, ", "))
	}

	return builder.String(), nil
}

func indexString(slice []string, str string) int {
	for i, s := range slice {
		if s == str {
			return i
		}
	}
	return -1
}

func saveStringToFile1(text string, filePath string, fileFormat string) error {
	formattedText := fmt.Sprintf(fileFormat, text)
	err := ioutil.WriteFile(filePath, []byte(formattedText), 0644)
	if err != nil {
		return err
	}
	return nil
}

func saveStringToFile(text string, filePath string, fileFormat string) error {
	formattedText := fmt.Sprintf(fileFormat, strings.ReplaceAll(text, "\uFEFF", ""))
	err := ioutil.WriteFile(filePath, []byte(formattedText), 0644)
	if err != nil {
		return err
	}
	return nil
}

func getPathInfo(path string) map[string]string {
	dir, file := filepath.Split(path)
	ext := filepath.Ext(file)

	return map[string]string{
		"dirname":   dir,
		"basename":  file[:len(file)-len(ext)],
		"extension": ext,
	}

	/**
	path := "/path/to/file.txt"
	    pathInfo := GetPathInfo(path)

	    fmt.Println(pathInfo["dirname"])   // 输出：/path/to/
	    fmt.Println(pathInfo["basename"])  // 输出：file
	    fmt.Println(pathInfo["extension"]) // 输出：.txt
	*/
}
