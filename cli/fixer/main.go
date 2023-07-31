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
)

const (
	// BaseProduct
	baseProductRgx = `^\tBaseProduct+$`

	baseProductGroupRgx = `^\tBaseProductGroup+$`

	// BaseAttachment
	baseAttachmentRgx = `^\tBaseAttachment+$`

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
	if err != nil {
		panic(err)
	}
	baseSalesPriceContent = string(fileContent)

	_ = editedOutDir
	_ = dir
}
