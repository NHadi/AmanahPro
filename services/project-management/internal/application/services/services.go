package services

type Services struct {
	ProjectUserService  *ProjectUserService
	ProjectService      *ProjectService
	ProjectRecapService *ProjectRecapService

	// Add ConsumerServices for RabbitMQ
	ProjectConsumer      *ConsumerService
	ProjectRecapConsumer *ConsumerService
	ProjectUserConsumer  *ConsumerService
}
