package main

import (
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/windows"
)

// MAPI Konstanten
const (
	MAPI_LOGON_UI = 0x00000001
	MAPI_DIALOG   = 0x00000008
	MAPI_TO       = 1
	SUCCESS_SUCCESS = 0
)

// MAPI Strukturen
type MapiFileDesc struct {
	UlReserved   uint32
	FlFlags      uint32
	NPosition    uint32
	LpszPathName *uint16
	LpszFileName *uint16
	LpFileType   uintptr
}

type MapiRecipDesc struct {
	UlReserved   uint32
	UlRecipClass uint32
	LpszName     *uint16
	LpszAddress  *uint16
	UlEIDSize    uint32
	LpEntryID    uintptr
}

type MapiMessage struct {
	UlReserved          uint32
	LpszSubject         *uint16
	LpszNoteText        *uint16
	LpszMessageType     *uint16
	LpszDateReceived    *uint16
	LpszConversationID *uint16
	FlFlags             uint32
	LpOriginator        uintptr
	NRecipCount         uint32
	LpRecips            uintptr
	NFileCount          uint32
	LpFiles             uintptr
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <FilePath> [Subject] [Recipient]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filePath := os.Args[1]
	var subject, recipient string
	if len(os.Args) > 2 {
		subject = os.Args[2]
	}
	if len(os.Args) > 3 {
		recipient = os.Args[3]
	}

	// MAPI DLL laden
	mapi32 := windows.NewLazyDLL("mapi32.dll")
	// procMAPISendMail := mapi32.NewProc("MAPISendMail")
	procMAPISendMail := mapi32.NewProc("MAPISendMailW")

	filePathUTF16, err := windows.UTF16PtrFromString(filePath)
	if err != nil {
		fmt.Printf("Fehler beim Pfad-Encoding: %v\n", err)
		os.Exit(1)
	}

	fileDesc := MapiFileDesc{
		UlReserved:   0,
		FlFlags:      0,
		NPosition:    0xFFFFFFFF,
		LpszPathName: filePathUTF16,
		LpszFileName: nil,
		LpFileType:   0,
	}

	var message MapiMessage
	message.NFileCount = 1
	message.LpFiles = uintptr(unsafe.Pointer(&fileDesc))

	if subject != "" {
		subjUTF16, _ := windows.UTF16PtrFromString(subject)
		message.LpszSubject = subjUTF16
	}

	if recipient != "" {
		recipUTF16, _ := windows.UTF16PtrFromString(recipient)
		recipDesc := MapiRecipDesc{
			UlReserved:   0,
			UlRecipClass: MAPI_TO,
			LpszName:     recipUTF16,
			LpszAddress:  recipUTF16,
		}
		message.NRecipCount = 1
		message.LpRecips = uintptr(unsafe.Pointer(&recipDesc))
	}

	// Aufruf von MAPISendMail
	ret, _, _ := procMAPISendMail.Call(
		0,
		0,
		uintptr(unsafe.Pointer(&message)),
		uintptr(MAPI_DIALOG|MAPI_LOGON_UI),
		0,
	)

	if ret == SUCCESS_SUCCESS {
		fmt.Println("Erfolgreich an MAPI uebergeben.")
	} else {
		fmt.Printf("MAPI Fehlercode: %d\n", ret)
	}
}