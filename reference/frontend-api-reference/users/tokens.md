# Tokens

## Available requests

* **`POST`**`/v1/me/tokens`

## The token object

```javascript
{
    "object": "token",
    "jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6Imluc18xbHlXRFppb2JyNjAwQUtVZVFEb1NsckVtb00iLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJjbGVyayIsImVtYWlsIjoiYnJhZGVuQGNsZXJrLmRldiIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJleHAiOjE2MjgyMjE1NzEsImZhbWlseV9uYW1lIjoiU2lkb3RpIiwiZ2l2ZW5fbmFtZSI6IkJyYWRlbiIsImlhdCI6MTYyODIyMTUwNiwiaXNzIjoiaHR0cHM6Ly9jbGVyay5jbGVyay5kZXYiLCJuYW1lIjoiQnJhZGVuIiwicGljdHVyZSI6Imh0dHBzOi8vaW1hZ2VzLmNsZXJrLmRldi91cGxvYWRlZC9pbWdfMXV4MTc3dkJYVmVVR240d1FYZm5yM2xTd0lnLmpwZWciLCJzdWIiOiJ1c2VyXzFvNHFmcU1lQ2tLNEtEOWZpQ1BLSVhjQTloOCIsInVwZGF0ZWRfYXQiOjE2MjgyMTg1Nzd9.b2oZwr8hW33dQC856k5Xv0j-SND3hI11r-j2I6BXkldw-yGnYiKq1GUAwC2F1ecFIaAvVJFHw-MYT34pQt6xMRY_InC2-K7CeDKUIuXzC8fglHU-a8WMp6ehON1MImGy5fuqyyWI6x8vY2TlaKjVipBWFkVDmqLrOWJuuzAePeMaYjN-zPkSOeD1ydiRbGf97wokTGx9hxyYq7C2vX2CLPOkb17f9Is92GcHlatH6Ac6Ptnd_JNy5bXGGjjPKwJgPYO2cCDlxamDgL-40IBJ8l_0IxJWN_PlShZQoT05xGdIJMcnje4WHlt0rm6nQYBPt3Cahz_gd4xoCAeYJEch0A"
}
```

{% api-method method="post" host="https://clerk.example.com" path="/v1/me/tokens" %}
{% api-method-summary %}
Create a token
{% endapi-method-summary %}

{% api-method-description %}
Create a short-lived JSON Web Token \(JWT\) structured to be compatible for the specified service.  If no service is supplied, you will get a JWT with the \`clerk\` payload.  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="service" type="string" required=false %}
\`clerk\`, \`firebase\`, \`hasura\`
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
    "object": "token",
    "jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6Imluc18xbHlXRFppb2JyNjAwQUtVZVFEb1NsckVtb00iLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJjbGVyayIsImVtYWlsIjoiYnJhZGVuQGNsZXJrLmRldiIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJleHAiOjE2MjgyMjE1NzEsImZhbWlseV9uYW1lIjoiU2lkb3RpIiwiZ2l2ZW5fbmFtZSI6IkJyYWRlbiIsImlhdCI6MTYyODIyMTUwNiwiaXNzIjoiaHR0cHM6Ly9jbGVyay5jbGVyay5kZXYiLCJuYW1lIjoiQnJhZGVuIiwicGljdHVyZSI6Imh0dHBzOi8vaW1hZ2VzLmNsZXJrLmRldi91cGxvYWRlZC9pbWdfMXV4MTc3dkJYVmVVR240d1FYZm5yM2xTd0lnLmpwZWciLCJzdWIiOiJ1c2VyXzFvNHFmcU1lQ2tLNEtEOWZpQ1BLSVhjQTloOCIsInVwZGF0ZWRfYXQiOjE2MjgyMTg1Nzd9.b2oZwr8hW33dQC856k5Xv0j-SND3hI11r-j2I6BXkldw-yGnYiKq1GUAwC2F1ecFIaAvVJFHw-MYT34pQt6xMRY_InC2-K7CeDKUIuXzC8fglHU-a8WMp6ehON1MImGy5fuqyyWI6x8vY2TlaKjVipBWFkVDmqLrOWJuuzAePeMaYjN-zPkSOeD1ydiRbGf97wokTGx9hxyYq7C2vX2CLPOkb17f9Is92GcHlatH6Ac6Ptnd_JNy5bXGGjjPKwJgPYO2cCDlxamDgL-40IBJ8l_0IxJWN_PlShZQoT05xGdIJMcnje4WHlt0rm6nQYBPt3Cahz_gd4xoCAeYJEch0A"
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

