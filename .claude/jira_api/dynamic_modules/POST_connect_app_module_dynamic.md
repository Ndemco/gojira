# POST /rest/atlassian-connect/1/app/module/dynamic
**operationId:** `DynamicModulesResource.registerModules_post`
**Summary:** Register modules

Registers a list of modules.

**[Permissions](#permissions) required:** Only Connect apps can make this request.

## Request Body
Content-Type: `application/json`
object:
  - `modules` (required): []ConnectModule

## Responses
- 200: Returned if the request is successful.
- 400: object:
  - `message`: string
- 401: object:
  - `message`: string
