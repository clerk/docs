# Getting started with Ruby

The source code of the Ruby SDK is hosted here: [https://github.com/clerkinc/clerk-sdk-ruby](https://github.com/clerkinc/clerk-sdk-ruby)

Add this line to your application's Gemfile:

```ruby
gem 'clerk-sdk-ruby', require: "clerk"
```

And then execute:

```text
$ bundle install
```

Or install it yourself as:

```text
$ gem install clerk-sdk-ruby
```

## Quick Start

First, you need to get an API key for a Clerk instance. This is done via the [Clerk dashboard](https://dashboard.clerk.dev/applications).

Then you can instantiate a `Clerk::SDK` instance and access all [Backend API](https://docs.clerk.dev/backend/backend-api-reference) endpoints. Here's a quick example:

```ruby
clerk = Clerk::SDK.new(api_key: "your_api_key")
# List all users
clerk.users.all
# Get your first user
user = clerk.users.all(limit: 1).first
# Extract their primary email address ID
email_id = user["primary_email_address_id"]
# Send them a welcome email
clerk.emails.create(
    email_address_id: email_id,
    from_email_name: "welcome",
    subject: "Welcome to MyApp",
    body: "Welcome to MyApp, #{user["first_name"]}",
)
```

