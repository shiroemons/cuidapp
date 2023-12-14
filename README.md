# cuidgenapp

The cuidgenapp is a simple web application for generating Collision-Resistant Unique Identifiers (CUIDs). Developed using Go and utilizing the `lucsky/cuid` library, it provides a reliable way to generate unique identifiers.

## Usage

### Using curl

You can easily use the cuidgenapp through command line using curl. Here's how:

- To receive a CUID as plain text, simply use:
  ```bash
  curl https://cuidgen.fly.dev/
  # => clq51c5uz00008ir2a9ht4ldz
  ```

- To get the CUID in JSON format, include the `Content-Type: application/json` header in your request:
  ```bash
  curl -H "Content-Type: application/json" https://cuidgen.fly.dev/
  # => {"cuid":"clq51ccrz00018ir28qdqzje3"}
  ```

## License

[MIT License](LICENSE).
