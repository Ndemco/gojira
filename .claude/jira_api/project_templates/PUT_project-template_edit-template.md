# PUT /rest/api/3/project-template/edit-template
**operationId:** `editTemplate`
**Summary:** Edit a custom project template

Edit custom template

This API endpoint allows you to edit an existing customised template.

***Note: Custom Templates are only supported for Jira Enterprise edition.***

## Request Body
Content-Type: `application/json`
object:
  - `templateDescription`: string
  - `templateGenerationOptions`: CustomTemplateOptions
  - `templateKey`: string
  - `templateName`: string

## Responses
- 200: any
