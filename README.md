# Energomera CLI
Energomera example CLI application written in Go

# Usage

To build it:
```bash
git clone https://github.com/peak-load/energomera-cli.git
cd energomera-cli
go get -v
go build 
```

To configure it edit config.json file located in the same directory

```json
{
     "Port": "/dev/serial0",
     "SleepInterval": 500,
     "Counters" : ["000001","000002"]
}
```

* Port - specify your port connected to rs-485 bus. Depends on your adapter model or OS (I'm using on Linux under RPi3) port can be different e.g. /dev/ttyS0 or /dev/USB0
* SleepInterval - in milliseconds, can be adjusted if you are getting errors due to long read time
* Counters - specify your power meter unique IDs. ID needs to be set beforehand using windows based utiliy as [described in manual here](https://shop.energomera.kharkov.ua/DOC/ASKUE-485/meter_settings_network_RS485.pdf)


To run it (with sample output):
```bash
./energomera-cli 
========== COUNTER 000001 ==========
phase1v: "214.989"
phase2v: "217.376"
phase3v: "221.558"
phase1a: "1.4998"
phase2a: "1.3667"
phase3a: "4.6107"
power: "1.2187"
phase1p: "0.2762"
phase2p: "0.0096"
phase3p: "0.9328"
freq: "49.98"
tarif1: "85095.7769548"
tarif2: "35949.3052429"
tarif3: "27860.3782999"
========== COUNTER 000002 ==========
phase1v: "223.819"
phase2v: "192.289"
phase3v: "248.016"
phase1a: "0.0128"
phase2a: "0.0222"
phase3a: "0.0128"
power: "0.0"
phase1p: "0.0"
phase2p: "0.0"
phase3p: "0.0"
freq: "50.01"
tarif1: "78997.0277689"
tarif2: "52564.3978779"
tarif3: "26404.3700133"
```

# Credits 
## Original Python code: 
* https://support.wirenboard.com/t/schityvanie-pokazanij-i-programmirovanie-elektroschetchika-energomera-se102m-po-rs-485/212                                                                                                                                               
* https://github.com/sj-asm/energomera

## Documentation / resources
* Manufacturer website (RU) http://www.energomera.ru
* Power meter users manual (RU) http://www.energomera.ru/documentations/ce102m_full_re.pdf
* Power meter basic setup guide (RU) https://shop.energomera.kharkov.ua/DOC/ASKUE-485/meter_settings_network_RS485.pdf

# License
MIT License, see [LICENSE](https://github.com/peak-load/energomera_exporter/blob/main/LICENSE)
