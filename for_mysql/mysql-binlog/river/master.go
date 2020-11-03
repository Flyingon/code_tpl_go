package river

import (
	"bytes"
	"os"
	"path"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/siddontang/go-log/log"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go/ioutil2"
)

type masterInfo struct {
	sync.RWMutex

	Name string `toml:"bin_name"`
	Pos  uint32 `toml:"bin_pos"`

	filePath     string
	lastSaveTime time.Time
}

func loadMasterInfo(dataDir string) (*masterInfo, error) {
	var m masterInfo

	if len(dataDir) == 0 {
		return &m, nil
	}

	m.filePath = path.Join(dataDir, "master.info")
	m.lastSaveTime = time.Now()

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	f, err := os.Open(m.filePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	} else if os.IsNotExist(err) {
		return &m, nil
	}
	defer f.Close()

	_, err = toml.DecodeReader(f, &m)
	return &m, err
}

func (m *masterInfo) Save(pos mysql.Position) error {
	log.Infof("save position %s", pos)

	m.Lock()
	defer m.Unlock()

	m.Name = pos.Name
	m.Pos = pos.Pos

	if len(m.filePath) == 0 {
		return nil
	}

	n := time.Now()
	if n.Sub(m.lastSaveTime) < time.Second {
		return nil
	}

	m.lastSaveTime = n
	var buf bytes.Buffer
	e := toml.NewEncoder(&buf)

	e.Encode(m)

	var err error
	if err = ioutil2.WriteFileAtomic(m.filePath, buf.Bytes(), 0644); err != nil {
		log.Errorf("canal save master info to file %s err %v", m.filePath, err)
	}

	return err
}

func (m *masterInfo) Position() mysql.Position {
	m.RLock()
	defer m.RUnlock()

	return mysql.Position{
		Name: m.Name,
		Pos:  m.Pos,
	}
}

func (m *masterInfo) Close() error {
	pos := m.Position()

	return m.Save(pos)
}
