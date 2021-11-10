package ast

func isLetter(ch uint16) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func digitVal(ch uint16) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}

func isDigit(ch uint16) bool {
	return '0' <= ch && ch <= '9'
}

const LEX_ERROR = 57346
const UNION = 57347
const SELECT = 57348
const INSERT = 57349
const UPDATE = 57350
const DELETE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const OFFSET = 57359
const FOR = 57360
const ALL = 57361
const DISTINCT = 57362
const AS = 57363
const EXISTS = 57364
const ASC = 57365
const DESC = 57366
const INTO = 57367
const DUPLICATE = 57368
const KEY = 57369
const DEFAULT = 57370
const SET = 57371
const LOCK = 57372
const VALUES = 57373
const LAST_INSERT_ID = 57374
const NEXT = 57375
const VALUE = 57376
const SHARE = 57377
const MODE = 57378
const SQL_NO_CACHE = 57379
const SQL_CACHE = 57380
const JOIN = 57381
const STRAIGHT_JOIN = 57382
const LEFT = 57383
const RIGHT = 57384
const INNER = 57385
const OUTER = 57386
const CROSS = 57387
const NATURAL = 57388
const USE = 57389
const FORCE = 57390
const ON = 57391
const ID = 57392
const HEX = 57393
const STRING = 57394
const INTEGRAL = 57395
const FLOAT = 57396
const HEXNUM = 57397
const VALUE_ARG = 57398
const LIST_ARG = 57399
const COMMENT = 57400
const NULL = 57401
const TRUE = 57402
const FALSE = 57403
const OR = 57404
const AND = 57405
const NOT = 57406
const BETWEEN = 57407
const CASE = 57408
const WHEN = 57409
const THEN = 57410
const ELSE = 57411
const END = 57412
const LE = 57413
const GE = 57414
const NE = 57415
const NULL_SAFE_EQUAL = 57416
const IS = 57417
const LIKE = 57418
const REGEXP = 57419
const IN = 57420
const SHIFT_LEFT = 57421
const SHIFT_RIGHT = 57422
const DIV = 57423
const MOD = 57424
const UNARY = 57425
const COLLATE = 57426
const BINARY = 57427
const INTERVAL = 57428
const JSON_EXTRACT_OP = 57429
const JSON_UNQUOTE_EXTRACT_OP = 57430
const CREATE = 57431
const ALTER = 57432
const DROP = 57433
const RENAME = 57434
const ANALYZE = 57435
const TABLE = 57436
const INDEX = 57437
const VIEW = 57438
const TO = 57439
const IGNORE = 57440
const IF = 57441
const UNIQUE = 57442
const USING = 57443
const SHOW = 57444
const DESCRIBE = 57445
const EXPLAIN = 57446
const DATE = 57447
const ESCAPE = 57448
const REPAIR = 57449
const OPTIMIZE = 57450
const TRUNCATE = 57451
const DATABASES = 57452
const TABLES = 57453
const VITESS_KEYSPACES = 57454
const VITESS_SHARDS = 57455
const VSCHEMA_TABLES = 57456
const INTEGER = 57457
const CHARACTER = 57458
const CURRENT_TIMESTAMP = 57459
const DATABASE = 57460
const CURRENT_DATE = 57461
const CURRENT_TIME = 57462
const LOCALTIME = 57463
const LOCALTIMESTAMP = 57464
const UTC_DATE = 57465
const UTC_TIME = 57466
const UTC_TIMESTAMP = 57467
const REPLACE = 57468
const CONVERT = 57469
const CAST = 57470
const GROUP_CONCAT = 57471
const SEPARATOR = 57472
const MATCH = 57473
const AGAINST = 57474
const BOOLEAN = 57475
const LANGUAGE = 57476
const WITH = 57477
const QUERY = 57478
const EXPANSION = 57479
const UNUSED = 57480

