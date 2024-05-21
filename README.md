# URL Shortener

This service helps you transform long, unwieldy URLs into shorter, more manageable links. Perfect for sharing on social media, emails, or anywhere the character count is limited.

## Features

1 ### Shorten URLs: Easily convert lengthy URLs into concise, easy-to-share short links.
2 ### Retrieve Original URLs: Access the original webpage from a shortened URL with a single click.
3 ### Top Domains: Discover the top 3 domains with the most shortened URLs, providing insights into user trends.

##Description
This URL shortener takes the hassle out of managing long URLs. It creates shorter, 6-character random strings from a predefined set of characters to represent the original URL. The service offers basic functionalities:

Shorten URLs: Generate shortened links for any valid URL.
Redirect: Clicking on a shortened URL automatically redirects you to the original webpage.
Installation

For developers who want to contribute or run the service locally, follow these steps:

Clone the repository:

git clone https://github.com/yourusername/url-shortener.git
Navigate to the directory:

cd url-shortener
Install dependencies:

go build -o url-shortener
Run the application:

./url-shortener
Creating a Short URL

Here's how to create a shortened URL using the provided API (for developers):

POST Request (using curl):

curl --location 'http://localhost:8000/url' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'url=https://example.com/{sample}'
GET Request (using curl) to access the original URL:

curl --location 'http://localhost:8000/url/{shorturl}'
Replace {shorturl} with the actual shortened URL you receive in the response.
