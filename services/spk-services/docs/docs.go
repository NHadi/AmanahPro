// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/spk": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new SPK",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKs"
                ],
                "summary": "Create SPK",
                "parameters": [
                    {
                        "description": "SPK Data",
                        "name": "spk",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SpkDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created SPK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/spk/filter": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Filter SPKs by organization ID, SPK ID, and project ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKs"
                ],
                "summary": "Filter SPKs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "project_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SPK"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/spk/{spk_id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an existing SPK",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKs"
                ],
                "summary": "Update SPK",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "SPK Data",
                        "name": "spk",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SpkDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated SPK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "SPK Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an SPK by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKs"
                ],
                "summary": "Delete SPK",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "SPK Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/spk/{spk_id}/sections": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new SPK Section",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKSections"
                ],
                "summary": "Create SPK Section",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "SPK Section Data",
                        "name": "section",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SpkSectionDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created SPK Section",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/spk/{spk_id}/sections/{section_id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an existing SPK Section",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKSections"
                ],
                "summary": "Update SPK Section",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Section ID",
                        "name": "section_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "SPK Section Data",
                        "name": "section",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SpkSectionDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated SPK Section",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "SPK Section Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an SPK Section by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKSections"
                ],
                "summary": "Delete SPK Section",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Section ID",
                        "name": "section_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "SPK Section Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/spk/{spk_id}/sections/{section_id}/details": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new SPK Detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKDetails"
                ],
                "summary": "Create SPK Detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Section ID",
                        "name": "section_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "SPK Detail Data",
                        "name": "detail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SpkDetailDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created SPK Detail",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/spk/{spk_id}/sections/{section_id}/details/{detail_id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an existing SPK Detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKDetails"
                ],
                "summary": "Update SPK Detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Section ID",
                        "name": "section_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Detail ID",
                        "name": "detail_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "SPK Detail Data",
                        "name": "detail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SpkDetailDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated SPK Detail",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "SPK Detail Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an SPK Detail by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SPKDetails"
                ],
                "summary": "Delete SPK Detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPK ID",
                        "name": "spk_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Section ID",
                        "name": "section_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Detail ID",
                        "name": "detail_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "SPK Detail Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.SpkDTO": {
            "type": "object",
            "required": [
                "ProjectId",
                "ProjectName",
                "Subject"
            ],
            "properties": {
                "Date": {
                    "$ref": "#/definitions/models.CustomDate"
                },
                "ProjectId": {
                    "type": "integer"
                },
                "ProjectName": {
                    "type": "string"
                },
                "SphId": {
                    "type": "integer"
                },
                "SpkId": {
                    "type": "integer"
                },
                "Subject": {
                    "type": "string"
                }
            }
        },
        "dto.SpkDetailDTO": {
            "type": "object",
            "required": [
                "Description",
                "Quantity",
                "TotalJasa",
                "TotalMaterial",
                "Unit",
                "UnitPriceJasa",
                "UnitPriceMaterial"
            ],
            "properties": {
                "Description": {
                    "type": "string"
                },
                "DetailId": {
                    "type": "integer"
                },
                "Quantity": {
                    "type": "number"
                },
                "TotalJasa": {
                    "type": "number"
                },
                "TotalMaterial": {
                    "type": "number"
                },
                "Unit": {
                    "type": "string"
                },
                "UnitPriceJasa": {
                    "type": "number"
                },
                "UnitPriceMaterial": {
                    "type": "number"
                }
            }
        },
        "dto.SpkSectionDTO": {
            "type": "object",
            "required": [
                "SectionTitle"
            ],
            "properties": {
                "SectionId": {
                    "type": "integer"
                },
                "SectionTitle": {
                    "type": "string"
                }
            }
        },
        "models.CustomDate": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "models.SPK": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "Creation timestamp",
                    "type": "string"
                },
                "createdBy": {
                    "description": "Created by user ID",
                    "type": "integer"
                },
                "date": {
                    "description": "Date of SPK",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.CustomDate"
                        }
                    ]
                },
                "deletedAt": {
                    "description": "Deletion timestamp",
                    "type": "string"
                },
                "deletedBy": {
                    "description": "Deleted by user ID",
                    "type": "integer"
                },
                "organizationId": {
                    "description": "Organization ID",
                    "type": "integer"
                },
                "projectId": {
                    "description": "Foreign key to Projects",
                    "type": "integer"
                },
                "projectName": {
                    "description": "Project name",
                    "type": "string"
                },
                "sections": {
                    "description": "Relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SPKSection"
                    }
                },
                "sphId": {
                    "description": "Foreign key to SPH",
                    "type": "integer"
                },
                "spkId": {
                    "description": "Primary key",
                    "type": "integer"
                },
                "subject": {
                    "description": "SPK subject",
                    "type": "string"
                },
                "totalJasa": {
                    "description": "Total Jasa Cost",
                    "type": "number"
                },
                "totalMaterial": {
                    "description": "Total Material Cost",
                    "type": "number"
                },
                "updatedAt": {
                    "description": "Update timestamp",
                    "type": "string"
                },
                "updatedBy": {
                    "description": "Updated by user ID",
                    "type": "integer"
                }
            }
        },
        "models.SPKDetail": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "Creation timestamp",
                    "type": "string"
                },
                "createdBy": {
                    "description": "Created by user ID",
                    "type": "integer"
                },
                "deletedAt": {
                    "description": "Deletion timestamp",
                    "type": "string"
                },
                "deletedBy": {
                    "description": "Deleted by user ID",
                    "type": "integer"
                },
                "description": {
                    "description": "Item description",
                    "type": "string"
                },
                "detailId": {
                    "description": "Primary key",
                    "type": "integer"
                },
                "organizationId": {
                    "description": "Organization ID",
                    "type": "integer"
                },
                "quantity": {
                    "description": "Item quantity",
                    "type": "number"
                },
                "sectionId": {
                    "description": "Foreign key to SPK Section",
                    "type": "integer"
                },
                "sphItemId": {
                    "description": "Reference to SPH Item (optional)",
                    "type": "integer"
                },
                "totalJasa": {
                    "description": "Total Jasa cost",
                    "type": "number"
                },
                "totalMaterial": {
                    "description": "Total Material cost",
                    "type": "number"
                },
                "unit": {
                    "description": "Unit of measurement",
                    "type": "string"
                },
                "unitPriceJasa": {
                    "description": "Unit price for Jasa",
                    "type": "number"
                },
                "unitPriceMaterial": {
                    "description": "Unit price for Material",
                    "type": "number"
                },
                "updatedAt": {
                    "description": "Update timestamp",
                    "type": "string"
                },
                "updatedBy": {
                    "description": "Updated by user ID",
                    "type": "integer"
                }
            }
        },
        "models.SPKSection": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "Creation timestamp",
                    "type": "string"
                },
                "createdBy": {
                    "description": "Created by user ID",
                    "type": "integer"
                },
                "deletedAt": {
                    "description": "Deletion timestamp",
                    "type": "string"
                },
                "deletedBy": {
                    "description": "Deleted by user ID",
                    "type": "integer"
                },
                "details": {
                    "description": "Relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SPKDetail"
                    }
                },
                "organizationId": {
                    "description": "Organization ID",
                    "type": "integer"
                },
                "sectionId": {
                    "description": "Primary key",
                    "type": "integer"
                },
                "sectionTitle": {
                    "description": "Section title",
                    "type": "string"
                },
                "sphSectionId": {
                    "type": "integer"
                },
                "spkId": {
                    "description": "Foreign key to SPK",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "Update timestamp",
                    "type": "string"
                },
                "updatedBy": {
                    "description": "Updated by user ID",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Provide your JWT token with \"Bearer \" prefix, e.g., \"Bearer \u003ctoken\u003e\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "SPK Management Services API",
	Description:      "This is the SPH Management Services API documentation for managing SPK, and reconciliations.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
