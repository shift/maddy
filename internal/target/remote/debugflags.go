//+build debugflags

/*
Maddy Mail Server - Composable all-in-one email server.
Copyright © 2019-2020 Max Mazurov <fox.cpp@disroot.org>, Maddy Mail Server contributors

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package remote

import (
	maddycli "github.com/foxcpp/maddy/internal/cli"
	"github.com/urfave/cli/v2"
)

func init() {
	maddycli.AddGlobalFlag(&cli.StringFlag{
		Name:        "debug.smtpport",
		Usage:       "SMTP port to use for connections in tests",
		Destination: &smtpPort,
	})
}
