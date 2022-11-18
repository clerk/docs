# Profile image

## Available requests

* **`POST`**`/v1/me/profile_image`
* **`DELETE`**` ``/v1/me/profile_image`

## The image object

```javascript
{
    "object": "image",
    "id": "img_1wL1cVE503unGZ05yyG5hZX7xdJ",
    "name": "img.jpeg",
    "public_url": "https://images.clerk.dev/uploaded/img_1wL1cVE503unGZ05yyG5hZX7xdJ.jpeg"
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/profile_image" method="post" summary="Update profile image" %}
{% swagger-description %}
Upload a new profile image for the current user.  Must use multipart/form-data with one image file.  It must be a jpg, png, gif, or webp image smaller than 10 MB.
{% endswagger-description %}

{% swagger-parameter in="body" name="file" type="object" %}
The image to upload.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="delete" path="/v1/me/profile_image" baseUrl="https://api.clerk.dev" summary="Delete profile image" %}
{% swagger-description %}
Remove a user's profile image.
{% endswagger-description %}

{% swagger-response status="200: OK" description="The user profile image was successfully deleted." %}
```javascript
{
    "object": "image",
    "id": "img_21Ufcy98STcA11s3QckIwtwHIES",
    "slug": "",
    "deleted": true
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="The user does not have a profile image." %}
```javascript
{
  "errors": [
    {
      "code": "resource_not_found",
      "long_message": "Resource not found",
      "message": "not found"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}
