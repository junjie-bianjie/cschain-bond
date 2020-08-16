--   债券品种维度表
drop table if exists bond_variety_dimension
create table bond_variety_dimension
(
    id int auto_increment,
    name varchar(255) null,
    parent_id int null,
    level int null,
    constraint bond_variety_dimension_pk
        primary key (id)
);

-- 债券回购品种维度表
drop table if exists bond_repurchase_variety_dimension
create table bond_repurchase_variety_dimension
(
    id int auto_increment,
    name varchar(255) null,
    parent_id int null,
    level int null,
    constraint bond_variety_dimension_pk
        primary key (id)
);

-- 债券交易记录表
drop table if exists bond_transaction
create table bond_transaction
(
    id                  int auto_increment
        primary key,
    nft_id              varchar(255) null,
    source_type         int          null,
    denom_id            varchar(255) null,
    owner               varchar(255) null,
    uri                 varchar(255) null,
    visible             tinyint(1)   null,
    amount              double       null,
    market              varchar(255) null,
    start_date          date         null,
    end_date            date         null,
    period_category     varchar(255) null,
    bond_category       int          null,
    repurchase_category int          null
);