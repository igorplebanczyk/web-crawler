# Web Crawler

A simple web crawler that prints the number of times each internal link appears on a given website. Useful for SEO purposes.

## Usage

```bash
  crawler --website <website> --max-concurrency <max-concurrency> --max-pages <max-pages>
```
```bash
  crawler -w <website> -c <max-concurrency> -p <max-pages>
```

## Installation

### Option 1: Install from GitHub Releases

To install using a precompiled binary from the GitHub releases:

* Go to the Releases page and download the latest binary for your operating system.

* After downloading, make the binary executable:
    ```bash
    chmod +x crawler
    ```

* Move the binary to a directory in your $PATH (e.g., /usr/local/bin):
    ```bash
    sudo mv crawler /usr/local/bin/
    ```

* Confirm that Cachprax is installed by running:
    ```bash
    crawler --help
    ```

### Option 2: Compile and Install from Source

To compile and install the application yourself:

* Ensure that you have Go installed by running:
    ```bash
    go version
    ```

* Clone the repository:
    ```bash
    git clone https://github.com/your-repo/cachprax.git
    cd web-crawler
    ```

* Build the application:
    ```bash
    go build -o crawler
    ```

* Move the compiled binary to a directory in your $PATH:
    ```bash
    sudo mv crawler /usr/local/bin/
    ```

* Confirm that Cachprax is installed by running:
    ```bash
    crawler --help
    ```

## Notes
* Originally built as part of a guided project on [boot.dev](https://www.boot.dev/courses/build-web-crawler-golang), but has since been improved.
* Built primarily as a learning exercise.