package dao

const (
	InsertSQL = `
	INSERT IGNORE zg_ag (code, date, jlr, jlrzb, zljlr, zljlrzb, cddjlr, cddjlrzb, ddjlr, ddjlrzb) VALUES (?,?,?,?,?,?,?,?,?,?)
	`

	InsertSDGDSQL = `
	INSERT IGNORE zg_ag_sdgd (code, date, gdmc, gdbh, gdcgsl, gdlb, gfxz, wzdn1, wzdn2, wzdn3,wzdn4) VALUES (?,?,?,?,?,?,?,?,?,?,?)
	`

	//in ('300041','300046')
	// SelectSQL = `
	// SELECT code,date,jlr,zljlr FROM zg_ag  where code in (?)
	// `
)
