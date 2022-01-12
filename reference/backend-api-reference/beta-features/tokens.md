# Tokens

Tokens requests allows you to create JWTs from your backend based on `jwt_templates` defined in your dashboard.

### Available requests

* **`POST`**` ``/v1/tokens`

{% swagger baseUrl="https://api.clerk.dev" path="/v1/tokens" method="post" summary="Create a token" %}
{% swagger-description %}
Creates a JSON Web Token(JWT) based on a JWT Template defined for your instance.
{% endswagger-description %}

{% swagger-parameter in="path" name="name" type="string" required="true" %}
The name of the JWT Template defined in your instance provider (e.g. 

`custom_hasura`

).
{% endswagger-parameter %}

{% swagger-parameter in="path" name="user_id" type="string" required="true" %}
The Clerk user ID which you would like to create a JWT for.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
{
    "object": "token",
    "jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6Imluc18xazVvd0NSWEx3c3huUmt6RGNoVm9XVG1YSTgiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE2NDIwMjk2NDIsImZpcnN0X25hbWUiOiJCb3NzIiwiaWF0IjoxNjQyMDI5NTgyLCJpc3MiOiJodHRwczovL2NsZXJrLmNsZXJrc3RhZ2UuZGV2IiwianRpIjoiYzU2YTUwN2Y3ODU3ZjkyNjBlNzQiLCJuYmYiOjE2NDIwMjk1NzcsInN1YiI6InVzZXJfMW41QVkyUlhKeUJHaEdMQWhkRWZmYU9DcEZKIn0.jMZjFadJiRsFS9uh0JqzmepAmbij52qKL0nEIzXWquTecRlWLgU2dmFt66jRLEN32SdV0ERyD4eyEKfN8L_judgVR_p38qXnKvgnE_pLmZpwjLjSDm05Ow081VH6Kd2PZuvMbW7DU6xAl5x2Fqj6QqlQGdicuEYcTKxzU5Bo3yJ8Heqnjaa-91VyRritF1hz6J3AVaePEAFS0aU2vdNvXABbup_Om2IfbT1Exr76k8btJKhQvGdkQPDd88V_5DIemmDvV76d4SKR9LE5fYXpGAM74Q1aSI5XNmvjf0aTR_hdoNFNAxjss_0geq3xdCCtLpKN8NMOh7beL4MWjBjifA"
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="The access token has expired but the provider hasn't provided us with a refresh token and so we cannot fetch a new access token." %}

{% endswagger-response %}
{% endswagger %}



