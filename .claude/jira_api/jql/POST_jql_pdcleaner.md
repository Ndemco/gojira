# POST /rest/api/3/jql/pdcleaner
**operationId:** `migrateQueries`
**Summary:** Convert user identifiers to account IDs in JQL queries

Converts one or more JQL queries with user identifiers (username or user key) to equivalent JQL queries with account IDs.

You may wish to use this operation if your system stores JQL queries and you want to make them GDPR-compliant. For more information about GDPR-related changes, see the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/).

**[Permissions](#permissions) required:** Permission to access Jira.

## Request Body
Content-Type: `application/json`
object:
  - `queryStrings`: []string

## Responses
- 200: object:
  - `queriesWithUnknownUsers`: []JQLQueryWithUnknownUsers
  - `queryStrings`: []string
- 400: string
- 401: string
