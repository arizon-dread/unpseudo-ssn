# Filter and match data containing SHA256 hashes with their clear text equivalents

The documentation exists in Swedish with Windows specific examples and in English with Linux specific examples. Scroll down for English and Linux.

Dokumentationen finns på svenska med Windows-specifika exempel och på engelska med Linux-specifika exempel. Skrolla ner för engelska och Linux.

## Svenska och Windows

### Förutsättningar

* Applikationen förväntar sig att exekveras från den mapp där följande filer finns:
* `input.txt` ska innehålla klartextdata, ett värde per rad, som motsvarar hashat data i följande fil,
* `hashed_data.txt` innehåller rader med en SHA256-hashad sträng någonstans på raden.
* Applikationen skapar `output.txt` när den körs med filtrerade rader från `hashed_data.txt`, bara de rader som har en hash som matchar innehållet i `input.txt` (när det hashas för jämförelse) kommer att komma med i `output.txt`.
* En `saltstring` är frivillig som första (`$1`) inparameter till programmet. Är klartextvärdet i indatat i `hashed_data.txt` hashat med en saltsträng måste den användas när programmet körs för att körningen ska kunna jämföra hashsträngar och få träff.
* Frivillig: `-d` flagga för "dirty"-läge (pga att känsligt data exponeras), gör att varje rad i `output.txt` då kommer få sitt klartextvärde istället för hashen.
* Frivillig: `-l` fragga för att lista klartextvärde (ssn) och hashat värde i nyckelvärdepar. Skapar filen `ssn_hash.txt` med klartextdata + motsvarande hash, ett par per rad. Blir då en nyckelfil för att kunna korsreferera innehållet i `output.txt`.

### Bygga och köra

* Utgivna binärer

    Se [Release-sidan](https://github.com/arizon-dread/unpseudo-ssn/releases) för mer information.

* Bygga

    `go build .` _kräver go 1.19 eller högre_

* Köra - exempel och utdata (förväntar sig att `unpseudo-ssn.exe`, `input.txt` och `hashed_data.txt` ligger i samma mapp, alla tre) byt ut hela "<saltsträng>" mot den faktiska saltsträngen, <> ska inte vara med i kommandot.

    Kommer bara att filtrera utdatarader där hash matchar inputdata + saltsträng:
    
    `.\unpseudo-ssn.exe <saltsträng>`

    Kommer ge filtrerade utdatarader med hash utbytt mot klartextvärde från `input.txt` ("dirty"-läge):
    
    `.\unpseudo-ssn.exe <saltsträng> -d`

    Kommer ge filtrerade rader och en separat nyckelfil med klartextvärde + matchande hash, en per rad:

    `.\unpseudo-ssn.exe <saltsträng> -l`

    Kommer ge filtrerade rader med klartextvärde i `output.txt` ("dirty"-läge) samt en separat nyckelfil med klartextvärde + matchande hash:

    `.\unpseudo-ssn.exe <saltsträng> -d -l`

    Alla dessa kan köras utan saltsträngen om det hashade datat är hashat utan saltsträng.

* Om du har Go installerat och har klonat ner repot så kan du köra koden direkt:

    `go run . [<saltsträng> -d -l]`

* Om du bara vill skapa en fil med klartextdata och deras korresponderande hash baserat på en specifik saltsträng kan du köra `.\pseudo-ssn.exe <saltsträng> -l` med en tom `hashed_data.txt`-fil fast med en `input.txt` innehållande rader med klartextdata. Du skulle då få en tom `output.txt` och en nyckelfil `ssn_hash.txt` som innehåller klartextdata + motsvarande hash, en per rad.

### Testdata

Innehållet i filen `hashed_data.txt` som följer med koden är hashat med följande saltsträng: 5e433dda-989b-11ed-98bf-00155d4e6602

## English and Linux

### Prerequisites

* Expects the binary to be executed from the folder where the following files reside.
* Expects `input.txt` to contain the cleartext data. The data to be hashed and compared to the hashed_data-lines.
* Expects `hashed_data.txt` to contain lines of data with SHA256 hashes somewhere on the line
* Outputs `output.txt` with only the lines containing unhashable strings (lines where the input.txt content when hashed, can be matched with the sha256 hash on the line) essentially a match-filter (except in "dirty" mode, see below).
* A `saltstring` is optional to the program as `$1` parameter.
* Optional: `-d` for "dirty" mode (as it exposes sensitive data), will replace the hash on each matched line with the clear text matched string
* Optional: `-l` for "list ssn hash", will output `ssn_hash.txt` containing clear text + corresponding hash, one on each line. 

### Building and running

* Released binaries

    See the [Releases page](https://github.com/arizon-dread/unpseudo-ssn/releases) for more info.

* Build

    `go build .` _requires go 1.19 or higher_

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

* If you only want to create a file with cleartext data and their hashes based on a specific saltstring, you could run `./pseudo-ssn <saltstring> -l` with an empty `hashed_data.txt` but with an `input.txt` containing data. You would then get an empty `output.txt` and a `ssn_hash.txt` file containing cleartext data and their SHA256 sum.

### Testing data
The `hashed_data.txt` content is hashed with this salt: 5e433dda-989b-11ed-98bf-00155d4e6602