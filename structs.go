package main

type User struct {
	UID      int16
	Username string
	Password string
	Email    string
}

type MapData struct {
	MapID      int16
	SongName   string
	ArtistName string
	Difficulty float32
}
