---
description: Allow your users to sign in using tokens
---

# Sign In Tokens

Sign in tokens are JWT tokens generated from the Backend API and are linked to a specific user.&#x20;

These tokens can be used to sign in users. More specifically, using a sign in token, a user can sign in without having to specify his credentials.

As mentioned above, the sign in tokens are linked to particular users and can only be used by them. A sign in token can be used exactly once.

After a token is generated from the Backend API, it can be used in the sign in object via the ticket strategy, i.e. using `ticket` as `strategy` and passing the generated token in the `ticket` property.&#x20;

### Available requests

* `POST /v1/sign_in_tokens`
* `POST /v1/sign_in_tokens/:id/revoke`

### Example sign in token schema

```json
{
    "object": "sign_in_token",
    "id": "sit_26Ed5ZqqJcOjRwecRQij2ZovDdG",
    "user_id": "user_26Ect5GuCCeaFWwSDiiKcgAGtVk",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzSW5TZWNvbmRzIjo1LCJleHAiOjE2NDY5OTI1MDEsImlpZCI6Imluc18yNkVja3R0TnJDamE3YTZQT0xINTVDQVBpZmQiLCJzaWQiOiJzaXRfMjZFZDVacXFKY09qUndlY1JRaWoyWm92RGRHIiwic3QiOiJzaWduX2luX3Rva2VuIn0.j6Gwl6g2QcAJ9AjRvG1k7aUrnMCyPU49hYgTlmDG9gD_8Yd7sxUepyDdCHRaDaABlWg-G3tUs09HRfdrAXM-4e6NwcEy_ak1LWkE3G6WVhPnlomwH7n7BsIbmoybf91Eel0XRlb33XdUVaWNaA_CH8INkVLtXfZWTorNsAN2-Es_6G-Jtz4Zvw8hZBtXQDMSlyl27rxohMvfefv-ffG6Kd0XsvT9yYj2kik5KcONMWO6XEPtMZRoHzMabnmPQbLrUPBmbnU_1UVFpxL0LfuOXlxbV3LIvuejmhNZZtR0ZwcbrAnXruof4KjmCK_QOpqShI3dTlyYTV18amy2se5oxA",
    "status": "pending",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```

{% swagger method="post" path="/v1/sign_in_tokens" baseUrl="https://api.clerk.dev" summary="Create sign in token" %}
{% swagger-description %}
Creates a new sign in token and associates it with the given user.

By default, sign in tokens expire in 30 days. You can optionally supply a different duration in seconds using the `expires_in_seconds` property.
{% endswagger-description %}

{% swagger-parameter in="body" name="user_id" type="string" required="true" %}
The id of the user that can use the newly created sign in token.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="expires_in_seconds" type="int" %}
Optional parameter to specify the life duration of the sign in token in seconds.



By default, the duration is 30 days.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="A new sign in token has been created" %}
```javascript
{
    "object": "sign_in_token",
    "id": "sit_26Ed5ZqqJcOjRwecRQij2ZovDdG",
    "user_id": "user_26Ect5GuCCeaFWwSDiiKcgAGtVk",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzSW5TZWNvbmRzIjo1LCJleHAiOjE2NDY5OTI1MDEsImlpZCI6Imluc18yNkVja3R0TnJDamE3YTZQT0xINTVDQVBpZmQiLCJzaWQiOiJzaXRfMjZFZDVacXFKY09qUndlY1JRaWoyWm92RGRHIiwic3QiOiJzaWduX2luX3Rva2VuIn0.j6Gwl6g2QcAJ9AjRvG1k7aUrnMCyPU49hYgTlmDG9gD_8Yd7sxUepyDdCHRaDaABlWg-G3tUs09HRfdrAXM-4e6NwcEy_ak1LWkE3G6WVhPnlomwH7n7BsIbmoybf91Eel0XRlb33XdUVaWNaA_CH8INkVLtXfZWTorNsAN2-Es_6G-Jtz4Zvw8hZBtXQDMSlyl27rxohMvfefv-ffG6Kd0XsvT9yYj2kik5KcONMWO6XEPtMZRoHzMabnmPQbLrUPBmbnU_1UVFpxL0LfuOXlxbV3LIvuejmhNZZtR0ZwcbrAnXruof4KjmCK_QOpqShI3dTlyYTV18amy2se5oxA",
    "status": "pending",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="Given user id does not correspond to an existing user" %}
```javascript
{
  "errors": [
    {
      "code": "resource_not_found",
      "long_message": "No user was found with id ...",
      "message": "not found"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="User id is missing from the request payload" %}
```javascript
{
  "errors": [
    {
      "code": "form_param_missing",
      "long_message": "user_id must be included.",
      "message": "is missing",
      "meta": {
        "param_name": "user_id"
      }
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="Non-positive value for expiration duration" %}
```javascript
{
  "errors": [
    {
      "code": "form_param_value_invalid",
      "long_message": "expires_in_seconds is invalid",
      "message": "is invalid",
      "meta": {
        "param_name": "expires_in_seconds"
      }
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="post" path="/v1/sign_in_tokens/:id/revoke" baseUrl="https://api.clerk.dev" summary="Revoke sign in token" %}
{% swagger-description %}
Revokes a pending sign in token.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}
The id of the sign in token to be revoked.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Sign in token has been revoked successfully" %}
```javascript
{
    "object": "sign_in_token",
    "id": "sit_26Ed5ZqqJcOjRwecRQij2ZovDdG",
    "user_id": "user_26Ect5GuCCeaFWwSDiiKcgAGtVk",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzSW5TZWNvbmRzIjo1LCJleHAiOjE2NDY5OTI1MDEsImlpZCI6Imluc18yNkVja3R0TnJDamE3YTZQT0xINTVDQVBpZmQiLCJzaWQiOiJzaXRfMjZFZDVacXFKY09qUndlY1JRaWoyWm92RGRHIiwic3QiOiJzaWduX2luX3Rva2VuIn0.j6Gwl6g2QcAJ9AjRvG1k7aUrnMCyPU49hYgTlmDG9gD_8Yd7sxUepyDdCHRaDaABlWg-G3tUs09HRfdrAXM-4e6NwcEy_ak1LWkE3G6WVhPnlomwH7n7BsIbmoybf91Eel0XRlb33XdUVaWNaA_CH8INkVLtXfZWTorNsAN2-Es_6G-Jtz4Zvw8hZBtXQDMSlyl27rxohMvfefv-ffG6Kd0XsvT9yYj2kik5KcONMWO6XEPtMZRoHzMabnmPQbLrUPBmbnU_1UVFpxL0LfuOXlxbV3LIvuejmhNZZtR0ZwcbrAnXruof4KjmCK_QOpqShI3dTlyYTV18amy2se5oxA",
    "status": "revoked",
    "created_at": 1638000669544,
    "updated_at": 1638000669544

```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="Token cannot be revoked because it's not in pending state" %}
```javascript
{
  "errors": [
    {
      "code": "sign_in_token_cannot_be_revoked_code",
      "long_message": "Sign in token cannot be revoked because its status is .... Only pending tokens can be revoked.",
      "message": "cannot revoke"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="Given sign in token was not found" %}
```javascript
{
  "errors": [
    {
      "code": "resource_not_found",
      "long_message": "Resource not found",
      "message": "not found"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}
