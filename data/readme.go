package data

// Default .gitignore file
var DefaultGitignore = `*~
lib/**/*
tmp`

// Default LICENSE file
var DefaultLicense = `Creative Commons Attribution 4.0 International License

This work is made available under the terms of the Creative Commons Attribution
4.0 International License. Please see http://creativecommons.org/licenses/by/4.0/
for more information.`

// Default README.md file
var DefaultReadme = `Bibliography
===========

This project uses **doc** package manager to manage citations and retrieval of
associated resources.

Fetch publication resources and save them to the ref folder:

    doc get

Fetch all resources and save them to the ref folder:

    doc get --all

Display the revision history for this project:

    doc log

For more information about the doc command, please visit the
[doc website](https://getdoc.io).


## License

This work is licensed under a [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/).
`