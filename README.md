# Description

`idgen` is a command line tool that allows you to generate
alphabetical sequential IDs, for whatever applications
may suit you (for instance, to generate ids for a static
file http server of some sort).

# Install

```
$ go get github.com/sugoiuguu/idgen
```

# Usage

```
$ idgen -h
Usage of idgen:
  -f string
        delete the specified key
  -k string
        the key of the id
  -p string
        the dir to save the ids in
```

## Example

```
$ alias idgen="idgen -p ~/.idgen"
$ id=$(idgen -k files)
$ echo wow it works > /var/www/static/$id.txt
.
.
.
$ idgen -k files -f a    # frees the id 'a'
$ id=$(idgen -k files)   # returns 'a' in $id
```
