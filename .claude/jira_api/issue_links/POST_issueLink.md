# POST /rest/api/3/issueLink
**operationId:** `linkIssues`
**Summary:** Create issue link

Creates a link between two issues. Use this operation to indicate a relationship between two issues and optionally add a comment to the from (outward) issue. To use this resource the site must have [Issue Linking](https://confluence.atlassian.com/x/yoXKM) enabled.

This resource returns nothing on the creation of an issue link. To obtain the ID of the issue link, use `https://your-domain.atlassian.net/rest/api/3/issue/[linked issue key]?fields=issuelinks`.

If the link request duplicates a link,

## Request Body
Content-Type: `application/json`
object:
  - `comment`: Comment
  - `inwardIssue` (required): LinkedIssue
  - `outwardIssue` (required): LinkedIssue
  - `type` (required): IssueLinkType

## Responses
- 201: any
- 400: Returned if the comment is not created. The response contains an error message indicating why the comment wasn't created. The issue link is also not created.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  issue linking is disabled.
 *  the user cannot view one or both of the issues. For example, the user doesn't have *Browse project* project permission for a project containing one of the issues.
 *  the user does not have *link issues* project permission.
 *  either of the link issues are not found.
 *  the issue link type is not found.
- 413: Returned if the per-issue limit for issue links has been breached.
