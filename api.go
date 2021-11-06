package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/jroimartin/gocui"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

var (
	mediaMap      map[string]Media
	nodesMap      map[string]Node
	volumeMap     map[string]Volume
	snapMap       map[string]Snapshot
	attachmentMap map[string]Attachment
	networkMap    map[string]Network
	jobsMap       map[int64]Job
	policyMap     map[string]Policy
)

func token() (string, error) {
	var jsonStr = []byte("{" + cognitoUserCreds + "," + cognitoUserName + "}")
	req, err := http.NewRequest("POST", gwUrl+"/signin", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t, err := jsonparser.GetString(body, "IdToken")
	//log.Println(t)
	if err != nil {
		return "", err
	}
	return t, nil
}

func nodes(g *gocui.Gui) {
	var msg string
	defer wg.Done()
	nodesMap = make(map[string]Node)
	for {
		select {
		case <-done:
			return
		case <-time.After(updateRateNodes * time.Second):
			var data []byte
			nodes := make([]Node, 0)
			erresp := getData("/nodes", &data)
			_ = json.Unmarshal(data, &nodes)
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("NODES")
				if err != nil {
					return err
				}
				v.Clear()
				if erresp != nil {
					fmt.Fprintln(v, erresp.Error())
				} else {

					for _, item := range nodes {
						switch item.State {
						case "online":
							msg = fmt.Sprintf("%s %s \033[32;1m%s\033[0m", item.Name, item.Status, item.State)
						default:
							msg = fmt.Sprintf("%s %s \033[31;7m%s\033[0m", item.Name, item.Status, item.State)
						}
						fmt.Fprintln(v, msg)
						nodesMap[item.Name] = item
					}
				}
				return nil
			})
		}
	}
}

func medias(g *gocui.Gui) {
	var msg string
	defer wg.Done()
	mediaMap = make(map[string]Media)
	for {
		select {
		case <-done:
			return
		case <-time.After(updateRateMedia * time.Second):
			var data []byte
			medias := make([]Media, 0)
			erresp := getData("/media", &data)
			_ = json.Unmarshal(data, &medias)
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("MEDIA")
				if err != nil {
					return err
				}
				v.Clear()
				if erresp != nil {
					fmt.Fprintln(v, erresp.Error())
				} else {
					sort.Slice(medias, func(i, j int) bool {
						return medias[i].Node < medias[j].Node
					})
					for _, item := range medias {
						switch item.Assignment {
						case "assigned":
							msg = fmt.Sprintf("%s %s %s \033[32;7m%s\033[0m", item.MediaID, item.Node, item.State, item.Assignment)
						case "used":
							msg = fmt.Sprintf("%s %s %s \033[30;4m%s\033[0m", item.MediaID, item.Node, item.State, item.Assignment)
						case "free":
							msg = fmt.Sprintf("%s %s %s \033[33;4m%s\033[0m", item.MediaID, item.Node, item.State, item.Assignment)
						default:
							msg = fmt.Sprintf("%s %s %s \033[31;7m%s\033[0m", item.MediaID, item.Node, item.State, item.Assignment)
						}
						fmt.Fprintln(v, msg)
						mediaMap[item.MediaID] = item
					}
				}
				return nil
			})
		}
	}
}

func volumes(g *gocui.Gui) {
	var msg string
	defer wg.Done()
	volumeMap = make(map[string]Volume)
	for {
		select {
		case <-done:
			return
		case <-time.After(updateRateVolumes * time.Second):
			var data []byte
			volumes := make([]Volume, 0)
			erresp := getData("/volumes", &data)
			_ = json.Unmarshal(data, &volumes)
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("VOLUMES")
				if err != nil {
					return err
				}
				v.Clear()
				if erresp != nil {
					fmt.Fprintln(v, erresp.Error())
				} else {
					for _, item := range volumes {
						switch item.State {
						case "online":
							msg = fmt.Sprintf("%s %s %s \033[32;1m%s\033[0m", item.VolumeID, item.Name, *item.Node, item.State)
						case "offline":
							msg = fmt.Sprintf("%s %s %s \033[31;7m%s\033[0m", item.VolumeID, item.Name, *item.Node, item.State)
						//case "used":
						//	msg = fmt.Sprintf("%s %s %s \033[30;4m%s\033[0m", item.MediaID, item.Node, item.State, item.Assignment)
						//case "free":
						//	msg = fmt.Sprintf("%s %s %s \033[33;4m%s\033[0m", item.MediaID, item.Node, item.State, item.Assignment)
						default:
							msg = fmt.Sprintf("%s %s %s \033[33;1m%s\033[0m", item.VolumeID, item.Name, *item.Node, item.State)
						}
						fmt.Fprintln(v, msg)
						volumeMap[item.VolumeID] = item
					}
				}
				return nil
			})
		}
	}
}

