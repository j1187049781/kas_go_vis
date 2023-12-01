create table if not exists addr
(	`time` DateTime NOT NULL comment 'time',
     `address` VARCHAR(255) not null comment 'address',
     `balance` DECIMAL(12,9) default '0.0'  comment 'balance'
)
UNIQUE key(`time`, `address`)
comment "addr"
PARTITION BY RANGE(time)
(
    PARTITION p1 VALUES LESS THAN ("2023-12-01"),

    PARTITION p01 VALUES LESS THAN ("2024-01-01"),
    PARTITION p02 VALUES LESS THAN ("2024-02-01"),
    PARTITION p03 VALUES LESS THAN ("2024-03-01"),
    PARTITION p04 VALUES LESS THAN ("2024-04-01"),
    PARTITION p05 VALUES LESS THAN ("2024-05-01"),
    PARTITION p06 VALUES LESS THAN ("2024-06-01"),
    PARTITION p07 VALUES LESS THAN ("2024-07-01"),
    PARTITION p08 VALUES LESS THAN ("2024-08-01"),
    PARTITION p09 VALUES LESS THAN ("2024-09-01"),
    PARTITION p10 VALUES LESS THAN ("2024-10-01"),
    PARTITION p11 VALUES LESS THAN ("2024-11-01"),
    PARTITION p12 VALUES LESS THAN ("2024-12-01"),

    PARTITION p21 VALUES LESS THAN ("2025-01-01"),
    PARTITION p22 VALUES LESS THAN ("2025-02-01"),
    PARTITION p23 VALUES LESS THAN ("2025-03-01"),
    PARTITION p24 VALUES LESS THAN ("2025-04-01"),
    PARTITION p25 VALUES LESS THAN ("2025-05-01"),
    PARTITION p26 VALUES LESS THAN ("2025-06-01"),
    PARTITION p27 VALUES LESS THAN ("2025-07-01"),
    PARTITION p28 VALUES LESS THAN ("2025-08-01"),
    PARTITION p29 VALUES LESS THAN ("2025-09-01"),
    PARTITION p20 VALUES LESS THAN ("2025-10-01"),
    PARTITION p21x VALUES LESS THAN ("2025-11-01"),
    PARTITION p22x VALUES LESS THAN ("2025-12-01"),

    PARTITION p51 VALUES LESS THAN ("2026-01-01"),
    PARTITION p52 VALUES LESS THAN ("2026-02-01"),
    PARTITION p53 VALUES LESS THAN ("2026-03-01"),
    PARTITION p54 VALUES LESS THAN ("2026-04-01"),
    PARTITION p55 VALUES LESS THAN ("2026-05-01"),
    PARTITION p56 VALUES LESS THAN ("2026-06-01"),
    PARTITION p57 VALUES LESS THAN ("2026-07-01"),
    PARTITION p58 VALUES LESS THAN ("2026-08-01"),
    PARTITION p59 VALUES LESS THAN ("2026-09-01"),
    PARTITION p50 VALUES LESS THAN ("2026-10-01"),
    PARTITION p51x VALUES LESS THAN ("2026-11-01"),
    PARTITION p52x VALUES LESS THAN ("2026-12-01")
    
)
Distributed BY Hash(`address`) BUCKETS 10
PROPERTIES (
    "replication_num" = "1"
);


create table if not exists addr_tag
(
     `address` VARCHAR(255) not null comment 'address',
     `name` DECIMAL(12,9) default '0.0'  comment 'name',
     `link` string  null comment 'link'
)
UNIQUE key(`address`, `name`)
comment "addr_tag"
Distributed BY Hash(`address`) BUCKETS 1
PROPERTIES (
    "replication_num" = "1"
);