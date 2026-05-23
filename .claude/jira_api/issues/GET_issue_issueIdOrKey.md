# GET /rest/api/3/issue/{issueIdOrKey}
**operationId:** `getIssue`
**Summary:** Get issue

Returns the details for an issue.

The issue is identified by its ID or key, however, if the identifier doesn't match an issue, a case-insensitive search and check for moved issues is performed. If a matching issue is found its details are returned, a 302 or other redirect is **not** returned. The issue key returned in the response is the key of the issue found.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](htt

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `fields` [query] []string — A list of fields to return for the issue. This parameter accepts a comma-separated list. Use it to retrieve a subset of 
- `fieldsByKeys` [query] boolean — Whether fields in `fields` are referenced by keys rather than IDs. This parameter is useful where fields have been added
- `expand` [query] string — Use [expand](#expansion) to include additional information about the issues in the response. This parameter accepts a co
- `properties` [query] []string — A list of issue properties to return for the issue. This parameter accepts a comma-separated list. Allowed values:
- `updateHistory` [query] boolean — Whether the project in which the issue is created is added to the user's **Recently viewed** project list, as shown unde
- `failFast` [query] boolean — Whether to fail the request quickly in case of an error while loading fields for an issue. For `failFast=true`, if one f

## Responses
- 200: object:
  - `changelog`: allOf(PageOfChangelogs)
  - `editmeta`: allOf(IssueUpdateMetadata)
  - `expand`: string
  - `fields`: object
  - `fieldsToInclude`: IncludedFields
  - `id`: string
  - `key`: string
  - `names`: object
  - `operations`: allOf(Operations)
  - `properties`: object
  - `renderedFields`: object
  - `schema`: object
  - `self`: string(uri)
  - `transitions`: []IssueTransition
  - `versionedRepresentations`: object
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue is not found or the user does not have permission to view it.
