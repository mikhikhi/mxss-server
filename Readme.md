# MXSS Server Side Sanitizer

This project demonstrates four different cross-site scripting (XSS) attack sanitization filters implemented with default configuration in a Dockerized environment. The filters are built using:

- sanitize-html (Node.js): A highly configurable XSS filter library.
- dompurify (Node.js): A DOM-based XSS sanitizer, useful for preventing client-side XSS attacks.
- bluemonday (Go): A fast and secure HTML sanitization library for Go.
- html_sanitizer (Python): A Python-based library for sanitizing HTML to prevent XSS attacks.

The project is containerized using Docker and orchestrated with Docker Compose. Each of the services runs in its own container, and they interact to sanitize a shared input and return a sanitized response.

## Getting Started

```
git clone https://github.com/your-username/xss-sanitizer-filters.git
cd xss-sanitizer-filters
docker-compose up --build
```

## Sending a Request

```
 curl 127.0.0.1:9090/js -d js='qwerty<img src=x onerror=alert(1)>'
```

## Example Output

```
go[bluemonday]: qwerty<img src="x">
js[sanitize-html]: qwerty
js[dompurify]: qwerty<img src="x">
python[html_sanitizer]: qwerty
```

