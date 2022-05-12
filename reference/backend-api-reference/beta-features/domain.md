# Domain

Manage the domain of a production instance

### Available requests

* **`PATCH`**`/v1/beta_features/domain`

{% hint style="info" %}
Updating the domain requires setting [DNS records](https://clerk.dev/docs/how-to/deploy-to-production#dns-records) again and deploying new [SSL certificates](https://clerk.dev/docs/how-to/deploy-to-production#deploy).
{% endhint %}

{% hint style="danger" %}
Updating the domain will invalidate all current user sessions for your application.
{% endhint %}

{% swagger method="patch" path="/v1/beta_features/domain" baseUrl="https://api.clerk.dev" summary="Update the domain of a production instance" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="body" name="home_url" required="true" %}
The new home URL of the production instance e.g. https://www.example.com
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="202: Accepted" description="" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}
