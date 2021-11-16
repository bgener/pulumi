package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/codegen/internal/utils"
	"github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
)

type programTest struct {
	Name           string
	Description    string
	Skip           codegen.StringSet
	ExpectNYIDiags codegen.StringSet
	SkipCompile    codegen.StringSet
}

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

var programTests = []programTest{
	{
		Name:           "aws-s3-folder",
		Description:    "AWS S3 Folder",
		ExpectNYIDiags: codegen.NewStringSet("python", "nodejs", "dotnet"),
		SkipCompile:    codegen.NewStringSet("go", "python", "nodejs"),
		// Blocked on python: TODO[pulumi/pulumi#8062]: Re-enable this test.
		// Blocked on go:
		//   TODO[pulumi/pulumi#8064]
		//   TODO[pulumi/pulumi#8065]
		// Blocked on nodejs: TODO[pulumi/pulumi#8063]
	},
	{
		Name:        "aws-eks",
		Description: "AWS EKS",
		SkipCompile: codegen.NewStringSet("nodejs"),
		// Blocked on nodejs: TODO[pulumi/pulumi#8067]
	},
	{
		Name:        "aws-fargate",
		Description: "AWS Fargate",
		Skip:        codegen.NewStringSet("nodejs", "dotnet", "python"),
	},
	{
		Name:        "aws-s3-logging",
		Description: "AWS S3 with logging",
		SkipCompile: codegen.NewStringSet("dotnet", "nodejs", "go"),
		// Blocked on dotnet: TODO[pulumi/pulumi#8069]
		// Blocked on nodejs: TODO[pulumi/pulumi#8068]
		// Flaky in go: TODO[pulumi/pulumi#8123]
	},
	{
		Name:        "aws-webserver",
		Description: "AWS Webserver",
		SkipCompile: codegen.NewStringSet("go"),
		// Blocked on go: TODO[pulumi/pulumi#8070]
	},
	{
		Name:        "azure-native",
		Description: "Azure Native",
		Skip:        codegen.NewStringSet("go"),
		// Blocked on TODO[pulumi/pulumi#8123]
		SkipCompile: codegen.NewStringSet("go", "nodejs"),
		// Blocked on go:
		//   TODO[pulumi/pulumi#8072]
		//   TODO[pulumi/pulumi#8073]
		//   TODO[pulumi/pulumi#8074]
		// Blocked on nodejs:
		//   TODO[pulumi/pulumi#8075]
	},
	{
		Name:        "azure-sa",
		Description: "Azure SA",
	},
	{
		Name:        "kubernetes-operator",
		Description: "K8s Operator",
	},
	{
		Name:        "kubernetes-pod",
		Description: "K8s Pod",
		SkipCompile: codegen.NewStringSet("go", "nodejs"),
		// Blocked on go:
		//   TODO[pulumi/pulumi#8073]
		//   TODO[pulumi/pulumi#8074]
		// Blocked on nodejs:
		//   TODO[pulumi/pulumi#8075]
	},
	{
		Name:        "kubernetes-template",
		Description: "K8s Template",
	},
	{
		Name:        "random-pet",
		Description: "Random Pet",
	},
	{
		Name:        "aws-resource-options",
		Description: "Resource Options",
		SkipCompile: codegen.NewStringSet("go"),
		// Blocked on go: TODO[pulumi/pulumi#8076]
	},
	{
		Name:        "aws-secret",
		Description: "Secret",
	},
	{
		Name:        "functions",
		Description: "Functions",
		SkipCompile: codegen.NewStringSet("go", "dotnet"),
		// Blocked on go: TODO[pulumi/pulumi#8077]
		// Blocked on dotnet:
		//   TODO[pulumi/pulumi#8078]
		//   TODO[pulumi/pulumi#8079]
	},
	{
		Name:        "output-funcs-aws",
		Description: "Output Versioned Functions",
		Skip:        codegen.NewStringSet("nodejs", "dotnet", "python"),
	},
}

