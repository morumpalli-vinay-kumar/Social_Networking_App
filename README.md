# Social Network API

This project is a Go-based RESTful API designed for managing user accounts, handling authentication, following/unfollowing users, and performing various user-related tasks in a social network context. The API is built using the Gin framework, GORM ORM for database management, and includes user validation, health checks, and more.

## Features

### User Registration and Login
- Register a new user with their personal, residential, and office details.
- Secure password storage using encryption.
- Token-based authentication using Bearer tokens.

### User Profile Management
- Update personal details (first name, last name, gender, etc.).
- Change passwords.
- Delete user accounts and associated data.

### Following/Unfollowing Users
- Follow or unfollow other users.
- Retrieve the list of users you follow or are followed by.

### Validation
- Input validation for user details, including email, password, gender, and marital status.
- Uses GORM validators and custom validation logic.

### Health Checks
- API health checks for checking service and endpoint availability.

## Technologies Used

- **Go (Golang)**: Core language used for development.
- **Gin**: A high-performance web framework used for routing and handling HTTP requests.
- **GORM**: ORM library for interacting with the PostgreSQL database.
- **PostgreSQL**: Database used for persistent data storage.

## Getting Started

1. **Clone the repository:**
    ```bash
    git clone https://github.com/yourusername/social-network-api.git
    cd social-network-api
    ```

2. **Install dependencies:**
    ```bash
    go mod tidy
    ```

3. **Set up the database:**
    - Update the database configuration in `config.yml` or `.env`.
    - Run database migrations using GORM.

4. **Start the API:**
    ```bash
    go run main.go
    ```

5. **Access the API:**
    - The API will be running on `http://localhost:8080`.


## Contributing

If you'd like to contribute to this project, please follow these steps:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or support, please contact [vinay.kumar@joshtechnologygroup.com](vinay.kumar@joshtechnologygroup.com).


