module mail-server

go 1.18

replace github.com/boisvertmathieu/mail-server/server => ./server

replace github.com/boisvertmathieu/mail-server/client => ./client

require (
	github.com/boisvertmathieu/mail-server/client v0.0.0-00010101000000-000000000000
	github.com/boisvertmathieu/mail-server/server v0.0.0-00010101000000-000000000000
)
