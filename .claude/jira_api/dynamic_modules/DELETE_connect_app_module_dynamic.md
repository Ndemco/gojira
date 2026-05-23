# DELETE /rest/atlassian-connect/1/app/module/dynamic
**operationId:** `DynamicModulesResource.removeModules_delete`
**Summary:** Remove modules

Remove all or a list of modules registered by the calling app.

**[Permissions](#permissions) required:** Only Connect apps can make this request.

## Parameters
- `moduleKey` [query] []string — The key of the module to remove. To include multiple module keys, provide multiple copies of this parameter.

## Responses
- 204: Returned if the request is successful.
- 401: object:
  - `message`: string
