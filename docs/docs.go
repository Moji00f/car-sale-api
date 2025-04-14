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
        "/v1/health/": {
            "get": {
                "description": "Health Check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "resultCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                },
                "validationErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/validation.ValidationError"
                    }
                }
            }
        },
        "validation.ValidationError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "property": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthBearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
