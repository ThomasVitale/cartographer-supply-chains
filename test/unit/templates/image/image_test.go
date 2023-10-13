package templates

import (
	"path/filepath"
	"testing"

	cartotesting "github.com/vmware-tanzu/cartographer/pkg/testing"
)

func TestKpackTemplate(t *testing.T) {

	testSuite := cartotesting.Suite{

		"template, workload and expected defined in files": {
			Given: cartotesting.Given{
				Template: &cartotesting.TemplateFile{
					Path: filepath.Join("deliverable", "regular-template", "template.yaml"),
				},
				Workload: &cartotesting.WorkloadFile{
					Path: filepath.Join("deliverable", "common-workload.yaml"),
				},
				SupplyChain: &cartotesting.MockSupplyChain{
					Params: &cartotesting.SupplyChainParamsObject{Params: params},
				},
			},
			Expect: &cartotesting.ExpectedFile{
				Path: filepath.Join("deliverable", "common-expectation.yaml"),
			},
		},
	}

	testSuite.Run(t)
}
