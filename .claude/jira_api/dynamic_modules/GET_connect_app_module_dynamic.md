# GET /rest/atlassian-connect/1/app/module/dynamic
**operationId:** `DynamicModulesResource.getModules_get`
**Summary:** Get modules

Returns all modules registered dynamically by the calling app.

**[Permissions](#permissions) required:** Only Connect apps can make this request.

## Responses
- 200: object:
  - `modules` (required): []ConnectModule
- 401: object:
  - `message`: string
