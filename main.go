package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/jroimartin/gocui"
)

const (
	httpTimeout        = 180
	updateRateNodes    = 5
	updateRateNetworks = 20
	updateRateVolumes  = 1
	updateRateMedia    = 3
	cognitoUserCreds   = `"email":"typeyours@volumez.com", "password":"typeyour!"`
	cognitoUserName    = `"name":"jenkins"`
	sioPublicDns       = "11-11-11-111-11.compute-1.amazonaws.com"
)

var (
	showHelp bool
	gwUrl    string
	client   *http.Client
	err      error
	tkn      string
)

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   true,
			MaxIdleConnsPerHost: -1,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
		Timeout: httpTimeout * time.Second,
	}
}

var (
	done = make(chan struct{})
	wg   sync.WaitGroup
)

func main() {
	flag.StringVarP(&gwUrl, "url", "u", "", "api gateway invoke url")
	flag.BoolVarP(&showHelp, "help", "h", false, "Show help message")
	flag.Parse()
	if showHelp {
		flag.Usage()
		return
	}
	if gwUrl == "" {
		flag.Usage()
		return
	}
	//log.Printf(gwUrl)
	tkn, _ = token()
	//var stdoutBuf bytes.Buffer
	//if err := sioSSH("tail /var/log/storingio/General.log",&stdoutBuf); err != nil {
	//	return
	//}
	//fmt.Printf("%s", stdoutBuf.String())
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Cursor = true
	g.Mouse = true
	g.SelFgColor = gocui.ColorGreen
	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	wg.Add(8)
	go networks(g)
	go nodes(g)
	go medias(g)
	go volumes(g)
	go jobs(g)
	go attachments(g)
	go snapshots(g)
	go policies(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	wg.Wait()
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("INFO", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "INFO"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorWhite
		v.Editable = true
		v.Wrap = true

		//v.Overwrite = true
		v.Frame = true
		v.Autoscroll = true

		fmt.Fprintln(v, fmt.Sprintf("gateway API: %s\nsioPublicDns: %s", gwUrl, sioPublicDns))
	}
	if v, err := g.SetView("NODES", 0, int(0.1*float32(maxY)), int(0.2*float32(maxX)), 1+int(0.3*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "NODES"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		//v.Editable = true
		v.Wrap = true
		v.Frame = true

		fmt.Fprintln(v, "l o a d i n g ...")
	}
	if v, err := g.SetView("NETWORKS", 0, 1+int(0.3*float32(maxY)), int(0.2*float32(maxX)), 3+int(0.4*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "NETWORKS"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		//v.Editable = true
		v.Wrap = true
		v.Frame = true
		fmt.Fprintln(v, "l o a d i n g ...")
	}
	if v, err := g.SetView("JOBS", 0, 3+int(0.4*float32(maxY)), int(0.2*float32(maxX)), maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "JOBS"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		v.Overwrite = true
		//v.Frame = true
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprintln(v, "l o a d i n g ...")
	}
	if v, err := g.SetView("MEDIA", int(0.2*float32(maxX)), int(0.1*float32(maxY)), int(0.7*float32(maxX))-1, int(0.5*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Title = "MEDIA"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		//v.Editable = true
		v.Wrap = true
		//v.Frame = true
		fmt.Fprintln(v, "l o a d i n g ...")
	}
	if v, err := g.SetView("ATTACHMENTS", int(0.7*float32(maxX)), int(0.1*float32(maxY)), maxX-1, int(0.3*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "ATTACHMENTS"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		//v.Editable = true
		v.Wrap = true
		//v.Frame = true
		fmt.Fprintln(v, "l o a d i n g ...")
	}
	if v, err := g.SetView("VOLUMES", int(0.2*float32(maxX)), int(0.5*float32(maxY))+1, int(0.7*float32(maxX))-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "VOLUMES"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		//v.Editable = true
		v.Wrap = true
		v.Frame = true
		fmt.Fprintln(v, "l o a d i n g ...")
	}
	if v, err := g.SetView("POLICIES", int(0.7*float32(maxX)), int(0.3*float32(maxY))+1, maxX-1, int(0.5*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "POLICIES"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		//v.Editable = true
		v.Wrap = true
		v.Frame = true
		fmt.Fprintln(v, "l o a d i n g ...")
	}
	if v, err := g.SetView("SNAPSHOTS", int(0.7*float32(maxX)), int(0.5*float32(maxY))+1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "SNAPSHOTS"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		//v.Editable = true
		v.Wrap = true
		v.Frame = true
		fmt.Fprintln(v, "l o a d i n g ...")
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	for _, n := range []string{"NODES", "NETWORKS", "MEDIA", "VOLUMES", "JOBS", "ATTACHMENTS", "SNAPSHOTS", "POLICIES"} {
		if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
			return err
		}
	}
	for _, n := range []string{"INFO"} {
		if err := g.SetKeybinding(n, gocui.MouseRight, gocui.ModNone, goBack); err != nil {
			return err
		}
	}
	if err := g.SetKeybinding("INFO", gocui.MouseLeft, gocui.ModNone, showFull); err != nil {
		return err
	}
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func showMsg(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	//if _, err := g.SetCurrentView(v.Name()); err != nil {
	//	return err
	//}
	if _, err := setCurrentViewOnTop(g, v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	switch v.Name() {
	case "NODES":
		key := strings.Split(l, " ")[0]
		if v, ok := nodesMap[key]; ok {
			l = v.Info()
		}
	case "MEDIA":
		key := strings.Split(l, " ")[0]
		if v, ok := mediaMap[key]; ok {
			l = v.Info()
		}
	case "VOLUMES":
		key := strings.Split(l, " ")[0]
		if v, ok := volumeMap[key]; ok {
			l = v.Info()
		}
	case "SNAPSHOTS":
		key := strings.Split(l, " ")[0]
		if v, ok := snapMap[key]; ok {
			l = v.Info()
		}
	case "JOBS":
		var stdoutBuf bytes.Buffer
		key := strings.Split(l, " ")[0]
		if err := sioSSH(fmt.Sprintf("tail /var/log/storingio/%s.log", key), &stdoutBuf); err != nil {
			l = err.Error()
		} else {
			l = stdoutBuf.String()
		}
	case "POLICIES":
		key := strings.Split(l, " ")[0]
		if v, ok := policyMap[key]; ok {
			l = v.Info()
		}
	case "ATTACHMENTS":
		key := strings.Split(l, " ")[0]
		if v, ok := attachmentMap[key]; ok {
			l = v.Info()
		}
	}
	v, err = g.View("INFO")
	if err != nil {
		return err
	}
	v.Clear()
	fmt.Fprintln(v, l)
	return nil
}

func showFull(g *gocui.Gui, v *gocui.View) error {
	//if err := g.DeleteView("msg"); err != nil {
	//	return err
	//}
	//v.SetOrigin()
	if _, err := setCurrentViewOnTop(g, v.Name()); err != nil {
		return err
	}

	return nil
}
func goBack(g *gocui.Gui, v *gocui.View) error {
	//if err := g.DeleteView("msg"); err != nil {
	//	return err
	//}
	//v.SetOrigin()
	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	} else {
		_, err := g.SetViewOnBottom(v.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	close(done)
	return gocui.ErrQuit
}
