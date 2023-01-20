# Replace hashed string with matching cleartext string

## Prerequisites

* Expects `input.txt` to contain the cleartext data.
* Expects `hashed_data` to contain lines of files with SHA256 hashes somewhere on the line
* Outputs `output.txt` with only the lines containing unhashable strings (lines where the input.txt content when hashed, can be matched with the sha256 hash on the line)
* A saltstring is optional to the program as $1 parameter.

## Building and running

* Build

    `go build .`

* Run

    `./unpseudo-ssn <saltstring>`

    or

    `go run . <saltstring>`

## Testing data
The hashed_data.txt content is hashed with this salt: 5e433dda-989b-11ed-98bf-00155d4e6602