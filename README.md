DOC, the research publication manager
=====================================

Discover, download, and share bibliographies and research content quickly, and
easily. Work as a team. Make your research reproducible with trivial effort.

This is a BETA release of the application and associated search infrastructure
for demonstration purposes.


## Dependencies

This project requires golang 1.9 or greater to be available in the environment.

Install all application dependencies:

    go get

Install fpm and associated libraries to support release tasks. See fpm
documentation at http://fpm.readthedocs.io/en/latest/installing.html


## Build

Build a statically compiled application binary for the current platforms:

    go build -o builds/doc

Build a statically compiled application binary for all target platforms:

    ./tasks/build


## Test

Run command unit tests:

    ./tasks/test
    

## Release

    
## Contributing

doc is a community driven, open source project. We accept pull requests for
new features and bug fixes. Please see the CODE_OF_CONDUCT.md file for rules
governing participation in the project.

    
## License

This project is made available under the MIT License. See the LICENSE file for
details.
