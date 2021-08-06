# Profile image

## Available requests

* **`POST`**`/v1/me/profile_image`

## The image object

```javascript
{
    "object": "image",
    "id": "img_1wL1cVE503unGZ05yyG5hZX7xdJ",
    "name": "img.jpeg",
    "public_url": "https://images.clerk.dev/uploaded/img_1wL1cVE503unGZ05yyG5hZX7xdJ.jpeg"
}
```

{% api-method method="post" host="https://clerk.example.com" path="/v1/me/profile\_image" %}
{% api-method-summary %}
Update profile image
{% endapi-method-summary %}

{% api-method-description %}
Upload a new profile image for the current user.  Must use multipart/form-data with one image file.  It must be a jpg, png, gif, or webp image smaller than 10 MB.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="file" type="object" required=true %}
The image to upload.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



