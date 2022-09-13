---
description: >-
  Learn how to prevent access to your application for certain emails, domain,
  phones or web3 wallets
---

# Blocklist Identifiers

### Overview

The _blocklist_ feature allows you to control who should not be allowed access to your application. It basically prevents sign ups and sign ins for certain set of identifiers, which can be email addresses, email domains, phone numbers or web3 wallets.&#x20;

In order to add a set of email addresses to your blocklist identifiers, simply pass an asterisk as the email inbox. The asterisk will act as a wildcard and add all inboxes for a certain email domain.&#x20;

Here's an example. If you add `*@clerk.dev` as your blocklist identifier, it means that email addresses from the `clerk.dev` email domain cannot sign up or sign in to your application. Email addresses from different domains can have access to it.

{% hint style="info" %}
In order to use the blocklist feature, it needs to be enabled for your instance. To enable it, use our Backend API to [update your instance settings](beta-features/instance-settings.md).\
\
Each instance of the app (_development, staging, production_) needs to be updated separately.
{% endhint %}

### Available requests

* **`GET`**` ``/v1/blocklist_identifiers`
* **`POST`**`/v1/blocklist_identifiers`
* **`DEL`**` ``/v1/blocklist_identifiers/:id`

### Example blocklist identifier schema

```javascript
{
	"object": "blocklist_identifier", 
	"id": "blid_1sOWqxBnLKkD7Alxv5EymdOKjue",
	"identifier": "+1655559004",
	"identifier_type": "phone_number",
 	"created_at": 1620370000,
 	"updated_at": 1620370000
}
```

{% swagger baseUrl="https://api.clerk.dev" path="/v1/blocklist_identifiers" method="get" summary="List all blocklist identifiers" %}
{% swagger-description %}
List all blocklist identifiers.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
[
	{
		"object":"blocvklist_identifier",
		"id":"blid_1sOWqxBnLKkD7Alxv5Eym0OKjue",
		"identifier":"user@example.com",
		"identifier_type":"email_address"
	},
	{
		"object":"blocklist_identifier",
		"id":"blid_1sOWqxBnLKkD7Alxv5EymdOKjue",
		"identifier":"+12025559999",
		"identifier_type":"phone_number"
	}
]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/blocklist_identifiers" method="post" summary="Create a blocklist identifier" %}
{% swagger-description %}
Add a new identifier to the blocklist.
{% endswagger-description %}

{% swagger-parameter in="body" name="identifier" type="string" required="true" %}
An email address, email domain, web3 wallet or a phone number in international format (E.164).

Instead of individual email addresses, you can also whitelist a whole email domain if you include it as `*@example.com`. In this case, everyone with an email address on `example.com` will be denied  access to your application.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
{
	"object": "blocklist_identifier", 
	"id": "blid_1sOWqxBnLKkD7Alxv5EymdOKjue",
	"identifier": "+1655559004",
	"identifier_type": "phone_number"
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/blocklist_identifiers/:id" method="delete" summary="Delete a blocklist identifier" %}
{% swagger-description %}
Remove an identifier from the blocklist.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
{
	"id": "blid_1sQqrOkmSTiXu8ho74r5lWuk3Zi",
	"object": "blocklist_identifier", 
	"deleted": true
}
```
{% endswagger-response %}
{% endswagger %}

