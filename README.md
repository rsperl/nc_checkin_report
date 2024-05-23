# NC Checkin Report

## Build
To build, run `build.sh`.

## Usage

A license from [unidoc](https://cloud.unidoc.io/dashboard?nextUrl=) is required. Set the value
of the license to `UNIDOC_LICENSE_API_KEY`.

Create a `config.yaml` such as

```yaml
infile: in.xlsx
# defaults to title with colons removed
outfile: output.docx
title: "NC Checkin: Physical Science Student Data"
headerStartsOnRow: 7
pointsPossible: 30
firstName: B
lastName: C
numberCorrect: E
percentCorrect: F
numberItemsAttempted: G
results:
  startsOnColumn: Q
  count: 10
```

Run as `./nc_checkin_report config.yaml`.