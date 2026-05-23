# POST /rest/api/3/webhook
**operationId:** `registerDynamicWebhooks`
**Summary:** Register dynamic webhooks

Registers webhooks.

**NOTE:** for non-public OAuth apps, webhooks are delivered only if there is a match between the app owner and the user who registered a dynamic webhook.

**[Permissions](#permissions) required:** Only [Connect](https://developer.atlassian.com/cloud/jira/platform/#connect-apps) and [OAuth 2.0](https://developer.atlassian.com/cloud/jira/platform/oauth-2-3lo-apps) apps can use this operation.

## Request Body
Content-Type: `application/json`
object:
  - `url` (required): string
  - `webhooks` (required): []WebhookDetails

## Responses
- 200: object:
  - `webhookRegistrationResult`: []RegisteredWebhook
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
