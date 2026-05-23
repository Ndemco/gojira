# GET /rest/api/3/issue/limit/report
**operationId:** `getIssueLimitReport`
**Summary:** Get issue limit report

Returns all issues breaching and approaching per-issue limits.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) is required for the project the issues are in. Results may be incomplete otherwise
 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `isReturningKeys` [query] boolean — Return issue keys instead of issue ids in the response.

## Responses
- 200: object:
  - `issuesApproachingLimit`: object
  - `issuesBreachingLimit`: object
  - `limits`: object
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to complete this request.
