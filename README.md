# SendMapiMail

[Deutsch](#deutsch) | [English](#english)

---

<a name="deutsch"></a>
## Deutsch

Ein leichtgewichtiges Befehlszeilenwerkzeug für Windows zum Öffnen einer neuen E-Mail mit Dateianhang im Standard-E-Mail-Programm (via Simple MAPI).

### Hintergrund & Zweck

Klassische `mailto:`-Hyperlinks können aus Sicherheitsgründen standardmäßig keine Dateianhänge übergeben. Viele Skripte (PowerShell, Batch) oder Drittanbieterprogramme (z. B. ERP-, DMS- oder Faktura-Software) müssen jedoch vorbereitete Dokumente (wie Rechnungen, Protokolle oder Berichte) direkt an das E-Mail-Programm des Anwenders übergeben, damit dieser vor dem Versand noch Text hinzufügen oder Anpassungen vornehmen kann.

**SendMapiMail** wurde primär entwickelt, um eine neue E-Mail mit Anhang im **Tobit David / Team Client** (oder anderen MAPI-kompatiblen Mailclients wie Microsoft Outlook, Mozilla Thunderbird etc.) interaktiv im Editorfenster zu öffnen.

### 32-Bit vs. 64-Bit MAPI

Windows-MAPI-Aufrufe (`mapi32.dll`) erfordern, dass die Bittiefe der aufrufenden Anwendung mit der Bittiefe des installierten E-Mail-Clients übereinstimmt:

- **`SendMapiMail32.exe` (32-Bit / x86):**  
  Erforderlich für **Tobit David / Team Client** sowie für 32-Bit-Versionen von Microsoft Outlook und anderen 32-Bit-Mailclients.
- **`SendMapiMail64.exe` (64-Bit / x64):**  
  Erforderlich für reine 64-Bit-Mailclients (z. B. 64-Bit-Installationen von Microsoft Office / Outlook).

> [!TIP]
> Für Tobit David / Team Client immer die **32-Bit-Version** (`SendMapiMail32.exe`) verwenden!

### Funktionen

- Öffnet das E-Mail-Fenster im Standard-MAPI-Client mit aktivem Eingabedialog (`MAPI_DIALOG | MAPI_LOGON_UI`).
- Übergabe eines Dateipfads als Anhang.
- Optionale Angabe von Empfänger-Adresse (`To`) und Betreff (`Subject`).
- Keine Abhängigkeit von externen Laufzeitumgebungen (reines Go-Binary).
- Ideal für den Aufruf aus Batch-Dateien, PowerShell, VBScript, Python oder externer Software.

### Syntax

```cmd
SendMapiMail32.exe <Dateipfad> [Betreff] [Empfaenger]
```
*(bzw. `SendMapiMail64.exe` bei 64-Bit-Mailclients)*

#### Parameter:
- `<Dateipfad>` *(erforderlich)*: Vollständiger oder relativer Pfad zur anzuhängenden Datei. Pfade mit Leerzeichen müssen in Anführungszeichen gesetzt werden.
- `[Betreff]` *(optional)*: Betreffzeile der E-Mail.
- `[Empfaenger]` *(optional)*: E-Mail-Adresse des Empfängers.

### Anwendungsbeispiele

#### 1. Nur Dateianhang öffnen:
```cmd
SendMapiMail32.exe "C:\Dokumente\Rechnung_12345.pdf"
```

#### 2. Anhang und Betreff:
```cmd
SendMapiMail32.exe "C:\Dokumente\Bericht.xlsx" "Monatsbericht"
```

#### 3. Anhang, Betreff und Empfänger:
```cmd
SendMapiMail32.exe "C:\Dokumente\Angebot.pdf" "Angebot Nr. 2026-04" "vertrieb@firma.de"
```

#### 4. Aufruf aus PowerShell:
```powershell
Start-Process -FilePath ".\SendMapiMail32.exe" -ArgumentList '"C:\Temp\Log.txt"', '"Systembericht"', '"admin@firma.local"'
```

### Kompilierung / Build

Das Projekt ist in [Go](https://go.dev/) geschrieben. Über das beiliegende Build-Skript [`make.bat`](make.bat) werden automatisch sowohl die 32-Bit- als auch die 64-Bit-Binary erzeugt:

```cmd
@echo off

set GOARCH=386
go build -o SendMapiMail32.exe .

set GOARCH=amd64
go build -o SendMapiMail64.exe .
```

Einfach in der Eingabeaufforderung ausführen:
```cmd
make.bat
```

Nach erfolgreichem Build stehen im Verzeichnis:
- `SendMapiMail32.exe` (für Tobit David & 32-Bit Clients)
- `SendMapiMail64.exe` (für 64-Bit Clients)

### Rückgabewerte / Fehlercodes

Das Programm gibt bei erfolgreicher Übergabe an den MAPI-Spooler `0` zurück. Bei Fehlern wird der MAPI-Fehlercode ausgegeben (z. B. `1` für `MAPI_USER_ABORT`, wenn der Benutzer das Fenster abbricht/schließt).

### Credits

Sowohl der Quellcode dieses Programms als auch das Anwendungs-Icon wurden mithilfe von **Google Gemini** generiert.

---

<a name="english"></a>
## English

A lightweight command-line tool for Windows to compose a new email with an attachment in the default email client via Simple MAPI.

### Background & Purpose

Standard `mailto:` URLs do not support file attachments due to security restrictions. However, scripts (PowerShell, Batch) and business software (such as ERP, CRM, DMS, or invoicing tools) often need to generate a document (e.g., invoices, logs, reports) and pass it to the user's default email client, allowing the user to review the message, edit content, and send it interactively.

**SendMapiMail** was primarily created to open a new message with an attachment in the **Tobit David / Team Client** (as well as other MAPI-compliant clients such as Microsoft Outlook or Mozilla Thunderbird).

### 32-Bit vs. 64-Bit MAPI

Windows MAPI calls (`mapi32.dll`) require the bitness of the calling executable to match the bitness of the installed email client:

- **`SendMapiMail32.exe` (32-bit / x86):**  
  Required for **Tobit David / Team Client** as well as 32-bit editions of Microsoft Outlook and other 32-bit mail clients.
- **`SendMapiMail64.exe` (64-bit / x64):**  
  Required for 64-bit mail clients (e.g., 64-bit Microsoft Office / Outlook installations).

> [!TIP]
> Always use the **32-bit version** (`SendMapiMail32.exe`) when working with Tobit David / Team Client!

### Features

- Launches the compose window in the default MAPI client with an interactive dialog (`MAPI_DIALOG | MAPI_LOGON_UI`).
- Attaches any file via its file path.
- Optional recipient (`To`) and subject line.
- Standalone executable with zero external dependencies.
- Perfect for integration into command line scripts, scheduled tasks, and third-party tools.

### Syntax

```cmd
SendMapiMail32.exe <FilePath> [Subject] [Recipient]
```
*(or `SendMapiMail64.exe` for 64-bit email clients)*

#### Arguments:
- `<FilePath>` *(required)*: Absolute or relative path to the file to attach. Enclose paths with spaces in double quotes.
- `[Subject]` *(optional)*: Subject text of the email.
- `[Recipient]` *(optional)*: Recipient email address.

### Usage Examples

#### 1. Open with attachment only:
```cmd
SendMapiMail32.exe "C:\Documents\Invoice_12345.pdf"
```

#### 2. Attachment and subject:
```cmd
SendMapiMail32.exe "C:\Documents\Report.xlsx" "Monthly Report"
```

#### 3. Attachment, subject, and recipient:
```cmd
SendMapiMail32.exe "C:\Documents\Quote.pdf" "Quote #2026-04" "sales@company.com"
```

#### 4. Calling from PowerShell:
```powershell
Start-Process -FilePath ".\SendMapiMail32.exe" -ArgumentList '"C:\Temp\Log.txt"', '"System Report"', '"admin@company.local"'
```

### Building from Source

The project is written in [Go](https://go.dev/). The provided [`make.bat`](make.bat) script compiles both the 32-bit and 64-bit binaries:

```cmd
@echo off

set GOARCH=386
go build -o SendMapiMail32.exe .

set GOARCH=amd64
go build -o SendMapiMail64.exe .
```

Run in terminal:
```cmd
make.bat
```

After building, two executables will be available:
- `SendMapiMail32.exe` (for Tobit David & 32-bit mail clients)
- `SendMapiMail64.exe` (for 64-bit mail clients)

### Return Codes

- Returns `0` (`SUCCESS_SUCCESS`) when successfully handed over to the MAPI subsystem.
- In case of failure or cancellation, the MAPI error code is printed (e.g., `1` for `MAPI_USER_ABORT` when the user closes the compose window without sending).

### Credits

Both the source code of this tool and the application icon were generated using **Google Gemini**.

