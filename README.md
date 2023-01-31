# Filter and match data containing SHA256 hashes with their clear text equivalents

## Prerequisites

* Expects the binary to be executed from the folder where the following files reside.
* Expects `input.txt` to contain the cleartext data. The data to be hashed and compared to the hashed_data-lines.
* Expects `hashed_data.txt` to contain lines of data with SHA256 hashes somewhere on the line
* Outputs `output.txt` with only the lines containing unhashable strings (lines where the input.txt content when hashed, can be matched with the sha256 hash on the line) essentially a match-filter (except in "dirty" mode, see below).
* A `saltstring` is optional to the program as `$1` parameter.
* Optional: `-d` for "dirty" mode (as it exposes sensitive data), will replace the hash on each matched line with the clear text matched string
* Optional: `-l` for "list ssn hash", will output `ssn_hash.txt` containing clear text + corresponding hash, one on each line. 

## Building and running

* Released binaries

    See the [Releases page](https://github.com/arizon-dread/unpseudo-ssn/releases) for more info

* Build

    `go build .`

* Run

    Will only filter the lines to output lines with hashes matching input data + saltstring:

    `./unpseudo-ssn <saltstring>`

    Will output filtered lines with hash replaced by cleartext string ("dirty" mode):

    `./unpseudo-ssn <saltstring> -d`

    Will output filtered lines and a separate file with cleartext + matching hash:
    
    `./unpseudo-ssn <saltstring> -l`

    Will output filtered lines with cleartext string ("dirty") and also a separate file with cleartext + hash:

    `./pseudo-ssn <saltstring> -d -l`

    Any of these can be run without the saltstring.

* If you have Go installed, you can run the code directly:

    `go run . [ <saltstring> -d -l]`

## Testing data
The hashed_data.txt content is hashed with this salt: 5e433dda-989b-11ed-98bf-00155d4e6602