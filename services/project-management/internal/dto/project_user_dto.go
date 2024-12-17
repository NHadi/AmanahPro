package dto

// ProjectUserDTO represents the data transfer object for Project User operations
type ProjectUserDTO struct {
	UserID      *int   `json:"UserID,omitempty"` // User ID (nullable)
	UserName    string `json:"UserName"`         // User Name
	OldUserName string `json:"OldUserName"`      // Old User Name (for ChangeUser)
	Role        string `json:"Role"`             // User Role
}
