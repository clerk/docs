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
Whether sign up is restricted to email addresses, phone numbers and usernames that are on the allowlist.  Defaults to `false`.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="from_email_address" type="string" %}
The local part of the email address from which authentication-related emails (e.g. OTP code, magic links) will be sent.



Only alphanumeric values are allowed.



Note that this value should contain _only the local part_ of the address. For example, assuming your domain is `example.com` and you want emails to be sent from `foo@example.com`. In that case, you should provide `foo` in this value, not `foo@example.com`. &#x20;
{% endswagger-parameter %}

{% swagger-parameter in="body" name="test_mode" type="boolean" %}
Toggles test mode for this instance. Defaults to 

`true`

 for development instances.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="progressive_sign_up" type="boolean" %}
Enable the Progressive Sign Up algorithm. Refer to the 

[docs](https://clerk.dev/docs/main-concepts/sign-up-flow#progressive-sign-up-beta)

 for more info.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="session_token_template" type="string" %}
The name of the JWT Template to augment your session tokens. To disable this, pass an empty string.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}
