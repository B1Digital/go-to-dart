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
	"os"
	"path"
	"regexp"
	"strings"
)

const (
	// BaseProduct
	baseProduct = `	BaseProduct`

	baseProductGroup = `	BaseProductGroup`

	// BaseAttachment
	baseAttachment = "	BaseAttachment"

	// BaseSalesPrice
	baseSalesPrice = `	BaseSalesPrice`

	// BaseRecordFields
	baseRecordFields = `	BaseRecordFields`

	// baseContent
	baseContent = `	BaseContent`
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
		// "BaseList.go":         true,
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

	err = os.RemoveAll(editedOutDir)
	checkError(err)
	err = os.MkdirAll(editedOutDir, 0666)
	checkError(err)

	fileContent, err := os.ReadFile(*inputDir + "/base_sales_price.go")
	checkError(err)
	baseSalesPriceContent = string(fileContent)

	fileContent, err = os.ReadFile(*inputDir + "/BaseAttachment.go")
	checkError(err)
	baseAttachmentContent = string(fileContent)

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

	err = os.WriteFile(path.Join(editedOutDir, "BaseProduct.go"), []byte(expandEmbedStruct(baseProductContent)), 0666)
	checkError(err)

	err = os.WriteFile(path.Join(editedOutDir, "BaseProductGroup.go"), []byte(expandEmbedStruct(baseProductGroupContent)), 0666)
	checkError(err)

	// base_sales_price.go
	err = os.WriteFile(path.Join(editedOutDir, "base_sales_price.go"), []byte(expandEmbedStruct(baseSalesPriceContent)), 0666)
	checkError(err)

	// BaseAttachment.go
	err = os.WriteFile(path.Join(editedOutDir, "BaseAttachment.go"), []byte(expandEmbedStruct(baseAttachmentContent)), 0666)
	checkError(err)

	// BaseContent.go
	err = os.WriteFile(path.Join(editedOutDir, "BaseContent.go"), []byte(expandEmbedStruct(baseContentContent)), 0666)
	checkError(err)

	// BaseRecordFields.go

	err = os.WriteFile(path.Join(editedOutDir, "BaseRecordFields.go"), []byte(expandEmbedStruct(baseRecordFieldsContent)), 0666)
	checkError(err)

	for _, f := range dir {
		if bases[f.Name()] || f.IsDir() {
			continue
		}
		fileContent, err = os.ReadFile(path.Join(*inputDir, f.Name()))
		checkError(err)
		fileStr := string(fileContent)
		fileStr = expandEmbedStruct(fileStr)
		err = os.WriteFile(path.Join(editedOutDir, f.Name()), []byte(expandEmbedStruct(fileStr)), 0666)
		checkError(err)

	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getStructContent(content string) string {
	const startStructBlock = ` struct {`
	lenStartStruct := len(startStructBlock)
	return content[strings.Index(content, startStructBlock)+lenStartStruct : strings.Index(content, `}`)]
}

func expandEmbedStruct(fileContent string) string {

	x := getStructContent(baseAttachmentContent)
	fileContent = strings.ReplaceAll(fileContent, baseAttachment, x)

	x = getStructContent(baseSalesPriceContent)
	fileContent = strings.ReplaceAll(fileContent, baseSalesPrice, x)

	x = getStructContent(baseRecordFieldsContent)
	fileContent = strings.ReplaceAll(fileContent, baseRecordFields, x)

	x = getStructContent(baseContentContent)
	fileContent = strings.ReplaceAll(fileContent, baseContent, x)

	x = getStructContent(baseProductGroupContent)
	fileContent = strings.ReplaceAll(fileContent, baseProductGroup, x)

	x = getStructContent(baseProductContent)
	fileContent = strings.ReplaceAll(fileContent, baseProduct, x)

	return fileContent
}
