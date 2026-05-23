# POST /rest/api/3/priorityscheme
**operationId:** `createPriorityScheme`
**Summary:** Create priority scheme

Creates a new priority scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `defaultPriorityId` (required): integer(int64)
  - `description`: string
  - `mappings`: allOf(PriorityMapping)
  - `name` (required): string
  - `priorityIds` (required): []integer(int64)
  - `projectIds`: []integer(int64)

## Responses
- 201: object:
  - `id`: string
  - `task`: allOf(TaskProgressBeanJsonNode)
- 202: object:
  - `id`: string
  - `task`: allOf(TaskProgressBeanJsonNode)
- 400: Returned if the request isn't valid.

**Mappings Validation Errors**

 *  ``The priorities with IDs [ID 1, ID 2, ...] require mapping. Please provide mappings in the 'in' mappings object, where these priorities are the keys with corresponding values.`` The listed priority ID(s) have not been provided as keys for ``in`` mappings but are required, add them to the mappings object.
- 401: Returned if the authentication credentials are incorrect.
- 403: Returned if the user doesn't have the necessary permissions.
- 409: Returned if an action with this priority scheme is still in progress.
