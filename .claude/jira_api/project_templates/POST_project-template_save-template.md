# POST /rest/api/3/project-template/save-template
**operationId:** `saveTemplate`
**Summary:** Save a custom project template

Save custom template

This API endpoint allows you to save a customised template

***Note: Custom Templates are only supported for Jira Enterprise edition.***

## Request Body
Content-Type: `application/json`
object:
  - `templateDescription`: string
  - `templateFromProjectRequest`: SaveProjectTemplateRequest
  - `templateName`: string

## Responses
- 200: object:
  - `projectTemplateKey`: ProjectTemplateKey
