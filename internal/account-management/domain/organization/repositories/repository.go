package repositories

type OrganizationRepositoryInterface interface {
	CreateOrganization(name string) (string, error)
	AddMemberInOrganization(organizationId string, userId string) error
}
