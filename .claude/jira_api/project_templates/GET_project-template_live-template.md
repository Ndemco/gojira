# GET /rest/api/3/project-template/live-template
**operationId:** `liveTemplate`
**Summary:** Gets a custom project template

Get custom template

This API endpoint allows you to get a live custom project template details by either templateKey or projectId

***Note: Custom Templates are only supported for Jira Enterprise edition.***

## Parameters
- `projectId` [query] string — optional - The \{@link String\} containing the project key linked to the custom template to retrieve
- `templateKey` [query] string — optional - The \{@link String\} containing the key of the custom template to retrieve

## Responses
- 200: object:
  - `archetype`: ProjectArchetype
  - `defaultBoardView`: string
  - `description`: string
  - `liveTemplateProjectIdReference`: integer(int64)
  - `name`: string
  - `projectTemplateKey`: ProjectTemplateKey
  - `snapshotTemplate`: object
  - `templateGenerationOptions`: CustomTemplateOptions
  - `type`: string
