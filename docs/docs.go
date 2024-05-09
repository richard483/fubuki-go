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
        "/gemini-history/history-data": {
            "get": {
                "description": "get all history data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gemini-history"
                ],
                "summary": "GetAllHistoryData",
                "responses": {}
            },
            "post": {
                "description": "create history data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gemini-history"
                ],
                "summary": "CreateHistoryData",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "CreateGeminiHistory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GeminiHistory"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "delete history data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gemini-history"
                ],
                "summary": "DeleteHistoryData",
                "parameters": [
                    {
                        "type": "string",
                        "description": "history ID to be deleted",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "patch": {
                "description": "update history data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gemini-history"
                ],
                "summary": "UpdateHistoryData",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "UpdateGeminiHistory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GeminiHistory"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/gemini-history/history-data/bulk": {
            "post": {
                "description": "create many history data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gemini-history"
                ],
                "summary": "CreateManyHistoryData",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "CreateManyGeminiHistory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/request.GeminiHistory"
                            }
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/gemini/chat": {
            "post": {
                "description": "chat action API",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gemini"
                ],
                "summary": "Chat",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "GeminiText",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GeminiText"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/gemini/prompt-text": {
            "post": {
                "description": "get prompt text result by prompt string",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gemini"
                ],
                "summary": "Prompt Text",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "GeminiText",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GeminiText"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/gemini/tune": {
            "get": {
                "description": "tune gemini model with history data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gemini"
                ],
                "summary": "TuneModel",
                "responses": {}
            }
        }
    },
    "definitions": {
        "request.GeminiHistory": {
            "type": "object",
            "required": [
                "answer",
                "question"
            ],
            "properties": {
                "answer": {
                    "type": "string"
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "request.GeminiText": {
            "type": "object",
            "required": [
                "text"
            ],
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
