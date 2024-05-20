URL Shortener
A simple URL shortener service that generates a short URL by taking a random 6-character string from a predefined character set.
Description
This URL shortener service allows users to shorten long URLs into shorter, more manageable links. The short URL is generated using a random 6-character string from a predefined character set. The service supports basic operations such as creating short URLs and redirecting to the original URLs.

Installation
Clone the repository:
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
Install the dependencies:
go build -o url-shortener
Run the application:
./url-shortener

Creating short url:
Use this curl command for post api :
curl --location 'http://localhost:8000/url' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'url= https://example.com/{sample}'

Use this curl command for get api :
curl --location 'http://localhost:8000/url/{shorturl}'

