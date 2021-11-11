# Instance Settings

Modify your instance's settings.

### Available requests

* **`PATCH`**`/v1/beta_features/instance_settings`

{% swagger baseUrl="https://api.clerk.dev" path="/v1/beta_features/instance_settings" method="patch" summary="Update instance settings" %}
{% swagger-description %}
Update this instances settings.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="restricted_to_allowlist" type="boolean" %}
Whether sign up is restricted to email addresses, phone numbers and usernames that are on the allowlist.  Defaults to `false`
{% endswagger-parameter %}

{% swagger-parameter in="body" name="session_time_to_abandon" type="integer" %}
Time in seconds that a session will be abandoned.  This usually occurs after expiration, and will completely remove the session from the client object. Defaults to 2592000 (1 month)
{% endswagger-parameter %}

{% swagger-parameter in="body" name="session_time_to_expire" type="integer" %}
Time in seconds that a session will expire.  Once a session expires, the user will be logged out.  A logged out user still exists on the client object, so that you can prompt them to sign in again.  Defaults to 604800 (1 week)
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}
