# Golang TERYT Database Processing & Vector Search

## Overview
This project is a Golang-based system for downloading, parsing, and exporting TERYT (National Official Register of Territorial Division of the Country) databases from the Polish Central Statistical Office. The goal is to enhance Golang expertise while developing a custom vector search engine from scratch, without third-party libraries.

## Features
- **Automated Data Retrieval**: Fetches the latest TERYT datasets from government sources.
- **Data Parsing & Transformation**: Converts raw XML/CSV data into structured formats like JSON and SQL.
- **Database Integration**: Supports importing parsed data into relational and NoSQL databases.
- **Custom Vector Search**: Implements a vector-based retrieval system for efficient data querying, developed entirely from scratch.
- **High-Performance Processing**: Utilizes Golang concurrency for optimized data handling.

## Technology Stack
- **Golang**: Core programming language for system-level data processing.
- **Database Support**: PostgreSQL, SQLite, or other suitable databases for structured storage.
- **Vector Search Development**: Custom implementation exploring different indexing and retrieval techniques.
- **Parallel Processing**: Efficient data handling using Goroutines.

## Project Goals
1. **Deepen Golang Knowledge**: Hands-on experience with concurrency, file handling, and data processing.
2. **Build a Custom Vector Search Engine**: Develop a vector search system without relying on third-party libraries.
3. **Data Engineering & Optimization**: Create an efficient pipeline for managing Polish administrative datasets.

## Development Status
This project is in the early development phase. The initial focus is on implementing data extraction and structuring, followed by designing and optimizing the custom vector search engine.

## Roadmap
- [X] Implement automated TERYT data retrieval
- [ ] Develop parsing logic for XML/CSV formats
- [ ] Support multiple data export formats (JSON, SQL, CSV)
- [ ] Implement basic database import functionality
- [ ] Design and prototype vector search algorithms
- [ ] Optimize search performance and indexing

## Usage of Downloaded Data
The official source for downloading and licensing of TERYT data can be found here: [eTERYT - Pobieranie danych](https://eteryt.stat.gov.pl/eTeryt/rejestr_teryt/udostepnianie_danych/baza_teryt/uzytkownicy_indywidualni/pobieranie/pobieranie.aspx?contrast=default)

## Contribution
Currently, this is a personal learning project, but contributions and discussions are welcome! If you have insights or suggestions on vector search implementations, feel free to reach out.

## License
MIT
