// +build darwin

package bsdp

import "github.com/insomniacslk/dhcp/dhcpv4"

// Options (occur as sub-options of DHCP option 43).
const (
	OptionMessageType                   dhcpv4.OptionCode = 1
	OptionVersion                       dhcpv4.OptionCode = 2
	OptionServerIdentifier              dhcpv4.OptionCode = 3
	OptionServerPriority                dhcpv4.OptionCode = 4
	OptionReplyPort                     dhcpv4.OptionCode = 5
	OptionBootImageListPath             dhcpv4.OptionCode = 6 // Not used
	OptionDefaultBootImageID            dhcpv4.OptionCode = 7
	OptionSelectedBootImageID           dhcpv4.OptionCode = 8
	OptionBootImageList                 dhcpv4.OptionCode = 9
	OptionNetboot1_0Firmware            dhcpv4.OptionCode = 10
	OptionBootImageAttributesFilterList dhcpv4.OptionCode = 11
	OptionShadowMountPath               dhcpv4.OptionCode = 128
	OptionShadowFilePath                dhcpv4.OptionCode = 129
	OptionMachineName                   dhcpv4.OptionCode = 130
)

// Versions
var (
	Version1_0 = []byte{1, 0}
	Version1_1 = []byte{1, 1}
)

// MessageType represents the different BSDP message types.
type MessageType byte

// BSDP Message types - e.g. LIST, SELECT, FAILED
const (
	MessageTypeList MessageType = iota + 1
	MessageTypeSelect
	MessageTypeFailed
)

// BootImageType represents the different BSDP boot image types.
type BootImageType byte

// Different types of BootImages - e.g. for different flavors of macOS.
const (
	BootImageTypeMacOS9 BootImageType = iota
	BootImageTypeMacOSX
	BootImageTypeMacOSXServer
	BootImageTypeHardwareDiagnostics
	// 0x4 - 0x7f are reserved for future use.
)

// OptionCodeToString maps BSDP OptionCodes to human-readable strings
// describing what they are.
var OptionCodeToString = map[dhcpv4.OptionCode]string{
	OptionMessageType:                   " Message Type",
	OptionVersion:                       " Version",
	OptionServerIdentifier:              " Server Identifier",
	OptionServerPriority:                " Server Priority",
	OptionReplyPort:                     " Reply Port",
	OptionBootImageListPath:             "", // Not used
	OptionDefaultBootImageID:            " Default Boot Image ID",
	OptionSelectedBootImageID:           " Selected Boot Image ID",
	OptionBootImageList:                 " Boot Image List",
	OptionNetboot1_0Firmware:            " Netboot 1.0 Firmware",
	OptionBootImageAttributesFilterList: " Boot Image Attributes Filter List",
	OptionShadowMountPath:               " Shadow Mount Path",
	OptionShadowFilePath:                " Shadow File Path",
	OptionMachineName:                   " Machine Name",
}
