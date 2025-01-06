package projectrole

type ProjectRoleResolver struct{}

func (ProjectRoleResolver) CreateProjectRole() (ProjectRole, error) {

	return ProjectRole{}, nil
}
