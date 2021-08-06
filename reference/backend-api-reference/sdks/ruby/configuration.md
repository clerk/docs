# Configuration

The SDK can be configured in three ways: environment variables, configuration singleton and constructor arguments. The priority goes like this:

1. Constructor arguments
2. Configuration object
3. Environment variables

If an argument is not provided, the configuration object is looked up, which falls back to the associated environment variable. Here's an example with all supported configuration settings their environment variable equivalents:

```ruby
Clerk.configure do |c|
  c.api_key = "your_api_key" # if omitted: ENV["CLERK_API_KEY"] - API calls will fail if unset
  c.base_url = "https://..." # if omitted: "https://api.clerk.dev/v1/"
  c.logger = Logger.new(STDOUT) # if omitted, no logging
  c.middleware_cache_store = Rails.cache # if omitted: no caching
end
```

You can customize each instance of the `Clerk::SDK` object by passing keyword arguments to the constructor:

```ruby
clerk = Clerk::SDK.new(
    api_key: "X",
    base_url: "Y",
    logger: Logger.new()
)
```

For full customization, you can instead pass a `Faraday` object directly, which will ignore all the other arguments, if passed:

```ruby
faraday = Faraday.new()
clerk = Clerk::SDK.new(connection: faraday)
```

Refer to the [Faraday documentation](https://lostisland.github.io/faraday/usage/#customizing-faradayconnection) for details.

## Environment variables

Here's a list of all environment variables the SDK uses:

| Variable name | Usage |
| :--- | :--- |
| `CLERK_API_BASE` | Overrides the default API base URL: https://api.clerk.dev/v1/ |
| `CLERK_API_KEY` | The API key of your instance \(required\) |
| `CLERK_SIGN_IN_URL` | Rails view helper: `clerk_sign_in_url` |
| `CLERK_SIGN_UP_URL` | Rails view helper: `clerk_sign_up_url` |
| `CLERK_USER_PROFILE_URL` | Rails view helper: `clerk_user_profile_url` |



