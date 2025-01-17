package partid

import "fmt"

// nolint:gochecknoglobals,lll
var (
	// Empty - Unused / Empty partition
	Empty = [16]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

	// LUKS - Linux LUKS (CA7D7CCB-63ED-4C53-861C-1742536059CC)
	LUKS = [16]byte{0xcb, 0x7c, 0x7d, 0xca, 0xed, 0x63, 0x53, 0x4c, 0x86, 0x1c, 0x17, 0x42, 0x53, 0x60, 0x59, 0xcc}

	// LinuxFS - Linux generic filesytem data (0FC63DAF-8483-4772-8E79-3D69D8477DE4)
	LinuxFS = [16]byte{0xaf, 0x3d, 0xc6, 0xf, 0x83, 0x84, 0x72, 0x47, 0x8e, 0x79, 0x3d, 0x69, 0xd8, 0x47, 0x7d, 0xe4}

	// LinuxRAID - Linux RAID data (A19D880F-05FC-4D3B-A006-743F0F84911E)
	LinuxRAID = [16]byte{0xf, 0x88, 0x9d, 0xa1, 0xfc, 0x5, 0x3b, 0x4d, 0xa0, 0x6, 0x74, 0x3f, 0xf, 0x84, 0x91, 0x1e}

	// LinuxRootX86 - Linux ia32 rootfs (44479540-F297-41B2-9AF7-D131D5F0458A)
	LinuxRootX86 = [16]byte{0x40, 0x95, 0x47, 0x44, 0x97, 0xf2, 0xb2, 0x41, 0x9a, 0xf7, 0xd1, 0x31, 0xd5, 0xf0, 0x45, 0x8a}

	// LinuxRootArm32 - Linux arm32 rootfs (69DAD710-2CE4-4E3C-B16C-21A1D49ABED3)
	LinuxRootArm32 = [16]byte{0x10, 0xd7, 0xda, 0x69, 0xe4, 0x2c, 0x3c, 0x4e, 0xb1, 0x6c, 0x21, 0xa1, 0xd4, 0x9a, 0xbe, 0xd3}

	// LinuxRootX86_64 - Linux x86_64 rootfs (4F68BCE3-E8CD-4DB1-96E7-FBCAF984B709)
	LinuxRootX86_64 = [16]byte{0xe3, 0xbc, 0x68, 0x4f, 0xcd, 0xe8, 0xb1, 0x4d, 0x96, 0xe7, 0xfb, 0xca, 0xf9, 0x84, 0xb7, 0x9}

	// LinuxRootArm64 - Linux/ arm64 rootfs (B921B045-1DF0-41C3-AF44-4C6F280D3FAE)
	LinuxRootArm64 = [16]byte{0x45, 0xb0, 0x21, 0xb9, 0xf0, 0x1d, 0xc3, 0x41, 0xaf, 0x44, 0x4c, 0x6f, 0x28, 0xd, 0x3f, 0xae}

	// LinuxDMCrypt - Linux dm-crypt data (7FFEC5C9-2D00-49B7-8941-3EA10A5586B7)
	LinuxDMCrypt = [16]byte{0xc9, 0xc5, 0xfe, 0x7f, 0x0, 0x2d, 0xb7, 0x49, 0x89, 0x41, 0x3e, 0xa1, 0xa, 0x55, 0x86, 0xb7}

	// LinuxReserved - Linux Reserved (8DA63339-0007-60C0-C436-083AC8230908)
	LinuxReserved = [16]byte{0x39, 0x33, 0xa6, 0x8d, 0x7, 0x0, 0xc0, 0x60, 0xc4, 0x36, 0x8, 0x3a, 0xc8, 0x23, 0x9, 0x8}

	// LinuxLVM - Linux LVM data (E6D6D379-F507-44C2-A23C-238F2A3DF928)
	LinuxLVM = [16]byte{0x79, 0xd3, 0xd6, 0xe6, 0x7, 0xf5, 0xc2, 0x44, 0xa2, 0x3c, 0x23, 0x8f, 0x2a, 0x3d, 0xf9, 0x28}

	// LinuxBoot - Linux /boot fs (BC13C2FF-59E6-4262-A352-B275FD6F7172)
	LinuxBoot = [16]byte{0xff, 0xc2, 0x13, 0xbc, 0xe6, 0x59, 0x62, 0x42, 0xa3, 0x52, 0xb2, 0x75, 0xfd, 0x6f, 0x71, 0x72}

	// LinuxSwap - Linux swap data (0657FD6D-A4AB-43C4-84E5-0933C84B4F4F)
	LinuxSwap = [16]byte{0x6d, 0xfd, 0x57, 0x6, 0xab, 0xa4, 0xc4, 0x43, 0x84, 0xe5, 0x9, 0x33, 0xc8, 0x4b, 0x4f, 0x4f}

	// LinuxHome - Linux /home fs (933AC7E1-2EB4-4F13-B844-0E14E2AEF915)
	LinuxHome = [16]byte{0xe1, 0xc7, 0x3a, 0x93, 0xb4, 0x2e, 0x13, 0x4f, 0xb8, 0x44, 0xe, 0x14, 0xe2, 0xae, 0xf9, 0x15}

	// LinuxSrv - Linux /srv fs (3B8F8425-20E0-4F3B-907F-1A25A76F98E8)
	LinuxSrv = [16]byte{0x25, 0x84, 0x8f, 0x3b, 0xe0, 0x20, 0x3b, 0x4f, 0x90, 0x7f, 0x1a, 0x25, 0xa7, 0x6f, 0x98, 0xe8}

	// MBR - MBR type (024DEE41-33E7-11D3-9D69-0008C781F39F)
	MBR = [16]byte{0x41, 0xee, 0x4d, 0x2, 0xe7, 0x33, 0xd3, 0x11, 0x9d, 0x69, 0x0, 0x8, 0xc7, 0x81, 0xf3, 0x9f}

	// BiosBoot - Bios Boot Partition (BBP) (21686148-6449-6E6F-744E-656564454649)
	BiosBoot = [16]byte{0x48, 0x61, 0x68, 0x21, 0x49, 0x64, 0x6f, 0x6e, 0x74, 0x4e, 0x65, 0x65, 0x64, 0x45, 0x46, 0x49}

	// EFI - EFI Partition (C12A7328-F81F-11D2-BA4B-00A0C93EC93B)
	EFI = [16]byte{0x28, 0x73, 0x2a, 0xc1, 0x1f, 0xf8, 0xd2, 0x11, 0xba, 0x4b, 0x0, 0xa0, 0xc9, 0x3e, 0xc9, 0x3b}

	// AtxReserved - F79962B9-24E6-9948-9F94-E6BFDAD2771F
	AtxReserved = [16]byte{0xb9, 0x62, 0x99, 0xf7, 0xe6, 0x24, 0x48, 0x99, 0x9f, 0x94, 0xe6, 0xbf, 0xda, 0xd2, 0x77, 0x1f}

	// AtxPBF - 01A3E19F-9FEA-ED47-92C2-E75639FF5601
	AtxPBF = [16]byte{0x9f, 0xe1, 0xa3, 0x01, 0xea, 0x9f, 0x47, 0xed, 0x92, 0xc2, 0xe7, 0x56, 0x39, 0xff, 0x56, 0x01}

	// AtxSBF - 01A3E19F-9FEA-ED47-92C2-E75639FF5602
	AtxSBF = [16]byte{0x9f, 0xe1, 0xa3, 0x01, 0xea, 0x9f, 0x47, 0xed, 0x92, 0xc2, 0xe7, 0x56, 0x39, 0xff, 0x56, 0x02}

	// AtxSignData - 8DEE18B1-77B5-1D25-AE89-2A252D1A422F
	AtxSignData = [16]byte{0xb1, 0x18, 0xee, 0x8d, 0xb5, 0x77, 0x25, 0x1d, 0xae, 0x89, 0x2a, 0x25, 0x2d, 0x1a, 0x42, 0x2f}

	// StoragedRaw - 26843217-D7A8-48E8-BBFC-6870C69BA060
	StoragedRaw = [16]byte{0x17, 0x32, 0x84, 0x26, 0xa8, 0xd7, 0xe8, 0x48, 0xbb, 0xfc, 0x68, 0x70, 0xc6, 0x9b, 0xa0, 0x60}

	// StoragedLVM - D5842A1E-DF14-4129-94DB-9C06DF842179
	StoragedLVM = [16]byte{0x1e, 0x2a, 0x84, 0xd5, 0x14, 0xdf, 0x29, 0x41, 0x94, 0xdb, 0x9C, 0x06, 0xdf, 0x84, 0x21, 0x79}
)

