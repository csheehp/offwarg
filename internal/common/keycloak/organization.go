package keycloak

type DomainRepresensation struct {
	Domains string `json:"domains"`
	Verfied bool   `json:"verfied"`
}

type OrganizationRepresentation struct {
	Name    string                 `json:"name"`
	Enabled bool                   `json:"enabled"`
	Domain  []DomainRepresensation `json:"domains"`
}

func NewOrganizationRepresentation(name string) *OrganizationRepresentation {
	return &OrganizationRepresentation{
		Name:    name,
		Enabled: true,
		Domain: []DomainRepresensation{
			{
				Domains: name,
				Verfied: false,
			},
		},
	}
}
