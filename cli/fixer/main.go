/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	Name    :    Emre Teke
	Date    :    Fri Jul 28 2023
	Time   :    14:49
	Notes       :

*/

package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

const (
	// BaseProduct
	baseProductRgx = `^\tBaseProduct$`

	baseProductGroupRgx = `^\tBaseProductGroup+$`

	// BaseAttachment
	baseAttachment = "\tBaseAttachment"

	// BaseSalesPrice
	baseSalesPriceRgx = `^\tBaseSalesPrice+$`

	// BaseRecordFields
	baseRecordFieldsRgx = `^\tBaseRecordFields+$`

	// baseContent
	baseContentRgx = `^\tBaseContent+$`
)

var (
	regexs   map[*regexp.Regexp]bool
	inputDir = flag.String("input", "models", "models dir")

	outputDir = flag.String("output", "out", "output directory")
)

var (
	// BaseProduct
	baseProductContent string

	baseProductGroupContent string

	baseAttachmentContent string

	baseSalesPriceContent string

	baseRecordFieldsContent string

	baseContentContent string
)

var (
	bases = map[string]bool{
		"base_sales_price.go": true,
		"BaseAttachment.go":   true,
		"BaseContent.go":      true,
		"BaseList.go":         true,
		"BaseProduct.go":      true,
		"BaseProductGroup.go": true,
		"BaseRecordFields.go": true,
	}
)

func main() {
	dir, err := os.ReadDir(*inputDir)
	if err != nil {
		panic(err)
	}

	editedOutDir := path.Join(*inputDir, "edited")

	err = os.MkdirAll(editedOutDir, 0666)
	checkError(err)

	fileContent, err := os.ReadFile(*inputDir + "/base_sales_price.go")
	checkError(err)
	baseSalesPriceContent = string(fileContent)

	fileContent, err = os.ReadFile(*inputDir + "/BaseAttachment.go")
	checkError(err)
	baseAttachmentContent = string(fileContent)

	//

	fileContent, err = os.ReadFile(*inputDir + "/BaseContent.go")
	checkError(err)
	baseContentContent = string(fileContent)

	//

	fileContent, err = os.ReadFile(*inputDir + "/BaseProduct.go")
	checkError(err)
	baseProductContent = string(fileContent)

	fileContent, err = os.ReadFile(*inputDir + "/BaseProductGroup.go")
	checkError(err)
	baseProductGroupContent = string(fileContent)

	fileContent, err = os.ReadFile(*inputDir + "/BaseRecordFields.go")
	checkError(err)
	baseRecordFieldsContent = string(fileContent)
	baseProductContent = strings.ReplaceAll(baseProductContent, baseAttachment, getStructContent(baseAttachmentContent))

	err = os.WriteFile(path.Join(editedOutDir, "BaseProduct.go"), []byte(baseAttachmentContent), 0666)
	checkError(err)
	fmt.Println(baseProductContent)

	_ = editedOutDir
	_ = dir
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getStructContent(content string) string {
	const startStructBlock = `struct {`
	lenStartStruct := len(startStructBlock)
	return content[strings.Index(content, startStructBlock)+lenStartStruct : strings.Index(content, `}`)]
}
