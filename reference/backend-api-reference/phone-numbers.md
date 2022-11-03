# Phone Numbers

This object represents a phone number in your instance.

### Available requests

* **`POST`**` ``/v1/phone_numbers`
* **`GET`**`  ``/v1/phone_numbers/:id`
* **`PATCH`**`/v1/phone_numbers/:id`
* **`DEL`**`  ``/v1/phone_numbers/:id`

### Example phone number schema

```json
{
    "id": "idn_2H2MoyJjUuDIQZLlafRppVIAG7T",
    "object": "phone_number",
    "phone_number": "+15555555555",
    "reserved_for_second_factor": false,
    "default_second_factor": false,
    "reserved": false,
    "verification": {
        "status": "verified",
        "strategy": "admin",
        "attempts": null,
        "expire_at": null
    },
    "linked_to": [],
    "backup_codes": null
}
```



{% swagger src="../../.gitbook/assets/COMB-server-api-oas3.yml" path="/v1/phone_numbers" method="post" %}
[COMB-server-api-oas3.yml](../../.gitbook/assets/COMB-server-api-oas3.yml)
{% endswagger %}

{% swagger src="../../.gitbook/assets/COMB-server-api-oas3.yml" path="/v1/phone_numbers/{phone_number_id}" method="get" %}
[COMB-server-api-oas3.yml](../../.gitbook/assets/COMB-server-api-oas3.yml)
{% endswagger %}

{% swagger src="../../.gitbook/assets/COMB-server-api-oas3.yml" path="/v1/phone_numbers/{phone_number_id}" method="patch" %}
[COMB-server-api-oas3.yml](../../.gitbook/assets/COMB-server-api-oas3.yml)
{% endswagger %}

{% swagger src="../../.gitbook/assets/COMB-server-api-oas3.yml" path="/v1/phone_numbers/{phone_number_id}" method="delete" %}
[COMB-server-api-oas3.yml](../../.gitbook/assets/COMB-server-api-oas3.yml)
{% endswagger %}
