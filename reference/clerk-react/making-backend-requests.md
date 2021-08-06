---
description: >-
  This guide helps you configure AJAX requests so they can be authenticated by
  your backend.
---

# Making backend requests

## Request requirements

### 1. Session management

#### Single-session application

In a single-session application, users can be signed into only one account at a time. Clerk handles session management automatically.

#### Multi-session application

In a multi-session application, users can be signed into multiple accounts simultaneously, and can switch between accounts using the user button. Your frontend must specify the active session ID when it makes requests to your backend. This ID should be passed through a query string parameter called **\_clerk\_session\_id**, which our SDKs are configured to read.

To retrieve the active session ID, call the useSession hook.

```jsx
import { useSession } from '@clerk/clerk-react';

const session = useSession();
const sessionId = session?.id;
```

### 2. Include cookies

This is automatic if your frontend and backend have the [same origin](https://developer.mozilla.org/en-US/docs/Web/Security/Same-origin_policy), but requires additional configuration for cross-origin requests.

In addition to following the examples below, you must configure your backend to respond with a special header for [Cross-Origin Resource Sharing](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) \(CORS\):

```http
Access-Control-Allow-Credentials: true
```

### 3. Use HTTPS in production

For production instances, your backend must have SSL configured.

## Examples

The examples below assume there is a variable in-scope named **sessionId** which contains the [active session ID](making-backend-requests.md#1-set-__clerk_session_id).

### fetch

The [fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch) is automatically included in modern browsers.

{% tabs %}
{% tab title="Same-origin" %}
```jsx
const reqUrl = `/foo?_clerk_session_id=${sessionId}`
const res = await fetch(reqUrl)
```
{% endtab %}

{% tab title="Cross-origin" %}
```javascript
const baseUrl = 'https://backend.yourdomain.com'
const reqUrl = `${baseUrl}foo?_clerk_session_id=${sessionId}`
const res = await fetch(reqUrl, {
  credentials: "include", // Include cookies
})
```
{% endtab %}
{% endtabs %}

### Axios

[Axios](https://github.com/axios/axios) is a third-party library that provides a more convenient wrapper over XHR. It must be imported before it can be used.

{% tabs %}
{% tab title="Same-origin" %}
```jsx
const reqUrl = `foo?_clerk_session_id=${sessionId}`
const res = await axios.get(reqURL);
```
{% endtab %}

{% tab title="Cross-origin" %}
```javascript
// Create a reusable axios instance
const backend = axios.create({
  baseURL: 'https://backend.yourdomain.com',
  withCredentials: true, // Include cookies
});

// Make the request
const reqUrl = `/foo?_clerk_session_id=${sessionId}`
const res = await backend.get(reqURL);
```
{% endtab %}
{% endtabs %}

### XHR

XHR is automatically included in modern browsers, but the API is quite tedious. While we have included the example here for illustrative purposes, we recommend using [fetch](making-backend-requests.md#fetch) or [Axios](making-backend-requests.md#axios) instead of XHR directly.

{% tabs %}
{% tab title="Same-origin" %}
```javascript
const reqUrl = `/foo?_clerk_session_id=${sessionId}`

let xhr = new XMLHttpRequest();
xhr.open('GET', reqUrl);
xhr.send();

xhr.onload = function() {
  const res = xhr.response
};

xhr.onerror = function() {
  alert("Request failed");
};
```
{% endtab %}
{% endtabs %}

