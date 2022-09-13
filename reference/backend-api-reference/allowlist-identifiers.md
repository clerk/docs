---
description: >-
  Learn how to restrict access to your application only to certain emails or
  phones
---

# Allowlist Identifiers

### Overview

The _allowlist_ feature allows you to control who can get access to your application. It basically restricts sign ups and sign ins only to a certain set of email addresses, email domains, phone numbers or web3 wallets that you define.&#x20;

Apart from individual email addresses, you can also whitelist whole email domains.

Whenever you add a single email address or phone number to your allowlist identifiers, you can also choose whether you want to notify this person by an invitation to their email address or phone number respectively.&#x20;

In order to add a set of email addresses to your allowlist identifiers, simply pass an asterisk as the email inbox. The asterisk will act as a wildcard and add all inboxes for a certain email domain.&#x20;

Here's an example. If you add `*@clerk.dev` as your allowlist identifier, it means that anybody with a `clerk.dev` email address can sign up for your application. Email addresses from different domains will not be able to sign up.

{% hint style="info" %}
In order to use the allowlist feature, it needs to be enabled for your instance. To enable it, use our Backend API to [update your instance settings](beta-features/instance-settings.md).\
\
Each instance of the app (_development, staging, production_) needs to be updated separately.
{% endhint %}

### Available requests

* **`GET`**` ``/v1/allowlist_identifiers`
* **`POST`**`/v1/allowlist_identifiers`
* **`DEL`**` ``/v1/allowlist_identifiers/:id`

### Example allowlist identifier schema

```javascript
{
	"object": "allowlist_identifier", 
	"id": "alid_1sOWqxBnLKkD7Alxv5EymdOKjue",
	"identifier": "+1655559004",
	"identifier_type": "phone_number",
 	"created_at": 1620370000,
 	"updated_at": 1620370000
}
```

{% swagger baseUrl="https://api.clerk.dev" path="/v1/allowlist_identifiers" method="get" summary="List all allowlist identifiers" %}
{% swagger-description %}
List all allowlist identifiers.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
[
	{
		"object":"allowlist_identifier",
		"id":"alid_1sOWqxBnLKkD7Alxv5Eym0OKjue",
		"identifier":"user@example.com",
		"identifier_type": "email_address"
	},
	{
		"object":"allowlist_identifier",
		"id":"alid_1sOWqxBnLKkD7Alxv5EymdOKjue",
		"identifier":"+12025559999",
		"identifier_type": "phone_number"
	}
]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/allowlist_identifiers" method="post" summary="Create an allowlist identifier" %}
{% swagger-description %}
Add a new identifier to the allowlist.
{% endswagger-description %}

{% swagger-parameter in="body" name="identifier" type="string" required="true" %}
An email address, web3 wallet or phone number in international format (E.164).

Instead of individual email addresses, you can also whitelist a whole email domain if you include it as `*@example.com`. In this case, everyone with an email address on `example.com` will have access to your application.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="notify" type="boolean" %}
Send an invitation via SMS (for phone identifier) or email (for email address identifier).

This parameter is not allowed if the identifier is a web3 wallet or an email domain.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
{
	"object": "allowlist_identifier", 
	"id": "alid_1sOWqxBnLKkD7Alxv5EymdOKjue",
	"identifier": "+1655559004",
	"identifier_type": "phone_number"
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/allowlist_identifiers/:id" method="delete" summary="Delete an allowlist identifier" %}
{% swagger-description %}
Remove an identifier from the allowlist.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
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

