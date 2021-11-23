# GO Fiber - CRUD - GORM - Mysql
## Folder Structure
    - database          | database config
        |- migration    | migration config
    - middleware        | middleware
    - public            | public asset
        |- images       | asset image
    - route             | route config
    - src               | bisnis logic
        |- book         | book => model, response, service
        |- user         | user => model, response, service
    - utils             | utilities

## Was Implemented
    1.  CRUD
    2.  Database Connection using GORM Mysql Driver
    3.  Database Migration
    4.  Middleware
    5.  Passing Header Response
    6.  Basic Routing
    7.  Grouping Routes
    8.  Route Static Asset
    9.  Handle File Upload and Store to asset directory
    10. Handle Delete File from specific directory
    11. Get default value from struct (return result only string)