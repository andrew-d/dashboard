# Intro

Dashboard is a simple, self-hosted, extensible dashboard for varying types of
input.  Logically, input is split into "sources", each of which consists of a
name, an input type, and the accompanying data.


## Sources

A "source" represents a single source of data - some logically distinct origin
that feeds data to the dashboard.  Each source has a unique name, and a "type",
representing the format of the data that will be fed in.  Sources can also have
additional settings, depending on the input type.

## Types

Currently implemented types are:

### Good/Not

The simplest form of source.  Input consists of a list of items, each with an
accompanying "good" boolean.  The JSON input format is:

```json
[
  {"name": "item 1", "good": true},
  {"name": "item 2", "good": false}
]
```

### Item Status

The data fed in consists of some number of items, each with an accompanying
status.  This is similar to the "Good/Not" format, except that it supports
arbitrary statuses.

```json
[
  {"name": "item 1", "status": "loading"},
  {"name": "item 2", "status": "processing"},
  {"name": "item 3", "status": "finished"}
]
```
