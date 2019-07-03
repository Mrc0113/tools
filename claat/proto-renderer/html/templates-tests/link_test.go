package htmltests

import (
	"go/build"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/googlecodelabs/tools/claat/proto-renderer/html"
	"github.com/googlecodelabs/tools/claat/proto-renderer/testing-utils"
	"github.com/googlecodelabs/tools/third_party"
)

const linkFileRelDir = "src/github.com/googlecodelabs/tools/claat/proto-renderer/html/templates-tests/testdata/InlineContent/google_weather.txt"

func TestRenderLinkTemplate(t *testing.T) {
	linkFileAbsDir := filepath.Join(build.Default.GOPATH, linkFileRelDir)
	weatherLinkBytes, err := ioutil.ReadFile(linkFileAbsDir)
	if err != nil {
		t.Errorf("Reading %#v outputted %#v", linkFileAbsDir, err)
		return
	}
	weatherLinkOutput := string(weatherLinkBytes[:])

	linkProto := testingutils.NewLink(
		"https://www.google.com/search?q=weather+in+nyc",
		testingutils.NewStylizedTextPlain("hey google,"),
		testingutils.NewStylizedTextStrong(" how's the"),
		testingutils.NewStylizedTextEmphasized(" weather in "),
		testingutils.NewStylizedTextStrongAndEmphasized("NYC today?"),
	)

	canonicalTests := []*testingutils.CanonicalRenderingBatch{
		{
			InProto: &tutorial.Link{},
			Out:     "",
			Ok:      false,
		},
		{
			InProto: testingutils.NewLink("only://link.does.not/work?#ok=false"),
			Out:     "",
			Ok:      false,
		},
		{
			InProto: linkProto,
			Out:     weatherLinkOutput,
			Ok:      true,
		},
	}
	testingutils.CanonicalRenderTestBatch(html.Render, canonicalTests, t)
}
