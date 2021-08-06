# Instance Settings

Modify your instance's settings.

### Available requests

* **`PATCH`**`/v1/beta_features/instance_settings`

{% api-method method="patch" host="https://api.clerk.dev" path="/v1/beta\_features/instance\_settings" %}
{% api-method-summary %}
Update instance settings
{% endapi-method-summary %}

{% api-method-description %}
Update this instances settings.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-form-data-parameters %}
{% api-method-parameter name="restricted\_to\_allowlist" type="boolean" required=false %}
Whether sign up is restricted to email addresses, phone numbers and usernames that are on the allowlist.  Defaults to \`false\`
{% endapi-method-parameter %}

{% api-method-parameter name="session\_time\_to\_abandon" type="integer" required=false %}
Time in seconds that a session will be abandoned.  This usually occurs after expiration, and will completely remove the session from the client object. Defaults to 2592000 \(1 month\)
{% endapi-method-parameter %}

{% api-method-parameter name="session\_time\_to\_expire" type="integer" required=false %}
Time in seconds that a session will expire.  Once a session expires, the user will be logged out.  A logged out user still exists on the client object, so that you can prompt them to sign in again.  Defaults to 604800 \(1 week\)
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

