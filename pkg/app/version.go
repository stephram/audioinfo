package app

const (
	UNKNOWN = "unknown"
)

//-X "$(PACKAGE)/pkg/app.Branch=$(BRANCH_NAME)" \
//-X "$(PACKAGE)/pkg/app.BuildDate=$(BUILD_DATE)" \
//-X "$(PACKAGE)/pkg/app.Commit=$(GIT_COMMIT)" \
//-X "$(PACKAGE)/pkg/app.Version=$(VERSION)" \
//-X "$(PACKAGE)/pkg/app.Author=$(AUTHOR)" \
//-X "$(PACKAGE)/pkg/app.AuthorEmail=$(AUTHOR_EMAIL)" \

var (
	Name        = UNKNOWN
	Product     = UNKNOWN
	Branch      = UNKNOWN
	BuildDate   = UNKNOWN
	Commit      = UNKNOWN
	Version     = UNKNOWN
	Author      = UNKNOWN
	AuthorEmail = UNKNOWN
)

type SourceDetails struct {
	Name        string `json:"name,omitempty"`
	Product     string `json:"product,omitempty"`
	Branch      string `json:"branch,omitempty"`
	BuildDate   string `json:"build_date,omitempty"`
	Commit      string `json:"commit,omitempty"`
	Version     string `json:"version,omitempty"`
	Author      string `json:"author,omitempty"`
	AuthorEmail string `json:"author_email,omitempty"`
}

func New() SourceDetails {
	return SourceDetails{
		Name:        Name,
		Product:     Product,
		Branch:      Branch,
		BuildDate:   BuildDate,
		Commit:      Commit,
		Version:     Version,
		Author:      Author,
		AuthorEmail: AuthorEmail,
	}
}
