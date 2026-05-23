# PUT /rest/atlassian-connect/1/migration/properties/{entityType}
**operationId:** `MigrationResource.updateEntityPropertiesValue_put`
**Summary:** Bulk update entity properties

Updates the values of multiple entity properties for an object, up to 50 updates per request. This operation is for use by Connect apps during app migration.

## Parameters
- `Atlassian-Transfer-Id` [header] (required) string(uuid) — The app migration transfer ID.
- `entityType` [path] (required) string — The type indicating the object that contains the entity properties.

## Request Body
Content-Type: `application/json`
[]object:
  - `entityId` (required): number
  - `key` (required): string
  - `value` (required): string

## Responses
- 200: Returned if the request is successful.
- 400: Returned if the request is not valid.
- 403: Returned if the authorisation credentials are incorrect or missing.
