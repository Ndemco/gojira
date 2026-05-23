# GET /rest/api/3/app/field/{fieldIdOrKey}/context/configuration
**operationId:** `getCustomFieldConfiguration`
**Summary:** Get custom field configurations

Returns a [paginated](#pagination) list of configurations for a custom field of a [type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field-type/) created by a [Forge app](https://developer.atlassian.com/platform/forge/).

The result can be filtered by one of these criteria:

 *  `id`.
 *  `fieldContextId`.
 *  `issueId`.
 *  `projectKeyOrId` and `issueTypeId`.

Otherwise, all configurations are returned.

**[Permissions](#permissions) required:** *Admini

## Parameters
- `fieldIdOrKey` [path] (required) string — The ID or key of the custom field, for example `customfield_10000`.
- `id` [query] []integer(int64) — The list of configuration IDs. To include multiple configurations, separate IDs with an ampersand: `id=10000&id=10001`. 
- `fieldContextId` [query] []integer(int64) — The list of field context IDs. To include multiple field contexts, separate IDs with an ampersand: `fieldContextId=10000
- `issueId` [query] integer(int64) — The ID of the issue to filter results by. If the issue doesn't exist, an empty list is returned. Can't be provided with 
- `projectKeyOrId` [query] string — The ID or key of the project to filter results by. Must be provided with `issueTypeId`. Can't be provided with `issueId`
- `issueTypeId` [query] string — The ID of the issue type to filter results by. Must be provided with `projectKeyOrId`. Can't be provided with `issueId`.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []ContextualConfiguration
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user is not a Jira admin or the request is not authenticated as from the app that provided the field.
- 404: Returned if the custom field is not found.
