package main

import (
    "os"
    "fmt"
    "flag"
    "encoding/json"

    "github.com/sug0/go-seq"
    "github.com/sug0/go-exit"
    "github.com/gofrs/flock"
)

var (
    keypath string
)

func main() {
    defer exit.Handler()

    var del, key, path string

    flag.StringVar(&path, "p", "", "The directory to save the IDs in.")
    flag.StringVar(&key, "k", "", "The key of the ID to generate.")
    flag.StringVar(&del, "f", "", "Delete the specified ID from a given key.")
    flag.Parse()

    if path == "" {
        exit.WithMsg(os.Stderr, 1, "%s: No given ID directory.", os.Args[0])
    }
    if key == "" {
        exit.WithMsg(os.Stderr, 1, "%s: No key given.", os.Args[0])
    }

    lockpath := fmt.Sprintf("%s%c%s.lock", path, os.PathSeparator, key)
    keypath = fmt.Sprintf("%s%c%s.json", path, os.PathSeparator, key)
    locker := flock.NewFlock(lockpath)

    locker.Lock()
    seq, err := openIds()
    if err != nil {
        locker.Unlock()
        exit.WithMsg(os.Stderr, 1, "%s: %s", os.Args[0], err)
    }
    defer locker.Unlock()
    defer saveIds(seq)

    if del != "" {
        if err = seq.Free([]byte(del)); err != nil {
            exit.WithMsg(os.Stderr, 1, "%s: Can't free the ID given: %s", os.Args[0], del)
        }
    } else {
        os.Stdout.Write(seq.Next())
        os.Stdout.Write([]byte{'\n'})
    }
}

func openIds() (*sequence.Seq, error) {
    f, err := os.Open(keypath)
    if err != nil {
        return sequence.NewSeq(), nil
    }
    defer f.Close()

    var seq sequence.Seq
    dec := json.NewDecoder(f)

    if err = dec.Decode(&seq); err != nil {
        return nil, err
    }
    return &seq, nil
}

func saveIds(seq *sequence.Seq) error {
    f, err := os.Create(keypath)
    if err != nil {
        return err
    }
    defer f.Close()

    enc := json.NewEncoder(f)
    if err = enc.Encode(seq); err != nil {
        return err
    }
    return nil
}
