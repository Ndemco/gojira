# POST /rest/api/3/project-template
**operationId:** `createProjectWithCustomTemplate`
**Summary:** Create custom project

Creates a project based on a custom template provided in the request.

The request body should contain the project details and the capabilities that comprise the project:

 *  `details` \- represents the project details settings
 *  `template` \- represents a list of capabilities responsible for creating specific parts of a project

A capability is defined as a unit of configuration for the project you want to create.

This operation is:

 *  [asynchronous](#async). Follow the `Location` link in

## Request Body
Content-Type: `application/json`
object:
  - `details`: CustomTemplatesProjectDetails
  - `template`: CustomTemplateRequestDTO

## Responses
- 303: any
