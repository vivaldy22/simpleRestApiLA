create database simple_rest_la;
use simple_rest_la;

create table tb_customer (
     customer_number varchar(50) primary key not null ,
     name varchar(50) not null
);

create table tb_account (
    account_number varchar(50) primary key not null ,
    customer_number varchar(50) not null ,
    balance int default 0,
    foreign key (customer_number) references tb_customer(customer_number)
);

insert into tb_customer
    values ('1001', 'Bob Martin'),
           ('1002', 'Linus Torvalds');

insert into tb_account
    values ('555001', '1001', '10000'),
           ('555002', '1002', '15000');