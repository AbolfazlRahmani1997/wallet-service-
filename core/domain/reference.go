package domain

import "github.com/oklog/ulid/v2"

type Reference struct {
	ID       ulid.ULID
	Token    string
	WalletId ulid.ULID
}
