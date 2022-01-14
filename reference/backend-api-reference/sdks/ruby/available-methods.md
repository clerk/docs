# Available methods

The Ruby SDK mirrors the [Backend API](../../). Here's a list of the instance methods on the `Clerk::SDK` and the API endpoint each one corresponds to:

| Instance method on SDK object | API endpoint/prefix                                          | Methods on resource object           |
| ----------------------------- | ------------------------------------------------------------ | ------------------------------------ |
| `allowlist_identifiers`       | [/v1/allowlist\_identifiers](../../allowlist-identifiers.md) | `all` `create` `delete`              |
| `allowlist`                   | [/v1/beta\_features/allowlist](../../beta-features/)         | `update`                             |
| `clients`                     | [/v1/clients](../../clients.md)                              | `all` `find` `verify_token`          |
| `emails`                      | [/v1/emails](../../emails.md)                                | `create`                             |
| `sessions`                    | [/v1/sessions](../../sessions.md)                            | `all` `find` `revoke` `verify_token` |
| `sms_messages`                | [/v1/sms\_messages](../../sms-messages.md)                   | `create`                             |
| `users`                       | [/v1/users](../../users.md)                                  | `all` `find` `update` `delete`       |

## Examples

All examples assume you have an instance of the `Clerk::SDK`:

```ruby
sdk = Clerk::SDK.new
```

### Allowlist identifiers

Get an array of all allowlist identifiers:

```ruby
sdk.allowlist_identifiers.all
```

Create a new allowlist identifier:

```ruby
sdk.allowlist_identifiers.create(identifier: "john@example.com", notify: true)
```

Delete an allowlist identifier:

```ruby
sdk.allowlist_identifiers.delete("alid_xyz")
```

### Allowlist

Toggle allowlist-only sign-ups on/off:

```ruby
sdk.allowlist.update(restricted_to_allowlist: true)
```

### Clients

Get a client by its ID:

```ruby
sdk.clients.find("client_xyz")
```

Get an array of all clients:

```ruby
sdk.clients.all
```

Verify the JWT and return the client:

```ruby
sdk.clients.verify_token("jwt")
```

### Emails

Send an email:

```ruby
sdk.emails.create(
    email_address_id: "ema_xyz", 
    from_email_name: "noreply", 
    subject: "Welcone",
    body: "<html>...</html>",
)
```

### Sessions

Get a session by its ID:

```ruby
sdk.sessions.find("sess_xyz")
```

Get an array of all sessions:

```ruby
sdk.sessions.all
```

Revoke a session:

```ruby
sdk.sessions.revoke("sess_xyz")
```

Verify the JWT of a specific session ID:

```ruby
sdk.sessions.verify_token("sess_xyz", "jwt")
```

### SMS Messages

Send an SMS:

```ruby
sdk.sms_messages.create(phone_number_id: "idn_xyz", message: "Welcome!")
```

### Users

Get an array of all users:

```ruby
sdk.users.all
```

Get an array of users, with filters:

```ruby
sdk.users.all(email_address: ["user1@example.com", "user2@example.com"])
```

Update a user:

```ruby
sdk.users.update("user_xyz", {first_name: "John"})
```

Delete a user:

```ruby
sdk.users.delete("user_xyz")
```

