package main

import (
	"encoding/xml"
	"github.com/giefferre/codemotion_soapclient/request"
	"github.com/giefferre/codemotion_soapclient/response"
	"fmt"
)

func main() {
	// instancing a request
	request, err := request.NewSoapRequest()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Request---> %#v \n\n", request)

	// set the request address, type and address
	request.SetRequest("http://www.webservicex.net/geoipservice.asmx", "GetGeoIp", "8.8.8.8")


	fmt.Printf("Request---> %#v \n\n", request)

	// run the request
	byteBody, err := request.Do()
	if err != nil {
		fmt.Println(err)
	}

	// get the response
	var response response.SoapGetGeoIpResponse

	// unmarshal
	err = xml.Unmarshal(byteBody, &response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("response---> %#v \n\n", response)

	// get the result object
	getGeoIpResult := response.Body.GetGeoIPResponse.GetGeoIPResponseResult

	fmt.Printf("getGeoIpResult---> %#v \n\n", getGeoIpResult)

	// printing the output
	if getGeoIpResult.ReturnCode == "1" {
		fmt.Println(
			fmt.Sprintf("Results for IP %s:\n", getGeoIpResult.IP),
			fmt.Sprintf("\t- Country:\t %s \n", getGeoIpResult.CountryName),
			fmt.Sprintf("\t- Country Code:\t %s \n", getGeoIpResult.CountryCode),
		)
	}

}