func snapshots(g *gocui.Gui) {
	defer wg.Done()
	snapMap = make(map[string]Snapshot)
	for {
		select {
		case <-done:
			return
		case <-time.After(updateRateVolumes * time.Second):
			var data []byte
			snapshots := make([]Snapshot, 0)
			erresp := getData("/snapshots", &data)
			_ = json.Unmarshal(data, &snapshots)
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("SNAPSHOTS")
				if err != nil {
					return err
				}
				v.Clear()
				if erresp != nil {
					fmt.Fprintln(v, erresp.Error())
				} else {
					for _, item := range snapshots {
						fmt.Fprintln(v, fmt.Sprintf("%s %s %s %s %s", item.SnapshotID, item.SnapName, item.VolumeName, item.Status, item.State))
						snapMap[item.SnapshotID] = item
					}
				}
				return nil
			})
		}
	}
}

func attachments(g *gocui.Gui) {
	defer wg.Done()
	attachmentMap = make(map[string]Attachment)
	for {
		select {
		case <-done:
			return
		case <-time.After(updateRateMedia * time.Second):
			var data []byte
			attachments := make([]Attachment, 0)
			erresp := getData("/attachments", &data)
			_ = json.Unmarshal(data, &attachments)
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("ATTACHMENTS")
				if err != nil {
					return err
				}
				v.Clear()
				if erresp != nil {
					fmt.Fprintln(v, erresp.Error())
				} else {
					for _, item := range attachments {
						fmt.Fprintln(v, fmt.Sprintf("%s %s %s %s %s", item.VolumeID, item.VolumeName, item.Node, item.Status, item.State))
						attachmentMap[item.VolumeID] = item
					}
				}
				return nil
			})
		}
	}
}
func networks(g *gocui.Gui) {
	defer wg.Done()
	networkMap = make(map[string]Network)
	for {
		select {
		case <-done:
			return
		case <-time.After(updateRateNetworks * time.Second):
			var data []byte
			networks := make([]Network, 0)
			erresp := getData("/networks", &data)
			_ = json.Unmarshal(data, &networks)
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("NETWORKS")
				if err != nil {
					return err
				}
				v.Clear()
				if erresp != nil {
					fmt.Fprintln(v, erresp.Error())
				} else {
					for _, item := range networks {
						fmt.Fprintln(v, fmt.Sprintf("%-12s %s %s", item.Name, item.Zone, item.Type))
						networkMap[item.Name] = item
					}
				}
				return nil
			})
		}
	}
}

func jobs(g *gocui.Gui) {
	defer wg.Done()
	jobsMap = make(map[int64]Job)
	for {
		select {
		case <-done:
			return
		case <-time.After(updateRateVolumes * time.Second):
			var data []byte
			jobs := make([]Job, 0)
			erresp := getData("/jobs", &data)
			_ = json.Unmarshal(data, &jobs)
			for i, j := 0, len(jobs)-1; i < j; i, j = i+1, j-1 {
				jobs[i], jobs[j] = jobs[j], jobs[i]
			}
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("JOBS")
				if err != nil {
					return err
				}
				v.Clear()
				if erresp != nil {
					fmt.Fprintln(v, erresp.Error())
				} else {
					for _, item := range jobs {
						fmt.Fprintln(v, fmt.Sprintf("%d %s %s %s", item.ID, item.Type, item.State, item.Status))
						jobsMap[item.ID] = item
					}
				}
				return nil
			})
		}
	}
}
func policies(g *gocui.Gui) {
	defer wg.Done()
	policyMap = make(map[string]Policy)
	for {
		select {
		case <-done:
			return
		case <-time.After(updateRateMedia * time.Second):
			var data []byte
			pols := make([]Policy, 0)
			erresp := getData("/policies", &data)
			_ = json.Unmarshal(data, &pols)

			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("POLICIES")
				if err != nil {
					return err
				}
				v.Clear()
				if erresp != nil {
					fmt.Fprintln(v, erresp.Error())
				} else {
					for _, item := range pols {
						fmt.Fprintln(v, fmt.Sprintf("%s", item.Name))
						policyMap[item.Name] = item
					}
				}
				return nil
			})
		}
	}
}

func getData(path string, ret *[]byte) error {
	Url, err := url.Parse(gwUrl)
	if err != nil {
		return err
	}
	Url.Path += path
	parameters := url.Values{}
	parameters.Add("wait", "true")
	Url.RawQuery = parameters.Encode()
	req, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", tkn)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	*ret, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("error %s , %s", string(*ret), strconv.Itoa(resp.StatusCode))
	}
	return nil
}

