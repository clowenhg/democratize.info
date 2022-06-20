package congress

import "time"

// PackageType is named as such to avoid collision with the `package` reserved word
type PackageType struct {
	PackageID    string    `json:"packageId"`
	LastModified time.Time `json:"lastModified"`
	PackageLink  string    `json:"packageLink"`
	DocClass     string    `json:"docClass"`
	Title        string    `json:"title"`
	Congress     string    `json:"congress"`
	DateIssued   string    `json:"dateIssued"`
}

type Packages []PackageType

type PackagesResult struct {
	Count        int         `json:"count"`
	Message      interface{} `json:"message"`
	NextPage     string      `json:"nextPage"`
	PreviousPage interface{} `json:"previousPage"`
	Packages     Packages    `json:"packages"`
}

func (p Packages) GetByID(id string) *PackageType {
	for _, pack := range p {
		if pack.PackageID == id {
			return &pack
		}
	}
	return nil
}
