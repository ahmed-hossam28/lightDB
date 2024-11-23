
# LightDB

**LightDB** is a lightweight, educational database project inspired by SQLite. It provides a simple implementation of a relational database system for learning purposes. The project is designed to demonstrate fundamental database management concepts, including memory management, page-based storage, persistent storage, and basic SQL query execution.

---

## Features

### Implemented Features
- **In-Memory Table Structure**: Data is stored in memory for fast operations and easier debugging.
- **Page-Based Storage**: Uses a fixed page size (4 KB) for row storage and retrieval.
- **Row Serialization**: Each row is serialized into a fixed-length byte array for structured storage.
- **Persistent Storage**: Data is now stored in a file, enabling disk-based storage similar to SQLite.
- **SQL Operations**:
  - `INSERT`: Add rows to the database.
  - `SELECT`: Retrieve rows from the database.
- **Dynamic Page Allocation**: Allocates additional pages as required when data grows.

### Planned Features
- [ ] **Advanced SQL Support**: Implement `UPDATE`, `DELETE`, and conditional queries.
- [ ] **Indexing with B-Trees**: Enhance query performance with B-tree-based indexing mechanisms.
- [ ] **Query Optimization**: Introduce techniques for faster query execution.
- [ ] **Concurrency and Transactions**: Add support for ACID compliance.

---

## Project Structure

```plaintext
lightDB
├── Makefile          # Build and run commands
├── README.md         # Project documentation
├── bin/              # Binary directory
│   └── lightdb       # Compiled binary
├── cmd/              # CLI entry point
│   └── lightdb-cli/
│       └── main.go   # Main function for the command-line interface
├── go.mod            # Dependency file
├── go.sum            # Dependency checksums
├── internal/         # Core database logic
│   ├── executor/     # SQL command executor
│   │   ├── metaCommand.go # Handles meta-commands (e.g., `.exit`)
│   │   └── statement.go   # Processes SQL statements
│   └── storage/      # Storage layer implementation
│       ├── db.go     # Database structure and entry point
│       ├── pager.go  # Handles page-level memory management
│       ├── row.go    # Row serialization and deserialization
│       └── table.go  # Table-level logic for managing rows and pages
├── temp/             # Temporary storage
│   └── test.db       # Sample database for testing
├── tests/            # Unit tests for core components
│   ├── db_test.go    # Database tests
│   ├── row_size_test.go 
│   └── test.go       # General tests
└── utils/            # Utility functions
    └── conv.go       # Conversion utilities (e.g., byte-to-struct)
```

---

## Usage

### Building the Project
Compile the `lightdb` CLI binary using the Makefile:
```bash
make build
```

### Running LightDB
Run the database CLI:
```bash
./bin/lightdb
```

### Interacting with LightDB
- **Meta-Commands**: Start with `.` (e.g., `.exit` to close the program).
- **SQL Commands**:
  - `INSERT ID USERNAME EMAIL` to add a row.
  - `SELECT` to retrieve all rows.

### Example Usage
```bash
db > INSERT 1 JohnDoe john@example.com
Executed.
db > INSERT 2 JaneSmith jane@example.com
Executed.
db > SELECT
(ID, Username, Email)
(1, JohnDoe, john@example.com)
(2, JaneSmith, jane@example.com)
Executed
db >
```

---

## Core Concepts

### Persistent Storage
- **File-Based Storage**: LightDB now persists data to a file using a custom binary format, ensuring data remains available across sessions.
- **File Format**: Data is stored as a sequence of pages, each containing serialized rows.
- **Recovery**: On startup, the database reads from the file to restore the state.

### Row Structure
- Each row has fixed fields:
  - `ID`: 4 bytes
  - `Username`: 32 bytes
  - `Email`: 255 bytes
- Rows are serialized into a fixed-length byte array for storage.

### Page-Based Storage
- **Page Size**: 4 KB.
- **Row Storage**: Rows are stored sequentially within a page.
- **Dynamic Growth**: Pages are added as needed when rows exceed current capacity.

---

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a feature branch.
3. Implement your changes.
4. Open a pull request.

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

---
