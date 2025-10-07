# wca-ranking

A regional ranking for World Cube Association official results. Hosted at [ranking.leinadium.dev](ranking.leinadium.dev)

This project is a rewrite of my previous project [ranking-wca](github.com/Leinadium/ranking-wca)

## Structure

This repository contains the code for the backend (`api`) and the database (`db`).

### API

The API was rewritten using hexagonal architecture. This architecture was not required for this project, as this API is pretty straight forward, but it was used for the purpose of studying.

The API external ports (user interface and database interface) used external tools to enforce an implementation:

The user interface is implemented using [`oapi-codegen`](github.com/oapi-codegen/oapi-codegen). A OpenAPI v3 file contains the definitions, and this tool generate Go interfaces that the API has to implement.

The database interface is implemented using [`sqlc`](https://sqlc.dev/). A `.sql` file contains all the database queries, and this tool generates Go interfaces that the API has to implement.
