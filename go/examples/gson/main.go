package main

import (
	"encoding/base64"
	"fmt"
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
const note = "eyJkZXRhaWxzIjpbXSwic291cmNlVGlja2V0Ijp7InRpbWVzdGFtcFV0YyI6MTY2NTU0Njc3NjQxNCwiYmV0cyI6W3sic3Rha2UiOnsidmFsdWUiOjEwMDAwMDAsInR5cGUiOiJ0b3RhbCJ9LCJpZCI6IjM5MzY4MWM0LWNiNGYtNDllOS04Y2I0LTFiOTlhNWU5NzBlOV8wIiwic2VsZWN0ZWRTeXN0ZW1zIjpbMV0sInNlbGVjdGlvblJlZnMiOlt7InNlbGVjdGlvbkluZGV4IjowfV19XSwidGlja2V0SWQiOiIzOTM2ODFjNC1jYjRmLTQ5ZTktOGNiNC0xYjk5YTVlOTcwZTkiLCJzZWxlY3Rpb25zIjpbeyJldmVudElkIjoiMzE4MTM3NjQiLCJpZCI6IjEvMTZfMzM2MDgyMi8xNzE1Iiwib2RkcyI6MTg3MDB9XSwic2VuZGVyIjp7ImN1cnJlbmN5IjoiSU5SIiwiY2hhbm5lbCI6ImludGVybmV0IiwiYm9va21ha2VySWQiOjM0OTU5LCJlbmRDdXN0b21lciI6eyJpcCI6IjEyNS4yMjguNDAuMjAzIiwibGFuZ3VhZ2VJZCI6ImVuIiwiaWQiOiIyNWM3ZWNmMC1kMGM4LTQ2MDMtOWRiNi1jODA4ZmMyYTg3NWMiLCJEZXZpY2VJZCI6IklrMXZlbWxzYkdFdk5TNHdJQ2hOWVdOcGJuUnZjMmc3SUVsdWRHVnNJRTFoWXlCUFV5QllJREV3WHpFMVh6Y3BJRUZ3Y0d4bFYyVmlTMmwwTHpVek55NHpOaUFvUzBoVVRVd3NJR3hwYTJVZ1IyVmphMjhwSUVOb2NtOXRaUzh4TURZdU1DNHdMakFnVTJGbVlYSnBMelV6Tnk0ek5pST0ifSwibGltaXRJZCI6Mjc0Nn0sInZlcnNpb24iOiIyLjMifSwiY2FuY2VsbGVkQ29kZSI6MCwiY3JlYXRlT3JkZXJJbmZvIjp7ImNvZGUiOjIwMCwic3VjY2Vzc2VzIjpbeyJvcmRlckluZm8iOnsiSUQiOjMzMiwiU2VsZWN0aW9ucyI6W3siSUQiOiIxIiwiRXZlbnRJRCI6IjMxODEzNzY0IiwiTWFya2V0SUQiOiIxNl8zMzYwODIyIiwiUHJvZHVjdElEIjoiMTcxNSIsIklzU2V0dGxlIjpmYWxzZSwiT2RkcyI6IjEuODgifV0sIk9kZHMiOiIxLjg4IiwiUG90ZW50aWFsV2lubmluZ3MiOiIxODgiLCJTdGFrZSI6IjEwMCIsIlR5cGUiOiJTaW5nbGUiLCJTZWxlY3RTeXN0ZW0iOlsxXSwiRmluaXNoIjpmYWxzZSwiVXNlcklEIjoiMjVjN2VjZjAtZDBjOC00NjAzLTlkYjYtYzgwOGZjMmE4NzVjIiwiQmFsYW5jZUNoYW5nZSI6Ii0xMDAiLCJOZXdCYWxhbmNlIjoiNDQwNzQxMzMuNDUiLCJDcmVhdGVBdCI6IjIwMjItMTAtMTJUMDM6NTI6NTguMTMyOTk4WiJ9fV0sImZhaWx1cmVzIjpbXSwic3RhdHVzIjoyMDB9fQ"
const noteStr = "{\"details\":[],\"sourceTicket\":{\"timestampUtc\":1665546776414,\"bets\":[{\"stake\":{\"value\":1000000,\"type\":\"total\"},\"id\":\"393681c4-cb4f-49e9-8cb4-1b99a5e970e9_0\",\"selectedSystems\":[1],\"selectionRefs\":[{\"selectionIndex\":0}]}],\"ticketId\":\"393681c4-cb4f-49e9-8cb4-1b99a5e970e9\",\"selections\":[{\"eventId\":\"31813764\",\"id\":\"1/16_3360822/1715\",\"odds\":18700}],\"sender\":{\"currency\":\"INR\",\"channel\":\"internet\",\"bookmakerId\":34959,\"endCustomer\":{\"ip\":\"125.228.40.203\",\"languageId\":\"en\",\"id\":\"25c7ecf0-d0c8-4603-9db6-c808fc2a875c\",\"DeviceId\":\"Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMDYuMC4wLjAgU2FmYXJpLzUzNy4zNiI=\"},\"limitId\":2746},\"version\":\"2.3\"},\"cancelledCode\":0,\"createOrderInfo\":{\"successes\":[{\"orderInfo\":{\"ID\":332,\"Selections\":[{\"ID\":\"1\",\"EventID\":\"31813764\",\"MarketID\":\"16_3360822\",\"ProductID\":\"1715\",\"IsSettle\":false,\"Odds\":\"1.88\"}],\"Odds\":\"1.88\",\"PotentialWinnings\":\"188\",\"Stake\":\"100\",\"Type\":\"Single\",\"SelectSystem\":[1],\"Finish\":false,\"UserID\":\"25c7ecf0-d0c8-4603-9db6-c808fc2a875c\",\"BalanceChange\":\"-100\",\"NewBalance\":\"44074133.45\",\"CreateAt\":\"2022-10-12T03:52:58.132998Z\"}}],\"failures\":[],\"status\":200}"

func main() {
	note, _ := base64.StdEncoding.DecodeString(note)
	value := gjson.Get(string(note), "createOrderInfo.code")
	value1 := gjson.Get(string(noteStr), "createOrderInfo.code")

	fmt.Println(string(note))
	fmt.Println("value:", value)
	fmt.Println("value1:", value1, len(value1.String()))
}