// Checks that a generated program is correct
//
// The arguments are to be read:
// (Testing environment, path to generated code, set of dependencies)
type CheckProgramOutput = func(*testing.T, string, codegen.StringSet)

// Generates a program from a pcl.Program
type GenProgram = func(program *pcl.Program) (map[string][]byte, hcl.Diagnostics, error)

type ProgramCodegenOptions struct {
	Language   string
	Extension  string
	OutputFile string
	Check      CheckProgramOutput
	GenProgram GenProgram
}

// TestProgramCodegen runs the complete set of program code generation tests against a particular
// language's code generator.
//
// A program code generation test consists of a PCL file (.pp extension) and a set of expected outputs
// for each language.
//
// The PCL file is the only piece that must be manually authored. Once the schema has been written, the expected outputs
// can be generated by running `PULUMI_ACCEPT=true go test ./..." from the `pkg/codegen` directory.
//nolint: revive
func TestProgramCodegen(
	t *testing.T,
	// language string,
	// genProgram func(program *pcl.Program) (map[string][]byte, hcl.Diagnostics, error
	testcase ProgramCodegenOptions,

) {
	ensureValidSchemaVersions(t)
	for _, tt := range programTests {
		t.Run(tt.Description, func(t *testing.T) {
			var err error
			if tt.Skip.Has(testcase.Language) {
				t.Skip()
				return
			}

			expectNYIDiags := tt.ExpectNYIDiags.Has(testcase.Language)

			testDir := filepath.Join(testdataPath, tt.Name+"-pp")
			pclFile := filepath.Join(testDir, tt.Name+".pp")
			testDir = filepath.Join(testDir, testcase.Language)
			err = os.MkdirAll(testDir, 0700)
			if err != nil && !os.IsExist(err) {
				t.Fatalf("Failed to create %q: %s", testDir, err)
			}

			contents, err := ioutil.ReadFile(pclFile)
			if err != nil {
				t.Fatalf("could not read %v: %v", pclFile, err)
			}

			expectedFile := filepath.Join(testDir, tt.Name+"."+testcase.Extension)
			expected, err := ioutil.ReadFile(expectedFile)
			if err != nil && os.Getenv("PULUMI_ACCEPT") == "" {
				t.Fatalf("could not read %v: %v", expectedFile, err)
			}

			parser := syntax.NewParser()
			err = parser.ParseFile(bytes.NewReader(contents), tt.Name+".pp")
			if err != nil {
				t.Fatalf("could not read %v: %v", pclFile, err)
			}
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)
			}

			program, diags, err := pcl.BindProgram(parser.Files, pcl.PluginHost(utils.NewHost(testdataPath)))
			if err != nil {
				t.Fatalf("could not bind program: %v", err)
			}
			if diags.HasErrors() {
				t.Fatalf("failed to bind program: %v", diags)
			}
			files, diags, err := testcase.GenProgram(program)
			assert.NoError(t, err)
			if expectNYIDiags {
				var tmpDiags hcl.Diagnostics
				for _, d := range diags {
					if !strings.HasPrefix(d.Summary, "not yet implemented") {
						tmpDiags = append(tmpDiags, d)
					}
				}
				diags = tmpDiags
			}
			if diags.HasErrors() {
				t.Fatalf("failed to generate program: %v", diags)
			}

			if os.Getenv("PULUMI_ACCEPT") != "" {
				err := ioutil.WriteFile(expectedFile, files[testcase.OutputFile], 0600)
				require.NoError(t, err)
			} else {
				assert.Equal(t, string(expected), string(files[testcase.OutputFile]))
			}
			if testcase.Check != nil && !tt.SkipCompile.Has(testcase.Language) {
				extraPulumiPackages := codegen.NewStringSet()
				for _, n := range program.Nodes {
					if r, isResource := n.(*pcl.Resource); isResource {
						pkg, _, _, _ := r.DecomposeToken()
						if pkg != "pulumi" {
							extraPulumiPackages.Add(pkg)
						}
					}
				}
				testcase.Check(t, expectedFile, extraPulumiPackages)
			}
		})
	}
}
