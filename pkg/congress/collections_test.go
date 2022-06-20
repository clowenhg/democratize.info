package congress_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/clowenhg/democratize.info.git/pkg/congress"
	"github.com/stretchr/testify/suite"
)

func (s *CollectionsSuite) TestGetByCode() {
	expected := congress.Collection{
		CollectionCode: "BILLS",
		CollectionName: "Congressional Bills",
		PackageCount:   241189,
		GranuleCount:   nil,
	}

	// Allowing for Table driven testing
	s.Run("BILLS", func() { s.ByCode("BILLS", expected) })
}

func (s *CollectionsSuite) ByCode(code string, expected congress.Collection) {
	c := s.collections.GetByCode(code)

	s.Assert().NotNilf(c, "Collection %s Not Found", code)
	s.Equal(expected.CollectionCode, c.CollectionCode)
	s.Equal(expected.CollectionName, c.CollectionName)
	s.Equal(expected.PackageCount, c.PackageCount)
	s.Equal(expected.GranuleCount, c.GranuleCount)
}

type CollectionsSuite struct {
	suite.Suite
	collections congress.Collections
}

func TestCollections(t *testing.T) {
	s := new(CollectionsSuite)

	suite.Run(t, s)
}

func (s *CollectionsSuite) SetupTest() {
	jsonBytes, err := ioutil.ReadFile("testdata/collections.json")
	if err != nil {
		s.FailNow("Reading test data\n", err)
	}

	var result congress.CollectionsResult
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		s.FailNow("Unmarshalling test data\n", err)
	}

	s.collections = result.Collections
}
