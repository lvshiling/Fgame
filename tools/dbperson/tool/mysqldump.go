package tool

import (
	"bytes"
	"fgame/fgame/tools/dbperson/model"
	"fmt"
	"os/exec"
)

func BackUpDb(db *model.DBConfigInfo, p_savePath string, p_dumpPath string) error {

	// dumpdb := getBackUpDbDump("", db, p_savePath)
	// fmt.Println(dumpdb)
	// c := exec.Command("cmd", "/C", dumpdb)
	arg := make([]string, 0)
	arg = append(arg, "-h"+db.Host)
	arg = append(arg, fmt.Sprintf("-P%d", db.Port))
	arg = append(arg, fmt.Sprintf("-u%s", db.UserName))
	arg = append(arg, fmt.Sprintf("-p%s", db.PassWord))
	arg = append(arg, "--set-gtid-purged=off")
	arg = append(arg, "--default-character-set=utf8mb4")
	arg = append(arg, "--skip-add-locks")
	arg = append(arg, "--skip-lock-tables")
	// arg = append(arg, "-t") --备份的时候将表结构也备份出来
	arg = append(arg, db.DBName)
	arg = append(arg, fmt.Sprintf("--result-file=%s", p_savePath))
	c := exec.Command(p_dumpPath, arg...)
	w := bytes.NewBuffer(nil)
	c.Stderr = w

	if err := c.Run(); err != nil {
		return fmt.Errorf("BackUpDb: %s\n", string(w.Bytes()))

	}

	return nil
}

func BackUpTable(db *model.DBConfigInfo, p_table string, p_where string, p_savePath string, p_dumpPath string) error {

	// dumpdb := getBackUpDbDump("", db, p_savePath)
	// fmt.Println(dumpdb)
	// c := exec.Command("cmd", "/C", dumpdb)
	arg := make([]string, 0)
	arg = append(arg, "-h"+db.Host)
	arg = append(arg, fmt.Sprintf("-P%d", db.Port))
	arg = append(arg, fmt.Sprintf("-u%s", db.UserName))
	arg = append(arg, fmt.Sprintf("-p%s", db.PassWord))
	arg = append(arg, "--default-character-set=utf8mb4")
	arg = append(arg, "--skip-add-locks")
	arg = append(arg, "--set-gtid-purged=off")
	arg = append(arg, "-t")
	arg = append(arg, "-c") //增加导出列名，避免两库顺序不一致
	arg = append(arg, "--skip-lock-tables")
	arg = append(arg, fmt.Sprintf("%s", db.DBName)) //--databases
	arg = append(arg, fmt.Sprintf("%s", p_table))   //--tables
	if len(p_where) > 0 {
		arg = append(arg, fmt.Sprintf("--where=%s", p_where))
	}
	arg = append(arg, fmt.Sprintf("--result-file=%s", p_savePath))
	c := exec.Command(p_dumpPath, arg...)
	w := bytes.NewBuffer(nil)
	c.Stderr = w

	if err := c.Run(); err != nil {
		return fmt.Errorf("BackUpTable: %s\n", string(w.Bytes()))
	}

	return nil
}

func ImportSqlPath(db *model.DBConfigInfo, p_sqlPath string, p_mysqlPath string) error {
	arg := make([]string, 0)
	arg = append(arg, "-h"+db.Host)
	arg = append(arg, fmt.Sprintf("-P%d", db.Port))
	arg = append(arg, fmt.Sprintf("-u%s", db.UserName))
	arg = append(arg, fmt.Sprintf("-p%s", db.PassWord))
	arg = append(arg, fmt.Sprintf("--database=%s", db.DBName)) //--databases
	arg = append(arg, "--default-character-set=utf8mb4")
	arg = append(arg, "--named-commands")
	arg = append(arg, fmt.Sprintf("--execute=source %s", p_sqlPath))
	// arg = append(arg, fmt.Sprintf("--init-command=source %s", p_sqlPath)) //--databases

	c := exec.Command(p_mysqlPath, arg...)
	w := bytes.NewBuffer(nil)
	c.Stderr = w

	if err := c.Run(); err != nil {
		fmt.Println(err)
		return fmt.Errorf("ImportSqlPath: %s\n", string(w.Bytes()))
	}

	return nil
}
