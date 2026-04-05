// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type ShortUrlDB struct {
	DSN string
}

type Config struct {
	rest.RestConf

	ShortUrlDB ShortUrlDB

	Sequence struct {
		DSN string
	}

	BaseString string

	ShortUrlBlackList []string

	ShortDomain string
}
