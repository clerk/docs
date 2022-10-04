---
description: Allow your users to sign in on behalf of other users
---

# Actor Tokens

Actor tokens are JWT tokens generated from the Backend API and are linked to a specific user.&#x20;

These tokens can be used to allow users to sign in on behalf of other users. These are particularly useful for features like user impersonation.

More specifically, an actor token has an actor who is the user who wants to sign in, and a target user who is the user whose account the actor user will have once they sign in.&#x20;

Actor tokens bypass all factors during sign in, even any second factors that the target user might have set up.

Given that actor tokens are particularly powerful, we have setup some limitations to make sure we limit any accidents as much as possible.

* An actor token can only be used once
* Creating a session with an actor token will terminate all other sessions on the same browser
* Sessions created via an actor token have an inactivity timeout of 10 minutes, and a fixed duration that can be configured when creating the actor token (by default it's 30 minutes)

After a token is generated from the Backend API, it can be used in the sign in object via the ticket strategy, i.e. using `ticket` as `strategy` and passing the generated token in the `ticket` property.&#x20;

### Available requests

* `POST /v1/actor_tokens`
* `POST /v1/actor_tokens/:id/revoke`

### Example actor token schema

```json
{
    "object": "actor_token",
    "id": "act_26Ed5ZqqJcOjRwecRQij2ZovDdG",
    "user_id": "user_26Ect5GuCCeaFWwSDiiKcgAGtVk",
    "actor": {
        "sub": "actor_id"
    },
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzSW5TZWNvbmRzIjo1LCJleHAiOjE2NDY5OTI1MDEsImlpZCI6Imluc18yNkVja3R0TnJDamE3YTZQT0xINTVDQVBpZmQiLCJzaWQiOiJzaXRfMjZFZDVacXFKY09qUndlY1JRaWoyWm92RGRHIiwic3QiOiJzaWduX2luX3Rva2VuIn0.j6Gwl6g2QcAJ9AjRvG1k7aUrnMCyPU49hYgTlmDG9gD_8Yd7sxUepyDdCHRaDaABlWg-G3tUs09HRfdrAXM-4e6NwcEy_ak1LWkE3G6WVhPnlomwH7n7BsIbmoybf91Eel0XRlb33XdUVaWNaA_CH8INkVLtXfZWTorNsAN2-Es_6G-Jtz4Zvw8hZBtXQDMSlyl27rxohMvfefv-ffG6Kd0XsvT9yYj2kik5KcONMWO6XEPtMZRoHzMabnmPQbLrUPBmbnU_1UVFpxL0LfuOXlxbV3LIvuejmhNZZtR0ZwcbrAnXruof4KjmCK_QOpqShI3dTlyYTV18amy2se5oxA",
    "status": "pending",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```

{% swagger method="post" path="/v1/actor_tokens" baseUrl="https://api.clerk.dev" summary="Create actor token" %}
{% swagger-description %}
Creates a new actor token for the given actor and target user.&#x20;

The actor payload accepts any JSON payload, which will end up in the session JWT token. The only requirement is that it includes a `sub` property, which will be the id of the actor user.

By default, actor tokens expire in 1 hour. You can optionally supply a different duration in seconds using the `expires_in_seconds` property.
{% endswagger-description %}

{% swagger-parameter in="body" name="user_id" type="string" required="true" %}
The id of the user that can use the newly created sign in token.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="actor" type="object" required="true" %}
The actor payload. It needs to include a `sub` property which should contain the id of the actor.

This whole payload will be also included in the JWT session token.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="expires_in_seconds" type="int" %}
Optional parameter to specify the life duration of the actor token in seconds.



By default, the duration is 1 hour.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="session_max_duration_in_seconds" type="int" %}
The maximum duration that the session which will be created by the generated actor token should last.

By default, the duration of a session created via an actor token, lasts 30 minutes.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="A new actor token has been created" %}
```javascript
{
    "object": "actor_token",
    "id": "act_26Ed5ZqqJcOjRwecRQij2ZovDdG",
    "user_id": "user_26Ect5GuCCeaFWwSDiiKcgAGtVk",
    "actor": {
        "sub": "actor_id"
    },
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

{% swagger-response status="422: Unprocessable Entity" description="Actor id is missing from the payload" %}
```javascript
{
  "errors": [
    {
      "code": "form_param_missing",
      "long_message": "actor.sub must be included.",
      "message": "is missing",
      "meta": {
        "param_name": "actor.sub"
      }
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

{% swagger-response status="422: Unprocessable Entity" description="Non-positive value for maximum session duration" %}
```javascript
{
  "errors": [
    {
      "code": "form_param_value_invalid",
      "long_message": "session_max_duration_in_seconds is invalid",
      "message": "is invalid",
      "meta": {
        "param_name": "session_max_duration_in_seconds"
      }
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="post" path="/v1/actor_tokens/:id/revoke" baseUrl="https://api.clerk.dev" summary="Revoke sign in token" %}
{% swagger-description %}
Revokes a pending actor token.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}
The id of the actor token to be revoked.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Actor token has been revoked successfully" %}
```javascript
{
    "object": "actor_token",
    "id": "act_26Ed5ZqqJcOjRwecRQij2ZovDdG",
    "user_id": "user_26Ect5GuCCeaFWwSDiiKcgAGtVk",
    "actor": {
        "sub": "actor_id"
    },
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzSW5TZWNvbmRzIjo1LCJleHAiOjE2NDY5OTI1MDEsImlpZCI6Imluc18yNkVja3R0TnJDamE3YTZQT0xINTVDQVBpZmQiLCJzaWQiOiJzaXRfMjZFZDVacXFKY09qUndlY1JRaWoyWm92RGRHIiwic3QiOiJzaWduX2luX3Rva2VuIn0.j6Gwl6g2QcAJ9AjRvG1k7aUrnMCyPU49hYgTlmDG9gD_8Yd7sxUepyDdCHRaDaABlWg-G3tUs09HRfdrAXM-4e6NwcEy_ak1LWkE3G6WVhPnlomwH7n7BsIbmoybf91Eel0XRlb33XdUVaWNaA_CH8INkVLtXfZWTorNsAN2-Es_6G-Jtz4Zvw8hZBtXQDMSlyl27rxohMvfefv-ffG6Kd0XsvT9yYj2kik5KcONMWO6XEPtMZRoHzMabnmPQbLrUPBmbnU_1UVFpxL0LfuOXlxbV3LIvuejmhNZZtR0ZwcbrAnXruof4KjmCK_QOpqShI3dTlyYTV18amy2se5oxA",
    "status": "revoked",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}

```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="Token cannot be revoked because it's not in pending state" %}
```javascript
{
  "errors": [
    {
      "code": "actor_token_cannot_be_revoked_code",
      "long_message": "Actor token cannot be revoked because its status is .... Only pending tokens can be revoked.",
      "message": "cannot revoke"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="Given actor token was not found" %}
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
