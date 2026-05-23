# POST /rest/api/3/redact
**operationId:** `redact`
**Summary:** Redact

Submit a job to redact issue field data. This will trigger the redaction of the data in the specified fields asynchronously.

The redaction status can be polled using the job id.

## Request Body
Content-Type: `application/json`
object:
  - `redactions`: []SingleRedactionRequest

## Responses
- 202: string(uuid)
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the user / app is not authorised to redact data
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