// Text gives human readable names
var Text = map[[16]byte]string{ // nolint:gochecknoglobals
	EFI:             "EFI",
	LUKS:            "LUKS",
	LinuxFS:         "Linux-FS",
	LinuxRAID:       "RAID",
	LinuxRootX86:    "Linux-root-x86",
	LinuxRootArm32:  "Linux-root-arm32",
	LinuxRootX86_64: "Linux-root-x86_64",
	LinuxRootArm64:  "Linux-root-arm64",
	LinuxDMCrypt:    "DM-crypt",
	LinuxReserved:   "Linux-Reserved",
	LinuxLVM:        "LVM",
	LinuxBoot:       "Linux-/boot",
	LinuxSwap:       "Linux-swap",
	LinuxHome:       "Linux-/home",
	LinuxSrv:        "Linux-/srv",
	MBR:             "MBR",
	BiosBoot:        "Bios-Boot",
	StoragedRaw:     "Storaged-Raw",
	StoragedLVM:     "Storaged-LVM",
	AtxPBF:          "Atomix-PBF",
	AtxReserved:     "Atomix-Reserved",
	AtxPBF:          "Atomix-PBF",
	AtxSBF:          "Atomix-SBF",
	AtxSignData:     "Atomix-SignData",
}

// nolint: gochecknoglobals,gomnd
var mapGPTToMBR = map[[16]byte]byte{
	Empty:     0x00,
	LinuxSwap: 0x82,
	LinuxFS:   0x83,
	LinuxLVM:  0x8E,
	LUKS:      0xE8,
}

// PartTypeToMBR - Convert a GPT Type to its MBR equivalent
func PartTypeToMBR(gptType [16]byte) (byte, error) {
	if val, ok := mapGPTToMBR[gptType]; ok {
		return val, nil
	}

	padded := true

	for i := 0; i < 15; i++ {
		if gptType[i] != 0 {
			padded = false
			break
		}
	}

	if padded {
		return gptType[15], nil
	}

	return 0, fmt.Errorf("unknown MBR type %v", gptType)
}
