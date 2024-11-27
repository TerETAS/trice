<div id="top">

# *Trice* over RTT

> _(Read only you are interested in)_
>
> Allows *Trice* over the debug probe without using a pin or UART.
>
> **This document needs a rework, currently it is a mess, sorry!**

<details><summary>Table of Contents</summary><ol><!-- TABLE OF CONTENTS START -->

<!-- 
Table of Contents Generation:
- Install vsCode extension "Markdown TOC" from dumeng 
- Use Shift-Ctrl-P "markdownTOC:generate" to get the automatic numbering.
- replace "<a id=" with "<a id=" 
-->

<!-- vscode-markdown-toc -->
* 1. [Preface](#preface)
* 2. [For the impatient (2 possibilities)](#for-the-impatient-(2-possibilities))
  * 2.1. [Start JLink commander and connect over TCP](#start-jlink-commander-and-connect-over-tcp)
    * 2.1.1. [Setup TCP4 server providing the trace data](#setup-tcp4-server-providing-the-trace-data)
  * 2.2. [Start using JLinkLogger](#start-using-jlinklogger)
* 3. [Segger Real Time Transfer (RTT)](#segger-real-time-transfer-(rtt))
* 4. [J-Link option](#j-link-option)
  * 4.1. [Convert a STM NUCLEO or DISCOVERY onboard ST-Link (valid for ST-Link v2 & v2.1, not for v3)](#convert-a-stm-nucleo-or-discovery-onboard-st-link-(valid-for-st-link-v2-&-v2.1,-not-for-v3))
    * 4.1.1. [First step (to do if some issues occur - otherwise you can skip it)](#first-step-(to-do-if-some-issues-occur---otherwise-you-can-skip-it))
    * 4.1.2. [Second step](#second-step)
  * 4.2. [Some SEGGER tools in short](#some-segger-tools-in-short)
    * 4.2.1. [JLink.exe](#jlink.exe)
    * 4.2.2. [JLinkRTTLogger.exe](#jlinkrttlogger.exe)
  * 4.3. [JLinkRTTClient.exe](#jlinkrttclient.exe)
  * 4.4. [JLinkRTTViewer.exe](#jlinkrttviewer.exe)
* 5. [Segger RTT](#segger-rtt)
* 6. [Segger J-Link SDK (~800 EUR) Option](#segger-j-link-sdk-(~800-eur)-option)
* 7. [Additional Notes (leftovers)](#additional-notes-(leftovers))
* 8. [Further development](#further-development)
* 9. [NUCLEO-F030R8 example](#nucleo-f030r8-example)
  * 9.1. [RTT with original on-board ST-LINK firmware](#rtt-with-original-on-board-st-link-firmware)
  * 9.2. [Change to J-LINK onboard firmware](#change-to-j-link-onboard-firmware)
  * 9.3. [RTT with J-LINK firmware on-board](#rtt-with-j-link-firmware-on-board)
* 10. [Possible issues](#possible-issues)
* 11. [OpenOCD with Darwin](#openocd-with-darwin)
* 12. [Links](#links)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

</div></ol></details><!-- TABLE OF CONTENTS END -->

##  1. <a id='preface'></a>Preface

This technique needs to be considered as experimental:

* RTT works good with a SEGGER J-Link debug probe but needs some closed source software components.
* Also ST-Link is usable for Trice logs, but maybe not parallel with debugging.
* Most investigations where done with a [NUCLEO64-STM32F030R8 evaluation board](https://www.st.com/en/evaluation-tools/nucleo-f030r8.html) which contains an on-board debug probe reflashed with a SEGGER J-Link OB software (see below).
  * When using very high Trice loads over RTT for a long time, sometimes an on-board J-Link (re-flashed ST-Link) could get internally into an inconsistent state (probably internal buffer overrun), what needs a power cycle then.
* You could consider RTT over open-OCD as an alternative.
* The default SEGGER up-buffer size is 1024 bytes, enough for most cases. If not, adapt it in your [SEGGER_RTT_Conf.h](../examples/F030R8_inst/Core/Inc/SEGGER_RTT_Conf.h) file.
* Possible: Parallel usage of RTT direct mode with UART deferred mode. You can define `TRICE_UARTA_MIN_ID` and `TRICE_UARTA_MAX_ID` inside triceConfig.h to log only a specific ID range over UARTA in deferred mode for example. ([#446](https://github.com/rokath/trice/issues/446))

<p align="right">(<a href="#top">back to top</a>)</p>

##  2. <a id='for-the-impatient-(2-possibilities)'></a>For the impatient (2 possibilities)

The default SEGGER tools only suport RTT channel 0.

###  2.1. <a id='start-jlink-commander-and-connect-over-tcp'></a>Start JLink commander and connect over TCP

* JLink.exe → `connect ⏎ ⏎ S ⏎` and keep it active.
  * You can control the target with `r[eset], g[o], h[alt]` and use other commands too.
  * ![./ref/JLink.exe.PNG](./ref/JLink.exe.PNG)
* Start in Git-Bash or s.th. similar: `trice l -p TCP4 -args localhost:19021`
* You may need a **trice** tool restart after firmware reload.


####  2.1.1. <a id='setup-tcp4-server-providing-the-trace-data'></a>Setup TCP4 server providing the trace data

This is just the SEGGER J-Link server here for demonstration, but if your target device has an TCP4 interface, you can replace this with your target server.

```bash
ms@DESKTOP-7POEGPB MINGW64 ~/repos/trice (master)
$ jlink
SEGGER J-Link Commander V7.92g (Compiled Sep 27 2023 15:36:46)
DLL version V7.92g, compiled Sep 27 2023 15:35:10

Connecting to J-Link via USB...O.K.
Firmware: J-Link STLink V21 compiled Aug 12 2019 10:29:20
Hardware version: V1.00
J-Link uptime (since boot): N/A (Not supported by this model)
S/N: 770806762
VTref=3.300V


Type "connect" to establish a target connection, '?' for help
J-Link>connect
Please specify device / core. <Default>: STM32G0B1RE
Type '?' for selection dialog
Device>
Please specify target interface:
  J) JTAG (Default)
  S) SWD
  T) cJTAG
TIF>s
Specify target interface speed [kHz]. <Default>: 4000 kHz
Speed>
Device "STM32G0B1RE" selected.


Connecting to target via SWD
InitTarget() start
SWD selected. Executing JTAG -> SWD switching sequence.
DAP initialized successfully.
InitTarget() end - Took 36.3ms
Found SW-DP with ID 0x0BC11477
DPv0 detected
CoreSight SoC-400 or earlier
Scanning AP map to find all available APs
AP[1]: Stopped AP scan as end of AP map has been reached
AP[0]: AHB-AP (IDR: 0x04770031)
Iterating through AP map to find AHB-AP to use
AP[0]: Core found
AP[0]: AHB-AP ROM base: 0xF0000000
CPUID register: 0x410CC601. Implementer code: 0x41 (ARM)
Found Cortex-M0 r0p1, Little endian.
FPUnit: 4 code (BP) slots and 0 literal slots
CoreSight components:
ROMTbl[0] @ F0000000
[0][0]: E00FF000 CID B105100D PID 000BB4C0 ROM Table
ROMTbl[1] @ E00FF000
[1][0]: E000E000 CID B105E00D PID 000BB008 SCS
[1][1]: E0001000 CID B105E00D PID 000BB00A DWT
[1][2]: E0002000 CID B105E00D PID 000BB00B FPB
Memory zones:
  Zone: "Default" Description: Default access mode
Cortex-M0 identified.
J-Link>
```

Now the TCP4 server is running and you can start the Trice tool as TCP4 client, which connects to the TCP4 server to receive the binary log data:

```bash
$ trice l -p TCP4 -args="127.0.0.1:19021" -til ../examples/G0B1_inst/til.json -li ../examples/G0B1_inst/li.json -d16 -pf none
```

In this **G0B1_inst** example we use the additional `-d16` and `-pf none` switches to decode the RTT data correctly.

**This is just a demonstration and test for the `-port TCP4` usage possibility**. Using RTT with J-Link is more easy possible as shown in the next point.

###  2.2. <a id='start-using-jlinklogger'></a>Start using JLinkLogger

* Start inside Git-Bash or s.th. similar: `trice l -p JLINK -args "-Device STM32F030R8 -if SWD -Speed 4000 -RTTChannel 0"`
  * Replace CLI details with your settings.
  * For **G01B_inst**: `trice l -p JLINK -args "-Device STM32G0B1RE -if SWD -Speed 4000 -RTTChannel 0" -d16 -pf none`
  * You can add the `-verbose` CLI switch for more details.
* You may **not** need a **trice** tool restart after firmware reload.
* For some reason the RTT technique does not work well with Darwin right now. The problem seems to be that the JLinkRTTLogger app cannot work correctly in the background. But there is a workaround:
  * In one terminal run `JLinkRTTLogger -Device STM32G0B1RE -if SWD -Speed 4000 -RTTChannel 0 myLogFile.bin` 
  * and in an other terminal execute `trice l -p FILE -args myLogFile.bin -pf none -d16`.

<p align="right">(<a href="#top">back to top</a>)</p>

##  3. <a id='segger-real-time-transfer-(rtt)'></a>Segger Real Time Transfer (RTT) 

* Prerequisite is a processor with memory background access support like ARM Cortex-M cores.
* If you can use a Segger J-Link or an STM ST-Link debug probe (ST Microelectronics eval boards have it) this is an easy and fast way to use *Trice* without any UART or other port.
* Detailed description can be found in document [UM08001_JLink.pdf](../third_party/segger.com/UM08001_JLink.pdf) in chapter 16 which is part of [https://www.segger.com/downloads/jlink/#J-LinkSoftwareAndDocumentationPack](https://www.segger.com/downloads/jlink/#J-LinkSoftwareAndDocumentationPack).
* Following examples are for Windows, but should work similar also on Linux and Darwin (MacOS).
* *Trice* can use the Segger RTT protocol in different ways.
  * Hardware paths:
    * Use [J-Link](https://www.segger.com/products/debug-probes/j-link/) or [J-Link OB (on-board)](https://www.segger.com/products/debug-probes/j-link/models/j-link-ob/).
      J-Link OB can be flashed to many ST Microelectronics evaluation boards (v2.0 link hardware) and for example is also usable with NXP and Atmel. For that you can also use a spare STM32 evaluation board (10 EUR) with jumper changes and breakout wires.
    * Use ST-Link with [gostlink](../third_party/goST/ReadMe.md).
      It uses only one USB endpoint so debugging and Trice output in parallel is not possible.
    * Use some other Debug-Probe with target memory access (support welcome)
  * RTT channel selection (on target and on host)
    * RECOMMENDED:
      * `trice l -p JLINK` or shorter `trice l` for STM32F030R8 (default port is JLINK) starts in background a `JLinkRTTLogger.exe` which connects to J-Link and writes to a logfile which in turn is read by the **trice** tool. On exit the `JLinkRTTLogger.exe` is killed automatically.
        * It expects a target sending messages over RTT channel **0** (zero). Chapter 16.3.3 in [UM08001_JLink.pdf](../third_party/segger.com/UM08001_JLink.pdf) refers to "Up-Channel 1" but this maybe is a typo and probably a 0 is mend. The `JLinkRTTLogger.exe` main advantage against other free available SEGGER tools is, that all bytes are transferred. Other SEGGER tools assume ASCII characters and use `FF 00` to `FF 0F` as a terminal switch command and filter that out causing *Trice* data disturbances.
        * It should be possible to start several instances on on different targets using `-SelectEmuBySN <SN>` inside the `-args` *Trice* CLI switch.
        * `JLinkRTTLogger` binaries for Linux & Darwin can be found at [https://www.segger.com/downloads/jlink/](https://www.segger.com/downloads/jlink/).
      * `trice l -p STLINK` starts in background a `trice/third_party/goST/stRttLogger.exe` which connects to ST-Link and writes to a logfile which in turn is read by the Trice tool. On exit the `stRttLogger.exe` is killed automatically. It expects a target sending messages over RTT channel 0 (other channels supported too but may not work).\
      It is possible to start several instances on different channels as well as on different targets. The source code is in [https://github.com/bbnote/gostlink](https://github.com/bbnote/gostlink) and should work also at least under Linux.
    * If you have the choice, prefer J-Link. It allows parallel debugging and Trice output.
    * The full `-args` string is normally required and depends on the used device. Example: `trice l -args="-Device STM32F070RB -if SWD -Speed 4000 -RTTChannel 0 -RTTSearchRanges 0x20000000_0x1000"`. The `-RTTSearchRanges` part is mostly optional.
    * Enter `trice h -log` and read info for `-args` switch:

```bash
        -args string
        Use to pass port specific parameters. The "default" value depends on the used port:
        port "COMn": default="", use "TARM" for a different driver. (For baud rate settings see -baud.)
        port "J-LINK": default="-Device STM32F030R8 -if SWD -Speed 4000 -RTTChannel 0 -RTTSearchRanges 0x20000000_0x1000",
                The -RTTSearchRanges "..." need to be written without extra "" and with _ instead of space.
                For args options see JLinkRTTLogger in SEGGER UM08001_JLink.pdf.
        port "ST-LINK": default="-Device STM32F030R8 -if SWD -Speed 4000 -RTTChannel 0 -RTTSearchRanges 0x20000000_0x1000",
                The -RTTSearchRanges "..." need to be written without extra "" and with _ instead of space.
                For args options see JLinkRTTLogger in SEGGER UM08001_JLink.pdf.
        port "BUFFER": default="0 0 0 0", Option for args is any byte sequence.
         (default "default")
 ```

<p align="right">(<a href="#top">back to top</a>)</p>

##  4. <a id='j-link-option'></a>J-Link option

* Prerequisite is a SEGGER J-Link debug probe or a development board with an on-board J-Link option.

###  4.1. <a id='convert-a-stm-nucleo-or-discovery-onboard-st-link-(valid-for-st-link-v2-&-v2.1,-not-for-v3)'></a>Convert a STM NUCLEO or DISCOVERY onboard ST-Link (valid for ST-Link v2 & v2.1, not for v3)

* Following steps describe the needed action for a ST Microelectronics evaluation board and windows - adapt them to your environment.
* It is always possible to turn back to the ST-Link OB firmware with the SEGGER `STLinkReflash.exe` tool but afterwards the ST-Link Upgrade tool should be used again to get the latest version. 

####  4.1.1. <a id='first-step-(to-do-if-some-issues-occur---otherwise-you-can-skip-it)'></a>First step (to do if some issues occur - otherwise you can skip it)

[Video](https://www.youtube.com/watch?app=desktop&v=g2Kf6RbdrIs)

See also [https://github.com/stlink-org/stlink](https://github.com/stlink-org/stlink)

* Get & install [STM32 ST-LINK utility](https://www.st.com/en/development-tools/stsw-link004.html)
* Run from default install location `"C:\Program Files (x86)\STMicroelectronics\STM32 ST-LINKUtility\ST-LINK Utility\ST-LinkUpgrade.exe"`)
* Enable checkbox `Change Type` and select radio button `STM32 Debug+Mass storage + VCP`. *The `STM32Debug+ VCP` won´t be detected by Segger reflash utility.*
  ![ST-LINK-Upgrade.PNG](./ref/ST-LINK-Upgrade.PNG)

####  4.1.2. <a id='second-step'></a>Second step

* Check [Converting ST-LINK On-Board Into a J-Link](https://www.segger.com/products/debug-probes/j-link/models/other-j-links/st-link-on-board/)
* Use `STLinkReflash.exe` to convert NUCLEO from ST-Link on-board to J-Link on-board. *`STM32 Debug+ VCP` won´t be detected by Segger reflash utility.*

###  4.2. <a id='some-segger-tools-in-short'></a>Some SEGGER tools in short

* Download [J-Link Software and Documentation Pack](https://www.segger.com/downloads/jlink/#J-LinkSoftwareAndDocumentationPack) and install.
  * You may need to add `C:\Program Files\SEGGER\JLink` to the %PATH% variable.
* Tested with [NUCLEO64-STM32F030R8 evaluation board](https://www.st.com/en/evaluation-tools/nucleo-f030r8.html).
* For example: Compile and flash `../_test/MDK-ARM_STM32F030R8` project.
  * Check in [../_test/MDK-ARM_STM32F030R8/Core/Inc/triceConfig.h](../_test/MDK-ARM_STM32F030R8/Core/Inc/triceConfig.h) if `#define TRICE_RTT_CHANNEL 0` is set as output option.

####  4.2.1. <a id='jlink.exe'></a>JLink.exe

* `JLink.exe` is the SEGGER J-Link commander. It starts the **J-Link driver/server** and one can connect to it
* Info found [here](https://gist.github.com/GaryLee/ecd8018d1ca046c1a40fcd265fa109c0):
  * J-Link Commander can be started with different command line options for test and automation
  * purposes. In the following, the command line options which are available for J-Link
  * Commander are explained. All command line options are case insensitive.
  * Command Explanation
  * -AutoConnect Automatically start the target connect sequence
  * -CommanderScript Passes a CommandFile to J-Link
  * -CommandFile Passes a CommandFile to J-Link
  * -Device Pre-selects the device J-Link Commander shall connect to
  * -ExitOnError Commander exits after error.
  * -If Pre-selects the target interface
  * -IP Selects IP as host interface
  * -JLinkScriptFile Passes a JLinkScriptFile to J-Link
  * -JTAGConf Sets IRPre and DRPre
  * -Log Sets logfile path
  * -RTTTelnetPort Sets the RTT Telnetport
  * -SelectEmuBySN Connects to a J-Link with a specific S/N over USB
  * -SettingsFile Passes a SettingsFile to J-Link
  * -Speed Starts J-Link Commander with a given initial speed
* Documentation: [https://wiki.segger.com/J-Link_Commander](https://wiki.segger.com/J-Link_Commander)
* If you run successful `jlink -device STM32F030R8 -if SWD -speed 4000 -autoconnect 1` the target is stopped.
  * To let in run you need manually execute `go` as command in the open jlink window.
  * To automate that create a text file named for example `jlink.go` containing the `go` command: `echo go > jlink.go` and do a `jlink -device STM32F030R8 -if SWD -speed 4000 -autoconnect 1 -CommandFile jlink.go`
* It is possible to see some output with Firefox (but not with Chrome?) for example: ![./ref/JLink19021.PNG](./ref/JLink19021.PNG).
* After closing the Firefox the **trice** tool can connect to it too:
  * Open a commandline and run:
    ```b
    trice log -port TCP4 -args localhost:19021
    ```
    * Trice output is visible
    * With `h`alt and `g`o inside the Jlink window the MCU can get haltes and released
    * It is possible in parallel to debug-step with a debugger (tested with ARM-MDK)
* ![./ref/JLinkServer.PNG](./ref/JLinkServer.PNG)
* **PLUS:**
  * Works reliable.
  * No file interface needed.
  * *Trice* can connect over TCP localhost:19021 and display logs over RTT channel 0.
  * The open `jlink` CLI can be handy to control the target: `[r]eset, [g]o. [h]alt`
  * No need to restart the **trice** tool after changed firmware download.
* **MINUS:**
  * Uses RTT up-channel 0 and therefore RTT up-channel 0 is not usable differently.
  * No down-channel usable.
  * Needs a separate manual start of the `jlink` binary with CLI parameters.
    * I would not recommend to automate that too, because this step is needed only once after PC power on.

####  4.2.2. <a id='jlinkrttlogger.exe'></a>JLinkRTTLogger.exe

* `JLinkRTTLogger.exe` is a CLI tool and connects via the SEGGER API to the target. It is usable for writing RTT channel 0 data from target into a file.
* **PLUS:**
  * Works reliable.
  * Is automatable.
  * Create file with raw log data: `JLinkRTTLogger.exe -Device STM32F030R8 -if SWD -Speed 4000 -RTTChannel 0 triceRaw.log`
    * It is possible to evaluate this file offline: `trice l -p FILE -args triceRaw.log`
    * ![./ref/TriceFILE.PNG](./ref/TriceFILE.PNG)
  * No need to restart the **trice** tool after changed firmware download.
* **MINUS:**
  * Logs in a file, so the **trice** tool needs to read from that file.
  * Maybe cannot write in a file as background process on Darwin.
* The **trice** tool can watch the output file and display the *Trices*: `trice log -port JLINK -args "-Device STM32F030R8 -if SWD -Speed 4000 -RTTChannel 0"
![./ref/JlinkLoggerTrice.PNG](./ref/JlinkLoggerTrice.PNG)

###  4.3. <a id='jlinkrttclient.exe'></a>JLinkRTTClient.exe

* `JLinkRTTClient.exe` can be used for simple text transmitting to the target, it also displays strings from target coming over channel 0. It is not used by the **trice** tool.
  * **PLUS:**
    * Target stimulation with proprietary protocol over RTT down-channel 0 possible.
  * **MINUS:**
    * Unfortunately it cannot run separately parallel to stimulate the target with any proprietary protocol because it connects to localhost:19021 and therefore blockades the only one possible connection.

###  4.4. <a id='jlinkrttviewer.exe'></a>JLinkRTTViewer.exe

* `JLinkRTTViewer.exe` is a GUI tool and connects via the SEGGER API to the target. It expects ASCII codes and is not used by the **trice** tool. The switching between the 16 possible terminals is done via `FF 00` ... `FF 0F`. These byte pairs can occur inside the Trice data.

<!---

* Start `"C:\Program Files (x86)\SEGGER\JLink\JLinkRTTViewer.exe"` and connect to the J-Link. You only need this as a running server to connect to.
  * Unfortunately the JLinkRTTViewer "steals" from time to time some Trice data packages and displays them as data garbage.
  * Better use JLink.exe or the *Segger J-Link SDK* instead.

-->

<p align="right">(<a href="#top">back to top</a>)</p>

##  5. <a id='segger-rtt'></a>Segger RTT

* The main advantages are:
  * Speed
  * No `TriceTransfer()` nor any interrupt is needed in the background
  * No UART or other output is needed
* This is, because automatically done by SeggerRTT. This way one can debug code as comfortable as with `printf()` but with all the TRICE advantages. Have a look here: ![SeggerRTTD.gif](./ref/JLINK-DebugSession.gif)
* Avoid Trice buffering inside target and write with TRICE macro directly into the RTT buffer (direct *Trice* mode = `#define TRICE_MODE 0` inside [triceConfig.h](../_test/MDK-ARM_STM32F030R8/Core/Inc/triceConfig.h)).
* Write the bytes per *Trice* directly (little time & some space overhead on target, but no changes on host side)
  
  ![triceBlockDiagramWithSeggerRTT.svg](./ref/triceBlockDiagramWithSeggerRTTD.svg)

<p align="right">(<a href="#top">back to top</a>)</p>

##  6. <a id='segger-j-link-sdk-(~800-eur)-option'></a>Segger J-Link SDK (~800 EUR) Option

* Segger offers a SeggerRTT SDK which allows to use more than just channel 0 and you can develop your own tooling with it.
* The `trice -port JLINK` is ok for usage **as is** right now. However if you wish more comfort check here:
* Question: [How-to-access-multiple-RTT-channels](https://forum.segger.com/index.php/Thread/6688-SOLVED-How-to-access-multiple-RTT-channels-from-Telnet/)
  * "Developer pack used to write your own program for the J-Link. Please be sure you agree to the terms of the associated license found on the Licensing Information tab before purchasing this SDK. You will benefit from six months of free email support from the time that this product is ordered."
* The main [Segger J-Link SDK](https://www.segger.com/products/debug-probes/j-link/technology/j-link-sdk/) disadvantage beside closed source and payment is: **One is not allowed to distribute binaries written with the SDK.** That makes it only interesting for company internal automatization.

<p align="right">(<a href="#top">back to top</a>)</p>

##  7. <a id='additional-notes-(leftovers)'></a>Additional Notes (leftovers)

* `Downloading RTT target package` from [https://www.segger.com/products/debug-probes/j-link/technology/about-real-time-transfer/](https://www.segger.com/products/debug-probes/j-link/technology/about-real-time-transfer/).
* Read the manual [UM08001_JLink.pdf](../third_party/segger.com/UM08001_JLink.pdf).
* Extract `../third_party/segger.com/SEGGER_RTT_V760g.zip` to `../third_party/segger.com/SEGGER_RTT`. Check for an update @ SEGGER.
* Add `SEGGER_RTTI.c` to target project

<p align="right">(<a href="#top">back to top</a>)</p>

##  8. <a id='further-development'></a>Further development



* Check OpenOCD!
  * Use OpenOCD and its built-in RTT feature. OpenOCD then starts a server on localhost:17001 where it dumps all RTT messages.
* The GoST project offers a way bypassing JLINK. Used -port STLINK instead.
* Maybe `libusb` together with `libjaylink` offer some options too.
* Checkout [https://github.com/deadsy/jaylink](https://github.com/deadsy/jaylink).
* `"C:\Program Files (x86)\SEGGER\JLink\JMem.exe"` shows a memory dump.

* Go to [https://libusb.info/](https://libusb.info/)
  * -> Downloads -> Latest Windows Binaries
  * extract `libusb-1.0.23` (or later version)

```b
libusb-1.0.23\examples\bin64> .\listdevs.exe
2109:2811 (bus 2, device 8) path: 6
1022:145f (bus 1, device 0)
1022:43d5 (bus 2, device 0)
0a12:0001 (bus 2, device 1) path: 13
1366:0105 (bus 2, device 10) path: 5
```

* Repeat without connected Segger JLink

```b
libusb-1.0.23\examples\bin64> .\listdevs.exe
2109:2811 (bus 2, device 8) path: 6
1022:145f (bus 1, device 0)
1022:43d5 (bus 2, device 0)
0a12:0001 (bus 2, device 1) path: 13
```

* In this case `1366:0105 (bus 2, device 10) path: 5` is missing, so `vid=1366`, `did=0105` as example
* On Windows install WSL2. The real Linux kernel is needed for full USB access.

<p align="right">(<a href="#top">back to top</a>)</p>

##  9. <a id='nucleo-f030r8-example'></a>NUCLEO-F030R8 example

Info: [https://www.st.com/en/evaluation-tools/nucleo-f030r8.html](https://www.st.com/en/evaluation-tools/nucleo-f030r8.html)

###  9.1. <a id='rtt-with-original-on-board-st-link-firmware'></a>RTT with original on-board ST-LINK firmware

* `#define TRICE_RTT_CHANNEL 0`:
* If you use a NUCLEO-F030R8 with the original ST-Link on board after firmware download enter: `trice l -p ST-LINK -args "-Device STM32F030R8 -if SWD -Speed 4000 -RTTChannel 0 -RTTSearchRanges 0x20000000_0x2000"`. After pressing the reset button output becomes visible: ![./ref/STRTT.PNG](./ref/STRTT.PNG)
* It works with both ST-Link variants (with or without mass storage device.)

###  9.2. <a id='change-to-j-link-onboard-firmware'></a>Change to J-LINK onboard firmware

 ![./ref/STLinkReflash.PNG](./ref/STLinkReflash.PNG)

###  9.3. <a id='rtt-with-j-link-firmware-on-board'></a>RTT with J-LINK firmware on-board

![./ref/J-LinkRTT.PNG](./ref/J-LinkRTT.PNG)

- Observations:
  - When pressing the black reset button, you need to restart the **trice** tool.
  - When restarting the Trice tool, a target reset occurs.
  - Other channel numbers than `0` seam not to work for some reason.

<p align="right">(<a href="#top">back to top</a>)</p>

##  10. <a id='possible-issues'></a>Possible issues

* These boards seem not to work reliable with RTT over J-Link on-board firmware.
  * NUCLEO-G071RB
  * NUCLEO_G031K8
* After flashing back the ST-LINK OB firmware with the SEGGER tool, it is recommended to use the ST tool to update the ST-LINK OB firmware. Otherwise issues could occur.

<p align="right">(<a href="#top">back to top</a>)</p>

##  11. <a id='openocd-with-darwin'></a>OpenOCD with Darwin

- OpenOCD on MacOS works out of the box after installing it.
- When using VS code with Cortex-Debug you cannot use OpenOCD at the same time.
- The `openocd.cfg` file is taylored to the flashed on-board J-Link adapter.

**Terminal 1:**

```bash
brew install open-ocd
...
cd ./trice/examples/G0B1_inst
openocd -f openocd.cfg
Open On-Chip Debugger 0.12.0
Licensed under GNU GPL v2
For bug reports, read
	http://openocd.org/doc/doxygen/bugs.html
srst_only separate srst_nogate srst_open_drain connect_deassert_srst

Info : Listening on port 6666 for tcl connections
Info : Listening on port 4444 for telnet connections
Info : J-Link STLink V21 compiled Aug 12 2019 10:29:20
Info : Hardware version: 1.00
Info : VTarget = 3.300 V
Info : clock speed 2000 kHz
Info : SWD DPIDR 0x0bc11477
Info : [stm32g0x.cpu] Cortex-M0+ r0p1 processor detected
Info : [stm32g0x.cpu] target has 4 breakpoints, 2 watchpoints
Info : starting gdb server for stm32g0x.cpu on 3333
Info : Listening on port 3333 for gdb connections
Info : rtt: Searching for control block 'SEGGER RTT'
Info : rtt: Control block found at 0x20001238
Info : Listening on port 9090 for rtt connections
Channels: up=1, down=0
Up-channels:
0: Terminal 1024 0
Down-channels:

Info : Listening on port 6666 for tcl connections
Info : Listening on port 4444 for telnet connections
```

**Terminal 2:**

```bash
ms@MacBook-Pro G0B1_inst % trice l -p TCP4 -args localhost:9090  -pf none -d16
Nov 14 17:32:33.319451  TCP4:       triceExamples.c    10        0_000  Hello! 👋🙂
Nov 14 17:32:33.319463  TCP4: 
Nov 14 17:32:33.319463  TCP4:         ✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨        
Nov 14 17:32:33.319463  TCP4:         🎈🎈🎈🎈  NUCLEO-G0B1RE   🎈🎈🎈🎈
Nov 14 17:32:33.319463  TCP4:         🍃🍃🍃🍃🍃🍃🍃🍃🍃🍃🍃🍃🍃🍃🍃🍃🍃        
Nov 14 17:32:33.319463  TCP4: 
Nov 14 17:32:33.319463  TCP4: 
Nov 14 17:32:33.406455  TCP4:       triceExamples.c    16        0_037 2.71828182845904523536 <- float number as string
Nov 14 17:32:33.505116  TCP4:       triceExamples.c    17        0_087 2.71828182845904509080 (double with more ciphers than precision)
Nov 14 17:32:33.607518  TCP4:       triceExamples.c    18        0_117 2.71828174591064453125 (float  with more ciphers than precision)
Nov 14 17:32:33.707851  TCP4:       triceExamples.c    19        0_146 2.718282 (default rounded float)
Nov 14 17:32:33.807685  TCP4:       triceExamples.c    20        0_175 A Buffer:
Nov 14 17:32:33.908202  TCP4:       triceExamples.c    21        0_204 32 2e 37 31 38 32 38 31 38 32 38 34 35 39 30 34 35 32 33 35 33 36 
Nov 14 17:32:34.007148  TCP4:       triceExamples.c    22        0_254 31372e32  31383238  34383238  34303935  35333235  
Nov 14 17:32:35.007949  TCP4:       triceExamples.c    23        0_301 ARemoteFunctionName(2e32)(3137)(3238)(3138)(3238)(3438)(3935)(3430)(3235)(3533)(3633)
Nov 14 17:32:35.112304  TCP4:       triceExamples.c    24              100 times a 16 byte long Trice messages, which not all will be written because of the TRICE_PROTECT:
Nov 14 17:32:35.307567  TCP4:       triceExamples.c    26        0_379 i=44444400 aaaaaa00
Nov 14 17:32:35.408257  TCP4:       triceExamples.c    27    0,000_002 i=44444400 aaaaaa00
Nov 14 17:32:35.509022  TCP4:       triceExamples.c    26        0_441 i=44444401 aaaaaa01
Nov 14 17:32:35.609439  TCP4:       triceExamples.c    27    0,000_002 i=44444401 aaaaaa01
Nov 14 17:32:35.710201  TCP4:       triceExamples.c    26        0_504 i=44444402 aaaaaa02
...
```

##  12. <a id='links'></a>Links

* [https://www.codeinsideout.com/blog/stm32/j-link-rtt/](https://www.codeinsideout.com/blog/stm32/j-link-rtt/) (A good explanation of SEGGER J-Link Realtime Transfer - Fast Debug protocol: - only suitable for ASCII transfer)
* [USB over WSL2?](https://twitter.com/beriberikix/status/1487127732190212102?s=20&t=NQVa27qvOqPi2uGz6pJNRA) (Maybe intersting for OpenOCD)
* https://kickstartembedded.com/2024/03/26/openocd-one-software-to-rule-debug-them-all/?amp=1
* https://mcuoneclipse.com/2021/10/03/visual-studio-code-for-c-c-with-arm-cortex-m-part-9-rtt/

<p align="right">(<a href="#top">back to top</a>)</p>
