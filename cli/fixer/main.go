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
	baseProductRgx = `^\tBaseProduct+$`

	baseProductGroupRgx = `^\tBaseProductGroup+$`

	// BaseAttachment
	baseAttachmentRgxs = `^\tBaseAttachment+$`

	// BaseSalesPrice
	baseSalesPriceRgx = `^\tBaseSalesPrice+$`

	// BaseRecordFields
	baseRecordFieldsRgx = `^\tBaseRecordFields+$`

	// baseContent
	baseContentRgx = `^\tBaseContent+$`
)

var (
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

func main() {
	dir, err := os.ReadDir(*inputDir)
	if err != nil {
		panic(err)
	}

	editedOutDir := path.Join(*inputDir, "edited")
	fileContent, err := os.ReadFile(*inputDir + "/base_sales_price.go")
	checkError(err)
	baseSalesPriceContent = string(fileContent)
	baseSalesPriceContent = getStructContent(baseSalesPriceContent)

	fileContent, err = os.ReadFile(*inputDir + "/BaseAttachment.go")
	checkError(err)
	baseAttachmentContent = string(fileContent)
	baseAttachmentContent = getStructContent(baseAttachmentContent)

	//

	fileContent, err = os.ReadFile(*inputDir + "/BaseContent.go")
	checkError(err)
	baseContentContent = string(fileContent)
	baseContentContent = getStructContent(baseContentContent)

	//

	fileContent, err = os.ReadFile(*inputDir + "/BaseProduct.go")
	checkError(err)
	baseProductContent = string(fileContent)
	baseProductContent = getStructContent(baseProductContent)

	fileContent, err = os.ReadFile(*inputDir + "/BaseProductGroup.go")
	checkError(err)
	baseProductGroupContent = string(fileContent)
	baseProductGroupContent = getStructContent(baseProductGroupContent)

	fileContent, err = os.ReadFile(*inputDir + "/BaseRecordFields.go")
	checkError(err)
	baseRecordFieldsContent = string(fileContent)
	baseRecordFieldsContent = getStructContent(baseRecordFieldsContent)

	baseAttachmentRgx, err := regexp.Compile(baseAttachmentRgxs)
	checkError(err)
	baseProductContent = string(baseAttachmentRgx.ReplaceAll([]byte(baseAttachmentContent), []byte(baseProductContent)))

	fmt.Println(baseProductContent)

	_ = baseAttachmentRgx
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
