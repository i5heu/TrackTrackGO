package main

import (
	"crypto/sha256"
	"encoding/hex"
)

var guestmode bool = true
var dblogin string = "USER:PASSWORD@/TrackTrackGO"
var personalpwdTMP string = "track"

var templatefolder string = "/home/her/CODE/TrackTrackGO/template"

/*############# END OF CONFIG ################*/

var foo2355523 = sha256.Sum256([]byte(personalpwdTMP))
var personalpwd = hex.EncodeToString(foo2355523[:])
