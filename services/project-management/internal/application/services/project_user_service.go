package services

import (
	"AmanahPro/services/project-management/common/messagebroker"
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"AmanahPro/services/project-management/internal/dto"
	"fmt"
	"log"
)

type ProjectUserService struct {
	projectUserRepo repositories.ProjectUserRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
	queueName       string
}

func NewProjectUserService(
	projectUserRepo repositories.ProjectUserRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	queueName string,
) *ProjectUserService {
	return &ProjectUserService{
		projectUserRepo: projectUserRepo,
		rabbitPublisher: rabbitPublisher,
		queueName:       queueName,
	}
}

func (s *ProjectUserService) Create(user *models.ProjectUser) error {
	model := &models.ProjectUser{
		ProjectID:      user.ProjectID,
		UserID:         user.UserID,
		Role:           user.Role,
		OrganizationID: user.OrganizationID,
	}
	log.Printf("Creating project user: %+v", user)
	if err := s.projectUserRepo.Create(model); err != nil {
		log.Printf("Error creating project user: %v", err)
		return fmt.Errorf("error creating project user: %w", err)
	}

	event := map[string]interface{}{
		"event":   "ProjectUserCreated",
		"payload": user,
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("Error publishing project user create event: %v", err)
		return fmt.Errorf("error publishing project user create event: %w", err)
	}

	log.Printf("Successfully created project user: %+v", user)
	return nil
}

func (s *ProjectUserService) Update(user *models.ProjectUser) error {
	model := &models.ProjectUser{
		ID:             user.ID,
		ProjectID:      user.ProjectID,
		UserID:         user.UserID,
		Role:           user.Role,
		OrganizationID: user.OrganizationID,
	}
	log.Printf("Updating project user: %+v", user)
	if err := s.projectUserRepo.Update(model); err != nil {
		log.Printf("Error updating project user: %v", err)
		return fmt.Errorf("error updating project user: %w", err)
	}

	event := map[string]interface{}{
		"event":   "ProjectUserUpdated",
		"payload": user,
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("Error publishing project user update event: %v", err)
		return fmt.Errorf("error publishing project user update event: %w", err)
	}

	log.Printf("Successfully updated project user: %+v", user)
	return nil
}

func (s *ProjectUserService) AssignUserToProject(userID, projectID int, organizationID *int, role *string) error {

	roleValue := "null"
	if role != nil {
		roleValue = *role
	}

	log.Printf("Assigning user %d to project %d with role %s", userID, projectID, roleValue)

	// Check if the user is already assigned
	existing, err := s.projectUserRepo.FindByUserAndProject(userID, projectID, organizationID)
	if err == nil && existing != nil {
		return fmt.Errorf("user %d is already assigned to project %d", userID, projectID)
	}

	// Assign user to project
	projectUser := &models.ProjectUser{
		UserID:         userID,
		ProjectID:      projectID,
		OrganizationID: organizationID,
		Role:           role,
	}
	if err := s.projectUserRepo.Create(projectUser); err != nil {
		return fmt.Errorf("failed to assign user to project: %w", err)
	}

	// Publish event to RabbitMQ
	event := map[string]interface{}{
		"event":   "UserAssignedToProject",
		"payload": projectUser,
	}
	if err := s.rabbitPublisher.PublishEvent("project_user_queue", event); err != nil {
		log.Printf("Failed to publish event: %v", err)
	}

	log.Printf("Successfully assigned user %d to project %d", userID, projectID)
	return nil
}

func (s *ProjectUserService) FindByUserAndProject(userID int, projectID int, organizationID *int) (*dto.ProjectUserDTO, error) {

	data, err := s.projectUserRepo.FindByUserAndProject(userID, projectID, organizationID)
	if err != nil {
		log.Printf("Error searching projects: %v", err)
		return nil, err
	}

	return data, nil
}
