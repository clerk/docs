# Allowlist Identifiers

This object represents an identifier \(email address or phone number\) which is allowed to sign-up. It is only used if you have enabled the "Allowlist" beta feature. 

### Available requests

* **`GET`** `/v1/allowlist_identifiers`
* **`POST`**`/v1/allowlist_identifiers`
* **`DEL`** `/v1/allowlist_identifiers/:id`

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

{% api-method method="get" host="https://api.clerk.dev" path="/v1/allowlist\_identifiers" %}
{% api-method-summary %}
List all allowlist identifiers
{% endapi-method-summary %}

{% api-method-description %}
List all allowlist identifiers.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://api.clerk.dev" path="/v1/allowlist\_identifiers" %}
{% api-method-summary %}
Create an allowlist identifier
{% endapi-method-summary %}

{% api-method-description %}
Add a new identifier to the allowlist.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-body-parameters %}
{% api-method-parameter name="notify" type="boolean" required=true %}
Send an invitation via SMS \(for phone identifier\) or email \(for email address identifier\).
{% endapi-method-parameter %}

{% api-method-parameter name="identifier" type="string" required=true %}
An email address or a phone number in international format \(E.164\)
{% endapi-method-parameter %}
{% endapi-method-body-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
	"object": "allowlist_identifier", 
	"id": "alid_1sOWqxBnLKkD7Alxv5EymdOKjue",
	"identifier": "+1655559004"
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="delete" host="https://api.clerk.dev" path="/v1/allowlist\_identifiers/:id" %}
{% api-method-summary %}
Delete an allowlist identifier
{% endapi-method-summary %}

{% api-method-description %}
Remove an identifier from the allowlist.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
	"id": "alid_1sQqrOkmSTiXu8ho74r5lWuk3Zi",
	"object": "allowlist_identifier", 
	"deleted": true
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



