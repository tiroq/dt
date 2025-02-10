# DevTree

DevTree is a command-line tool written in Go that recursively walks your project directory, printing the directory structure and file contents (with delimiters) to an output file. It respects your `.gitignore` and always ignores the `.git` folder, making it ideal for development projects.

## Features

- Recursively walks the current directory.
- Outputs the directory tree and file contents with clear delimiters.
- Respects `.gitignore` rules.
- Ignores the `.git` folder by default.
- Allows customization of the output file via a CLI flag.

## Installation

1. Ensure [Go](https://golang.org/dl/) is installed.
2. Clone the repository:

    ```bash
    git clone https://github.com/tiroq/dt.git
    cd dt
    ```

3. Build and install:

    ```bash
    make build
    make install
    ```


## Usage

Build the application:

```bash
make build
```

Run with default settings:

```bash
./dt
```

Specify a custom output file:

```bash
./dt -output=custom_structure.txt
```

Install:

```bash
make install
```

## AI Integration Note

The generated output file is structured to be easily parsed by AI tools, enabling automated analysis or navigation of your project structure.

## License

Distributed under the MIT License. See LICENSE for more information.

## Contributing

Contributions are welcome! Open an issue or submit a pull request to help improve DevTree.