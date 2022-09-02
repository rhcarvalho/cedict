match($0, "^^#! date=(.+)$", m) {
    print m[1]
    exit
}
