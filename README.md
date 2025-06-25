# URL Shortener Frontend

This project is a simple URL shortener web application that allows users to input a long URL and generate a shortened version. The generated URLs will automatically expire after 6 months.

## Features

- Input field for entering the long URL
- Button to generate the short URL
- Display area for the generated short URL
- Notification about the expiration of generated URLs

## Project Structure

```
url-shortener-frontend
├── src
│   ├── index.html        # Main HTML document
│   ├── styles
│   │   └── main.css      # CSS styles for the website
│   ├── scripts
│   │   └── app.js        # JavaScript functionality for URL shortening
│   └── assets
│       └── favicon.ico    # Favicon for the website
├── package.json          # NPM configuration file
└── README.md             # Project documentation
```

## Getting Started

To run this project locally, follow these steps:

1. Clone the repository:
   ```
   git clone <repository-url>
   ```

2. Navigate to the project directory:
   ```
   cd url-shortener-frontend
   ```

3. Open the `src/index.html` file in your web browser to view the application.

## Usage

1. Enter the long URL in the input field.
2. Click the "Generate Short URL" button.
3. The shortened URL will be displayed below the button.
4. Please note that the generated URLs will be removed automatically in 6 months.

## License

This project is open-source and available under the [MIT License](LICENSE).