package main

import (
	"encoding/json"
	"log"
)

func main() {
	var jsonMap map[string]interface{}

	var logText interface{}

	logText = `{
		"host": {
			"ip": "10.234.3.144",
			"hostname": "PTTDIGI-TH-MN-ENCOA-FL5-FW3260-1"
		},
		"log": {
			"syslog": {
				"severity": {
					"code": 6,
					"name": "Informational"
				},
				"facility": {
					"code": 1,
					"name": "user-level"
				},
				"priority": 14
			}
		},
		"message": "LEEF:1.0|Palo Alto Networks|PAN-OS Syslog Integration|10.1.5-h1|allow|cat=TRAFFIC|ReceiveTime=2022/06/30 16:24:23|SerialNumber=016401006591|Type=TRAFFIC|Subtype=end|devTime=Jun 30 2022 09:24:23 GMT|src=10.0.96.46|dst=10.224.242.69|srcPostNAT=0.0.0.0|dstPostNAT=0.0.0.0|RuleName=PTTEP-FireEye-HX|usrName=|SourceUser=|DestinationUser=|Application=web-browsing|VirtualSystem=vsys1|SourceZone=Vl3071-pttep-wgdp|DestinationZone=Vl2471-untrust|IngressInterface=ae1.3071|EgressInterface=ae1.2471|LogForwardingProfile=PTTDIGI-SIEM|SessionID=2084224|RepeatCount=1|srcPort=60400|dstPort=80|srcPostNATPort=0|dstPostNATPort=0|Flags=0x1c|proto=tcp|action=allow|totalBytes=994|dstBytes=461|srcBytes=533|totalPackets=10|StartTime=2022/06/30 16:24:05|ElapsedTime=0|URLCategory=not-resolved|sequence=7102777242082691407|ActionFlags=0x0|SourceLocation=10.0.0.0-10.255.255.255|DestinationLocation=10.0.0.0-10.255.255.255|dstPackets=4|srcPackets=6|SessionEndReason=tcp-fin|DeviceGroupHierarchyL1=0|DeviceGroupHierarchyL2=0|DeviceGroupHierarchyL3=0|DeviceGroupHierarchyL4=0|vSrcName=|DeviceName=PTTDIGI-TH-MN-ENCOA-FL5-FW3260-|ActionSource=from-policy|SrcUUID=|DstUUID=|TunnelID=0|MonitorTag=|ParentSessionID=0|ParentStartTime=|TunnelType=N/A",
		"@version": "1",
		"event": {
			"original": "<14>Jun 30 16:24:23 PTTDIGI-TH-MN-ENCOA-FL5-FW3260-1 LEEF:1.0|Palo Alto Networks|PAN-OS Syslog Integration|10.1.5-h1|allow|cat=TRAFFIC|ReceiveTime=2022/06/30 16:24:23|SerialNumber=016401006591|Type=TRAFFIC|Subtype=end|devTime=Jun 30 2022 09:24:23 GMT|src=10.0.96.46|dst=10.224.242.69|srcPostNAT=0.0.0.0|dstPostNAT=0.0.0.0|RuleName=PTTEP-FireEye-HX|usrName=|SourceUser=|DestinationUser=|Application=web-browsing|VirtualSystem=vsys1|SourceZone=Vl3071-pttep-wgdp|DestinationZone=Vl2471-untrust|IngressInterface=ae1.3071|EgressInterface=ae1.2471|LogForwardingProfile=PTTDIGI-SIEM|SessionID=2084224|RepeatCount=1|srcPort=60400|dstPort=80|srcPostNATPort=0|dstPostNATPort=0|Flags=0x1c|proto=tcp|action=allow|totalBytes=994|dstBytes=461|srcBytes=533|totalPackets=10|StartTime=2022/06/30 16:24:05|ElapsedTime=0|URLCategory=not-resolved|sequence=7102777242082691407|ActionFlags=0x0|SourceLocation=10.0.0.0-10.255.255.255|DestinationLocation=10.0.0.0-10.255.255.255|dstPackets=4|srcPackets=6|SessionEndReason=tcp-fin|DeviceGroupHierarchyL1=0|DeviceGroupHierarchyL2=0|DeviceGroupHierarchyL3=0|DeviceGroupHierarchyL4=0|vSrcName=|DeviceName=PTTDIGI-TH-MN-ENCOA-FL5-FW3260-|ActionSource=from-policy|SrcUUID=|DstUUID=|TunnelID=0|MonitorTag=|ParentSessionID=0|ParentStartTime=|TunnelType=N/A"
		},
		"@timestamp": "2022-06-30T16:24:23Z",
		"KV": "LEEF:1.0|Palo Alto Networks|PAN-OS Syslog Integration|10.1.5-h1|allow|cat=TRAFFIC|ReceiveTime=2022/06/30 16:24:23|SerialNumber=016401006591|Type=TRAFFIC|Subtype=end|devTime=Jun 30 2022 09:24:23 GMT|src=10.0.96.46|dst=10.224.242.69|srcPostNAT=0.0.0.0|dstPostNAT=0.0.0.0|RuleName=PTTEP-FireEye-HX|usrName=|SourceUser=|DestinationUser=|Application=web-browsing|VirtualSystem=vsys1|SourceZone=Vl3071-pttep-wgdp|DestinationZone=Vl2471-untrust|IngressInterface=ae1.3071|EgressInterface=ae1.2471|LogForwardingProfile=PTTDIGI-SIEM|SessionID=2084224|RepeatCount=1|srcPort=60400|dstPort=80|srcPostNATPort=0|dstPostNATPort=0|Flags=0x1c|proto=tcp|action=allow|totalBytes=994|dstBytes=461|srcBytes=533|totalPackets=10|StartTime=2022/06/30 16:24:05|ElapsedTime=0|URLCategory=not-resolved|sequence=7102777242082691407|ActionFlags=0x0|SourceLocation=10.0.0.0-10.255.255.255|DestinationLocation=10.0.0.0-10.255.255.255|dstPackets=4|srcPackets=6|SessionEndReason=tcp-fin|DeviceGroupHierarchyL1=0|DeviceGroupHierarchyL2=0|DeviceGroupHierarchyL3=0|DeviceGroupHierarchyL4=0|vSrcName=|DeviceName=PTTDIGI-TH-MN-ENCOA-FL5-FW3260-|ActionSource=from-policy|SrcUUID=|DstUUID=|TunnelID=0|MonitorTag=|ParentSessionID=0|ParentStartTime=|TunnelType=N/A",
		"service": {
			"type": "system"
		},
		"dst_hostname": "adada",
		"username": "adsad"
	}`
	err := json.Unmarshal([]byte(logText.(string)), &jsonMap)
	if err != nil {
		log.Println(err)
	}
	_, ok1 := jsonMap["dst_hostname"]
	_, ok2 := jsonMap["username"]
	if ok1 || ok2 {
		log.Println("yesy")
	}
	for k, _ := range jsonMap {
		log.Println(k)
	}
}
