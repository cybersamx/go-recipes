# Go Time Format

Most programming languages use an alphabetic layout eg. yyyy-MM-dd to represent the time format. In Go, a time format is based on a layout that references this "" date `Mon Jan 2 15:04:05 MST 2006`, in which each time component is represented by a magical value in the reference date.

## Go Code

### Parse Time

Parse date string `26 Jul 2020 13:47:56 PDT` to a Go time struct value.

```go
datestr := "26 Jul 2020 13:47:56 PDT"
layout := "2 Jan 2006 15:04:05 MST"
date, err := time.Parse(layout, datestr)
```

or use a pre-defined layout.

```go
datestr = "2021-01-27T12:08:38.093906-08:00"
layout = time.RFC3339Nano
date, err = time.Parse(layout, datestr)
```

### Print Time

Print Go time struct value as `Sun, 26 Jul 2020 13:47:56 PDT`.

```go
date := time.Date(2020, 7, 26, 13, 47, 56, 0, time.UTC)
layout := "Mon, 2 Jan 2006 15:04:05 MST"
fmt.Println(date.Format(layout))
```

## Layouts

So let's suppose we have the following timestamp `26 Jul 2020 13:47:56`. Here are the notations you would to display the timestamp in the formats you want:

### Full Datetime Layouts

| Expected output                  | Go notation                          | Standard | Go Const         |
|----------------------------------|--------------------------------------|----------|------------------|
| 26 Jul 20 13:47 PDT              | 02 Jan 06 15:04 MST                  | RFC822   | time.RFC822      |
| 26 Jul 20 13:47 -0700            | 02 Jan 06 15:04 -0700                | RFC822   | time.RFC822Z     |
| Sun, 26 Jul 2020 13:47:56 PDT    | Mon, 2 Jan 2006 15:04:05 MST         | RFC 1123 | time.RFC1123     |
| Sun, 26 Jul 2020 13:47:56 -0700  | Mon, 2 Jan 2006 15:04:05 -0700       | RFC 1123 | time.RFC1123Z    |
| 2020-07-26T13:47:56-0700         | 2006-01-02T15:04:05Z07:00	          | RFC3339  | time.RFC3339     |
| 2020-07-26T13:47:56.093906-07:00 | 2006-01-02T15:04:05.999999999Z07:00  | RFC3339  | time.RFC3339Nano |

When the datetime string representation does not include the timezone, the UTC timezone is assumed. For example, the following will be parsed as `2020-07-26 13:47:56 +0000 UTC`.

| Expected output       |
|-----------------------|
| 2020-07-26T13:47:56   |
| 26 Jul 2020 13:47:56  |

### Date-Only Layouts

| Expected output | Go notation      | Java notation | C notation | Standard |
|-----------------|------------------|---------------|------------|----------|
| 2020-07-26      | 2006-01-02       | yyyy-MM-dd    | %F         | ISO 8601 |
| 20200726        | 20060102         | yyyyMMdd      | %Y%m%d     | ISO 8601 |
| July 26, 2020   | January 02, 2006 | MMMM dd, yyyy | %B %d, %Y  |          |
| 26 Jul, 2020    | 02 Jan, 2006     | dd MMM, yyyy  | %d %B, %Y  |          |
| 07/26/20        | 01/02/06         | MM/dd/yy      | %D         |          |
| 07/26/2020      | 01/02/2006       | MM/dd/yyyy    | %m/%d/%Y   |          |
| Sunday          | Monday           | EEEE          | %A         |          |
| Sun             | Mon              | EEE           | %a         |          |

### Time-Only Layout

| Expected output | Go notation      | Java notation | C notation | Standard |
|-----------------|------------------|---------------|------------|----------|
| 13:47           | 15:04            | HH:mm	     | %R         |          |
| 13:47:56        | 15:04:05         | HH:mm:ss	     | %T         | ISO 8601 |
| 1:47 PM         | 3:04 PM	         | K:mm a	     | %l:%M %p   |          |
| 01:47:56 PM     | 03:04:05 PM	     | KK:mm:ss a	 | %r         |          |

## Time Zone

You can define the time zone of a time value using the `time.Location` struct. Typically, you will call the `LoadLocation` function with a passed name, which corresponds to the TZ database name found in the IANA Time Zone database, eg. America/Los_Angeles. See [list of TZ database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for details. The `LoadLocation` function also accepts "Local" as a valid name and loads the local location.

## Reference

* [Go: format a time or date](https://programming.guide/go/format-parse-string-time-date-example.html)
