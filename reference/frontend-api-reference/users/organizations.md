# Backup codes

{% hint style="info" %}
Backup codes is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Available requests

* **`POST`**`/v1/me/backup codes`

## The Backup Code object

```json
{
  "object": "backup_code",
  "id": "bcode_21Ufcy98STcA11s3QckIwtwHIES",
  "codes": ["12345678", "abcdefgh"],
  "created_at": 1438000669544,
  "updated_at": 1438000669544
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/backup_codes" method="post" summary="Create backup codes" %}
{% swagger-description %}
Generates a fresh new set of backup codes for the user.
{% endswagger-description %}

{% swagger-response status="200" description="Creates a new set successfully." %}

{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions" %}

{% endswagger-response %}
{% endswagger %}
