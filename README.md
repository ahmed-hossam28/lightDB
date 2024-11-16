# LightDB

**LightDB** is a lightweight, educational project inspired by SQLite. It aims to provide a simple implementation of a relational database system for learning purposes. Currently, it supports basic `SELECT` and `INSERT` operations, with more features planned for future development.

---

## Features (Work in Progress)

- **In-Memory Database**: LightDB stores data in-memory for fast operations.
- **Fixed-Size Pages**: Implements a 4 KB page size for row storage.
- **Row-Based Storage**: Rows are serialized and deserialized for structured data storage.
- **Table Management**: Ability to create tables and insert rows.
- **Basic SQL Operations**: Supports `INSERT` and `SELECT` operations.

---

## Current Status

### What's Implemented
- **In-Memory Table Structure**: Tables can hold rows in memory, and each table is divided into fixed-size pages (4 KB each).
- **Row Serialization**: Basic row format with `ID`, `Username`, and `Email` fields is supported.
- **Dynamic Page Allocation**: Pages are allocated dynamically as needed to store rows.
- **SQL Support**: Currently supports `INSERT` to add rows, and `SELECT` to retrieve rows.

### Planned Features
- [ ] Support for `UPDATE` and `DELETE` operations.
- [ ] Disk-based storage for persistence.
- [ ] Advanced memory management (e.g., caching).
- [ ] Query optimization and indexing.

---

## Project Structure

### File Layout

- **`main.go`**: Entry point for the database application.
- **`row.go`**: Contains functions for row serialization and deserialization.
- **`table.go`**: Manages table structure and page allocation.
- **`page.go`**: Handles page-based memory and row slot calculations.

---

### Core Concepts

1. **Row Structure**  
   - Each row consists of:
     - `ID`: 4 bytes
     - `Username`: 32 bytes
     - `Email`: 255 bytes
   - Rows are serialized into a fixed-length byte array for easy storage in pages.

2. **Pages**  
   - Fixed page size: 4 KB.
   - Each page can hold a certain number of rows, defined by `ROWS_PER_PAGE`.
   - Pages are allocated dynamically as needed when rows are
