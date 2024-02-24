create table orders (
        id varchar(36) unique, 
        company_id varchar(36) not null, 
        created_by varchar(36),
        title text, 
        created_at timestamp default CURRENT_TIMESTAMP, 
        updated_at timestamp default CURRENT_TIMESTAMP,
        deleted_at timestamp default null, 
        closed_at timestamp default null,
        price float default 0, 
        prime_cost float default 0,
        primary key (id)
);
create index index_orders_id_company_id
        on orders (id, company_id);
create index index_orders_id_company_id_created_by
        on orders (id, company_id, created_by);

create table estimates (
        id varchar(36) unique, 
        order_id varchar(36) REFERENCES orders.id, 
        company_id varchar(36) not null,
        created_at timestamp default CURRENT_TIMESTAMP, 
        updated_at timestamp default CURRENT_TIMESTAMP
);
create index index_estimates_id_order_id
        on estimates (id, order_id, company_id);

create table estimate_positions (
        id varchar(36) primary key, 
        estimate_id varchar(36) not null, 
        company_id varchar(36) not null, 
        created_at timestamp default CURRENT_TIMESTAMP, 
        updated_at timestamp default CURRENT_TIMESTAMP,
        amount float default 0,
        price float default 0, 
        total_price float default 0, 
        prime_cost float default 0, 
        prime_cost_total float default 0
)
create index index_estimate_positions_id_estimate_id
        on estimate_positions (id, estimate_id);
create index index_estimate_positions_id_estimate_id_company_id
        on estimate_positions (id, estimate_id, company_id);