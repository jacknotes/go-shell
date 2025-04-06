package dao

const (
	InsertSQL = `
	INSERT IGNORE zg_ag (code, date, jlr, jlrzb, zljlr, zljlrzb, cddjlr, cddjlrzb, ddjlr, ddjlrzb) VALUES (?,?,?,?,?,?,?,?,?,?)
	`

	SelectSQL = `
	SELECT code,date,jlr,zljlr FROM zg_ag  
	`
)
