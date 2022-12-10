package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateProjectPage is the resolver for the CreateProjectPage field.
func (r *mutationResolver) CreateProjectPage(ctx context.Context) (models.ProjectPageResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	userID, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return &models.Error{Message: "status unauthorized"}, err
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}
	newProjectPage, createProjectPageErr := r.projectPageDelegate.CreateProjectPage(userID)
	if createProjectPageErr != nil {
		err := createProjectPageErr
		return &models.Error{Message: err.Error()}, err
	}
	return &newProjectPage, nil
}

// UpdateProjectPage is the resolver for the UpdateProjectPage field.
func (r *mutationResolver) UpdateProjectPage(ctx context.Context, input models.UpdateProjectPage) (models.ProjectPageResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return &models.Error{Message: "status unauthorized"}, err
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	updateProjectPageInput := &models.ProjectPageHTTP{
		ProjectID:     input.ProjectID,
		ProjectPageID: input.ProjectPageID,
		Instruction:   input.Instruction,
		Notes:         input.Notes,
		Title:         input.Title,
		IsShared:      input.IsShared,
	}

	updatedProjectPage, updateProjectPageErr := r.projectPageDelegate.UpdateProjectPage(updateProjectPageInput)
	if updateProjectPageErr != nil {
		err := updateProjectPageErr
		return &models.Error{Message: err.Error()}, err
	}
	return updatedProjectPage, nil
}

// DeleteProjectPage is the resolver for the DeleteProjectPage field.
func (r *mutationResolver) DeleteProjectPage(ctx context.Context, projectID string) (*models.DeletedProjectPage, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedProjectPage{ProjectPageID: ""}, err
	}
	_, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return &models.DeletedProjectPage{ProjectPageID: ""}, err
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.DeletedProjectPage{ProjectPageID: ""}, err
	}
	deleteProjectPageErr := r.projectPageDelegate.DeleteProjectPage(projectID)
	if deleteProjectPageErr != nil {
		err := errors.New("baq request")
		return &models.DeletedProjectPage{ProjectPageID: ""}, err
	}
	return &models.DeletedProjectPage{ProjectPageID: projectID}, nil
}

// GetProjectPageByID is the resolver for the GetProjectPageById field.
func (r *queryResolver) GetProjectPageByID(ctx context.Context, projectPageID string) (models.ProjectPageResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return &models.Error{Message: "internal server error"}, err
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	projectPageHttp, getProjectPageByIDErr := r.projectPageDelegate.GetProjectPageById(projectPageID)
	if getProjectPageByIDErr != nil {
		err := getProjectPageByIDErr
		return &models.Error{Message: err.Error()}, err
	}
	return &projectPageHttp, nil
}

// GetAllProjectPagesByUserID is the resolver for the GetAllProjectPagesByUserID field.
func (r *queryResolver) GetAllProjectPagesByUserID(ctx context.Context, userID string) (models.ProjectPagesResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	userID, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return &models.Error{Message: "status unauthorized"}, err
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}
	projectPages, _, getAllProjectPagesErr := r.projectPageDelegate.GetAllProjectPagesByUserId(userID, "0", "0")
	if getAllProjectPagesErr != nil {
		err := getAllProjectPagesErr
		return &models.Error{Message: err.Error()}, err
	}

	return &models.ProjectPageHTTPList{
		ProjectPages: projectPages,
	}, nil
}

// GetAllProjectPagesByAccessToken is the resolver for the GetAllProjectPagesByAccessToken field.
func (r *queryResolver) GetAllProjectPagesByAccessToken(ctx context.Context, page string, pageSize string) (models.ProjectPagesResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	userId, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := userIdentityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}
	projectPages, countRows, getAllProjectPagesErr := r.projectPageDelegate.GetAllProjectPagesByUserId(userId, page, pageSize)
	if getAllProjectPagesErr != nil {
		err := getAllProjectPagesErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.ProjectPageHTTPList{
		ProjectPages: projectPages,
		CountRows:    countRows,
	}, nil
}
