# GWM (Golang Web Modules)

Solution collection for quickly setting up
Golang web applications (Currently especially for Gin framework).

## Why GWM?

GWM maintains the best practices and patterns for building web applications in Go.
If painlessly building a web application in Go is your goal, GWM is the solution.
It provides a solid foundation for your application,
allowing you to focus on building features rather than boilerplate code.

### Application Core

- **Configuration Store** depends on [viper](https://github.com/spf13/viper) for managing application configuration.
- **Slog Logger** with contextual logging and support (using [tint](github.com/lmittmann/tint) for colorized output).

### Clients

- **MongoDB Driver** for MongoDB database operations.
- **Redis Clients** for Redis operations and other Redis-based services (For example, distributed locks & caching).

### Utilities

- **Gin Utils** for Gin framework, including middleware and helpers.

See more in [documentation](https://slhmy.github.io/go-webmods).
