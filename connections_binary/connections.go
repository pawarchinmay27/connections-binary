package main 
import "fmt"
import "github.com/drael/GOnetstat"
import "encoding/json"
func main() {
	tcp_data := GOnetstat.Tcp()
	output_tcp,err:=json.MarshalIndent(tcp_data,"","  ")
	fmt.Println(output_tcp)
	fmt.Println(err)
}