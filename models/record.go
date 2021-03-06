package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Dbrecord struct {
	Id         int64
	Schema     string `orm:size(50)`
	Object     string `orm:size(50)`
	Operation  string `orm:size(50)`
	Isbackup   string `orm:size(10)`
	Content    string `orm:size(2000)`
	Attachment string `orm:size(100)`
	Comment    string `orm:szie(100)`
	Operater   string `orm:size(20)`
	Created    string `orm:size(19)`
}

type Apprecord struct {
	Id        int64
	Group     string `orm:size(50)`
	Operation string `orm:size(50)`
	Appname   string `orm:size(100)`
	Disthost  string `orm:size(100)`
	Isauto    string `orm:size(20)`
	Applicant string `orm:size(50)`
	Operater  string `orm:size(20)`
	Content   string `orm:size(2000)`
	Created   string `orm:size(19)`
}

type Faultrecord struct {
	Id        int64
	Question  string `orm:size(50)`
	InterNum  string `orm:size(30)`
	Starttime string `orm:size(19)`
	Endtime   string `orm:size(19)`
	Level     string `orm:size(20)`
	System    string `orm:size(20)`
	Appname   string `orm:size(50)`
	Category  string `orm:size(20)`
	Issolute  string `orm:size(10)`
	Operater  string `orm:size(30)`
	Desc      string `orm:size(1000)`
	Solution  string `orm:size(1000)`
	Effection string `orm:size(1000)`
	Analysis  string `orm:size(1000)`
	Nextstep  string `orm:size(1000)`
}

func init() {
	orm.RegisterModel(new(Dbrecord), new(Apprecord), new(Faultrecord))
}

func GetDBRecordCount() (int64, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	total, err := o.QueryTable("dbrecord").All(&dbRecs)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetDBRecords(currPage, pageSize int) ([]*Dbrecord, int64, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	total, err := o.QueryTable("dbrecord").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&dbRecs)
	if err != nil {
		return nil, 0, err
	}
	return dbRecs, total, err
}

func AddDBRecord(schema, object, operation, backup, content, attachment, comment, operater string) error {
	o := orm.NewOrm()
	record := &Dbrecord{
		Schema:     schema,
		Object:     object,
		Operation:  operation,
		Isbackup:   backup,
		Content:    content,
		Attachment: attachment,
		Comment:    comment,
		Operater:   operater,
		Created:    time.Now().String()[:19],
	}
	_, err := o.Insert(record)
	return err
}

func DeleteDBRecord(id string) error {
	o := orm.NewOrm()
	rid, err := strconv.ParseInt(id, 10, 64)
	dbRec := &Dbrecord{
		Id: rid,
	}
	_, err = o.Delete(dbRec)
	if err != nil {
		return err
	}
	return nil
}

func DBRecordDetail(id string) (*Dbrecord, error) {
	o := orm.NewOrm()
	rid, err := strconv.ParseInt(id, 10, 64)
	dbRec := &Dbrecord{}
	err = o.QueryTable("dbrecord").Filter("id", rid).One(dbRec)
	return dbRec, err
}

func SearchDBRecCount(schema string) (int64, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	total, err := o.QueryTable("dbrecord").Filter("schema__icontains", schema).All(&dbRecs)
	return total, err
}

func SearchDBRecBySchema(currPage, pageSize int, schema string) ([]*Dbrecord, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	_, err := o.QueryTable("dbrecord").Filter("schema__icontains", schema).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&dbRecs)
	return dbRecs, err
}

func GetAppRecordCount() (int64, error) {
	o := orm.NewOrm()
	appRecs := make([]*Apprecord, 0)
	total, err := o.QueryTable("apprecord").All(&appRecs)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetAppRecords(currPage, pageSize int) ([]*Apprecord, int64, error) {
	o := orm.NewOrm()
	appRecs := make([]*Apprecord, 0)
	total, err := o.QueryTable("apprecord").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&appRecs)
	if err != nil {
		return nil, 0, err
	}
	return appRecs, total, err
}

func AddAppRecord(group, operation, appname, disthost, isauto, applicant, content, operater string) error {
	o := orm.NewOrm()
	record := &Apprecord{
		Group:     group,
		Operation: operation,
		Appname:   appname,
		Disthost:  disthost,
		Isauto:    isauto,
		Applicant: applicant,
		Content:   content,
		Operater:  operater,
		Created:   time.Now().String()[:19],
	}
	_, err := o.Insert(record)
	return err
}

func DeleteAppRecord(id string) error {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	appRec := &Apprecord{
		Id: aid,
	}
	_, err = o.Delete(appRec)
	if err != nil {
		return err
	}
	return nil
}

func AppRecordDetail(id string) (*Apprecord, error) {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	appRec := &Apprecord{}
	err = o.QueryTable("apprecord").Filter("id", aid).One(appRec)
	return appRec, err
}

func SearchAppRecCount(appname string) (int64, error) {
	o := orm.NewOrm()
	appRecs := make([]*Apprecord, 0)
	total, err := o.QueryTable("apprecord").Filter("appname__icontains", appname).All(&appRecs)
	return total, err
}

