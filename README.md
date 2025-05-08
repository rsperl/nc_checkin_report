# NC Checkin Report

## Build

To build, run `build.sh`.

## Usage

A license from [unidoc](https://cloud.unidoc.io/dashboard?nextUrl=) is required.
Set the value of the license to `UNIDOC_LICENSE_API_KEY`.

On MacOS, this would be

```shell
export UNIDOC_LICENSE_API_KEY=kfksjkadksadfkjsadkfjsakdfjskjfksajfksajdfksjdf
```

On Windows,

```shell
set UNIDOC_LICENSE_API_KEY=kfksjkadksadfkjsadkfjsakdfjskjfksajfksajdfksjdf
```

Remove **ALL** blank columns and rows. The library used to read the spreadsheet
omits blank rows, so leaving them in means you will not get all the data. Blank
columns after the `startsOnColumn` are not accounted for in the code, so they
must be removed as well. Refer to
[./sample_spreadsheet.xlsx](./sample_spreadsheet.xlsx) to see how the
spreadsheet should look.

Create a `config.yaml` such as

```yaml
# Name of the workbook containing data
# Note -- only the first sheet in the workbook will be read
infile: in.xlsx

# optional: defaults to title with colons removed
# outfile: output.docx

# Will be used as the title of each page.
# Also will be used to name the output word document if outfile is not
# specified.
title: "NC Checkin: Physical Science Student Data"

# The numerical row the header starts on
headerStartsOnRow: 7

# optional: The column containing the core.
core: E

# The number of points possible for this assessment
pointsPossible: 30

# The column containing the first name
firstName: B

# The column containing the last name
lastName: C

# The column containing the number correct
numberCorrect: F

# The column containing the percent correct
percentCorrect: G

# The column containing the number of items attempted
numberItemsAttempted: H

results:

  # The column where results start
  startsOnColumn: L

  # The number of result columns to expect
  # There should be NO blank columns in between the first and last
  # result column
  count: 13
```

Run as `./nc_checkin_report config.yaml`.