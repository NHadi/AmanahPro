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
