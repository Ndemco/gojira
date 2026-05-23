# GET /rest/api/3/fieldconfigurationscheme/mapping
**operationId:** `getFieldConfigurationSchemeMappings`
**Summary:** Get field configuration issue type items

Deprecated, use [ Field schemes](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-field-schemes/#api-group-field-schemes) which supports field association schemes.

Returns a [paginated](#pagination) list of field configuration issue type items.

Only items used in classic projects are returned.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `fieldConfigurationSchemeId` [query] []integer(int64) — The list of field configuration scheme IDs. To include multiple field configuration schemes separate IDs with ampersand:

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []FieldConfigurationIssueTypeItem
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if no field configuration schemes are found.
