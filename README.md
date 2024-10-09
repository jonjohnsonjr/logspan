# logspan

https://x.com/kellabyte/status/1843436926214590633

## Install

```
go install github.com/jonjohnsonjr/logspan@latest
```

## Usage

Use with [`trot`](https://github.com/jonjohnsonjr/trot) or [`retrace`](https://github.com/jonjohnsonjr/retrace).

```
for i in $(seq 4); do echo -n $(date) && echo -n "," && sleep $i && echo -n $(date) && echo -n "," && echo "this is log number $i"; done | logspan | trot > trace.html && open trace.html
```

This assumes your input is formatted as:

```
<start>,<end>,<text2>
<start>,<end>,<text2>
```

Where `<start>` and `<end>` are both in the same format as `$(date)` output.
If that's not correct, fork it and fix it.
