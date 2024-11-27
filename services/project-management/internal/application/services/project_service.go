package services

import (
	"AmanahPro/services/project-management/common/messagebroker"
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"AmanahPro/services/project-management/internal/dto"
	"fmt"
	"log"
)

type ProjectService struct {
	projectRepo     repositories.ProjectRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
	queueName       string
}

func NewProjectService(
	projectRepo repositories.ProjectRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	queueName string,
) *ProjectService {
	return &ProjectService{
		projectRepo:     projectRepo,
		rabbitPublisher: rabbitPublisher,
		queueName:       queueName,
	}
}

func (s *ProjectService) Create(project *models.Project) error {
	log.Printf("Creating project: %+v", project)
	if err := s.projectRepo.Create(project); err != nil {
		log.Printf("Error creating project: %v", err)
		return fmt.Errorf("error creating project: %w", err)
	}

	event := map[string]interface{}{
		"event":   "Created",
		"payload": project,
		"meta": map[string]interface{}{
			"idField": "ProjectID", // Specify the primary key field in the metadata
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("Error publishing project create event: %v", err)
		return fmt.Errorf("error publishing project create event: %w", err)
	}

	log.Printf("Successfully created project: %+v", project)
	return nil
}

func (s *ProjectService) Update(project *models.Project) error {
	log.Printf("Updating project: %+v", project)
	if err := s.projectRepo.Update(project); err != nil {
		log.Printf("Error updating project: %v", err)
		return fmt.Errorf("error updating project: %w", err)
	}

	event := map[string]interface{}{
		"event":   "Updated",
		"payload": project,
		"meta": map[string]interface{}{
			"idField": "ProjectID", // Specify the primary key field in the metadata
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("Error publishing project update event: %v", err)
		return fmt.Errorf("error publishing project update event: %w", err)
	}

	log.Printf("Successfully updated project: %+v", project)
	return nil
}

func (s *ProjectService) Delete(id int) error {
	log.Printf("Deleting project with ID: %d", id)
	if err := s.projectRepo.Delete(id); err != nil {
		log.Printf("Error deleting project: %v", err)
		return fmt.Errorf("error deleting project: %w", err)
	}

	event := map[string]interface{}{
		"event":   "Deleted",
		"payload": map[string]int{"id": id},
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("Error publishing project delete event: %v", err)
		return fmt.Errorf("error publishing project delete event: %w", err)
	}

	log.Printf("Successfully deleted project with ID: %d", id)
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
