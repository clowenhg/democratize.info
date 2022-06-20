package congress

type Collection struct {
	CollectionCode string      `json:"collectionCode"`
	CollectionName string      `json:"collectionName"`
	PackageCount   int         `json:"packageCount"`
	GranuleCount   interface{} `json:"granuleCount" omitempty:"true"`
}

type Collections []Collection

type CollectionsResult struct {
	Collections Collections `json:"collections"`
}

func (c Collections) GetByCode(code string) *Collection {
	for _, collection := range c {
		if collection.CollectionCode == code {
			return &collection
		}
	}

	return nil
}
