---
description: Manage the redirect URLs whitelist for each instance
---

# Redirect URLs

Redirect URLs endpoints are used to whitelist URLs for native application authentication flows such as OAuth sign-ins and sign-ups in [React Native](../clerk-react/) and [Expo](../clerk-expo.md).

For native applications, Clerk ensures that security critical nonces will be passed only to whitelisted URLs when the OAuth flow is complete in native browsers or webviews.

### Available requests

* `POST /v1/redirect_urls`
* `GET /v1/redirect_urls`
* `GET /v1/redirect_urls/:id`
* `DELETE /v1/redirect_urls/:id`

### Example redirect URL schema

```json
{
  "object": "redirect_url",
  "id": "ru_28eW1GeqywLZSzoBoaHJ79pkPR6",
  "url": "my-app://oauth-callback",
  "created_at": 1651577314987,
  "updated_at": 1651577314987
}
```

{% swagger method="post" path="/v1/redirect_urls" baseUrl="https://api.clerk.dev" summary="Add a redirect URL" %}
{% swagger-description %}
Add a new redirect URL to the whitelist of the instance.
{% endswagger-description %}

{% swagger-parameter in="body" name="url" type="url" required="true" %}
The full url value
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
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

{% swagger-response status="422: Unprocessable Entity" description="Invalid url value" %}
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
{% endswagger %}

{% swagger method="get" path="/v1/redirect_urls" baseUrl="https://api.clerk.dev" summary="List all redirect urls" %}
{% swagger-description %}
Lists all whitelisted redirect_urls for the instance
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="A list of redirect urls" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/v1/redirect_urls/:id" baseUrl="https://api.clerk.dev" summary="Retrieve a redirect URL" %}
{% swagger-description %}
Retrieve the details of a redirect URL
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="String" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="path" required="true" %}
The id of the redirect URL
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="delete" path="/v1/redirect_urls/:id" baseUrl="https://api.clerk.dev" summary="Delete a redirect URL" %}
{% swagger-description %}
Remove the selected redirect URL from the whitelist of the instance
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="path" required="true" %}
The id of the redirect URL
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}
