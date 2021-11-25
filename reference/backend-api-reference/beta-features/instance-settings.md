# Instance Settings

Modify your instance's settings.

### Available requests

* **`PATCH`**`/v1/beta_features/instance_settings`

{% swagger baseUrl="https://api.clerk.dev" path="/v1/beta_features/instance_settings" method="patch" summary="Update instance settings" %}
{% swagger-description %}
Update this instances settings.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="restricted_to_allowlist" type="boolean" %}
Whether sign up is restricted to email addresses, phone numbers and usernames that are on the allowlist.  Defaults to `false`
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}