func SearchAppRecByName(currPage, pageSize int, appname string) ([]*Apprecord, error) {
	o := orm.NewOrm()
	appRecs := make([]*Apprecord, 0)
	_, err := o.QueryTable("apprecord").Filter("appname__icontains", appname).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&appRecs)
	return appRecs, err
}

func GetFaultRecordCount(quest string) (int64, error) {
	o := orm.NewOrm()
	var faultRecs []*Faultrecord
	var total int64
	var err error
	faultRecs = make([]*Faultrecord, 0)
	if len(quest) > 0 {
		total, err = o.QueryTable("faultrecord").Filter("question", quest).All(&faultRecs)
		if err != nil {
			return 0, err
		}
	} else {
		total, err = o.QueryTable("faultrecord").All(&faultRecs)
		if err != nil {
			return 0, err
		}
	}
	return total, err
}

func GetFaultRecords(currPage, pageSize int, quest string) ([]*Faultrecord, int64, error) {
	o := orm.NewOrm()
	var faultRecs []*Faultrecord
	var total int64
	var err error
	faultRecs = make([]*Faultrecord, 0)
	if len(quest) > 0 {
		total, err = o.QueryTable("faultrecord").Filter("question", quest).OrderBy("-starttime").Limit(pageSize, (currPage-1)*pageSize).All(&faultRecs)
		if err != nil {
			return nil, 0, err
		}
	} else {
		total, err = o.QueryTable("faultrecord").OrderBy("-starttime").Limit(pageSize, (currPage-1)*pageSize).All(&faultRecs)
		if err != nil {
			return nil, 0, err
		}
	}

	return faultRecs, total, err
}

func AddFaultRecord(question, level, system, appname, category, issolu, operater, starttime, endtime, solution, effection, analysis, nextstep string) error {
	o := orm.NewOrm()
	//now := time.Now().String()
	//beego.Info(endtime);
	interNum := "IE" + endtime[:4] + endtime[5:7] + endtime[8:10] + endtime[11:13] + endtime[14:16] + endtime[17:19]
	//interNum := "IE" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19]
	beego.Info(interNum);
	record := &Faultrecord{
		InterNum:  interNum,
		Question:  question,
		Level:     level,
		System:    system,
		Appname:   appname,
		Category:  category,
		Issolute:  issolu,
		Operater:  operater,
		Starttime: starttime,
		Endtime:   endtime,
		Solution:  solution,
		Effection: effection,
		Analysis:  analysis,
		Nextstep:  nextstep,
	}
	_, err := o.Insert(record)
	if err == nil {
		res, err := o.Raw("update question set fault_count=fault_count+1 where name= ?", question).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			fmt.Println("update question row affected nums: ", num)
		}
	}

	return err
}

func DeleteFaultRecord(id string) error {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	faultRec := &Faultrecord{
		Id: aid,
	}
	_, err = o.Delete(faultRec)
	if err != nil {
		return err
	}
	return nil
}

func FaultRecordDetail(id string) (*Faultrecord, error) {
	o := orm.NewOrm()
	fid, err := strconv.ParseInt(id, 10, 64)
	faultRec := &Faultrecord{}
	err = o.QueryTable("faultrecord").Filter("id", fid).One(faultRec)
	return faultRec, err
}

func SearchFaultRecCount(cate, quest string) (int64, error) {
	o := orm.NewOrm()
	var faultRecs []*Faultrecord
	var total int64
	var err error
	faultRecs = make([]*Faultrecord, 0)
	if len(quest) > 0 {
		total, err = o.QueryTable("faultrecord").Filter("question", quest).Filter("category__icontains", cate).All(&faultRecs)
	} else {
		total, err = o.QueryTable("faultrecord").Filter("category__icontains", cate).All(&faultRecs)
	}
	return total, err
}

func SearchFaultRecByName(currPage, pageSize int, cate, quest string) ([]*Faultrecord, error) {
	o := orm.NewOrm()
	var faultRecs []*Faultrecord
	var err error
	faultRecs = make([]*Faultrecord, 0)
	if len(quest) > 0 {
		_, err = o.QueryTable("faultrecord").Filter("question", quest).Filter("category__icontains", cate).OrderBy("-starttime").Limit(pageSize, (currPage-1)*pageSize).All(&faultRecs)
	} else {
		_, err = o.QueryTable("faultrecord").Filter("category__icontains", cate).OrderBy("-starttime").Limit(pageSize, (currPage-1)*pageSize).All(&faultRecs)
	}
	return faultRecs, err
}

func QueryFaultExport() (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}
	defer conn.Close()

	rows, err := conn.Query("select inter_num,starttime,endtime,level,system,appname,category,issolute,operater,solution,effection,analysis,nextstep from faultrecord")
	if err != nil {
		return &result, columns, total
	}
	defer rows.Close()
	columns, err = rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		var row []string
		_ = rows.Scan(scans...)
		for _, col := range values {
			row = append(row, string(col))
		}
		total = total + 1
		result[total] = row
	}

	return &result, columns, total
}
