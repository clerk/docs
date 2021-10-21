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

