package dataManager

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"log"
	"../entities"
	"../config"
)
var userCount int
var like string

type ManagerDb struct {
	db *sql.DB
}
var mdb *ManagerDb

func init() {
	mdb = GetInstance()
}

func (md *ManagerDb) GetCount(filter string) int{
	if userCount == 0 || filter != like {
		like = filter
		log.Println("init count")
		sql:="SELECT count(`id_resume`) FROM `resume`"
		if filter != "" {
			sql+=" WHERE LOWER(position) LIKE LOWER(?)"
		}

		stmt, err  := mdb.db.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		if filter != "" {
			row := stmt.QueryRow("%"+filter+"%")
			err = row.Scan(&userCount)
		}else {
			mdb.db.QueryRow(sql).Scan(&userCount)

		}

	}
	return  userCount
}

func (md *ManagerDb) IsExist(patch string) bool{
	var count int
	mdb.db.QueryRow("SELECT `id` FROM `file` where `patch` = ?",patch).Scan(&count)
	return  count != 0
}
func (md *ManagerDb) IsExistByResume(idResume string) bool{
	var count int
	 mdb.db.QueryRow("SELECT `id` FROM `file` where `id_resume` = ?",idResume).Scan(&count)
	return  count != 0
}

func (md *ManagerDb) GetRowsWithFiles(limit int, offset int,filter string)([]*entities.User,error) {

	var(
		rows *sql.Rows
		err error
	)
	sql:= `
SELECT
	r.id_resume AS id,
	r.first_name,
	r.last_name,
	r.email,
	r.phone,
	r.salary,
	r.position,
COALESCE((SELECT f.patch  FROM file as f  WHERE f.id_resume = r.id_resume ),'') as  patch 
FROM 
(select id_resume,first_name,last_name,email,phone,salary,id_file,position FROM resume 
`
	if filter!="" {
		sql+="WHERE LOWER(position) LIKE LOWER(?)"
	}


	sql+=
` ORDER BY id_resume LIMIT ? OFFSET ?) as r
 LEFT JOIN file AS f
ON f.id = r.id_file
`

	stmt, err  := mdb.db.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer stmt.Close()
	if filter != "" {
		rows, err = stmt.Query("%"+filter+"%",limit,offset)
	}else{
		rows, err = stmt.Query(limit,offset)
	}

	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer rows.Close()
	items := make([]*entities.User, limit)
	key:= 0
	for rows.Next() {
		item := new(entities.User)
		err := rows.Scan(&item.Id,&item.FirstName,&item.LastName,&item.Email,&item.Phone,&item.Salary,&item.Position,&item.Path)
		if err != nil {
			log.Fatal(err)
		}
		items[key] = item
		key++
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil,err
	}
	return items,err
}

func GetInstance() *ManagerDb {
	var err error
	var db *sql.DB
	if mdb == nil {
		params:= config.GetConf().Db
		db, err = sql.Open("mysql", params.DbUser+":"+params.DbPass+"@/"+params.DbName)
		if err != nil {
			panic(err)
		}
		mdb = &ManagerDb{
			db: db,
		}
		return mdb
	}
	return mdb
}