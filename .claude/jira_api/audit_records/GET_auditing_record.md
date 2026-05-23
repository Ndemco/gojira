# GET /rest/api/3/auditing/record
**operationId:** `getAuditRecords`
**Summary:** Get audit records

Returns a list of audit records. The list can be filtered to include items:

 *  where each item in `filter` has at least one match in any of these fields:
    
     *  `summary`
     *  `category`
     *  `eventSource`
     *  `objectItem.name` If the object is a user, account ID is available to filter.
     *  `objectItem.parentName`
     *  `objectItem.typeName`
     *  `changedValues.changedFrom`
     *  `changedValues.changedTo`
     *  `remoteAddress`
    
    For example, if `filter` cont

## Parameters
- `offset` [query] integer(int32) — The number of records to skip before returning the first result.
- `limit` [query] integer(int32) — The maximum number of results to return.
- `filter` [query] string — The strings to match with audit field content, space separated.
- `from` [query] string — The date and time on or after which returned audit records must have been created. If `to` is provided `from` must be be
- `to` [query] string — The date and time on or before which returned audit results must have been created. If `from` is provided `to` must be a

## Responses
- 200: object:
  - `limit`: integer(int32)
  - `offset`: integer(int32)
  - `records`: []AuditRecordBean
  - `total`: integer(int64)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
