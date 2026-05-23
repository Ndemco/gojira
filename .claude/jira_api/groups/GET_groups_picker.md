# GET /rest/api/3/groups/picker
**operationId:** `findGroups`
**Summary:** Find groups

Returns a list of groups whose names contain a query string. A list of group names can be provided to exclude groups from the results.

The primary use case for this resource is to populate a group picker suggestions list. To this end, the returned object includes the `html` field where the matched query term is highlighted in the group name with the HTML strong tag. Also, the groups list is wrapped in a response object that contains a header for use in the picker, specifically *Showing X of Y m

## Parameters
- `accountId` [query] string — This parameter is deprecated, setting it does not affect the results. To find groups containing a particular user, use [
- `query` [query] string — The string to find in group names.
- `exclude` [query] []string — As a group's name can change, use of `excludeGroupIds` is recommended to identify a group.  
- `excludeId` [query] []string — A group ID to exclude from the result. To exclude multiple groups, provide an ampersand-separated list. For example, `ex
- `maxResults` [query] integer(int32) — The maximum number of groups to return. The maximum number of groups that can be returned is limited by the system prope
- `caseInsensitive` [query] boolean — Whether the search for groups should be case insensitive.
- `userName` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/

## Responses
- 200: object:
  - `groups`: []FoundGroup
  - `header`: string
  - `total`: integer(int32)
