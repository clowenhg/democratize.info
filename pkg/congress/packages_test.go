package congress_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/clowenhg/democratize.info.git/pkg/congress"
	"github.com/stretchr/testify/suite"
)

func (s *PackagesSuite) TestGetByID() {
	expectedTime, err := time.Parse(time.RFC3339, "2022-06-20T04:23:42Z")
	s.Assert().Nil(err)

	expected := congress.PackageType{
		PackageID:    "BILLS-117hr8086ih",
		LastModified: expectedTime,
		PackageLink:  "https://api.govinfo.gov/packages/BILLS-117hr8086ih/summary",
		DocClass:     "hr",
		Title:        "Conservation Jobs Act of 2022",
		Congress:     "117",
		DateIssued:   "2022-06-15",
	}

	// Allowing for Table driven testing
	s.Run("BILLS-117hr8086ih", func() { s.ByID("BILLS-117hr8086ih", expected) })
}

func (s *PackagesSuite) ByID(id string, expected congress.PackageType) {
	p := s.packages.GetByID(id)

	s.Assert().NotNilf(p, "Package %s Not Found", id)
	s.Equal(expected.PackageID, p.PackageID)
	s.Equal(expected.LastModified, p.LastModified)
	s.Equal(expected.PackageLink, p.PackageLink)
	s.Equal(expected.DocClass, p.DocClass)
	s.Equal(expected.Title, p.Title)
	s.Equal(expected.Congress, p.Congress)
	s.Equal(expected.DateIssued, p.DateIssued)
}

type PackagesSuite struct {
	suite.Suite
	packages congress.Packages
}

func TestPackages(t *testing.T) {
	s := new(PackagesSuite)

	suite.Run(t, s)
}

func (s *PackagesSuite) SetupTest() {
	jsonBytes, err := ioutil.ReadFile("testdata/packages.json")
	if err != nil {
		s.FailNow("Reading test data\n", err)
	}

	var result congress.PackagesResult
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		s.FailNow("Unmarshalling test data\n", err)
	}

	s.packages = result.Packages
}