var keywords = map[string]int{
	"accessible":          UNUSED,
	"add":                 UNUSED,
	"against":             AGAINST,
	"all":                 ALL,
	"alter":               ALTER,
	"analyze":             ANALYZE,
	"and":                 AND,
	"as":                  AS,
	"asc":                 ASC,
	"asensitive":          UNUSED,
	"before":              UNUSED,
	"between":             BETWEEN,
	"bigint":              UNUSED,
	"binary":              BINARY,
	"blob":                UNUSED,
	"boolean":             BOOLEAN,
	"both":                UNUSED,
	"by":                  BY,
	"call":                UNUSED,
	"cascade":             UNUSED,
	"case":                CASE,
	"cast":                CAST,
	"change":              UNUSED,
	"character":           CHARACTER,
	"check":               UNUSED,
	"collate":             COLLATE,
	"column":              UNUSED,
	"condition":           UNUSED,
	"constraint":          UNUSED,
	"continue":            UNUSED,
	"convert":             CONVERT,
	"create":              CREATE,
	"cross":               CROSS,
	"current_date":        CURRENT_DATE,
	"current_time":        CURRENT_TIME,
	"current_timestamp":   CURRENT_TIMESTAMP,
	"current_user":        UNUSED,
	"cursor":              UNUSED,
	"database":            DATABASE,
	"databases":           DATABASES,
	"day_hour":            UNUSED,
	"day_microsecond":     UNUSED,
	"day_minute":          UNUSED,
	"day_second":          UNUSED,
	"date":                DATE,
	"dec":                 UNUSED,
	"declare":             UNUSED,
	"default":             DEFAULT,
	"delayed":             UNUSED,
	"delete":              DELETE,
	"desc":                DESC,
	"describe":            DESCRIBE,
	"deterministic":       UNUSED,
	"distinct":            DISTINCT,
	"distinctrow":         UNUSED,
	"div":                 DIV,
	"double":              UNUSED,
	"drop":                DROP,
	"duplicate":           DUPLICATE,
	"each":                UNUSED,
	"else":                ELSE,
	"elseif":              UNUSED,
	"enclosed":            UNUSED,
	"end":                 END,
	"escape":              ESCAPE,
	"escaped":             UNUSED,
	"exists":              EXISTS,
	"exit":                UNUSED,
	"explain":             EXPLAIN,
	"expansion":           EXPANSION,
	"false":               FALSE,
	"fetch":               UNUSED,
	"float":               UNUSED,
	"float4":              UNUSED,
	"float8":              UNUSED,
	"for":                 FOR,
	"force":               FORCE,
	"foreign":             UNUSED,
	"from":                FROM,
	"fulltext":            UNUSED,
	"generated":           UNUSED,
	"get":                 UNUSED,
	"grant":               UNUSED,
	"group":               GROUP,
	"group_concat":        GROUP_CONCAT,
	"having":              HAVING,
	"high_priority":       UNUSED,
	"hour_microsecond":    UNUSED,
	"hour_minute":         UNUSED,
	"hour_second":         UNUSED,
	"if":                  IF,
	"ignore":              IGNORE,
	"in":                  IN,
	"index":               INDEX,
	"infile":              UNUSED,
	"inout":               UNUSED,
	"inner":               INNER,
	"insensitive":         UNUSED,
	"insert":              INSERT,
	"int":                 UNUSED,
	"int1":                UNUSED,
	"int2":                UNUSED,
	"int3":                UNUSED,
	"int4":                UNUSED,
	"int8":                UNUSED,
	"integer":             INTEGER,
	"interval":            INTERVAL,
	"into":                INTO,
	"io_after_gtids":      UNUSED,
	"is":                  IS,
	"iterate":             UNUSED,
	"join":                JOIN,
	"key":                 KEY,
	"keys":                UNUSED,
	"kill":                UNUSED,
	"language":            LANGUAGE,
	"last_insert_id":      LAST_INSERT_ID,
	"leading":             UNUSED,
	"leave":               UNUSED,
	"left":                LEFT,
	"like":                LIKE,
	"limit":               LIMIT,
	"linear":              UNUSED,
	"lines":               UNUSED,
	"load":                UNUSED,
	"localtime":           LOCALTIME,
	"localtimestamp":      LOCALTIMESTAMP,
	"lock":                LOCK,
	"long":                UNUSED,
	"longblob":            UNUSED,
	"longtext":            UNUSED,
	"loop":                UNUSED,
	"low_priority":        UNUSED,
	"master_bind":         UNUSED,
	"match":               MATCH,
	"maxvalue":            UNUSED,
	"mediumblob":          UNUSED,
	"mediumint":           UNUSED,
	"mediumtext":          UNUSED,
	"middleint":           UNUSED,
	"minute_microsecond":  UNUSED,
	"minute_second":       UNUSED,
	"mod":                 MOD,
	"mode":                MODE,
	"modifies":            UNUSED,
	"natural":             NATURAL,
	"next":                NEXT,
	"not":                 NOT,
	"no_write_to_binlog":  UNUSED,
	"null":                NULL,
	"numeric":             UNUSED,
	"offset":              OFFSET,
	"on":                  ON,
	"optimize":            OPTIMIZE,
	"optimizer_costs":     UNUSED,
	"option":              UNUSED,
	"optionally":          UNUSED,
	"or":                  OR,
	"order":               ORDER,
	"out":                 UNUSED,
	"outer":               OUTER,
	"outfile":             UNUSED,
	"partition":           UNUSED,
	"precision":           UNUSED,
	"primary":             UNUSED,
	"procedure":           UNUSED,
	"query":               QUERY,
	"range":               UNUSED,
	"read":                UNUSED,
	"reads":               UNUSED,
	"read_write":          UNUSED,
	"real":                UNUSED,
	"references":          UNUSED,
	"regexp":              REGEXP,
	"release":             UNUSED,
	"rename":              RENAME,
	"repair":              REPAIR,
	"repeat":              UNUSED,
	"replace":             REPLACE,
	"require":             UNUSED,
	"resignal":            UNUSED,
	"restrict":            UNUSED,
	"return":              UNUSED,
	"revoke":              UNUSED,
	"right":               RIGHT,
	"rlike":               REGEXP,
	"schema":              UNUSED,
	"schemas":             UNUSED,
	"second_microsecond":  UNUSED,
	"select":              SELECT,
	"sensitive":           UNUSED,
	"separator":           SEPARATOR,
	"set":                 SET,
	"share":               SHARE,
	"show":                SHOW,
	"signal":              UNUSED,
	"smallint":            UNUSED,
	"spatial":             UNUSED,
	"specific":            UNUSED,
	"sql":                 UNUSED,
	"sqlexception":        UNUSED,
	"sqlstate":            UNUSED,
	"sqlwarning":          UNUSED,
	"sql_big_result":      UNUSED,
	"sql_cache":           SQL_CACHE,
	"sql_calc_found_rows": UNUSED,
	"sql_no_cache":        SQL_NO_CACHE,
	"sql_small_result":    UNUSED,
	"ssl":                 UNUSED,
	"starting":            UNUSED,
	"stored":              UNUSED,
	"straight_join":       STRAIGHT_JOIN,
	"table":               TABLE,
	"tables":              TABLES,
	"terminated":          UNUSED,
	"then":                THEN,
	"tinyblob":            UNUSED,
	"tinyint":             UNUSED,
	"tinytext":            UNUSED,
	"to":                  TO,
	"trailing":            UNUSED,
	"trigger":             UNUSED,
	"true":                TRUE,
	"truncate":            TRUNCATE,
	"undo":                UNUSED,
	"union":               UNION,
	"unique":              UNIQUE,
	"unlock":              UNUSED,
	"update":              UPDATE,
	"usage":               UNUSED,
	"use":                 USE,
	"using":               USING,
	"utc_date":            UTC_DATE,
	"utc_time":            UTC_TIME,
	"utc_timestamp":       UTC_TIMESTAMP,
	"values":              VALUES,
	"varbinary":           UNUSED,
	"varchar":             UNUSED,
	"varcharacter":        UNUSED,
	"varying":             UNUSED,
	"virtual":             UNUSED,
	"view":                VIEW,
	"vitess_keyspaces":    VITESS_KEYSPACES,
	"vitess_shards":       VITESS_SHARDS,
	"vschema_tables":      VSCHEMA_TABLES,
	"when":                WHEN,
	"where":               WHERE,
	"while":               UNUSED,
	"with":                WITH,
	"write":               UNUSED,
	"xor":                 UNUSED,
	"year_month":          UNUSED,
	"zerofill":            UNUSED,
}
