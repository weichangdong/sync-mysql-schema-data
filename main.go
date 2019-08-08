package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/hidu/mysql-schema-sync/internal"
)

//user:password@tcp(127.0.0.1:3306)/test
var configPath = flag.String("conf", "./config.json", "json config file path")
var sync = flag.Bool("sync", false, "sync shcema change to dest db")
var drop = flag.Bool("drop", false, "drop fields,index,foreign key")
var syncData = flag.Bool("sync_data", true, "sync source db table data  to dest db table")
var syncDataTruncate = flag.Bool("sync_data_truncate", false, "is need truncate  source db table data  to dest db table")

var source = flag.String("source", "", "mysql dsn source,eg: test@(10.10.0.1:3306)/test\n\twhen it is not empty ignore [-conf] param")
var dest = flag.String("dest", "", "mysql dsn dest,eg test@(127.0.0.1:3306)/imis")
var tables = flag.String("tables", "", "table names to check\n\teg : product_base,order_*")
var tablesIGNORE = flag.String("tables_ignore", "", "table names to ignore check and ignore sync data\n\teg : product_base,order_*")
var mailTo = flag.String("mail_to", "", "overwrite config's email.to")

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	df := flag.Usage
	flag.Usage = func() {
		df()
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "mysql schema && data sync tools "+internal.Version)
		fmt.Fprintln(os.Stderr, internal.AppURL+"\n")
	}
}

var cfg *internal.Config

func main() {
	flag.Parse()
	if *source == "" {
		cfg = internal.LoadConfig(*configPath)
	} else {
		cfg = new(internal.Config)
		cfg.SourceDSN = *source
		cfg.DestDSN = *dest
	}
	cfg.Sync = *sync
	cfg.Drop = *drop
	cfg.SyncData = *syncData
	cfg.SyncDataTruncate = *syncDataTruncate

	if *mailTo != "" && cfg.Email != nil {
		cfg.Email.To = *mailTo
	}

	if cfg.Tables == nil {
		cfg.Tables = []string{}
	}
	if cfg.TablesIGNORE == nil {
		cfg.TablesIGNORE = []string{}
	}
	if *tables != "" {
		_ts := strings.Split(*tables, ",")
		for _, _name := range _ts {
			_name = strings.TrimSpace(_name)
			if _name != "" {
				cfg.Tables = append(cfg.Tables, _name)
			}
		}
	}
	if *tablesIGNORE != "" {
		_ts := strings.Split(*tablesIGNORE, ",")
		for _, _name := range _ts {
			_name = strings.TrimSpace(_name)
			if _name != "" {
				cfg.TablesIGNORE = append(cfg.TablesIGNORE, _name)
			}
		}
	}
	defer (func() {
		if err := recover(); err != nil {
			log.Println(err)
			log.Println(fullStack())
			cfg.SendMailFail(fmt.Sprintf("%s", err))
			log.Fatalln("exit")
		}
	})()

	cfg.Check()
	if *syncData == false {
		internal.CheckSchemaDiff(cfg)
	} else {
		internal.SyncTableData(cfg)
	}

}

func fullStack() string {
	var buf [2 << 11]byte
	runtime.Stack(buf[:], true)
	return string(buf[:])
}
