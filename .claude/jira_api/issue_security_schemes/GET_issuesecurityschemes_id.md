# GET /rest/api/3/issuesecurityschemes/{id}
**operationId:** `getIssueSecurityScheme`
**Summary:** Get issue security scheme

Returns an issue security scheme along with its security levels.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).
 *  *Administer Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for a project that uses the requested issue security scheme.

## Parameters
- `id` [path] (required) integer(int64) — The ID of the issue security scheme. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) oper

## Responses
- 200: object:
  - `defaultSecurityLevelId`: integer(int64)
  - `description`: string
  - `id`: integer(int64)
  - `levels`: []SecurityLevel
  - `name`: string
  - `self`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the administrator permission and the scheme is not used in any project where the user has administrative permission.
