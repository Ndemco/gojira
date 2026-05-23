# PUT /rest/api/3/issue/unarchive
**operationId:** `unarchiveIssues`
**Summary:** Unarchive issue(s) by issue keys/ID

Enables admins to unarchive up to 1000 issues in a single request using issue ID/key, returning details of the issue(s) unarchived in the process and the errors encountered, if any.

**Note that:**

 *  you can't unarchive subtasks directly, only through their parent issues
 *  you can only unarchive issues from software, service management, and business projects

**[Permissions](#permissions) required:** Jira admin or site admin: [global permission](https://confluence.atlassian.com/x/x4dKLg)

*

## Request Body
Content-Type: `application/json`
object:
  - `issueIdsOrKeys`: []string

## Responses
- 200: object:
  - `errors`: Errors
  - `numberOfIssuesUpdated`: integer(int64)
- 400: any
- 401: any
- 403: any
- 412: any
