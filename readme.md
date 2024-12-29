Packages Used:
- context: The context package is used to manage deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines. Here, it’s used for managing the connection timeout to MongoDB.

- fmt: The fmt package provides I/O formatting functions. It's used to format strings and print error messages.

- log: The log package provides simple logging functions to print logs to standard output. Here, it’s used for logging connection and error messages.

- time: The time package provides functions to work with time-related operations. It's used here for setting timeouts and sleeping between retry attempts.

- go.mongodb.org/mongo-driver/mongo: This is the official MongoDB driver for Go. It provides the required functions to connect to MongoDB, ping it, and perform CRUD operations.

- go.mongodb.org/mongo-driver/mongo/options: This package provides options for configuring various MongoDB operations, such as client connection settings.
