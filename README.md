Social Network App API
This project is a Go-based RESTful API designed for managing user accounts, handling authentication, following/unfollowing users, and performing various user-related tasks in a social network context. The API is built using the Gin framework, GORM ORM for database management, and includes user validation, health checks, and more.

Features
User Registration and Login:

Register a new user with their personal, residential, and office details.
Secure password storage using encryption.
Token-based authentication using Bearer tokens.
User Profile Management:

Update personal details (first name, last name, gender, etc.).
Change passwords.
Delete user accounts and associated data.
Following/Unfollowing Users:

Follow or unfollow other users.
Retrieve the list of users you follow or are followed by.
Validation:

Input validation for user details, including email, password, gender, and marital status.
Uses GORM validators and custom validation logic.
Health Checks:

API health checks for checking service and endpoint availability.

Technologies Used
Go (Golang): Core language used for development.
Gin: A high-performance web framework used for routing and handling HTTP requests.
GORM: ORM library for interacting with the PostgreSQL database.
PostgreSQL: Database used for persistent data storage.
