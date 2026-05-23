# GET /rest/api/3/bulk/issues/fields
**operationId:** `getBulkEditableFields`
**Summary:** Get bulk editable fields

Use this API to get a list of fields visible to the user to perform bulk edit operations. You can pass single or multiple issues in the query to get eligible editable fields. This API uses pagination to return responses, delivering 50 fields at a time.

**[Permissions](#permissions) required:**

 *  Global bulk change [permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-global-permissions/).
 *  Browse [project permission](https://support.atlassian.com/jira-cloud-admi

## Parameters
- `issueIdsOrKeys` [query] (required) string — The IDs or keys of the issues to get editable fields from.
- `searchText` [query] string — (Optional)The text to search for in the editable fields.
- `endingBefore` [query] string — (Optional)The end cursor for use in pagination.
- `startingAfter` [query] string — (Optional)The start cursor for use in pagination.

## Responses
- 200: object:
  - `endingBefore`: string
  - `fields`: []IssueBulkEditField
  - `startingAfter`: string
- 400: object:
  - `errors`: []ErrorMessage
- 401: object:
  - `errors`: []ErrorMessage
- 403: object:
  - `errors`: []ErrorMessage
- 404: object:
  - `errors`: []ErrorMessage
