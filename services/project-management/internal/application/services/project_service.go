package services

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"AmanahPro/services/project-management/internal/dto"
	"fmt"
	"log"

	"github.com/NHadi/AmanahPro-common/messagebroker"
)

type ProjectService struct {
	projectRepo     repositories.ProjectRepository
	projectUserRepo repositories.ProjectUserRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
	queueName       string
}

func NewProjectService(
	projectRepo repositories.ProjectRepository,
	projectUserRepo repositories.ProjectUserRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	queueName string,
) *ProjectService {
	return &ProjectService{
		projectRepo:     projectRepo,
		projectUserRepo: projectUserRepo,
		rabbitPublisher: rabbitPublisher,
		queueName:       queueName,
	}
}

func (s *ProjectService) GetByID(id int) (*models.Project, error) {
	log.Printf("Fetching Project with ID: %d", id)

	sph, err := s.projectRepo.GetByID(id, false)
	if err != nil {
		log.Printf("Error fetching sph: %v", err)
		return nil, fmt.Errorf("failed to fetch sph: %w", err)
	}

	return sph, nil
}

// CreateProject creates a new project
func (s *ProjectService) CreateProject(project *models.Project, traceID string) error {
	log.Printf("Creating project: %+v", project)

	// Save to the database
	if err := s.projectRepo.Create(project); err != nil {
		log.Printf("TraceID: %s - Error creating project: %v", traceID, err)
		return fmt.Errorf("error creating project: %w", err)
	}

	// Publish event for centralized logging
	event := map[string]interface{}{
		"event":   "Created",
		"payload": project,
		"meta": map[string]interface{}{
			"traceId": traceID,
			"action":  "CREATE",
			"idField": "ProjectID",
			"userId":  project.CreatedBy,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("TraceID: %s - Error publishing create event for ProjectID: %d, %v", traceID, project.ProjectID, err)
	}

	log.Printf("TraceID: %s - Successfully created project: %+v", traceID, project)
	return nil
}

// UpdateProject updates an existing project
func (s *ProjectService) UpdateProject(project *models.Project, traceID string) error {
	log.Printf("TraceID: %s - Updating project: %+v", traceID, project)

	// Update in the database
	if err := s.projectRepo.Update(project); err != nil {
		log.Printf("TraceID: %s - Error updating project: %v", traceID, err)
		return fmt.Errorf("error updating project: %w", err)
	}

	// Publish event for centralized logging
	event := map[string]interface{}{
		"event":   "Updated",
		"payload": project,
		"meta": map[string]interface{}{
			"traceId": traceID,
			"action":  "UPDATE",
			"idField": "ProjectID",
			"userId":  project.UpdatedBy,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("TraceID: %s - Error publishing update event for ProjectID: %d, %v", traceID, project.ProjectID, err)
	}

	if err := s.rabbitPublisher.PublishEvent(fmt.Sprintf("%s_breakdown", s.queueName), event); err != nil {
		log.Printf("TraceID: %s - Error publishing update event for ProjectID: %d, %v", traceID, project.ProjectID, err)
	}

	log.Printf("TraceID: %s - Successfully updated project: %+v", traceID, project)
	return nil
}

// DeleteProject deletes a project by ID
func (s *ProjectService) DeleteProject(projectID int, traceID string, userID int) error {
	log.Printf("TraceID: %s - Deleting project with ID: %d", traceID, projectID)

	// Delete from the database
	if err := s.projectRepo.Delete(projectID); err != nil {
		log.Printf("TraceID: %s - Error deleting project: %v", traceID, err)
		return fmt.Errorf("error deleting project: %w", err)
	}

	// Publish event for centralized logging
	event := map[string]interface{}{
		"event":   "Deleted",
		"payload": map[string]int{"ProjectID": projectID},
		"meta": map[string]interface{}{
			"traceId": traceID,
			"action":  "DELETE",
			"idField": "ProjectID",
			"userId":  userID,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("TraceID: %s - Error publishing delete event for ProjectID: %d, %v", traceID, projectID, err)
	}

	log.Printf("TraceID: %s - Successfully deleted project with ID: %d", traceID, projectID)
	return nil
}

func (s *ProjectService) SearchProjectsByOrganization(organizationID int, query string) ([]dto.ProjectDTO, error) {
	log.Printf("Searching projects for organization %d with query: %s", organizationID, query)

	projects, err := s.projectRepo.SearchProjectsByOrganization(organizationID, query)
	if err != nil {
		log.Printf("Error searching projects: %v", err)
		return nil, err
	}

	return projects, nil
}

func (s *ProjectService) SearchProjectUsersByProject(projectID int, organizationID int) ([]models.ProjectUser, error) {
	log.Printf("Searching project users for projectID %d", projectID)

	projects, err := s.projectUserRepo.GetByProjectID(projectID)
	if err != nil {
		log.Printf("Error searching projects: %v", err)
		return nil, err
	}

	return projects, nil
}

func (s *ProjectService) SearchProjectUsersById(ID int, organizationID int) (*models.ProjectUser, error) {
	log.Printf("Searching project user with ID %d and OrganizationID %d", ID, organizationID)

	// Retrieve the project user by ID
	projectUser, err := s.projectUserRepo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching project user by ID %d: %v", ID, err)
		return nil, fmt.Errorf("failed to fetch project user: %w", err)
	}

	// Check if the retrieved user belongs to the specified organization
	if projectUser == nil || (projectUser.OrganizationID != nil && *projectUser.OrganizationID != organizationID) {
		log.Printf("Project user ID %d does not belong to OrganizationID %d", ID, organizationID)
		return nil, fmt.Errorf("project user not found or does not belong to the organization")
	}

	log.Printf("Successfully retrieved project user ID %d for OrganizationID %d", ID, organizationID)
	return projectUser, nil
}

// AssignUser adds a user to a project
func (s *ProjectService) AssignUser(projectID int, userID *int, userName, role string, assignedBy, organizationID int) error {
	log.Printf("Assigning User '%s' with UserID: %v to ProjectID: %d", userName, userID, projectID)

	projectUser := &models.ProjectUser{
		ProjectID:      projectID,
		UserID:         userID,
		UserName:       userName,
		Role:           role,
		CreatedBy:      &assignedBy,
		OrganizationID: &organizationID,
	}

	if err := s.projectUserRepo.Create(projectUser); err != nil {
		log.Printf("Failed to assign user: %v", err)
		return fmt.Errorf("failed to assign user to project: %w", err)
	}

	log.Printf("Successfully assigned User '%s' with UserID: %v to ProjectID: %d", userName, userID, projectID)
	return nil
}

// UnAssignUser removes a user from a project
func (s *ProjectService) UnAssignUser(projectID int, userID int) error {
	log.Printf("Unassigning User ID '%d' from ProjectID: %d", userID, projectID)

	// Fetch users assigned to the project
	projectUsers, err := s.projectUserRepo.GetByProjectID(projectID)
	if err != nil {
		log.Printf("Failed to fetch users for unassigning: %v", err)
		return fmt.Errorf("failed to fetch users: %w", err)
	}

	// Find the user to remove by UserID
	var userToRemove *models.ProjectUser
	for _, pu := range projectUsers {
		if pu.ID == userID { // Compare UserID (integer)
			userToRemove = &pu
			break
		}
	}

	// Check if user was found
	if userToRemove == nil {
		return fmt.Errorf("user ID '%d' not assigned to the project", userID)
	}

	// Perform the delete operation
	if err := s.projectUserRepo.Delete(userToRemove.ID); err != nil {
		log.Printf("Failed to unassign user ID '%d': %v", userID, err)
		return fmt.Errorf("failed to unassign user ID '%d' from project: %w", userID, err)
	}

	log.Printf("Successfully unassigned User ID '%d' from ProjectID: %d", userID, projectID)
	return nil
}

// ChangeUser updates an assigned user's details
func (s *ProjectService) ChangeUser(projectID int, oldUserName, newUserName, role string, updatedBy int) error {
	log.Printf("Changing User '%s' to '%s' for ProjectID %d", oldUserName, newUserName, projectID)

	// Fetch users assigned to the project
	projectUsers, err := s.projectUserRepo.GetByProjectID(projectID)
	if err != nil {
		log.Printf("Failed to fetch users for change: %v", err)
		return fmt.Errorf("failed to fetch users: %w", err)
	}

	// Find the user to update
	var userToUpdate *models.ProjectUser
	for _, pu := range projectUsers {
		if pu.UserName == oldUserName {
			userToUpdate = &pu
			break
		}
	}

	if userToUpdate == nil {
		return fmt.Errorf("user '%s' not found in the project", oldUserName)
	}

	// Update user details
	userToUpdate.UserName = newUserName
	userToUpdate.Role = role
	userToUpdate.UpdatedBy = &updatedBy

	if err := s.projectUserRepo.Update(userToUpdate); err != nil {
		log.Printf("Failed to change user: %v", err)
		return fmt.Errorf("failed to update user assignment: %w", err)
	}

	log.Printf("Successfully changed User '%s' to '%s' for ProjectID %d", oldUserName, newUserName, projectID)
	return nil
}

// UpdateProject updates an existing project
func (s *ProjectService) UpdateProjectUser(project *models.ProjectUser) error {
	log.Printf("Updating project User: %+v", project)

	// Update in the database
	if err := s.projectUserRepo.Update(project); err != nil {
		log.Printf("Error updating project user: %v", err)
		return fmt.Errorf("error updating project: %w", err)
	}

	// Step 2: Update the Category in ProjectFinancial for matching ProjectUserID and ProjectID
	if err := s.projectUserRepo.UpdateCategoryByProjectUser(project.ProjectID, project.ID, project.UserName); err != nil {
		log.Printf("Error updating financial categories for ProjectUserID %d: %v", project.ID, err)
		return fmt.Errorf("error updating financial categories: %w", err)
	}

	log.Printf("Successfully updated project user: %+v", project)
	return nil
}

// GetCombinedFinancialData retrieves financial summary data along with details for each project.
func (s *ProjectFinancialService) GetProjectFinancialSPVSummary(userID int) ([]dto.ProjectFinancialSPVSummaryDTO, error) {
	// Fetch summaries
	summaries, err := s.projectFinancialRepo.GetProjectFinancialSPVSummary(userID)
	if err != nil {
		log.Printf("Error fetching project financial summaries: %v", err)
		return nil, err
	}

	log.Println("Successfully combined project financial data")
	return summaries, nil
}
