//    The Source file for the Scrape command in the toolbox.
//    Copyright (C) 2026  solosrc.com
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <https://www.gnu.org/licenses/>.

package scrape

import (
	"encoding/json"
	"fmt"

	"github.com/solosrc/goaty-oh/toolbox/scraper"
	"github.com/spf13/cobra"
)

var url string

var Command = &cobra.Command{
	Use:   "scrape",
	Short: "Scrapes the web for card details",
	Long:  `Runs a Web Scraper that retrieves card details and saves it a the database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cards, err := scraper.Visit(url)
		if err != nil {
			return err
		}
		for _, card := range cards {
			bytes, err := json.Marshal(card)
			if err != nil {
				return err
			}
			fmt.Println(string(bytes))
		}
		return nil
	},
}

func init() {
	Command.PersistentFlags().StringVar(&url, "url", "", "the URL to scrape")
}
