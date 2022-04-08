package automaticshit

import (
	"automaticshit/common/config"
	"automaticshit/common/context"
	"automaticshit/common/safego"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

// IAutoMaticShit
type IAutoMaticShit interface {
	// GetCurShit 获取当前值班名单
	GetCurShit() []string
}

type Shits struct {
	Shit  [][]string `json:"shit"`  // 排班
	Index int        `json:"index"` // 第几个
}

type AutomaticShitMgr struct {
	file  *os.File
	shits *Shits
}

func NewAutomaticShitMgr(ctx context.IContext, dataPath string) (mgr *AutomaticShitMgr, err error) {
	file, err := os.OpenFile(dataPath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	mgr = &AutomaticShitMgr{
		file:  file,
		shits: &Shits{},
	}
	err = mgr.readDataFromFile()
	if err != nil {
		return
	}
	safego.SafeGo(ctx, func() {
		mgr.saveData(ctx)
	})
	return
}

func (a *AutomaticShitMgr) Close() {
	a.file.Close()
}

func (a *AutomaticShitMgr) readDataFromFile() (err error) {
	buf, err := ioutil.ReadAll(a.file)
	if err != nil {
		return
	}
	if len(buf) <= 0 {
		return
	}
	return json.Unmarshal(buf, a.shits)
}

func (a *AutomaticShitMgr) saveData(ctx context.IContext) {
	tc := time.NewTicker(time.Second)
	defer tc.Stop()
	for {
		select {
		case <-tc.C:
			if _, err := a.file.Seek(0, 0); err != nil {
				ctx.Error("file seek error", err.Error())
				break
			}
			buf, err := json.Marshal(a.shits)
			if err != nil {
				ctx.Error("json marshal error", err.Error())
				break
			}
			if _, err = a.file.Write(buf); err != nil {
				ctx.Error("file write error", err.Error())
			}
		}
	}
}

func (a *AutomaticShitMgr) GetCurShit() []string {
	cfg := config.GetConfig()
	if a.shits.Index >= len(a.shits.Shit) || len(a.shits.Shit) != int(cfg.PerUserNum) {
		now := time.Now()
		days := getDaysOfMonth(now.Year(), now.Month())
		a.shits.Shit = AutomaticShit(cfg.People, int(cfg.PerUserNum), int(cfg.Num), days)
		a.shits.Index = 0
	}
	shits := a.shits.Shit[a.shits.Index]
	a.shits.Index++
	return shits
}

// AutomaticShit 自动排班
// people 人员
// perUserNum 每天排班的人数
// num 从第几个开始排
func AutomaticShit(people []string, perUserNum, num int, days int) (shit [][]string) {
	totalNum := days * perUserNum
	tmp := make([]string, totalNum)
	shit = make([][]string, 0, days)
	for i := 0; i < totalNum; i++ {
		if num < len(people) {
			tmp[i] = people[num]
		} else {
			num = 0
			tmp[i] = people[num]
		}
		num++
	}
	for i := 0; i < totalNum; i += perUserNum {
		shitTmp := make([]string, 0, perUserNum)
		for j := 0; j < perUserNum; j++ {
			shitTmp = append(shitTmp, tmp[i+j])
		}
		shit = append(shit, shitTmp)
	}
	return
}
