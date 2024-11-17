// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/auth/login": {
            "post": {
                "description": "Authenticate a user and return a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Get the health status of the API",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Health check endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
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
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/samples": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a list of samples with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "samples"
                ],
                "summary": "List samples",
                "parameters": [
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 100,
                        "description": "Limit for pagination",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ListSampleResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new sample",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "samples"
                ],
                "summary": "Sample create",
                "parameters": [
                    {
                        "description": "Sample information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SampleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SampleResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/samples/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get details of a sample",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "samples"
                ],
                "summary": "Get a sample by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sample ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SampleResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.LoginRequest": {
            "description": "LoginRequest is a struct that represents the request of login",
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "request.SampleDetail": {
            "description": "Sample detail information",
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "minimum": 1
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "price": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "request.SampleRequest": {
            "description": "Sample information",
            "type": "object",
            "required": [
                "int_val",
                "sample_detail_required",
                "string_val"
            ],
            "properties": {
                "array_val": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "email": {
                    "type": "string",
                    "example": "test@example.com"
                },
                "id": {
                    "description": "refs: https://github.com/swaggo/swag#example-value-of-struct",
                    "type": "string"
                },
                "int_val": {
                    "type": "integer",
                    "minimum": 1
                },
                "sample_detail_not_required": {
                    "$ref": "#/definitions/request.SampleDetail"
                },
                "sample_detail_required": {
                    "$ref": "#/definitions/request.SampleDetail"
                },
                "string_val": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                }
            }
        },
        "response.ErrorResponse": {
            "description": "Error response structure",
            "type": "object",
            "properties": {
                "details": {},
                "message": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "response.ListSampleResponse": {
            "description": "Sample list information",
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "samples": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.SampleResponse"
                    }
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "response.LoginResponse": {
            "description": "LoginResponse is a struct that represents the response of login",
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "response.SampleResponse": {
            "description": "Sample information",
            "type": "object",
            "properties": {
                "array_val": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "int_val": {
                    "type": "integer"
                },
                "string_val": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Go REST Clean API with Chi",
	Description:      "This is a sample server for a Go REST API using clean architecture.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
