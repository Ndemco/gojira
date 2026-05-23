# POST /rest/api/3/uiModifications
**operationId:** `createUiModification`
**Summary:** Create UI modification

Creates a UI modification. UI modification can only be created by Forge apps.

Each app can define up to 3000 UI modifications. Each UI modification can define up to 1000 contexts. The same context can be assigned to maximum 100 UI modifications.

**Context types:**

 *  **Jira contexts:** For Jira view types, use `projectId` and `issueTypeId`. One field can act as a wildcard. Supported Jira views:
    
     *  `GIC` \- Jira global issue create
     *  `IssueView` \- Jira issue view
     *  `Iss

## Request Body
Content-Type: `application/json`
object:
  - `contexts`: []UiModificationContextDetails
  - `data`: string
  - `description`: string
  - `name` (required): string

## Responses
- 201: object:
  - `id` (required): string
  - `self` (required): string
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the request is not from a Forge app.
- 404: object:
  - `details`: object
  - `errorMessages`: []string
  - `errors`: object
