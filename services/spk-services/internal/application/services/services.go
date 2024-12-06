package services

import (
	commonServices "github.com/NHadi/AmanahPro-common/services"
)

type Services struct {
	SPKService        *SpkService
	AuditTrailService *commonServices.AuditTrailService
}
