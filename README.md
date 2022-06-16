# csv_unique_jsArr
## Usage
```bash
$ go build
```
```
$ ./csv_unique_jsArr <CSV Path>
```
---
## Example
### CSV content
```csv
2022/4/1 (金),12:00,21:30
2022/4/4 (月),11:00,19:15
2022/4/6 (水),12:00,21:30
2022/4/8 (金),12:00,21:30
```
  
### Output
*VERRRRRYYYYY SIMPLE*
```js
{ Since: 12.00, Until: 21.5, Days: 3 },
{ Since: 11, Until: 19.25, Days: 1 },
```
