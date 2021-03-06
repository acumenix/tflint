package main

import (
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"

	"github.com/terraform-linters/tflint/tools/utils"
)

type providerMeta struct {
	ResourceNames []string
}

var awsProvider = aws.Provider().(*schema.Provider)

func main() {
	providerMeta := &providerMeta{}

	for k, v := range awsProvider.ResourcesMap {
		if _, ok := v.Schema["tags"]; ok {
			providerMeta.ResourceNames = append(providerMeta.ResourceNames, k)
		}
	}

	sort.Strings(providerMeta.ResourceNames)

	templateFile := "../rules/awsrules/aws_resource_missing_tags.go.tmpl"
	providerFile := "../rules/awsrules/aws_resource_missing_tags.go"
	utils.GenerateFile(
		providerFile,
		templateFile,
		providerMeta,
	)
}
