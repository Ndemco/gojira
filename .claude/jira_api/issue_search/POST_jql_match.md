# POST /rest/api/3/jql/match
**operationId:** `matchIssues`
**Summary:** Check issues against JQL

Checks whether one or more issues would be returned by one or more JQL queries.

**[Permissions](#permissions) required:** None, however, issues are only matched against JQL queries where the user has:

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.

## Request Body
Content-Type: `application/json`
object:
  - `issueIds` (required): []integer(int64)
  - `jqls` (required): []string

## Responses
- 200: object:
  - `matches` (required): []IssueMatchesForJQL
- 400: Returned if `jqls` exceeds the maximum number of JQL queries or `issueIds` exceeds the maximum number of issue IDs.