func sioSSH(cmd string, stdoutBuf *bytes.Buffer) error {
	//pemBytes, err := ioutil.ReadFile("/home/alexgra/.ssh/automation-kp.pem")
	//fmt.Printf("%s", pemBytes)
	pemBytes := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAqCx7EW4guG9KeD4dscm7hiF+N/IEH+DclfxAcpcBAtR8F6I7
iUZqbfZaKfnjwTk6heeqEwbMZIMxCvbWK3LKGl356Krw00mumQQ8dZ+klD1gbQH6
2vlnMUDV/MV/xVMIQpC3cjQ0jyUcF0YfQ3XK1us6rUFqZAjpFzBoBxtyUHZt1Lj/
LtjjedTB8S8YnybDTGEi1H+CU3q548YT1DhCNX2yPzy46hT/SSCqRG5EbWwWVnQ/
irRzLW3yDBq44K6FtZM2yfskSv9eXeCi2N7WIrOGIRFT6jCrjlXvHJP8FgJctnNk
b/DHxMr6ZS5i1eLfY12hG56mgf5rILgKchQsJQIDAQABAoIBAHO5EAKhfoCLjHoL
fFF/2Ltmxrzmm7H4ALJwv0Ra5oY1AyMcLs26l7gNQmQKIZAvuja0gFLRZcpwgEnk
KuIA/lOAgVx6bHdoB24h/RyPeyfKyFSafS07W6gHzng+yzpUdaWggatjgxtRVPAq
/45jOu4DNgBMuFIX05VyaNMjLtlwVufHwwIS/FG1CjwFvoYOVoKo2YmERw1DK/Jr
MjnSLkFwqyqt9VHAMn99C/tc8KnSTo2uvsU3z+vtHxbJFIr41FxcSGCPYWn5+2E1
0AkORvdJyGjc11GU0XZU3FdeT9z+s7O9hzOFG+FsJr2C0ZSFRK/FdawaeaPxweir
tuVFbc0CgYEA3QwNRptPrGnvjc+3W/GvcDey4B8IZpFxKOx6uxozZLfZgz8uVIag
fCKP8FR78jZhpKUSghKWC1TbpqeAq5TA0POarZXVSdQo85O8guT2iek7e0ofL6P1
Aew3KZkuEtzfTUWvQKggRghuYL857pJ+7+BYHJtFpynvCOFXGIED2xcCgYEAwsQg
LE0t+koiwA9TRGENsK7hFpUHqUG+ZokMF/qIJ4WsmrONH5hEjIHvA0nmNP+l79V3
ce+WF2Xf3DXpC98X+PkTfXzshmb95XAy+q48b91vTVCz9WfPRy6nvs7PB6vHiBy9
MpLYCCClMTpoVMXefNfbuTqEr7iLPOGka5S/iCMCgYEAjDds5HD4pUG9t5MfmK9C
vkhWq1yEE6wGwBLh93WzTBxjWaHmXa/YdWXnMGgnB4n/flVH3EK18xItExYFxNFj
Tih44cu9tEtkfr4kQlPDH9BW7uohxjKW5FVW2IhWdZit/XJKrRT5A/OtMKmcsf0z
kC4bNmo4UMWE33kxqlWMgJkCgYB0YKy4zAVFITdSe9XNbhC4Gkb1L2e8g0Q6EHnh
ehoRQ5a3ecJBtsJ/EsS2ulmMIZYNkQgmVHri0ETLWItARLYWVv6GZTcPuErN5hUQ
JTyHu1DeafKeGMGKTx58rSaX9tTrSADlT0k20grjN3tP7EvdXT41l/ng5eyNHGca
wW8Q8QKBgCFKHXvSnieYFPDWZNiugLGkbak0I8EopHcIHpaVQrBZZhhGK74nfTxM
3jsEOvAKpdO5BFHswLDobmsktBWUpcstAJ/z9V1GqBpmZA/nGuw5P9408nkVRvzm
qX2cQIZT3k5MlEXUHEztM/UZxI17SRQYhZD7STgVw85UO3qPpVzC
-----END RSA PRIVATE KEY-----`
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey([]byte(pemBytes))
	if err != nil {
		log.Fatalf("parse key failed:%v", err)
	}
	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", sioPublicDns+":22", config)
	if err != nil {
		log.Fatalf("dial failed:%v", err)
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("session failed:%v", err)
	}
	defer session.Close()
	//var stdoutBuf bytes.Buffer
	session.Stdout = stdoutBuf
	err = session.Run(cmd)
	if err != nil {
		log.Fatalf("Run failed:%v", err)
	}
	//log.Printf(">%s", stdoutBuf)
	return nil
}
