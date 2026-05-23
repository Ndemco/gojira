# GET /rest/api/3/application-properties
**operationId:** `getApplicationProperty`
**Summary:** Get application property

Returns all application properties or an application property.

If you specify a value for the `key` parameter, then an application property is returned as an object (not in an array). Otherwise, an array of all editable application properties is returned. See [Set application property](#api-rest-api-3-application-properties-id-put) for descriptions of editable properties.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `key` [query] string — The key of the application property.
- `permissionLevel` [query] string — The permission level of all items being returned in the list.
- `keyFilter` [query] string — When a `key` isn't provided, this filters the list of results by the application property `key` using a regular expressi

## Responses
- 200: []object:
  - `allowedValues`: []string
  - `defaultValue`: string
  - `desc`: string
  - `example`: string
  - `id`: string
  - `key`: string
  - `name`: string
  - `type`: string
  - `value`: string
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
