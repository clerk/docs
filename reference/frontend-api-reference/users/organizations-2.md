# TOTP

{% hint style="info" %}
TOTP (Authenticator App) is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Available requests

* **`POST`**`/v1/me/totp`
* **`POST`**`/v1/me/totp/attempt_verification`
* **`DEL`**`/v1/me/totp`

## The TOTP object

```json
{
  "object": "totp",
  "id": "totp_21Ufcy98STcA11s3QckIwtwHIES",
  "secret": "6JHMXLT5H3WOF5PSHX7ALXXFX7NLTRBJ",
  "uri": "otpauth://totp/docs:docs@email.com?algorithm=SHA1&digits=6&issuer=docs&period=30&secret=6JHMXLT5H3WOF5PSHX7ALXXFX7NLTRBJ",
  "verified": true,
  "backup_codes": ["12345678", "abcdefgh"],
  "created_at": 1438000669544,
  "updated_at": 1438000669544
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/totp" method="post" summary="Creates a new TOTP object" %}
{% swagger-description %}
Creates a new unverified TOTP object that seeks verification before it is enabled
{% endswagger-description %}

{% swagger-response status="200" description="TOTP created successfully." %}

{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="TOTP already enabled." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions." %}

{% endswagger-response %}
{% endswagger %}

{% swagger method="post" path="/v1/me/totp/attempt_verification" baseUrl="https://clerk.example.com" summary="Attempt verification of the TOTP" %}
{% swagger-description %}
Attempts to verify the existing TOTP, by providing a valid code. If the operation is successful, TOTP is been enabled for the user.
{% endswagger-description %}

{% swagger-response status="200: OK" description="TOTP successfully verified and enabled." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="TOTP already enabled." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="TOTP not found." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="The provided code is invalid." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="delete" path="/v1/me/totp" baseUrl="https://clerk.example.com" summary="Delete TOTP" %}
{% swagger-description %}
Deletes and disables the TOTP of the user.
{% endswagger-description %}

{% swagger-response status="200: OK" description="TOTP deleted successfully." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="TOTP not enabled." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}
