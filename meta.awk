match($0, "^^#! (date|time|entries)=(.+)$", m) {
    a[m[1]] = m[2]
}

/^[^#]/ {
    print a["date"]","a["time"]","a["entries"]
    exit
}
