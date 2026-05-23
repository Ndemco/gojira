# GET /rest/api/3/field
**operationId:** `getFields`
**Summary:** Get fields

Returns system and custom issue fields according to the following rules:

 *  Fields that cannot be added to the issue navigator are always returned.
 *  Fields that cannot be placed on an issue screen are always returned.
 *  Fields that depend on global Jira settings are only returned if the setting is enabled. That is, timetracking fields, subtasks, votes, and watches.
 *  Fields that are not associated to any used field configurations or screens are not returned.
 *  For all other fields, th

## Responses
- 200: []object:
  - `clauseNames`: []string
  - `custom`: boolean
  - `id`: string
  - `key`: string
  - `name`: string
  - `navigable`: boolean
  - `orderable`: boolean
  - `schema`: allOf(JsonTypeBean)
  - `scope`: allOf(Scope)
  - `searchable`: boolean
- 401: Returned if the authentication credentials are incorrect or missing.
