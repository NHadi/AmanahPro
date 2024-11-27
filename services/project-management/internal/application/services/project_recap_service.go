package services

import (
	"AmanahPro/services/project-management/common/messagebroker"
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"fmt"
	"log"
)

type ProjectRecapService struct {
	projectRecapRepo repositories.ProjectRecapRepository
	rabbitPublisher  *messagebroker.RabbitMQPublisher
	queueName        string
}

func NewProjectRecapService(
	projectRecapRepo repositories.ProjectRecapRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	queueName string,
) *ProjectRecapService {
	return &ProjectRecapService{
		projectRecapRepo: projectRecapRepo,
		rabbitPublisher:  rabbitPublisher,
		queueName:        queueName,
	}
}

func (s *ProjectRecapService) Create(recap *models.ProjectRecap) error {
	log.Printf("Creating project recap: %+v", recap)
	if err := s.projectRecapRepo.Create(recap); err != nil {
		log.Printf("Error creating project recap: %v", err)
		return fmt.Errorf("error creating project recap: %w", err)
	}

	event := map[string]interface{}{
		"event":   "ProjectRecapCreated",
		"payload": recap,
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("Error publishing project recap create event: %v", err)
		return fmt.Errorf("error publishing project recap create event: %w", err)
	}

	log.Printf("Successfully created project recap: %+v", recap)
	return nil
}

func (s *ProjectRecapService) Update(recap *models.ProjectRecap) error {
	log.Printf("Updating project recap: %+v", recap)
	if err := s.projectRecapRepo.Update(recap); err != nil {
		log.Printf("Error updating project recap: %v", err)
		return fmt.Errorf("error updating project recap: %w", err)
	}

	event := map[string]interface{}{
		"event":   "ProjectRecapUpdated",
		"payload": recap,
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("Error publishing project recap update event: %v", err)
		return fmt.Errorf("error publishing project recap update event: %w", err)
	}

	log.Printf("Successfully updated project recap: %+v", recap)
	return nil
}

func (s *ProjectRecapService) Delete(id int) error {
	log.Printf("Deleting project recap with ID: %d", id)
	if err := s.projectRecapRepo.Delete(id); err != nil {
		log.Printf("Error deleting project recap: %v", err)
		return fmt.Errorf("error deleting project recap: %w", err)
	}

	event := map[string]interface{}{
		"event":   "ProjectRecapDeleted",
		"payload": map[string]int{"id": id},
	}
	if err := s.rabbitPublisher.PublishEvent(s.queueName, event); err != nil {
		log.Printf("Error publishing project recap delete event: %v", err)
		return fmt.Errorf("error publishing project recap delete event: %w", err)
	}

	log.Printf("Successfully deleted project recap with ID: %d", id)
	return nil
}
