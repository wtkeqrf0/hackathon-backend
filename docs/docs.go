// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Contact us",
            "url": "https://github.com/while-act/hackathon-backend/issues/new/choose",
            "email": "matvey-sizov@mail.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.DomainName}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/session": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns detail information about me",
                "tags": [
                    "Session"
                ],
                "summary": "Get detail data about the user by session",
                "responses": {
                    "200": {
                        "description": "Info about session",
                        "schema": {
                            "$ref": "#/definitions/dao.Me"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "401": {
                        "description": "User isn't logged in",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Get cookie and delete them from db",
                "tags": [
                    "Session"
                ],
                "summary": "Delete cookie session",
                "responses": {
                    "200": {
                        "description": "deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Compare the user's password with an existing user's password. If it matches, create session of this user",
                "tags": [
                    "Auth"
                ],
                "summary": "Sign in by password",
                "parameters": [
                    {
                        "description": "User's email, password",
                        "name": "SignIn",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user's session",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Data is not valid",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Compare the user's password with an existing user's password. If it matches, create session of the user. If the user does not exist, create new user",
                "tags": [
                    "Auth"
                ],
                "summary": "Sign up by password",
                "parameters": [
                    {
                        "description": "User's email, password, firstName, lastName, inn",
                        "name": "SignUp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user's session",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Data is not valid",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/calc": {
            "post": {
                "description": "Returns PDF file, gotten from body",
                "tags": [
                    "Calc"
                ],
                "summary": "Generate PDF from body",
                "parameters": [
                    {
                        "description": "Completed application form",
                        "name": "from",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.History"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "PDF file"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/calc/save": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Saves given values to user's history",
                "tags": [
                    "Calc"
                ],
                "summary": "Save calc data to history",
                "parameters": [
                    {
                        "description": "Completed application form",
                        "name": "from",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.History"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "401": {
                        "description": "User isn't logged in",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/calc/{industry}": {
            "get": {
                "description": "Returns detail information about industry",
                "tags": [
                    "Calc"
                ],
                "summary": "Get data about industry",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Industry Branch",
                        "name": "industry",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Info about industry",
                        "schema": {
                            "$ref": "#/definitions/dao.Industry"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/company": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns information about company by session",
                "tags": [
                    "Company"
                ],
                "summary": "Get data about company by session",
                "responses": {
                    "200": {
                        "description": "Info about company",
                        "schema": {
                            "$ref": "#/definitions/dao.Company"
                        }
                    },
                    "401": {
                        "description": "User isn't logged in",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates information about company by INN",
                "tags": [
                    "Company"
                ],
                "summary": "Update data about company",
                "parameters": [
                    {
                        "description": "Company",
                        "name": "updCompany",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateCompany"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "401": {
                        "description": "User isn't logged in",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/email/send-code": {
            "post": {
                "description": "Generates secret code and sends it to specified email",
                "tags": [
                    "Auth"
                ],
                "summary": "Send code to specified email",
                "parameters": [
                    {
                        "description": "User's email",
                        "name": "SignIn",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Email"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user's session",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Data is not valid",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/user": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates user's info",
                "tags": [
                    "User"
                ],
                "summary": "Update user's data",
                "parameters": [
                    {
                        "description": "Fields to update",
                        "name": "updFields",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Updated"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "401": {
                        "description": "User isn't logged in",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/user/email": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates user's email",
                "tags": [
                    "User"
                ],
                "summary": "Update user's email",
                "parameters": [
                    {
                        "description": "New email with password",
                        "name": "updEmail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateEmail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Updated"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "401": {
                        "description": "User isn't logged in",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/user/password": {
            "patch": {
                "description": "Updates user's password",
                "tags": [
                    "User"
                ],
                "summary": "Update user's password",
                "parameters": [
                    {
                        "description": "Email with new password",
                        "name": "updPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Updated"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        },
        "/user/{history_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns PDF file got from user history",
                "tags": [
                    "User"
                ],
                "summary": "Generate PDF file from user history",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Unique id from history",
                        "name": "history_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "PDF file"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "401": {
                        "description": "User isn't logged in",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MyError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dao.Company": {
            "type": "object",
            "required": [
                "inn"
            ],
            "properties": {
                "inn": {
                    "type": "string",
                    "example": "7707083893"
                },
                "name": {
                    "type": "string",
                    "maxLength": 150,
                    "minLength": 2,
                    "example": "ООО 'Парк'"
                },
                "website": {
                    "type": "string",
                    "example": "https://www.rusprofile.ru"
                }
            }
        },
        "dao.Industry": {
            "type": "object",
            "required": [
                "avgSalary",
                "avgSalaryCad",
                "avgWorkersNum",
                "avgWorkersNumCad"
            ],
            "properties": {
                "avgSalary": {
                    "type": "number",
                    "example": 72.7825875
                },
                "avgSalaryCad": {
                    "type": "number",
                    "example": 95.54196489
                },
                "avgWorkersNum": {
                    "type": "number",
                    "example": 1.208
                },
                "avgWorkersNumCad": {
                    "type": "number",
                    "example": 1243
                }
            }
        },
        "dao.Me": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "name",
                "role"
            ],
            "properties": {
                "biography": {
                    "type": "string",
                    "example": "I'd like to relax"
                },
                "city": {
                    "type": "string",
                    "example": "Москва"
                },
                "country": {
                    "type": "string",
                    "example": "Россия"
                },
                "email": {
                    "type": "string",
                    "example": "myemail@gmail.com"
                },
                "fatherName": {
                    "type": "string",
                    "example": "Ivanovich"
                },
                "firstName": {
                    "type": "string",
                    "example": "Ivan"
                },
                "lastName": {
                    "type": "string",
                    "example": "Ivanov"
                },
                "name": {
                    "type": "string",
                    "example": "user94"
                },
                "position": {
                    "type": "string",
                    "example": "Director"
                },
                "role": {
                    "type": "string",
                    "example": "USER"
                }
            }
        },
        "dto.Email": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "myemail@gmail.com"
                }
            }
        },
        "dto.History": {
            "type": "object",
            "required": [
                "accountingServices",
                "companyName",
                "constructionFacilitiesArea",
                "districtTitle",
                "equipmentType",
                "facilityType",
                "fullTimeEmployees",
                "industryBranch",
                "landArea",
                "organizationType",
                "other",
                "patent"
            ],
            "properties": {
                "accountingServices": {
                    "type": "boolean",
                    "example": true
                },
                "companyName": {
                    "type": "string",
                    "maxLength": 150,
                    "minLength": 2,
                    "example": "ООО 'Парк'"
                },
                "constructionFacilitiesArea": {
                    "type": "number",
                    "maximum": 0,
                    "example": 50
                },
                "districtTitle": {
                    "type": "string",
                    "example": "ВАО"
                },
                "equipmentType": {
                    "type": "string",
                    "example": "Токарные станки"
                },
                "facilityType": {
                    "description": "TODO",
                    "type": "string",
                    "example": "idk"
                },
                "fullTimeEmployees": {
                    "type": "integer",
                    "maximum": 0,
                    "example": 50
                },
                "industryBranch": {
                    "type": "string",
                    "example": "Авиационная_промышленность"
                },
                "landArea": {
                    "type": "number",
                    "maximum": 0,
                    "example": 120
                },
                "organizationType": {
                    "type": "string",
                    "example": "ООО"
                },
                "other": {
                    "type": "string",
                    "example": "I want some cookies"
                },
                "patent": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "dto.SignIn": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "myemail@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4,
                    "example": "bob126"
                }
            }
        },
        "dto.SignUp": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "biography": {
                    "type": "string",
                    "maxLength": 1024,
                    "example": "I'd like to relax"
                },
                "city": {
                    "type": "string",
                    "example": "Москва"
                },
                "company": {
                    "$ref": "#/definitions/dao.Company"
                },
                "country": {
                    "type": "string",
                    "example": "Россия"
                },
                "email": {
                    "type": "string",
                    "example": "myemail@gmail.com"
                },
                "fatherName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 2,
                    "example": "Ivanovich"
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 2,
                    "example": "Ivan"
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 2,
                    "example": "Ivanov"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4,
                    "example": "bob126"
                },
                "position": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "Director"
                }
            }
        },
        "dto.UpdateCompany": {
            "type": "object",
            "properties": {
                "inn": {
                    "type": "string",
                    "example": "7707083893"
                },
                "name": {
                    "type": "string",
                    "maxLength": 150,
                    "minLength": 2,
                    "example": "ООО 'Парк'"
                },
                "website": {
                    "type": "string",
                    "example": "https://www.rusprofile.ru"
                }
            }
        },
        "dto.UpdateEmail": {
            "type": "object",
            "required": [
                "newEmail",
                "password"
            ],
            "properties": {
                "newEmail": {
                    "type": "string",
                    "example": "myemail@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4,
                    "example": "mob126"
                }
            }
        },
        "dto.UpdatePassword": {
            "type": "object",
            "required": [
                "code",
                "email",
                "newPassword"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "example": "N1OSP"
                },
                "email": {
                    "type": "string",
                    "example": "myemail@gmail.com"
                },
                "newPassword": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4,
                    "example": "mob126"
                }
            }
        },
        "dto.UpdateUser": {
            "type": "object",
            "properties": {
                "biography": {
                    "type": "string",
                    "maxLength": 1024,
                    "example": "I'd like to relax"
                },
                "city": {
                    "type": "string",
                    "example": "Москва"
                },
                "country": {
                    "type": "string",
                    "example": "Россия"
                },
                "fatherName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 2,
                    "example": "Ivanovich"
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 2,
                    "example": "Ivan"
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 2,
                    "example": "Ivanov"
                },
                "position": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "Director"
                }
            }
        },
        "errs.MyError": {
            "description": "All native errors must be this type",
            "type": "object",
            "properties": {
                "advice": {
                    "type": "string",
                    "example": "Try to send request later"
                },
                "message": {
                    "type": "string",
                    "example": "Exception was occurred"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "session_id",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "While.act API",
	Description:      "It's an API interacting with While.act using Golang",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
