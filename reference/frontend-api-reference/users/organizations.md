---
description: Retrieve information regarding the organizations the user is a member of.
---

# Organizations

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Available requests

* **`GET`**`/v1/me/organizations`

## The organization object

```json
{
  "object": "organization",
  "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
  "logo_url": "https://images.clerk.services/default-logo.png",
  "name": "Acme Inc",
  "public_metadata": {},
  "slug": "acme-inc",
  "role": "basic_member",
  "created_at": 1438000669544,
  "updated_at": 1438000669544
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/organizations" method="get" summary="List organizations" %}
{% swagger-description %}
Returns a list of the organizations that the current user is a member of.&#x20;

The organization membership role is included.

The organizations will be ordered by descending membership creation date, which means that the organization that the user joined most recently will be first in the list.

Results can be paginated by supplying a combination of `limit` and `offset` query parameters.
{% endswagger-description %}

{% swagger-parameter in="query" name="limit" type="Number" %}
Applies a limit to the number of organizations returned. Can be used for paginating the results together with 

`offset`

. Must be an integer greater than zero.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="offset" type="Number" %}
Skip the first 

`offset`

 results when paginating. Needs to be an integer greater than zero. Use it together with 

`limit`

 to skip the last retrieved organizations.
{% endswagger-parameter %}

{% swagger-response status="200" description="Organization list retrieved successfully." %}
```javascript
[
  {
    "object": "organization",
    "id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
    "logo_url": "https://images.clerk.services/default-logo.png",
    "name": "Acme Inc",
    "public_metadata": {},
    "role": "basic_member",
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
  },
  {
    "object": "organization",
    "id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
    "logo_url": null,
    "name": "Foozbaz",
    "public_metadata": {},
    "role": "admin",
    "slug": "foozbaz",
    "created_at": 1612556316784,
    "updated_at": 1612556316784
  }
]
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to list organizations." %}
```javascript
{
  "errors": [
    {
      "code": "organizations_not_enabled_in_instance",
      "long_message": "The organizations feature is not enabled for your instance. If you want to try it out, contact us at support@clerk.dev.",
      "message": "access denied"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}
