# PUT /rest/api/3/project/{projectIdOrKey}/properties/{propertyKey}
**operationId:** `setProjectProperty`
**Summary:** Set project property

Sets the value of the [project property](https://developer.atlassian.com/cloud/jira/platform/storing-data-without-a-database/#a-id-jira-entity-properties-a-jira-entity-properties). You can use project properties to store custom data against the project.

The value of the request body must be a [valid](http://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Administer

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `propertyKey` [path] (required) string — The key of the project property. The maximum length is 255 characters.

## Request Body
Content-Type: `application/json`
any

## Responses
- 200: any
- 201: any
- 400: Returned if the project key or id is invalid.
- 401: Returned if the authentication credentials are incorrect.
- 403: Returned if the user does not have permission to administer the project.
- 404: Returned if the project is not found.
