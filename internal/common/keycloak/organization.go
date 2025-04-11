package keycloak

type DomainRepresensation struct {
	Name     string `json:"name"`
	Verified bool   `json:"verified"`
}

type OrganizationRepresentation struct {
	Name    string                 `json:"name"`
	Enabled bool                   `json:"enabled"`
	Domains []DomainRepresensation `json:"domains"`
}

func NewOrganizationRepresentation(name string) *OrganizationRepresentation {
	return &OrganizationRepresentation{
		Name:    name,
		Enabled: true,
		Domains: []DomainRepresensation{
			{
				Name:     name,
				Verified: false,
			},
		},
	}
}
