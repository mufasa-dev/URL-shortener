## 🚀 URL Shortener
A simple and efficient URL shortener built with Go and SQLite. This application allows users to shorten long URLs into concise links, making them easier to share and manage.

## Features
🔗 Shorten long URLs quickly

🛠 Store and manage shortened URLs using SQLite

📡 Redirect users to the original URL

📊 Track URL usage statistics (optional)

🚀 Fast and lightweight

## Installation
#### Clone the repository:

git clone https://github.com/your-user/url-shortener.git
cd url-shortener
Install dependencies:

go mod tidy
#### Run the application:

go run main.go

## Configuration
The application uses SQLite as its database.

You can modify the settings in config.json to customize behavior.

## Example Usage
Once the server is running, you can shorten URLs with an API request:

curl -X POST -H "Content-Type: application/json" -d '{"url": "https://example.com"}' http://localhost:8080/shorten
This will return a shortened URL, which can be used to redirect users.

## Contributing
Feel free to contribute to this project by submitting pull requests or issues.

## License
This project is licensed under the MIT License.
