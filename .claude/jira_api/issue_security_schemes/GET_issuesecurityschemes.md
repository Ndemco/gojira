# GET /rest/api/3/issuesecurityschemes
**operationId:** `getIssueSecuritySchemes`
**Summary:** Get issue security schemes

Returns all [issue security schemes](https://confluence.atlassian.com/x/J4lKLg).

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: object:
  - `issueSecuritySchemes`: []SecurityScheme
- 401: Returned if the authentication credentials are incorrect.
- 403: Returned if the user does not have permission to administer issue security schemes.
