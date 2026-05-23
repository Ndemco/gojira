# GET /rest/api/3/jql/autocompletedata/suggestions
**operationId:** `getFieldAutoCompleteForQueryString`
**Summary:** Get field auto complete suggestions

Returns the JQL search auto complete suggestions for a field.

Suggestions can be obtained by providing:

 *  `fieldName` to get a list of all values for the field.
 *  `fieldName` and `fieldValue` to get a list of values containing the text in `fieldValue`.
 *  `fieldName` and `predicateName` to get a list of all predicate values for the field.
 *  `fieldName`, `predicateName`, and `predicateValue` to get a list of predicate values containing the text in `predicateValue`.

This operation can be

## Parameters
- `fieldName` [query] string — The name of the field.
- `fieldValue` [query] string — The partial field item name entered by the user.
- `predicateName` [query] string — The name of the [ CHANGED operator predicate](https://confluence.atlassian.com/x/hQORLQ#Advancedsearching-operatorsrefer
- `predicateValue` [query] string — The partial predicate item name entered by the user.

## Responses
- 200: object:
  - `results`: []AutoCompleteSuggestion
- 400: Returned if an invalid combination of parameters is passed.
- 401: Returned if the authentication credentials are incorrect.
