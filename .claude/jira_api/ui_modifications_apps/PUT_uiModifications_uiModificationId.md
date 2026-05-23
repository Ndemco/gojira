# PUT /rest/api/3/uiModifications/{uiModificationId}
**operationId:** `updateUiModification`
**Summary:** Update UI modification

Updates a UI modification. UI modification can only be updated by Forge apps.

Each UI modification can define up to 1000 contexts. The same context can be assigned to maximum 100 UI modifications.

**Context types:**

 *  **Jira contexts:** For Jira view types, use `projectId` and `issueTypeId`. One field can act as a wildcard. Supported Jira views:
    
     *  `GIC` \- Jira global issue create
     *  `IssueView` \- Jira issue view
     *  `IssueTransition` \- Jira issue transition
 *  **Jira

## Parameters
- `uiModificationId` [path] (required) string — The ID of the UI modification.

## Request Body
Content-Type: `application/json`
object:
  - `contexts`: []UiModificationContextDetails
  - `data`: string
  - `description`: string
  - `name`: string

## Responses
- 204: any
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the request is not from a Forge app.
- 404: object:
  - `details`: object
  - `errorMessages`: []string
  - `errors`: object
