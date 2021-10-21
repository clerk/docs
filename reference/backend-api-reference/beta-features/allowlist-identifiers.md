# Allowlist Identifiers

This object represents an identifier (email address or phone number) which is allowed to sign-up. It is only used if you have enabled the "Allowlist" beta feature.&#x20;

### Available requests

* **`GET `**`/v1/allowlist_identifiers`
* **`POST`**`/v1/allowlist_identifiers`
* **`DEL `**`/v1/allowlist_identifiers/:id`

### Example allowlist identifier schema

```javascript
{
	"object": "allowlist_identifier", 
	"id": "alid_1sOWqxBnLKkD7Alxv5EymdOKjue",
	"identifier": "+1655559004",
  "created_at": 1620370000,
  "updated_at": 1620370000
}
```

{% swagger baseUrl="https://api.clerk.dev" path="/v1/allowlist_identifiers" method="get" summary="List all allowlist identifiers" %}
{% swagger-description %}
List all allowlist identifiers.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
[
	{
		"object":"allowlist_identifier",
		"id":"alid_1sOWqxBnLKkD7Alxv5Eym0OKjue",
		"identifier":"user@example.com"
	},
	{
		"object":"allowlist_identifier",
		"id":"alid_1sOWqxBnLKkD7Alxv5EymdOKjue",
		"identifier":"+12025559999"
	}
]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/allowlist_identifiers" method="post" summary="Create an allowlist identifier" %}
{% swagger-description %}
Add a new identifier to the allowlist.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="notify" type="boolean" %}
Send an invitation via SMS (for phone identifier) or email (for email address identifier).
{% endswagger-parameter %}

{% swagger-parameter in="body" name="identifier" type="string" %}
An email address or a phone number in international format (E.164)
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
{
	"object": "allowlist_identifier", 
	"id": "alid_1sOWqxBnLKkD7Alxv5EymdOKjue",
	"identifier": "+1655559004"
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/allowlist_identifiers/:id" method="delete" summary="Delete an allowlist identifier" %}
{% swagger-description %}
Remove an identifier from the allowlist.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
{
	"id": "alid_1sQqrOkmSTiXu8ho74r5lWuk3Zi",
	"object": "allowlist_identifier", 
	"deleted": true
}
```
{% endswagger-response %}
{% endswagger %}

