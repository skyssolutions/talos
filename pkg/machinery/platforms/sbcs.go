// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package platforms

import (
	"github.com/blang/semver/v4"

	"github.com/siderolabs/talos/pkg/machinery/imager/quirks"
)

// SBC describes a Single Board Computer configuration.
type SBC struct {
	Name string

	// For Talos < 1.7.
	BoardName string

	// For Talos 1.7+
	OverlayName  string
	OverlayImage string

	Label         string
	Documentation string

	MinVersion semver.Version
}

// DiskImagePath returns the path to the disk image for the SBC.
func (s SBC) DiskImagePath(talosVersion string) string {
	if quirks.New(talosVersion).SupportsOverlay() {
		return "metal-arm64.raw.xz"
	}

	return "metal-" + s.BoardName + "-arm64.raw.xz"
}

// SBCs returns a list of supported Single Board Computers.
func SBCs() []SBC {
	return []SBC{
		{
			Name: "rpi_generic",

			BoardName: "rpi_generic",

			OverlayName:  "rpi_generic",
			OverlayImage: "skyssolutions/sbc-raspberrypi",

			Label: "Raspberry Pi Series (fork)",
		},
		{
			Name: "rpi_5",

			BoardName: "rpi_5",

			OverlayName:  "rpi_5",
			OverlayImage: "skyssolutions/sbc-raspberrypi",

			Label: "Raspberry Pi 5",
		},
	}
}
