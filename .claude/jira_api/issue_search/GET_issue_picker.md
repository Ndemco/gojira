# GET /rest/api/3/issue/picker
**operationId:** `getIssuePickerResource`
**Summary:** Get issue picker suggestions

Returns lists of issues matching a query string. Use this resource to provide auto-completion suggestions when the user is looking for an issue using a word or string.

This operation returns two lists:

 *  `History Search` which includes issues from the user's history of created, edited, or viewed issues that contain the string in the `query` parameter.
 *  `Current Search` which includes issues that match the JQL expression in `currentJQL` and contain the string in the `query` parameter.

Thi

## Parameters
- `query` [query] string — A string to match against text fields in the issue such as title, description, or comments.
- `currentJQL` [query] string — A JQL query defining a list of issues to search for the query term. Note that `username` and `userkey` cannot be used as
- `currentIssueKey` [query] string — The key of an issue to exclude from search results. For example, the issue the user is viewing when they perform this qu
- `currentProjectId` [query] string — The ID of a project that suggested issues must belong to.
- `showSubTasks` [query] boolean — Indicate whether to include subtasks in the suggestions list.
- `showSubTaskParent` [query] boolean — When `currentIssueKey` is a subtask, whether to include the parent issue in the suggestions if it matches the query.

## Responses
- 200: object:
  - `sections`: []IssuePickerSuggestionsIssueType
- 401: Returned if the authentication credentials are incorrect or missing.
