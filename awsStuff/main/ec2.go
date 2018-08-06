package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"fmt"
	"os"
	"regexp"
	"github.com/aws/aws-sdk-go/aws"

)

type Ec2Struct struct {
	service *ec2.EC2;
}


func initEC2Service(region string) Ec2Struct {
	// Load session from shared config
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	// Create new ec2 service client
	ec2Svc := ec2.New(sess)
	return Ec2Struct{ec2Svc};
}

func (ec2Struct Ec2Struct) getRegionNames() []string{
	ec2Svc := ec2Struct.service;
	regionsList:=[]string{};
	resp,err:=ec2Svc.DescribeRegions(nil)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	for idx:=range resp.Regions{
		regionStruct:=resp.Regions[idx]
		regionsList=append(regionsList, *regionStruct.RegionName)
	}
	fmt.Println(regionsList)
	return regionsList
}

func (ec2Struct Ec2Struct) getNameIDMap() map[string]map[string]string{
	//get service
	ec2Svc := ec2Struct.service;
	//create container for storing
	instanceMap := map[string]map[string]string{};
	//Call to get detailed information on each instance
	resp, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	for idx:= range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			instanceId := *inst.InstanceId
			instanceMap[instanceId]=map[string]string{}
			tags := inst.Tags
			for _, tag := range tags {
				layerMatch, _ := regexp.MatchString("opsworks:layer:.*", *tag.Key)
				if layerMatch {
					instanceMap[instanceId]["Layer"] = *tag.Value
				}
				nameMatch, _ := regexp.MatchString("Name", *tag.Key)
				if nameMatch {
					instanceMap[instanceId]["Name"] = *tag.Value
				}
				clusterMatch, _ := regexp.MatchString("opsworks:stack", *tag.Key)
				if clusterMatch {
					instanceMap[instanceId]["Cluster"] = *tag.Value
				}
			}
		}
	}
	return instanceMap;
}

func main() {
	ec2 := initEC2Service("us-east-2")
	ec2.getRegionNames()
	//getEc2NameIDMap()
}
